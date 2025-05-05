package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

const (
	EAI_OldWH_ChainID      = "43338" // workerHub version 0 , other chain workerHub v1
	BaseChainID            = "8453"
	BaseChainIDInt         = 8453
	BitAiChainID           = "222671"
	DAGIChainID            = "222672"
	HermesChain            = "45762"
	ArbitrumChainID        = "42161"
	DuckChainID            = "5545"
	PolygonChainID         = "137"
	ZkSyncChainID          = "324"
	ZkSyncChainIDInt       = 324
	EthereumChainID        = "1"
	BscChainID             = "56"
	AbstractTestnetChainID = "11124"
	SubtensorEVMChainID    = "964"
	SubtensorEVMChainIDInt = 964
	SolanaChainID          = "1111"
	SolanaModelID          = "990001"
	IPFSPrefix             = "ipfs://"
	TronChainID            = "728126428"
)

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  *ResultResponse `json:"result"`
	ID      int             `json:"id"`
}

type ResultResponse struct {
	Hash          string  `json:"hash"`
	BlockHash     *string `json:"blockHash"`
	BlockNumber   *string `json:"blockNumber"`
	Number        *string `json:"number"`
	ChainID       string  `json:"chainId"`
	L1BlockNumber *string `json:"l1BlockNumber"`
}

var config *Config

func init() {
	file, err := os.Open("configs/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	config = &v
}

func GetConfig() *Config {
	return config
}

type MissionTokensConfig struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func GetMissionTokenConfig() ([]MissionTokensConfig, error) {
	mapToken := map[string]map[string]string{}
	for _, item := range MISSION_TOKEN_CONFIGS {
		mapToken[item["symbol"]] = item
	}
	// jsonbody, err := json.Marshal(MISSION_TOKEN_CONFIGS)
	// if err != nil {
	// 	return nil, errs.NewError(err)
	// }

	v := []MissionTokensConfig{}
	for _, item := range BINANCE_TOKENS {
		i := MissionTokensConfig{
			Symbol:   item,
			Name:     item,
			ImageUrl: "",
		}
		if v, ok := mapToken[item]; ok {
			i.Name = v["name"]
			i.ImageUrl = v["image_url"]
		}
		v = append(v, i)
	}

	return v, nil
}

type AppChainConfig struct {
	NetworkID                      uint64         `json:"network_id"`
	ExplorerUrl                    string         `json:"explorer_url"`
	RpcUrl                         string         `json:"rpc_url"`
	IsBtcL1                        bool           `json:"is_btc_l1"`
	ZkSync                         bool           `json:"zk_sync"`
	AgentContractAddress           string         `json:"agent_contract_address"`
	AgentTokenAccountAddress       string         `json:"agent_token_account_address"`
	AgentAdminAddress              string         `json:"agent_admin_address"`
	EaiContractAddress             string         `json:"eai_contract_address"`
	MemePoolAddress                string         `json:"meme_pool_address"`
	UniswapFactoryContractAddress  string         `json:"uniswap_factory_contract_address"`
	UniswapPositionMamangerAddress string         `json:"uniswap_position_mamanger_address"`
	Weth9ContractAddress           string         `json:"weth9_contract_address"`
	MinGasPrice                    numeric.BigInt `json:"min_gas_price"`
	PaymasterAddress               string         `json:"paymaster_address"`
	PaymasterToken                 string         `json:"paymaster_token"`
	PaymasterFeeZero               bool           `json:"paymaster_fee_zero"`
}

type Config struct {
	Env                      string `json:"env"`
	Port                     int    `json:"port"`
	Debug                    bool   `json:"debug"`
	Job                      bool   `json:"job"`
	DbURL                    string `json:"db_url"`
	EncryptAuthenKey         string `json:"encrypt_authen_key"`
	InternalApiKey           string `json:"internal_api_key"`
	TokenTwiterID            string `json:"token_twiter_id"`
	TokenTwiterIdForInternal string `json:"token_twiter_id_for_internal"`
	CoinMarketCapApiKey      string `json:"coin_market_cap_api_key"`
	Core                     struct {
		Url string `json:"url"`
	} `json:"core"`
	BlockchainUtils struct {
		Url string `json:"url"`
	} `json:"blockchain_util"`
	DeepResearch struct {
		Url string `json:"url"`
	} `json:"deep_research"`
	Btc struct {
		Network  string `json:"network"`
		BcyToken string `json:"bcy_token"`
		QnUrl    string `json:"qn_url"`
	} `json:"btc"`
	Redis *struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
		Prefix   string `json:"prefix"`
	} `json:"redis"`
	Twitter struct {
		OauthClientId                   string `json:"oauth_client_id"`
		OauthClientSecret               string `json:"oauth_client_secret"`
		Token                           string `json:"token"`
		RedirectUri                     string `json:"redirect_uri"`
		ConsumerKey                     string `json:"consumer_key"`
		ConsumerSecret                  string `json:"consumer_secret"`
		AccessToken                     string `json:"access_token"`
		AccessSecret                    string `json:"access_secret"`
		OauthClientIdForTwitterData     string `json:"oauth_client_id_for_twitter_data"`
		OauthClientSecretForTwitterData string `json:"oauth_client_secret_for_twitter_data"`
		TokenForTwitterData             string `json:"token_for_twitter_data"`
		AccessTokenForTwitterData       string `json:"access_token_for_twitter_data"`
		AccessSecretForTwitterData      string `json:"access_secret_for_twitter_data"`
	} `json:"twitter"`
	GsStorage struct {
		CredentialsFile string `json:"credentials_file"`
		BucketName      string `json:"bucket_name"`
		Url             string `json:"url"`
	} `json:"google_storage"`
	AiImageApiKey string `json:"ai_image_api_key"`
	CMCApiKey     string `json:"cmc_api_key"`
	Ai            struct {
		ChatUrl         string `json:"chat_url"`
		ApiKey          string `json:"api_key"`
		ApiKeyMacStudio string `json:"api_key_mac_studio"`
		ModelName       string `json:"model_name"`
	} `json:"ai"`
	Lighthouse struct {
		Apikey string `json:"apikey"`
	} `json:"lighthouse"`
	AiDojoBackend struct {
		Url            string `json:"url"`
		ApiKey         string `json:"api_key"`
		MentionApiKey  string `json:"mention_api_key"`
		MentionNewFlow bool   `json:"MentionNewFlow"`
	} `json:"ai_dojo_backend"`
	EternalAiAgentInfoId              uint   `json:"eternal_ai_agent_info_id"`
	VideoAiAgentInfoId                uint   `json:"video_ai_agent_info_id"`
	DetectVideoLLMVersion             string `json:"detect_video_llm_version"`
	VideoTelegramKey                  string `json:"video_telegram_key"`
	VideoActivitiesTelegramAlert      string `json:"video_activities_telegram_alert"`
	VideoFailSyntaxTelegramAlert      string `json:"video_fail_syntax_telegram_alert"`
	GenMagicVideoPrompt               bool   `json:"gen_magic_video_prompt"`
	MagicVideoActivitiesTelegramAlert string `json:"magic_video_activities_telegram_alert"`
	NobullshitAgentInfoId             uint   `json:"nobullshit_agent_info_id"`
	LaunchpadAgentInfoId              uint   `json:"launchpad_agent_info_id"`
	HiroUrl                           string `json:"hiro_url"`
	OpenseaAPIKey                     string `json:"opensea_api_key"`
	TaApiKey                          string `json:"ta_api_key"`
	DelegateCash                      struct {
		Url    string `json:"url"`
		ApiKey string `json:"api_key"`
	} `json:"delegate_cash"`
	MoralisApiKey        string `json:"moralis_api_key"`
	RapidApiKey          string `json:"rapid_api_key"`
	GenerateImageUrl     string `json:"generate_image_url"`
	GenerateGifImageUrl  string `json:"generate_gif_image_url"`
	AgentOffchainChatUrl string `json:"agent_offchain_chat_url"`
	AgentOffchain        struct {
		Url    string `json:"url"`
		ApiKey string `json:"api_key"`
	} `json:"agent_offchain"`
	EternalaiBridgesUrl string `json:"eternalai_bridges_url"`
	Telebot             struct {
		Tracker struct {
			Botkey          string `json:"botkey"`
			ChatID          int64  `json:"chat_id"`
			MessageThreadID int    `json:"message_thread_id"`
		} `json:"tracker"`
		Alert struct {
			Botkey          string `json:"botkey"`
			ChatID          int64  `json:"chat_id"`
			MessageThreadID int    `json:"message_thread_id"`
		} `json:"alert"`
		TradeAnalytics struct {
			Botkey string `json:"botkey"`
		} `json:"trade_analytics"`
	} `json:"telebot"`
	ToolLists struct {
		FarcasterPost  string `json:"farcaster_post"`
		FarcasterReply string `json:"farcaster_reply"`
		TradeNews      string `json:"trade_news"`
		TradeAnalytic  string `json:"trade_analytic"`
		PostTwitter    string `json:"post_twitter"`
	} `json:"tool_lists"`
	Tron struct {
		RpcUrl string `json:"rpc_url"`
		ApiKey string `json:"api_key"`
	} `json:"tron"`
	PrivateKeys                 map[string]string            `json:"private_keys"`
	Networks                    map[string]map[string]string `json:"networks"`
	AdminAutoCreateAgentAddress string                       `json:"admin_auto_create_agent_address"`
	SecretKey                   string                       `json:"secret_key"`
	WalletDeployAIKB20          string                       `json:"wallet_deploy_aik_b20"`
	HiddenNetworkId             string                       `json:"hidden_network_id"`
	RagApi                      string                       `json:"rag_api"`
	ListTestToolSet             string                       `json:"list_test_tool_set"`
	LuckyMoneyAdminAddress      string                       `json:"lucky_money_admin_address"`
	LuckyMoneyAdminAddressSol   string                       `json:"lucky_money_admin_address_sol"`
	WebhookUrl                  string                       `json:"webhook_url"`
	KnowledgeBaseConfig         KnowledgeBaseConfig          `json:"knowledge_base_config"`
	GetResultInferUrl           string                       `json:"get_result_infer_url"`
	AgentDeployer               struct {
		Url    string `json:"url"`
		ApiKey string `json:"api_key"`
	} `json:"agent_deployer"`
	SampleTwitterApp struct {
		OauthClientId     string `json:"oauth_client_id"`
		OauthClientSecret string `json:"oauth_client_secret"`
		RedirectUri       string `json:"redirect_uri"`
	} `json:"sample_twitter_app"`
	InfraTwitterApp struct {
		OauthClientId     string `json:"oauth_client_id"`
		OauthClientSecret string `json:"oauth_client_secret"`
		RedirectUri       string `json:"redirect_uri"`
		InfraAuthUri      string `json:"infra_auth_uri"`
		NetworkID         uint64 `json:"network_id"`
		AgentAddress      string `json:"agent_address"`
		WorkerAddress     string `json:"worker_address"`
		AgentID           uint   `json:"agent_id"`
		Fee               uint   `json:"fee"`
	} `json:"infra_twitter_app"`
	Clanker struct {
		IsCreateToken            bool    `json:"is_create_token"`
		RequestorAddress         string  `json:"requestor_address"`
		ApiKey                   string  `json:"api_key"`
		ApiUrl                   string  `json:"api_url"`
		CreatorRewardsPercentage float64 `json:"creator_rewards_percentage"`
		CreatorRewardsAdmin      string  `json:"creator_rewards_admin"`
	} `json:"clanker"`
	Privy struct {
		AppID     string   `json:"app_id"`
		AppSecret string   `json:"app_secret"`
		AppIDList []string `json:"app_id_list"`
	} `json:"privy"`
	PrivyEx map[string]string `json:"privy_ex"`
	Robot   struct {
		TokenAdminAddress string `json:"token_admin_address"`
		TokenSupply       uint64 `json:"token_supply"`
	} `json:"robot"`
	LighthouseUploadBinaryUrl string `json:"lighthouse_upload_binary_url"`
}

func (cf *Config) ExistsedConfigKey(networkID uint64, name string) bool {
	networkIDStr := fmt.Sprintf("%d", networkID)
	n, ok := cf.Networks[networkIDStr]
	if !ok {
		return false
	}
	v, ok := n[name]
	if !ok {
		return false
	}
	if v == "" {
		return false
	}
	return true
}

func (cf *Config) GetConfigKeyString(networkID uint64, name string) string {
	networkIDStr := fmt.Sprintf("%d", networkID)
	n, ok := cf.Networks[networkIDStr]
	if !ok {
		panic(fmt.Sprintf("not found %d %s", networkID, name))
	}
	v, ok := n[name]
	if !ok {
		panic(fmt.Sprintf("not found %d %s", networkID, name))
	}
	if v == "" {
		panic(fmt.Sprintf("not found %d %s", networkID, name))
	}
	return v
}

func (cf *Config) GetConfigKeyBool(networkID uint64, name string) bool {
	networkIDStr := fmt.Sprintf("%d", networkID)
	n, ok := cf.Networks[networkIDStr]
	if !ok {
		panic(fmt.Sprintf("not found %d %s", networkID, name))
	}
	v, ok := n[name]
	if !ok {
		return false
	}
	if v == "" {
		return false
	}
	rs, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return rs
}

type KnowledgeBaseConfig struct {
	EnableSimulation          bool   `json:"enable_simulation"`
	QueryServiceUrl           string `json:"query_service_url"`
	ToolCallServiceUrl        string `json:"tool_call_service_url"`
	DirectServiceUrl          string `json:"direct_service_url"`
	KbChatTopK                int    `json:"kb_chat_top_k"`
	KBTelegramKey             string `json:"kb_telegram_key"`
	KBErrorTelegramAlert      string `json:"kb_error_telegram_alert"`
	KBActivitiesTelegramAlert string `json:"kb_activities_telegram_alert"`
	BackendWallet             string `json:"backend_wallet"`
	OnchainAPIKey             string `json:"onchain_api_key"`
	OnChainUrl                string `json:"on_chain_url"`
}
