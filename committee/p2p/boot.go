package p2p

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	ipfs_datastore "github.com/ipfs/go-datastore/sync"
	crdt "github.com/ipfs/go-ds-crdt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
	"github.com/phayes/freeport"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"log"
	"time"
)

type LibP2PManager interface {
	Host() host.Host // Expose host from manager
	// Start starts the libp2p manager.
	Start(ctx context.Context, bootstrapPeers []string) error
	DoSomething()
	DoSomethingElse() bool
}

type libP2PManagerImpl struct {
	host              host.Host
	dht               *dual.DHT
	announcementTopic *pubsub.Topic
	pubsub            *pubsub.PubSub
	pubSubBroadcaster crdt.Broadcaster
	globalDS          datastore.Batching
	datastoreDs       *crdt.Datastore
	discovery         discovery.Discovery
}

const dbTopic = "crdt_db"

var RebroadcastingInterval = time.Minute

// NewLibP2PManager creates a new libp2p manager.
// listenHost is the host to listen on.
//
// TODO: we need to figure out how this works across multiple nodes.
func NewLibP2PManager(ctx context.Context, auth signer.Signer) (LibP2PManager, error) {
	l := &libP2PManagerImpl{}
	_, err := l.setupHost(ctx, auth.PrivKey()) // call createHost function
	if err != nil {
		return nil, err
	}

	l.globalDS = ipfs_datastore.MutexWrap(datastore.NewMapDatastore())

	return l, nil
}

// Host returns the host from the manager.
// TODO: consider not exposing the host directly.
func (l *libP2PManagerImpl) Host() host.Host {
	return l.host
}

func (l *libP2PManagerImpl) DoSomething() {
	var err error
	for i := 500 - 1; i >= 0; i-- {
		err = l.datastoreDs.Put(context.Background(), datastore.NewKey(gofakeit.Word()), []byte("test"))
	}

	err = l.datastoreDs.Put(context.Background(), datastore.NewKey("test"), []byte("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}

	err = l.datastoreDs.Sync(context.Background(), datastore.NewKey("/"))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (l *libP2PManagerImpl) DoSomethingElse() bool {
	fmt.Println(len(l.pubsub.ListPeers(dbTopic)))
	val, err := l.datastoreDs.Get(context.Background(), datastore.NewKey("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}

	return val != nil
}

func (l *libP2PManagerImpl) setupHost(ctx context.Context, privKeyWrapper crypto.PrivKey) (host.Host, error) {
	port, _ := freeport.GetFreePort()
	// Create a new libp2p host
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	if err != nil {
		return nil, errors.Wrap(err, "create multi addr")
	}

	ds := ipfs_datastore.MutexWrap(datastore.NewMapDatastore())

	opts := ipfslite.Libp2pOptionsExtra
	opts = append(opts, libp2p.Ping(true))
	// todo: setup datastore
	// TODO: add eth connection gater: https://github.com/dTelecom/p2p-realtime-database/blob/main/gater.go
	l.host, l.dht, err = ipfslite.SetupLibp2p(ctx, privKeyWrapper, nil, []multiaddr.Multiaddr{sourceMultiAddr}, ds, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	return l.host, nil
}

func (l *libP2PManagerImpl) Start(ctx context.Context, bootstrapPeers []string) error {
	// setup ipfs
	peers, err := makePeers(bootstrapPeers)
	if err != nil {
		return err
	}

	ipfs, err := ipfslite.New(ctx, l.globalDS, nil, l.host, l.dht, &ipfslite.Config{})
	go l.Discover(ctx, l.host, l.dht, dbTopic)

	l.pubsub, err = pubsub.NewGossipSub(ctx, l.host)
	err = l.pubsub.RegisterTopicValidator(dbTopic, func(ctx context.Context, from peer.ID, msg *pubsub.Message) bool {
		return true
	})
	if err != nil {
		return fmt.Errorf("could not register topic validator: %w", err)
	}

	ipfs.Bootstrap(peers)
	for _, p := range peers {
		l.host.ConnManager().TagPeer(p.ID, "keep", 100)
	}

	if err != nil {
		return fmt.Errorf("could not create pubsub: %w", err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("pubsub peers: ", len(l.pubsub.ListPeers(dbTopic)))
			fmt.Println("global peers: ", len(l.host.Peerstore().Peers()))
		}
	}()

	l.pubSubBroadcaster, err = crdt.NewPubSubBroadcaster(ctx, l.pubsub, dbTopic)
	if err != nil {
		return err
	}

	crdtOpts := crdt.DefaultOptions()
	crdtOpts.Logger = logging.Logger("p2p_logger")

	crdtOpts.RebroadcastInterval = RebroadcastingInterval
	crdtOpts.MaxBatchDeltaSize = 1

	crdtOpts.PutHook = func(k datastore.Key, v []byte) {
		fmt.Printf("[%s] Added: [%s] -> %s\n", time.Now().Format(time.RFC3339), k, string(v))
		// TODO: some validation goes here
	}
	// TODO: this probably never gets called
	crdtOpts.DeleteHook = func(k datastore.Key) {
		fmt.Printf("[%s] Removed: [%s]\n", time.Now().Format(time.RFC3339), k)
	}
	crdtOpts.RebroadcastInterval = time.Second

	l.datastoreDs, err = crdt.New(l.globalDS, datastore.NewKey(dbTopic), ipfs, l.pubSubBroadcaster, crdtOpts)
	if err != nil {
		return err
	}

	err = l.datastoreDs.Sync(ctx, datastore.NewKey("/"))

	return nil
}

func (l *libP2PManagerImpl) Discover(ctx context.Context, h host.Host, dht *dual.DHT, rendezvous string) {
	routingDiscovery := routing.NewRoutingDiscovery(dht)
	util.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			peers, err := util.FindPeers(ctx, routingDiscovery, rendezvous)
			if err != nil {
				log.Fatal(err)
			}

			for _, p := range peers {
				if p.ID == h.ID() {
					continue
				}
				if h.Network().Connectedness(p.ID) != network.Connected {
					_, err = h.Network().DialPeer(ctx, p.ID)
					fmt.Printf("Connected to peer %s\n", p.ID)
					if err != nil {
						continue
					}
				}
			}
		}
	}
}
func makePeers(peers []string) ([]peer.AddrInfo, error) {
	var p []peer.AddrInfo
	for _, addr := range peers {
		maddr, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, fmt.Errorf("could not create multiaddr: %w", err)
		}
		info, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			return nil, fmt.Errorf("could not create peer info: %w", err)
		}
		p = append(p, *info)
	}
	return p, nil
}
