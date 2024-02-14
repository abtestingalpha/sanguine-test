package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	_ "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	pubsubrouter "github.com/libp2p/go-libp2p-pubsub-router"
	record "github.com/libp2p/go-libp2p-record"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"log"
	"strconv"
	"time"
)

type LibP2PManager interface {
	Host() host.Host // Expose host from manager
	// Start starts the libp2p manager.
	Start(ctx context.Context, bootstrapPeers []string) error
	DoSomething()
}

type libP2PManagerImpl struct {
	host              host.Host
	announcementTopic *pubsub.Topic
	psRouter          *pubsubrouter.PubsubValueStore
	myI               int
}

// NewLibP2PManager creates a new libp2p manager.
// listenHost is the host to listen on.
//
// TODO: we need to figure out how this works across multiple nodes
func NewLibP2PManager(auth signer.Signer) (LibP2PManager, error) {
	h, err := createHost(auth.PrivKey()) // call createHost function
	if err != nil {
		return nil, err
	}

	return &libP2PManagerImpl{host: h}, nil
}

// Host returns the host from the manager.
// TODO: consider not exposing the host directly
func (l *libP2PManagerImpl) Host() host.Host {
	return l.host
}

func createHost(privKeyWrapper crypto.PrivKey) (host.Host, error) {
	// Create a new libp2p host
	h, err := libp2p.New(libp2p.Identity(privKeyWrapper), libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	return h, nil
}

var i = 0

func (l *libP2PManagerImpl) Start(ctx context.Context, bootstrapPeers []string) error {
	// Connect to the bootstrap peers
	for _, peerAddr := range bootstrapPeers {
		err := connectToPeer(ctx, l.host, peerAddr)
		if err != nil {
			log.Printf("Warning: Could not connect to bootstrap peer %s: %v", peerAddr, err)
		} else {
			log.Printf("Connected to bootstrap peer %s", peerAddr)
		}
	}
	// TODO: initialize peer discovery w/ dht
	// https://github.com/libp2p/go-libp2p/blob/66a20a8f530ed09baae8200c92ddbba161a3b5c0/examples/pubsub/basic-chat-with-rendezvous/main.go#L51

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, l.host)
	if err != nil {
		return fmt.Errorf("could not create pubsub: %w", err)
	}

	// todo add:
	// pubsubrouter.WithDatastore() and use our db.
	l.psRouter, err = pubsubrouter.NewPubsubValueStore(ctx, l.host, ps, l, pubsubrouter.WithRebroadcastInterval(time.Second))
	if err != nil {
		return fmt.Errorf("could not create pubsub value store: %w", err)
	}

	cpI := i

	go func() {
		time.Sleep(time.Second)
		err = l.psRouter.PutValue(ctx, fmt.Sprintf("%d", cpI), []byte("testfag"))
		if err != nil {

		}
	}()
	l.myI = i
	i++

	return nil
}

func (l *libP2PManagerImpl) DoSomething() {
	yo, err := l.psRouter.GetValue(context.Background(), "test-0")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(yo))
}

func (l *libP2PManagerImpl) Validate(key string, value []byte) error {
	theirI, _ := strconv.Atoi(key)
	if theirI != l.myI {
		fmt.Printf("validate: %s, (my i is %d)\n", key, l.myI)
	}
	// TODO: validate sig, etc
	return nil
}

func (l *libP2PManagerImpl) Select(key string, values [][]byte) (int, error) {
	// TODO: implement me
	return 0, nil
}

var _ record.Validator = &libP2PManagerImpl{}

func connectToPeer(ctx context.Context, h host.Host, multiAddrString string) error {
	maddr, err := multiaddr.NewMultiaddr(multiAddrString)
	if err != nil {
		return err
	}

	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		return err
	}

	h.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)
	return h.Connect(ctx, *info)
}
