package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/usecase/agent_info"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/usecase/appconfig"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/repository"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/usecase/knowledge"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/drivers/mysql"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/secret_manager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/aidojo"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/bridgeapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/btcapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/clanker"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/coingecko"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/coinmarketcap"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/core"
	deepresearch "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/deep_research"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/dexscreener"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/googlestorage"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/moralis"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/openai"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/opensea"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/privy"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/pumfun"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/rapid"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/taapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/zkapi"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Service struct {
	jobRunMap   map[string]bool
	jobMutex    sync.Mutex
	jobDisabled bool
	// config
	conf *configs.Config
	// clients
	rdb             *redis.Client
	coreClient      *core.Client
	gsClient        *googlestorage.Client
	openais         map[string]*openai.OpenAI
	ethApiMap       map[uint64]*ethapi.Client
	zkApiMap        map[uint64]*zkapi.Client
	trxApi          *trxapi.Client
	rapid           *rapid.Rapid
	blockchainUtils *blockchainutils.Client
	deepResearch    *deepresearch.Client
	btcAPI          *btcapi.Client
	pumfunAPI       *pumfun.Client
	cmc             *coinmarketcap.CoinMarketCap
	cgc             *coingecko.CoinGeckoAPI
	twitterAPI      *twitter.Client
	twitterWrapAPI  *twitter.Client
	dojoAPI         *aidojo.AiDojoBackend
	bridgeAPI       *bridgeapi.BridgeApi
	dexscreener     *dexscreener.DexScreenerAPI
	openseaService  *opensea.OpenseaService
	taapi           *taapi.TaApi
	clanker         *clanker.Client
	privyClient     *privy.Client
	moralisClient   *moralis.Client
	// daos
	dao *daos.DAO

	KnowledgeUsecase ports.IKnowledgeUsecase
	AppConfigUseCase ports.IAppConfigUseCase
	AgentInfoUseCase ports.IAgentInfoUseCase
}

func NewService(conf *configs.Config) *Service {
	s := &Service{
		jobRunMap: map[string]bool{},
		jobMutex:  sync.Mutex{},
		//
		conf: conf,
		//
		rdb: redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
			DB:       conf.Redis.Db,
		}),
		coreClient: &core.Client{
			BaseURL: conf.Core.Url,
		},
		gsClient: googlestorage.InitClient(conf.GsStorage.CredentialsFile, conf.GsStorage.BucketName),
		openais: map[string]*openai.OpenAI{
			"Agent": openai.NewAgentAI(conf.Ai.ApiKey),
			"Lama":  openai.NewOpenAI(conf.Ai.ChatUrl, conf.Ai.ApiKeyMacStudio, conf.Ai.ModelName),
		},
		ethApiMap: map[uint64]*ethapi.Client{},
		zkApiMap:  map[uint64]*zkapi.Client{},
		trxApi: &trxapi.Client{
			RpcURL:  conf.Tron.RpcUrl,
			ApiURL:  "https://api.trongrid.io",
			GrpcURL: "grpc.trongrid.io:50051",
			APIKey:  conf.Tron.ApiKey,
		},
		rapid: rapid.NewRapid(conf.RapidApiKey),
		blockchainUtils: &blockchainutils.Client{
			BaseURL: conf.BlockchainUtils.Url,
		},
		deepResearch: &deepresearch.Client{
			BaseURL: conf.DeepResearch.Url,
		},
		btcAPI: &btcapi.Client{
			Chain:             "btc",
			Network:           conf.Btc.Network,
			Token:             conf.Btc.BcyToken,
			QNUrl:             conf.Btc.QnUrl,
			SdkUrl:            "",
			BlockstreamUrl:    "https://blockstream.info",
			MempoolUrl:        "https://mempool.space",
			HirosoUrl:         "https://api.hiro.so",
			BlockchainInfoUrl: "https://blockchain.info",
		},
		pumfunAPI: &pumfun.Client{
			BaseUrl: "https://frontend-api.pump.fun",
		},
		cmc: coinmarketcap.NewCoinMarketCap(conf.CMCApiKey),
		cgc: coingecko.NewCoinGeckoAPI(),
		// daos
		dao: &daos.DAO{},
		twitterAPI: twitter.NewClient(conf.Twitter.Token, conf.Twitter.ConsumerKey, conf.Twitter.ConsumerSecret,
			conf.Twitter.AccessToken, conf.Twitter.AccessSecret,
			conf.Twitter.OauthClientId, conf.Twitter.OauthClientSecret,
			conf.Twitter.RedirectUri,
		),
		twitterWrapAPI: twitter.NewTwitterWrapClient(conf.Twitter.TokenForTwitterData),
		dojoAPI:        aidojo.NewAiDojoBackend(conf.AiDojoBackend.Url, conf.AiDojoBackend.ApiKey),
		bridgeAPI:      bridgeapi.NewBridgeApi(conf.EternalaiBridgesUrl),
		dexscreener:    dexscreener.NewDexScreenerAPI(),
		openseaService: opensea.NewOpensea(conf.OpenseaAPIKey),
		taapi:          taapi.NewTaApi(conf.TaApiKey),
		clanker:        clanker.NewClankerClient(conf.Clanker.ApiKey, conf.Clanker.ApiUrl),
		privyClient:    privy.NewPrivyClient(conf.Privy.AppID, conf.Privy.AppSecret),
		moralisClient: &moralis.Client{
			APIKey: conf.MoralisApiKey,
		},
	}

	gormDB := mysql.NewDefaultMysqlGormConn(nil, s.conf.DbURL, s.conf.Debug)
	knowledgeBaseRepo := repository.NewKnowledgeBaseRepository(gormDB)
	knowledgeBaseFileRepo := repository.NewKnowledgeBaseFileRepository(gormDB)
	agentInfoKnowledgeBaseRepo := repository.NewAgentInfoKnowledgeBaseRepository(gormDB)
	agentInfoRepo := repository.NewAgentInfoRepository(gormDB)

	secretKey := conf.SecretKey
	var googleSecretKey string
	if utils.IsEnvProduction(conf.Env) {
		key, err := secret_manager.GetGoogleSecretKey(context.Background(), secretKey)
		googleSecretKey = key
		if err != nil {
			logger.Error("", "GetGoogleSecretKey", zap.Error(err))
		}
	} else {
		googleSecretKey = secretKey
	}

	s.KnowledgeUsecase = knowledge.NewKnowledgeUsecase(
		knowledge.WithRepos(
			knowledgeBaseRepo, knowledgeBaseFileRepo,
			agentInfoKnowledgeBaseRepo, agentInfoRepo,
		),
		knowledge.WithSecretKey(googleSecretKey),
		knowledge.WithEthApiMap(s.ethApiMap),
		knowledge.WithNetworks(conf.Networks),
		knowledge.WithTrxApi(s.trxApi),
		knowledge.WithRagApi(conf.RagApi),
		knowledge.WithLighthousekey(conf.Lighthouse.Apikey),
		knowledge.WithWebhookUrl(conf.WebhookUrl),
		knowledge.WithNotiBot(
			s.conf.KnowledgeBaseConfig.KBTelegramKey,
			s.conf.KnowledgeBaseConfig.KBActivitiesTelegramAlert,
			s.conf.KnowledgeBaseConfig.KBErrorTelegramAlert,
		),
		knowledge.WithConfig(s.conf),
	)
	appConfigRepo := repository.NewAppConfigRepository(gormDB)
	s.AppConfigUseCase = appconfig.NewAppConfigUseCase(appConfigRepo)
	s.AgentInfoUseCase = agent_info.NewAgentInfoUseCase(agentInfoRepo)
	InitTeleVideoActivitiesAlert(s.conf.VideoTelegramKey, s.conf.VideoActivitiesTelegramAlert)
	InitTeleMagicVideoActivitiesAlert(s.conf.VideoTelegramKey, s.conf.MagicVideoActivitiesTelegramAlert)
	return s
}

func (s *Service) GetAddressPrk(address string) string {
	var prkHex string
	var err error
	for k, v := range s.conf.PrivateKeys {
		if strings.EqualFold(k, address) {
			prkHex = v
			break
		}
	}
	if prkHex == "" {
		prkHex, err = s.coreClient.GetAddressPrk(
			address,
		)
		if err != nil {
			panic(err)
		}
	}
	return prkHex
}
func (s *Service) GetTwitterWrapAPI() *twitter.Client {
	return s.twitterWrapAPI
}
func (s *Service) JobRunCheck(ctx context.Context, jobId string, jobFunc func() error) error {
	s.jobMutex.Lock()
	isRun := s.jobRunMap[jobId]
	jobDisabled := s.jobDisabled
	s.jobMutex.Unlock()
	if !isRun && !jobDisabled {
		s.jobMutex.Lock()
		s.jobRunMap[jobId] = true
		s.jobMutex.Unlock()
		defer func() {
			s.jobMutex.Lock()
			s.jobRunMap[jobId] = false
			s.jobMutex.Unlock()
			if rval := recover(); rval != nil {
				err := errs.NewError(errors.New(fmt.Sprint(rval)))
				stacktrace := err.(*errs.Error).Stacktrace()
				fmt.Println(time.Now(), jobId, "panic", err.Error(), stacktrace)
			}
		}()
		fmt.Println(time.Now(), jobId, "begin")
		err := jobFunc()
		if err != nil {
			err = errs.NewError(err)
			stacktrace := err.(*errs.Error).Stacktrace()
			fmt.Println(time.Now(), jobId, "error", err.Error(), stacktrace)
			return err
		} else {
			fmt.Println(time.Now(), jobId, "end")
		}
		return err
	}
	return nil
}

func (s *Service) JobRun(ctx context.Context, jobName string, duration time.Duration, jobFunc func() error) {
	s.jobMutex.Lock()
	isRun := s.jobRunMap[jobName]
	s.jobMutex.Unlock()
	if !isRun {
		s.jobMutex.Lock()
		s.jobRunMap[jobName] = true
		s.jobMutex.Unlock()
		go func() {
			for {
				fmt.Println(time.Now(), jobName, "begin")
				err := func() error {
					defer func() {
						if rval := recover(); rval != nil {
							err := errs.NewError(errors.New(fmt.Sprint(rval)))
							stacktrace := err.(*errs.Error).Stacktrace()
							fmt.Println(time.Now(), jobName, "panic", err.Error(), stacktrace)
						}
					}()
					err := jobFunc()
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				}()
				if err != nil {
					err = errs.NewError(err)
					stacktrace := err.(*errs.Error).Stacktrace()
					fmt.Println(time.Now(), jobName, "error", err.Error(), stacktrace)
				} else {
					fmt.Println(time.Now(), jobName, "end")
				}
				time.Sleep(duration)
			}
		}()
	}
}

func (s *Service) VerifyAddressSignature(ctx context.Context, networkID uint64, address string, message string, signature string) error {
	err := s.GetEthereumClient(ctx, networkID).ValidateMessageSignature(
		message,
		signature,
		address,
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GetTokenMarketPrice(tx *gorm.DB, symbol string) *big.Float {
	cachedKey := fmt.Sprintf(`GetTokenMarketPrice_%s`, symbol)
	tokenPrice := big.NewFloat(0)
	_ = s.GetRedisCachedWithKey(cachedKey, &tokenPrice)
	if tokenPrice.Cmp(big.NewFloat(0)) <= 0 {
		tkPrice, _, err := s.dao.GetTokenMarketPrice(tx, symbol)
		if err != nil {
			return big.NewFloat(0)
		}
		tokenPrice = &tkPrice.Float
		_ = s.SetRedisCachedWithKey(cachedKey, tokenPrice, 5*time.Minute)
	}
	return tokenPrice
}

func (s *Service) GetMapTokenPrice(ctx context.Context) map[string]*big.Float {
	cachedKey := `AgentGetMapTokenPrice`
	mapTokenPrice := map[string]*big.Float{}
	err := s.GetRedisCachedWithKey(cachedKey, &mapTokenPrice)
	if err != nil {
		mapTokenPrice["BTC"] = s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "BTC")
		mapTokenPrice["ETH"] = s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "ETH")
		mapTokenPrice["BVM"] = s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "BVM")
		mapTokenPrice["EAI"] = s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "EAI")
		mapTokenPrice["SOL"] = s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "SOL")
		s.SetRedisCachedWithKey(cachedKey, mapTokenPrice, 1*time.Minute)
	}
	return mapTokenPrice
}

func (s *Service) GetDao() *daos.DAO {
	return s.dao
}
