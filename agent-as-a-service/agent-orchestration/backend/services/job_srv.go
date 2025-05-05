package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jasonlvhit/gocron"
)

func (s *Service) JobUpdateMarketPrice(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobUpdateMarketPrice",
		func() error {
			eaiPrice, err := s.cmc.GetQuotesLatest([]string{"31401"})
			if err != nil {
				return err
			}
			_ = s.CreateOrUpdateTokenPrice(ctx, "EAI", numeric.NewBigFloatFromFloat(&eaiPrice["31401"].Quote.USD.Price.Float))

			// BTC - ETH price
			mapSymbol := map[string]string{
				"BTCUSDT": "BTC",
				"ETHUSDT": "ETH",
				"SOLUSDT": "SOL",
				"BNBUSDT": "BNB",
			}

			symbol := `["BTCUSDT","ETHUSDT","SOLUSDT","BNBUSDT"]`
			prices, err := helpers.GetListExternalPrice(symbol)
			if err != nil {
				return err
			}

			for _, price := range prices {
				err = s.CreateOrUpdateTokenPrice(ctx, mapSymbol[price.Symbol], numeric.NewBigFloatFromString(price.Price))
				if err != nil {
					fmt.Println(err)
				}

			}

			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateOrUpdateTokenPrice(ctx context.Context, symbol string, price numeric.BigFloat) error {
	filters := map[string][]interface{}{}
	filters["symbol = ?"] = []interface{}{symbol}
	tokenPrice, err := s.dao.FirstTokenPrice(daos.GetDBMainCtx(ctx), filters, map[string][]interface{}{}, true)
	if err != nil {
		return errs.NewError(err)
	}

	isCreated := false
	if tokenPrice == nil {
		tokenPrice = &models.TokenPrice{}
		tokenPrice.Symbol = symbol
		isCreated = true
	}
	tokenPrice.NetworkID = 0
	tokenPrice.Price = price
	tokenPrice.UpdatedAt = time.Now()

	if isCreated {
		err = s.dao.Create(daos.GetDBMainCtx(ctx), tokenPrice)
		if err != nil {
			return errs.NewError(err)
		}
	} else {
		err = s.dao.Save(daos.GetDBMainCtx(ctx), tokenPrice)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) BlockchainDurations(ctx context.Context) ([]uint, error) {
	chains, err := s.dao.FindBlockScanInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"duration > ?": {0},
		},
		map[string][]interface{}{},
		[]string{},
		0,
		100,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	durationMap := map[uint]bool{}
	for _, v := range chains {
		durationMap[v.Duration] = true
	}
	durations := []uint{}
	for duration := range durationMap {
		durations = append(durations, duration)
	}
	return durations, nil
}

func (s *Service) JobEnabledDB(ctx context.Context) error {
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AppConfig{}).
		Where("network_id = ?", models.GENERTAL_NETWORK_ID).
		Where("name = ?", "job_enabled").
		UpdateColumn("value", "1").Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) DisableJobs() {
	s.jobMutex.Lock()
	s.jobDisabled = true
	s.jobMutex.Unlock()
}

func (s *Service) RunJobs(ctx context.Context) error {
	gocron.Every(360).Second().Do(func() {
		defer func() {
			if rval := recover(); rval != nil {
				err := errs.NewError(errors.New(fmt.Sprint(rval)))
				stacktrace := err.(*errs.Error).Stacktrace()
				fmt.Println(time.Now(), err.Error(), stacktrace)
			}
		}()
		s.KnowledgeUsecase.WatchWalletChange(context.Background())
	})
	gocron.Every(30).Second().Do(func() {
		defer func() {
			if rval := recover(); rval != nil {
				err := errs.NewError(errors.New(fmt.Sprint(rval)))
				stacktrace := err.(*errs.Error).Stacktrace()
				fmt.Println(time.Now(), err.Error(), stacktrace)
			}
		}()
		s.KnowledgeUsecase.ScanKnowledgeBaseStatusPaymentReceipt(context.Background())
	})

	// scan events by chain
	var networkIDs []uint64
	for networkIDStr := range s.conf.Networks {
		networkID, err := strconv.ParseUint(networkIDStr, 10, 64)
		if err != nil {
			panic(err)
		}
		if networkID > 0 {
			networkIDs = append(networkIDs, networkID)
		}
	}
	for i := range 4 {
		gocron.Every(30).Second().
			From(helpers.GetNextScheduleTime(30*time.Second, time.Duration(i)*7*time.Second)).
			Do(
				func(rangeIndex int) {
					for idx, networkID := range networkIDs {
						if idx%4 == rangeIndex {
							fmt.Printf("ScanEventsByChainRange_%d_%d\n", rangeIndex, networkID)
							s.ScanEventsByChain(context.Background(), networkID)
						}
					}
				},
				i,
			)
	}

	// agent mint nft
	gocron.Every(1).Minute().Do(
		func() error {
			s.JobAgentMintNft(context.Background())
			s.JobRetryAgentMintNft(context.Background())
			s.JobRetryAgentMintNftError(context.Background())
			s.JobAgentStart(context.Background())
			return nil
		},
	)
	gocron.Every(1).Minute().Do(
		func() error {
			s.JobAgentTwinTrain(context.Background())
			return nil
		},
	)

	gocron.Every(1).Minute().Do(s.JobUpdateTwitterAccessToken, context.Background())
	gocron.Every(1).Minute().Do(s.JobCreateTokenInfo, context.Background())
	gocron.Every(7).Minute().Do(s.JobUpdateTokenPriceInfo, context.Background())
	// gocron.Every(5).Minute().Do(s.JobUpdateTrendingTokens, context.Background())

	// generate video
	gocron.Every(5).Minutes().Do(s.JobScanAgentTwitterPostForGenerateVideo, context.Background())
	gocron.Every(15).Seconds().Do(s.JobAgentTwitterPostSubmitVideoInfer, context.Background())
	gocron.Every(30).Seconds().Do(s.JobAgentTwitterScanResultGenerateVideo, context.Background())
	gocron.Every(2).Minutes().Do(s.JobAgentTwitterPostCreateClankerToken, context.Background())
	gocron.Every(1).Minutes().Do(s.JobAgentTwitterPostGenerateVideo, context.Background())
	//gocron.Every(1).Minutes().Do(s.JobAgentTwitterScanResultGenerateVideoMagicPrompt, context.Background())

	// trading analyze
	gocron.Every(5).Minute().Do(s.JobScanAgentTwitterPostForTA, context.Background())
	gocron.Every(1).Minute().Do(s.JobAgentTwitterPostTA, context.Background())

	// lucky moneys
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyActionExecuted, context.Background())
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyCollectPost, context.Background())
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyProcessUserReward, context.Background())

	gocron.Every(1).Minute().Do(s.JobUpdateOffchainAutoOutputForMission, context.Background())
	gocron.Every(5).Minute().Do(s.JobUpdateOffchainAutoOutput, context.Background())
	gocron.Every(30).Minute().Do(s.JobUpdateOffchainAutoOutput3Hour, context.Background())
	gocron.Every(5).Minute().Do(s.JobAgentSnapshotPostStatusInferRefund, context.Background())

	gocron.Every(1).Minute().Do(
		func() {
			s.JobAgentSnapshotPostCreate(context.Background())
		},
	)
	gocron.Every(10).Second().Do(
		func() {
			s.JobAgentSnapshotPostActionExecuted(context.Background())
		},
	)
	gocron.Every(5).Minute().Do(
		func() {
			s.JobAgentSnapshotPostActionDupplicated(context.Background())
			s.JobAgentSnapshotPostActionCancelled(context.Background())
		},
	)
	//
	gocron.Every(5).Minute().Do(
		func() {
			s.JobUpdateMarketPrice(context.Background())
			s.JobUpdateMemeUsdPrice(context.Background())
		},
	)
	gocron.Every(1).Day().
		From(helpers.GetNextScheduleTime(24*time.Hour, 0*time.Hour)).
		Do(s.AgentDailyReport, context.Background())

	gocron.Every(5).Minute().
		From(helpers.GetNextScheduleTime(5*time.Minute, 0*time.Hour)).
		Do(s.JobAgentTeleAlertTopupNotAction, context.Background())

	gocron.Every(15).Minute().Do(
		func() {
			s.JobScanTwitterLiked(context.Background())
		},
	)

	gocron.Every(30).Second().Do(
		func() {
			s.JobRunCheck(
				context.Background(),
				"JobAgentLiquidity",
				func() error {
					s.JobAgentDeployToken(context.Background())
					s.JobRetryAgentDeployToken(context.Background())
					s.JobMemeAddPositionInternal(context.Background())
					s.JobMemeRemovePositionInternal(context.Background())
					s.JobCheckMemeReachMarketCap(context.Background())
					s.JobMemeAddPositionUniswap(context.Background())
					s.JobRetryAddPool1(context.Background())
					s.JobRetryAddPool2(context.Background())
					s.JobMemeBurnPositionUniswap(context.Background())
					return nil
				},
			)
		},
	)

	gocron.Every(1).Minutes().Do(
		func() {
			s.JobCreateAgentKnowledgeBase(context.Background())
		},
	)

	// scan robot sale wallet balance
	gocron.Every(30).Second().Do(func() {
		s.JobRobotScanBalanceSOL(context.Background())
	})

	<-gocron.Start()
	return nil
}
