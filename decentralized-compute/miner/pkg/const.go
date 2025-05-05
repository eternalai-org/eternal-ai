package pkg

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"
	"math/big"
	"time"
)

var (
	Line         = "----------------------------------------------------------------------------\n"
	RootNodeTxt  = "Neurons/Solo"
	ErrorFillOut = `Please back to the "` + RootNodeTxt + `" and use "` + COMMAND_CONFIG + `" to fill out all information first.`
)

const (
	CHAIN_BASE string = "8453"
)

const TimeToWating time.Duration = 2
const (
	MODE_MINER              = "miner"
	MODE_VALIDATOR          = "validator"
	MINER_SERVICE_NAME      = "service_miner"
	API_SERVICE_NAME        = "service_api"
	REDIS_PUBSUB            = "redis_pubsub"
	MINER_SERVICE_CONTRACTS = "contracts"
	MINER_SERVICE_OLLAMA    = "ollama"
	MINER_SERVICE_HARDHAT   = "hardhat"
	PLATFROM_INTEL          = "Intel_Chip"
	PLATFROM_APPLE_SILLICON = "Apple_Silicon"
	LOG_INFO_FILE           = "%s/env/info.json"
	ENV_FILE                = "%s/env/config.env"
	ENTRY_POINT_FILE        = "%s/env/entrypoint.sh"
	ENV_LOCAL_MINERS_FILE   = "%s/env/config_local_%d.env"
	ENV_SAMPLE_FILE         = "%s/env/sample.env"
	ENV_SOL_FOLDER          = "sol"

	ENV_SMART_CONTRACT_V1 = "smart-contract-v1"
	ENV_SMART_CONTRACT_V2 = "smart-contract-v2"

	ENV_CONTRACT_V2_PATH             = "%s/" + ENV_SOL_FOLDER + "/" + ENV_SMART_CONTRACT_V2
	ENV_CONTRACT_V2_AUTO_DEPLOY      = ENV_CONTRACT_V2_PATH + "/scripts/autoDeploy.ts"
	ENV_CONTRACT_V2_DEPLOY_SH        = ENV_CONTRACT_V2_PATH + "/deploy.sh"
	ENV_CONTRACT_V2_ENV              = ENV_CONTRACT_V2_PATH + "/.env"
	ENV_CONTRACT_V2_DEPLOYED_ADDRESS = ENV_CONTRACT_V2_PATH + "/deployedAddressesV2.json"

	ENV_CONTRACT_V1_PATH             = "%s/" + ENV_SOL_FOLDER + "/" + ENV_SMART_CONTRACT_V1
	ENV_CONTRACT_V1_AUTO_DEPLOY      = ENV_CONTRACT_V1_PATH + "/scripts/autoDeploy.ts"
	ENV_CONTRACT_V1_DEPLOY_SH        = ENV_CONTRACT_V1_PATH + "/deploy.sh"
	ENV_CONTRACT_V1_ENV              = ENV_CONTRACT_V1_PATH + "/.env"
	ENV_CONTRACT_V1_DEPLOYED_ADDRESS = ENV_CONTRACT_V1_PATH + "/deployedAddresses.json"

	LOCAL_CHAIN_INFO      = "%s/env/local_contracts.json"
	ENV_FOLDER            = "%s/env"
	LOCAL_CHAIN_GAS_PRICE = 2000_000_000
	LOCAL_CHAIN_GAS_LIMIT = 30000000
)

const (
	COMMAND_SETUP           = "1"
	COMMAND_SETUP_AUTOMATIC = "1"
	COMMAND_SETUP_MANUAL    = "2"

	COMMAND_EXIT   = "q"
	COMMAND_BACK   = "q"
	COMMAND_CONFIG = "config"

	COMMAND_CLUSTER_CREATE_VERSION      = "version"
	COMMAND_CLUSTER_CREATE_TYPE         = "type"
	COMMAND_CLUSTER_CREATE_MODEL_NAME   = "model-name"
	COMMAND_CLUSTER_CREATE_MIN_HARDWARE = "min-hardware"

	COMMAND_CREATE_GROUP_NAME        = "name"
	COMMAND_CREATE_GROUP_CLUSTER_IDS = "cluster-ids"

	PLATFORM                      = "PLATFORM"
	COMMAND_INFER                 = "2"
	COMMAND_INFER_PROMPT          = "prompt"
	COMMAND_LOCAL_SET_WEAI        = "set-weai"
	COMMAND_LOCAL_PRIV_KEY        = "private key"
	COMMAND_LOCAL_CHAIN_RPC       = "rpc"
	COMMAND_LOCAL_PUBSUB          = "pubsub"
	COMMAND_LOCAL_CHAIN_ID        = "chainID"
	COMMAND_LOCAL_GAS_PRICE       = "gasPrice"
	COMMAND_LOCAL_GAS_LIMIT       = "gasLimit"
	COMMAND_LOCAL_MODEL_NAME      = "modelName"
	COMMAND_LOCAL_RUN_POD_URL     = "run_pod_url"
	COMMAND_LOCAL_RUN_POD_API_KEY = "run_pod_api_key"
	COMMAND_LOCAL_CONTRACT_NAME   = "contractName"
	COMMAND_LOCAL_START_CONFIG    = "1"
	COMMAND_LOCAL_START_HARDHAT   = "2"
	// COMMAND_LOCAL_START_OLLAMA    = "3"
	COMMAND_LOCAL_DEPLOY_CONTRACT = "3"
	COMMAND_LOCAL_START_MINERS    = "4"
	COMMAND_LOCAL_START_APIS      = "5"

	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION = "collection_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER      = "gpu_manager_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_LOAD_BALANCER    = "model_load_balancer_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER  = "prompt_scheduler_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI             = "weai_address"

	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION_V1 = "collection_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI_V1             = "wrapped_token_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER_V1  = "worker_hub_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER_V1      = "staking_hub_address"
	COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1           = "hybrid_model_address"
	API_URL                                                  = "http://localhost:8004/v1/chat/completions"
	//staking
	MIN_STAKE           = 25000
	BLOCK_PER_EPOCH     = 600
	REWARD_PER_EPOCH    = 1
	UNSTAK_DEPLAY_TIME  = 907200 //907200 blocks = 21 days (blocktime = 2)
	PENALTY_DURATION    = 0
	FINE_PERCENTAGE     = 0
	MIN_FEE_TO_USE      = 0
	STREAM_DATA_CHANNEL = "stream_channel"
)

var SupportedContracts = []string{
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI,
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION,
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER,
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_LOAD_BALANCER,
	COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER,
}

type Command struct {
	Parent        *Command
	Key           string
	Name          string
	Help          string
	Default       string
	Required      bool
	VerifyInArray []string
	Children      []*Command
	Function      func(reader *bufio.Reader, node *Command)
}

func CreateSeed(params string, requestID string) uint64 {
	seed := hex.EncodeToString([]byte(params + requestID))
	h := sha256.New()
	h.Write([]byte(seed))
	bs := h.Sum(nil)
	seedHex := hex.EncodeToString(bs)
	i := new(big.Int)
	i.SetString(seedHex, 16)

	// check if the seed is too large for uint64
	if i.BitLen() > 64 {
		i = i.Mod(i, new(big.Int).SetUint64(math.MaxUint64))
	}

	return i.Uint64()
}

func Copy(in interface{}, out interface{}) error {
	_b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(_b, &out)
}
