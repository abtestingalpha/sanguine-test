package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniConfig "github.com/synapsecns/sanguine/services/omnirpc/config"
	"gopkg.in/yaml.v2"
)

type chainConfig struct {
	Name        string `yaml:"name"`
	ChainID     uint32 `yaml:"chain-id"`
	Port        int    `yaml:"port"`
	Deployments map[string]deploymentConfig
}

func (c chainConfig) getChainClient(port uint16) (chainClient chain.Chain, err error) {
	chainURL := fmt.Sprintf("http://localhost:%d/rpc/%d", port, c.ChainID)
	return chain.NewFromURL(context.Background(), chainURL)
}

type dockerComposeConfig struct {
	Services map[string]chainConfig `yaml:"services"`
}

type deploymentConfig struct {
	Name            string
	ContractAddress string
	Contract        interface{}
}

func (d *deploymentConfig) loadContract(chainClient chain.Chain) (err error) {
	// TODO: respect context
	ctx := context.Background()
	switch d.Name {
	case "PingPongClient":
		d.Contract, err = evm.NewPingPongClientContract(ctx, chainClient, common.HexToAddress(d.ContractAddress))
		if err != nil {
			return err
		}
	default:
		err = fmt.Errorf("unknown contract %s", d.Name)
	}
	return err
}

func getChainConfigs(dockerPath string) (chainConfigs map[uint32]chainConfig, err error) {
	// Read the Docker Compose YAML file.
	dockerComposePath := fmt.Sprintf("%s/%s", dockerPath, dockerComposeFile)
	data, err := os.ReadFile(dockerComposePath)
	if err != nil {
		return chainConfigs, err
	}

	// Parse the YAML data into a dockerComposeConfig struct.
	var dockerComposeConfig dockerComposeConfig
	err = yaml.Unmarshal(data, &dockerComposeConfig)
	if err != nil {
		return chainConfigs, err
	}

	chainConfigs = map[uint32]chainConfig{}
	for _, chainConfig := range dockerComposeConfig.Services {
		if chainConfig.ChainID > 0 {
			chainConfigs[chainConfig.ChainID] = chainConfig
		}
	}
	return chainConfigs, err
}

func loadOmniRPCConfig(dockerPath string) (omniRPCConfig omniConfig.Config, err error) {
	omniRPCPath := fmt.Sprintf("%s/config/%s", dockerPath, omnirpcConfig)
	data, err := os.ReadFile(omniRPCPath)
	if err != nil {
		return
	}
	return omniConfig.UnmarshallConfig(data)
}

func loadDeployments(contractName, deploymentPath string, chainConfigs map[uint32]chainConfig, omniRPCConfig omniConfig.Config) (err error) {
	for chainID, chainConfig := range chainConfigs {
		contractABIPath := fmt.Sprintf("%s/%s/%s", deploymentPath, chainConfig.Name, contractName)
		abiFile, err := os.Open(contractABIPath)
		if err != nil {
			return err
		}
		defer abiFile.Close()

		abiBytes, err := io.ReadAll(abiFile)
		if err != nil {
			return err
		}

		var data map[string]interface{}
		err = json.Unmarshal(abiBytes, &data)
		if err != nil {
			return err
		}

		deployment := deploymentConfig{Name: contractName}
		var ok bool
		deployment.ContractAddress, ok = data["address"].(string)
		if !ok {
			return fmt.Errorf("could not find address for %s", contractName)
		}

		chainClient, err := chainConfig.getChainClient(omniRPCConfig.Port)
		if err != nil {
			return err
		}
		err = deployment.loadContract(chainClient)
		if err != nil {
			return err
		}

		chainConfig.Deployments[contractName] = deployment
		chainConfigs[chainID] = chainConfig
	}
	return nil
}

func getSummitChainID(dockerPath string) (summitChainID uint32, err error) {
	executorPath := fmt.Sprintf("%s/config/%s", dockerPath, executorConfig)
	executorConfig, err := execConfig.DecodeConfig(executorPath)
	if err != nil {
		return 0, err
	}
	return executorConfig.SummitChainID, nil
}

func getSigner(privateKey string) (signer *localsigner.Signer, err error) {
	localWallet, err := wallet.FromHex(privateKey)
	if err != nil {
		return
	}
	signer = localsigner.NewSigner(localWallet.PrivateKey())
	return signer, nil
}

func getMessageRoutes(chainConfigs map[uint32]chainConfig, summitChainID uint32, numRoutes int) (routes [][2]uint32, err error) {
	chainIDs := []uint32{}
	for chainID := range chainConfigs {
		if chainID == summitChainID {
			continue
		}
		chainIDs = append(chainIDs, chainID)
	}
	for i, chainID := range chainIDs {
		if len(chainIDs) >= numRoutes {
			return routes, nil
		}
		origin := chainID
		destination := chainIDs[(i+1)%len(chainIDs)]
		if origin == destination {
			return nil, fmt.Errorf("cannot generate unique origin / destination pair")
		}
		routes = append(routes, [2]uint32{origin, destination})
	}
	return routes, nil
}

func watchEvents(ctx context.Context, chainCfg chainConfig, contractName string) (err error) {
	subs := []event.Subscription{}

	switch contractName {
	case "PingPongClient":
		contract, ok := chainCfg.Deployments[contractName].Contract.(domains.PingPongClientContract)
		if !ok {
			return fmt.Errorf("could not cast contract")
		}

		// Watch sent events.
		pingSentChan := make(chan *pingpongclient.PingPongClientPingSent, eventBufferSize)
		sentSub, err := contract.WatchPingSent(ctx, pingSentChan)
		if err != nil {
			return err
		}
		defer sentSub.Unsubscribe()
		subs = append(subs, sentSub)
		go func() {
			for {
				select {
				case event := <-pingSentChan:
					fmt.Printf("Ping sent: %+v\n", event)
				}
			}
		}()

		// Watch received events.
		pongReceivedChan := make(chan *pingpongclient.PingPongClientPongReceived, eventBufferSize)
		receivedSub, err := contract.WatchPongReceived(ctx, pongReceivedChan)
		if err != nil {
			return err
		}
		defer receivedSub.Unsubscribe()
		subs = append(subs, receivedSub)
		go func() {
			for {
				select {
				case event := <-pingSentChan:
					fmt.Printf("Pong received: %+v\n", event)
				}
			}
		}()
	default:
		return fmt.Errorf("unknown contract %s", contractName)
	}

	// Listen for subscription errors.
	for _, s := range subs {
		sub := s
		go func() {
			subErr := <-sub.Err()
			if subErr != nil {
				fmt.Printf("Error in subscription: %w", subErr)
			}
		}()
	}
	return nil
}

var dockerComposeFile = "docker-compose.yml"
var omnirpcConfig = "omnirpc.yaml"
var executorConfig = "executor-config.yml"

const eventBufferSize = 1000

func main() {
	var dockerPath string
	var deploymentPath string
	var privateKey string
	flag.StringVar(&dockerPath, "d", "", "path to devnet docker files")
	flag.StringVar(&deploymentPath, "c", "", "path to contract deployments")
	flag.StringVar(&privateKey, "p", "", "private key")
	flag.Parse()
	if len(dockerPath) == 0 {
		panic("expected docker path to be set (use -d flag)")
	}
	if len(deploymentPath) == 0 {
		panic("expected deployment path to be set (use -c flag)")
	}
	if len(privateKey) == 0 {
		panic("expected private key to be set (use -p flag)")
	}

	// TODO: respect context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load the chain configs.
	chainConfigs, err := getChainConfigs(dockerPath)
	if err != nil {
		panic(err)
	}

	// Load the omnirpc config.
	omniRPCConfig, err := loadOmniRPCConfig(dockerPath)
	if err != nil {
		panic(err)
	}

	// Load the deployments.
	err = loadDeployments("PingPongClient", deploymentPath, chainConfigs, omniRPCConfig)
	if err != nil {
		panic(err)
	}

	// Load the summit chain id.
	summitChainID, err := getSummitChainID(dockerPath)
	if err != nil {
		panic(err)
	}

	// Load the private key.
	signer, err := getSigner(privateKey)
	if err != nil {
		panic(err)
	}

	// Get routes.
	routes, err := getMessageRoutes(chainConfigs, summitChainID, 1)
	if err != nil {
		panic(err)
	}

	// Listen for messages.
	contractName := "PingPongClient"
	for _, chainCfg := range chainConfigs {
		watchEvents(ctx, chainCfg, contractName)
	}

	// Send messages.
	for _, route := range routes {
		contract, ok := chainConfigs[route[0]].Deployments[contractName].Contract.(domains.PingPongClientContract)
		if !ok {
			panic("could not cast contract")
		}
		destPingPongAddr := common.HexToAddress(chainConfigs[route[1]].Deployments[contractName].ContractAddress)
		err = contract.DoPing(ctx, signer, route[1], destPingPongAddr, 0)
		if err != nil {
			panic(err)
		}
	}
}
