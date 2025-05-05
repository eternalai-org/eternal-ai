package config

import (
	"errors"
	"fmt"
	"os"
	"solo/pkg"
	"solo/pkg/eth"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	IPFSPrefix = "ipfs://"
)

type Chain struct {
	Name      string            `json:"name"`
	ID        string            `json:"id"`
	Contracts map[string]string `json:"contracts"`
	ModelID   uint32
	RPC       string
}

type Config struct {
	Rpc                      string
	PubSubURL                string
	Account                  string
	WorkerAddress            string
	StakingHubAddress        string
	WorkerHubAddress         string
	ApiUrl                   string
	ApiKey                   string
	LighthouseKey            string
	ModelAddress             string
	ChainID                  string
	Erc20Address             string
	DebugMode                bool
	ClusterID                string
	ModelName                string
	Platform                 string
	ModelCollectionAddress   string
	ModelLoadBalancerAddress string
	NodeID                   string
}

func ChainConfig() map[string]Chain {
	resp := make(map[string]Chain)
	resp[pkg.CHAIN_BASE] = Chain{
		Name: "Base",
		ID:   pkg.CHAIN_BASE,
		Contracts: map[string]string{
			pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER:      "0x14A008005cfa25621dD48E958EA33d14dd519d0d",
			pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_LOAD_BALANCER:    "0x812c7F05f12B1FF14AED93751D4B0576e4020806",
			pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER:  "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
			pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI:             "0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7",
			pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION: "0x8DF579e2907FC45Ed477DF72480A87C404703a8F",
		},
		ModelID: uint32(700050),
		RPC:     "https://mainnet.base.org",
	}

	return resp
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := godotenv.Overload(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cfg.PubSubURL = os.Getenv("PUBSUB_URL")
	cfg.Rpc = os.Getenv("CHAIN_RPC")
	cfg.Account = os.Getenv("ACCOUNT_PRIV")
	cfg.StakingHubAddress = os.Getenv("STAKING_HUB_ADDRESS")
	cfg.WorkerHubAddress = os.Getenv("WORKER_HUB_ADDRESS")
	cfg.ApiUrl = os.Getenv("API_URL")
	cfg.ApiKey = os.Getenv("API_KEY")
	cfg.LighthouseKey = os.Getenv("LIGHT_HOUSE_API_KEY")
	cfg.ModelAddress = os.Getenv("MODEL_ADDRESS")
	cfg.ChainID = os.Getenv("CHAIN_ID")
	cfg.ClusterID = os.Getenv("CLUSTER_ID")
	cfg.ModelName = os.Getenv("MODEL_NAME")
	cfg.Erc20Address = os.Getenv("ERC20_ADDRESS")
	cfg.Platform = os.Getenv("PLATFORM")
	dmode := os.Getenv("DEBUG_MODE")
	nodeID := os.Getenv("NODE_ID")
	modelCollectionAddress := os.Getenv("COLLECTION_ADDRESS")
	modelLoadBalancerAddress := os.Getenv("MODEL_LOAD_BALANCER_ADDRESS")
	if dmode != "" {
		dmodeBool, errP := strconv.ParseBool(dmode)
		if errP == nil {
			cfg.DebugMode = dmodeBool
		}
	}

	_, ad, err := eth.GetAccountInfo(cfg.Account)
	if err != nil {
		return nil, err
	}

	cfg.WorkerAddress = ad.Hex()
	cfg.ModelLoadBalancerAddress = modelLoadBalancerAddress
	cfg.ModelCollectionAddress = modelCollectionAddress
	cfg.NodeID = nodeID
	return cfg, nil
}

func (cfg *Config) Verify() error {
	// validate
	/*
		if cfg.LighthouseKey == "" {
			return errors.New("Lighthouse key is missing. Let's configure it now.")
		}*/

	if cfg.ApiUrl == "" {
		return errors.New("API URL is missing. Let's configure it now.")
	}

	/*
		if cfg.ApiKey == "" {
			return errors.New("API KEY is missing. Let's configure it now.")
		}*/
	return nil
}
