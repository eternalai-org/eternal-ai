package services

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/coinmarketcap"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) DexSpotPairsLatest(ctx context.Context, quoteAssetSymbol, networkSlug string) (*coinmarketcap.DexSpotPairsLatestResp, error) {
	return s.cmc.DexSpotPairsLatest(quoteAssetSymbol, networkSlug)
}

func (s *Service) DexPairsTradeLatest(ctx context.Context, contractAddress, networkSlug string) (*coinmarketcap.DexPairsTradeLatestResp, error) {
	return s.cmc.DexPairsTradeLatest(contractAddress, networkSlug)
}

func (s *Service) DexScreenInfo(ctx context.Context, contractAddress string) (interface{}, error) {
	coinInfo, err := s.dexscreener.SearchPairs(contractAddress)
	if err != nil {
		return nil, errs.NewError(err)
	}
	var resp struct {
		ChainId     string `json:"chain_id"`
		DexId       string `json:"dex_id"`
		PriceNative string `json:"price_native"`
		PriceUsd    string `json:"price_usd"`
		Volume      *struct {
			H24 float64 `json:"h24"`
			H6  float64 `json:"h6"`
			H1  float64 `json:"h1"`
			M5  float64 `json:"m5"`
		} `json:"volume_usd"`
		PriceChange *struct {
			H24 float64 `json:"h24"`
			H6  float64 `json:"h6"`
			H1  float64 `json:"h1"`
			M5  float64 `json:"m5"`
		} `json:"price_change_percent"`
		Fdv       uint64 `json:"fdv_usd"`
		MarketCap uint64 `json:"market_cap_usd"`
	}

	if coinInfo != nil {
		resp.ChainId = coinInfo.ChainId
		resp.DexId = coinInfo.DexId
		resp.PriceNative = coinInfo.PriceNative
		resp.PriceUsd = coinInfo.PriceUsd
		resp.Volume = coinInfo.Volume
		resp.PriceChange = coinInfo.PriceChange
		resp.Fdv = coinInfo.Fdv
		resp.MarketCap = coinInfo.MarketCap
		return resp, nil
	}

	return nil, errs.NewError(errs.ErrTokenNotFound)
}

func (s *Service) GetAgentWallet(tx *gorm.DB, networkID uint64, agentInfoID uint) (*models.AgentWallet, error) {
	agentWallet, err := s.dao.FirstAgentWallet(
		tx,
		map[string][]interface{}{
			"network_id = ?":    {networkID},
			"agent_info_id = ?": {agentInfoID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentWallet == nil {
		agent, err := s.dao.FirstAgentInfoByID(
			tx,
			agentInfoID,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}
		var address string
		var cdpWalletID string
		switch networkID {
		case models.ETHEREUM_CHAIN_ID,
			models.BASE_CHAIN_ID,
			models.ARBITRUM_CHAIN_ID,
			models.BSC_CHAIN_ID:
			{
				address = agent.TipEthAddress
			}
		case models.BTC_CHAIN_ID:
			{
				address = agent.TipBtcAddress
			}
		case models.SOLANA_CHAIN_ID:
			{
				address = agent.TipSolAddress
			}
		}
		if address == "" {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
		agentWallet = &models.AgentWallet{
			NetworkID:   networkID,
			AgentInfoID: agentInfoID,
			Address:     address,
			CdpWalletID: cdpWalletID,
		}
		err = s.dao.Create(tx, agentWallet)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return agentWallet, nil
}

func (s *Service) GetTokenQuoteLatestForSolana(ctx context.Context, mint string) (any, error) {
	m, err := s.dao.FirstAgentTradeToken(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":    {models.SOLANA_CHAIN_ID},
			"token_address = ?": {mint},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil || m.CmcId == "" {
		return nil, errs.NewError(err)
	}
	if m == nil {
		return nil, nil
	}
	resp, err := s.cmc.GetQuotesLatest([]string{m.CmcId})
	if err != nil {
		return nil, errs.NewError(err)
	}
	return resp[m.CmcId], nil
}

func (s *Service) GetSolanaTokenDecimals(mint string) (int, error) {
	var decimals int
	err := s.RedisCached(
		fmt.Sprintf("GetSolanaTokenDecimals_%s", mint),
		true,
		999999*time.Hour,
		&decimals,
		func() (interface{}, error) {
			mintInfo, err := s.blockchainUtils.SolanaTokenInfo(mint)
			if err != nil {
				return nil, errs.NewError(err)
			}
			if mintInfo == nil || mintInfo.Data == nil || mintInfo.Data.Parsed == nil || mintInfo.Data.Parsed.Info == nil {
				return nil, errs.NewError(errs.ErrBadRequest)
			}
			return mintInfo.Data.Parsed.Info.Decimals, nil
		},
	)
	if err != nil {
		return 0, errs.NewError(err)
	}
	return decimals, nil
}

func (s *Service) AgentWalletCreatePumpFunMeme(ctx context.Context, networkID uint64, agentContractID string, req *serializers.AdminCreatePumpfunMemeReq) error {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent == nil || agent.TwitterInfoID == 0 {
		return errs.NewError(errs.ErrBadRequest)
	}
	wallet, err := s.GetAgentWallet(daos.GetDBMainCtx(ctx), models.SOLANA_CHAIN_ID, agent.ID)
	if err != nil {
		return errs.NewError(err)
	}
	m := &models.AgentWalletAction{
		NetworkID:     models.SOLANA_CHAIN_ID,
		AgentInfoID:   agent.ID,
		AgentWalletID: wallet.ID,
		ActionInput: helpers.ConvertJsonString(map[string]interface{}{
			"action_input": req,
		}),
	}
	err = s.dao.Create(
		daos.GetDBMainCtx(ctx),
		m,
	)
	if err != nil {
		return errs.NewError(err)
	}
	resp, err := s.blockchainUtils.SolanaCreatePumpfunToken(
		&blockchainutils.SolanaCreatePumpfunTokenReq{
			Address:     wallet.Address,
			Name:        req.Name,
			Symbol:      req.Symbol,
			Description: req.Description,
			Twitter:     fmt.Sprintf("https://x.com/%s", agent.TwitterUsername),
			Telegram:    "",
			Website:     "",
			Amount:      req.Amount,
			ImageBase64: req.ImageBase64,
		},
	)
	if err != nil {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", err.Error()).
			UpdateColumn("status", models.AgentWalletActionStatusError).Error
		if err != nil {
			return errs.NewError(err)
		}
	} else {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", helpers.ConvertJsonString(map[string]interface{}{
				"action_output": resp,
			})).
			UpdateColumn("status", models.AgentWalletActionStatusDone).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) AgentWalletTradePumpFunMeme(ctx context.Context, networkID uint64, agentContractID string, req *serializers.AdminTradePumpfunMemeReq) error {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent == nil || agent.TwitterInfoID == 0 {
		return errs.NewError(errs.ErrBadRequest)
	}
	wallet, err := s.GetAgentWallet(daos.GetDBMainCtx(ctx), models.SOLANA_CHAIN_ID, agent.ID)
	if err != nil {
		return errs.NewError(err)
	}
	var snapshotPostID, snapshotMissionID uint
	var toolSet models.ToolsetType
	if req.RefID != "" {
		snapshotPost, _ := s.dao.FirstAgentSnapshotPost(
			daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"infer_tx_hash = ?": {req.RefID},
			},
			map[string][]interface{}{
				"AgentSnapshotMission": {},
			},
			[]string{},
		)
		if snapshotPost != nil {
			snapshotPostID = snapshotPost.ID
			snapshotMissionID = snapshotPost.AgentSnapshotMissionID
			if snapshotPost.AgentSnapshotMission != nil {
				toolSet = snapshotPost.AgentSnapshotMission.ToolSet
			}
		}
	}
	m := &models.AgentWalletAction{
		NetworkID:     models.SOLANA_CHAIN_ID,
		AgentInfoID:   agent.ID,
		AgentWalletID: wallet.ID,
		ActionType:    "trade_pumpfun",
		ActionInput: helpers.ConvertJsonString(map[string]interface{}{
			"action_input": req,
		}),
		RefID:                  req.RefID,
		Toolset:                string(toolSet),
		AgentSnapshotMissionID: snapshotMissionID,
		AgentSnapshotPostID:    snapshotPostID,
	}
	err = s.dao.Create(
		daos.GetDBMainCtx(ctx),
		m,
	)
	if err != nil {
		return errs.NewError(err)
	}
	coinInfo, err := s.pumfunAPI.GetPumpFunCoinInfo(req.Mint)
	if err != nil {
		return errs.NewError(err)
	}
	pool := "pump"
	if coinInfo.RaydiumPool != "" {
		pool = "raydium"
	}
	resp, err := s.blockchainUtils.SolanaTradePumpfunToken(
		&blockchainutils.SolanaTradePumpfunTokenReq{
			Address: wallet.Address,
			Action:  req.Action,
			Mint:    req.Mint,
			Amount:  req.Amount,
			Pool:    pool,
		},
	)
	if err != nil {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", err.Error()).
			UpdateColumn("status", models.AgentWalletActionStatusError).Error
		if err != nil {
			return errs.NewError(err)
		}
	} else {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", helpers.ConvertJsonString(map[string]interface{}{
				"action_output": resp,
			})).
			UpdateColumn("status", models.AgentWalletActionStatusDone).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) AgentWalletTradeRaydiumToken(ctx context.Context, networkID uint64, agentContractID string, req *serializers.AdminTradePumpfunMemeReq) error {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent == nil || agent.TwitterInfoID == 0 {
		return errs.NewError(errs.ErrBadRequest)
	}
	wallet, err := s.GetAgentWallet(daos.GetDBMainCtx(ctx), models.SOLANA_CHAIN_ID, agent.ID)
	if err != nil {
		return errs.NewError(err)
	}
	var snapshotPostID, snapshotMissionID uint
	var toolSet models.ToolsetType
	if req.RefID != "" {
		snapshotPost, _ := s.dao.FirstAgentSnapshotPost(
			daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"infer_tx_hash = ?": {req.RefID},
			},
			map[string][]interface{}{
				"AgentSnapshotMission": {},
			},
			[]string{},
		)
		if snapshotPost != nil {
			snapshotPostID = snapshotPost.ID
			snapshotMissionID = snapshotPost.AgentSnapshotMissionID
			if snapshotPost.AgentSnapshotMission != nil {
				toolSet = snapshotPost.AgentSnapshotMission.ToolSet
			}
		}
	}
	m := &models.AgentWalletAction{
		NetworkID:     models.SOLANA_CHAIN_ID,
		AgentInfoID:   agent.ID,
		AgentWalletID: wallet.ID,
		ActionType:    "trade_raydium",
		ActionInput: helpers.ConvertJsonString(map[string]interface{}{
			"action_input": req,
		}),
		RefID:                  req.RefID,
		Toolset:                string(toolSet),
		AgentSnapshotMissionID: snapshotMissionID,
		AgentSnapshotPostID:    snapshotPostID,
	}
	err = s.dao.Create(
		daos.GetDBMainCtx(ctx),
		m,
	)
	if err != nil {
		return errs.NewError(err)
	}
	var inputMint, outputMint string
	var decimalsIn int
	switch req.Action {
	case "buy":
		{
			inputMint = "So11111111111111111111111111111111111111112"
			outputMint = req.Mint
			decimalsIn = 9
		}
	case "sell":
		{
			inputMint = req.Mint
			outputMint = "So11111111111111111111111111111111111111112"
			decimalsMint, err := s.GetSolanaTokenDecimals(req.Mint)
			if err != nil {
				return errs.NewError(err)
			}
			decimalsIn = decimalsMint
		}
	default:
		{
			return errs.NewError(errs.ErrBadContent)
		}
	}
	resp, err := s.blockchainUtils.SolanaTradeRaydiumToken(
		&blockchainutils.SolanaTradeRaydiumTokenReq{
			Address:    wallet.Address,
			InputMint:  inputMint,
			OutputMint: outputMint,
			Slippage:   0.5,
			Amount:     uint64(req.Amount * math.Pow10(decimalsIn)),
		},
	)
	if err != nil {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", err.Error()).
			UpdateColumn("status", models.AgentWalletActionStatusError).Error
		if err != nil {
			return errs.NewError(err)
		}
	} else {
		err = daos.GetDBMainCtx(ctx).Model(m).
			UpdateColumn("action_output", helpers.ConvertJsonString(map[string]interface{}{
				"action_output": resp,
			})).
			UpdateColumn("status", models.AgentWalletActionStatusDone).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	err = s.MigrateWalletActionForID(ctx, m.ID)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentWalletGetSolanaTokenBalances(ctx context.Context, networkID uint64, agentContractID string) ([]*serializers.SolanaTokenBalanceResp, error) {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agent == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	wallet, err := s.GetAgentWallet(daos.GetDBMainCtx(ctx), models.SOLANA_CHAIN_ID, agent.ID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	balances, err := s.blockchainUtils.SolanaGetTokenBalances(wallet.Address)
	if err != nil {
		return nil, errs.NewError(err)
	}
	resps := []*serializers.SolanaTokenBalanceResp{}
	for _, v := range balances {
		if v.TokenAmount.UIAmount > 0 {
			resps = append(resps, &serializers.SolanaTokenBalanceResp{
				IsNative: v.IsNative,
				Mint:     v.Mint,
				Amount:   v.TokenAmount.UIAmount,
			})
		}
	}
	return resps, nil
}

func (s *Service) GetPumpFunTrades(ctx context.Context, mint string, page int, limit int) ([]*serializers.PumpFunTradeResp, error) {
	tknDecimals, err := s.GetSolanaTokenDecimals(mint)
	if err != nil {
		return nil, errs.NewError(err)
	}
	trades, err := s.pumfunAPI.GetPumpFunTrades(mint, page, limit)
	if err != nil {
		return nil, errs.NewError(err)
	}
	resps := []*serializers.PumpFunTradeResp{}
	for _, v := range trades {
		resps = append(resps, &serializers.PumpFunTradeResp{
			Signature:   v.Signature,
			Mint:        v.Mint,
			SolAmount:   float64(v.SolAmount) / math.Pow10(9),
			TokenAmount: float64(v.TokenAmount) / math.Pow10(tknDecimals),
			IsBuy:       v.IsBuy,
			Timestamp:   v.Timestamp,
		})
	}
	return resps, nil
}

func (s *Service) GetPumpFunTokenPrice(ctx context.Context, mint string) (float64, error) {
	solPrice, _ := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), string(models.BaseTokenSymbolSOL)).Float64()
	price, err := func() (float64, error) {
		coinInfo, err := s.pumfunAPI.GetPumpFunCoinInfo(mint)
		if err != nil {
			return 0, errs.NewError(err)
		}
		if coinInfo == nil || coinInfo.Mint == "" {
			return 0, errs.NewError(errs.ErrBadRequest)
		}
		tknDecimals, err := s.GetSolanaTokenDecimals(mint)
		if err != nil {
			return 0, errs.NewError(err)
		}
		return coinInfo.UsdMarketCap / (float64(coinInfo.TotalSupply) / math.Pow10(tknDecimals)) / solPrice, nil
	}()
	if err != nil {
		price, err = func() (float64, error) {
			m, err := s.dao.FirstAgentTradeToken(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"network_id = ?":    {models.SOLANA_CHAIN_ID},
					"token_address = ?": {mint},
				},
				map[string][]interface{}{},
				[]string{},
			)
			if err != nil || m.CmcId == "" {
				return 0, errs.NewError(fmt.Errorf("coin not found"))
			}
			resp, err := s.cmc.GetQuotesLatest([]string{m.CmcId})
			if err != nil {
				return 0, errs.NewError(err)
			}
			if resp[m.CmcId] == nil {
				return 0, errs.NewError(fmt.Errorf("coin not found"))
			}
			usdPrice, _ := resp[m.CmcId].Quote.USD.Price.Float64()
			return usdPrice / solPrice, nil
		}()
		if err != nil {
			return 0, errs.NewError(err)
		}
	}
	return price, nil
}

func (s *Service) GetSolanaDataChart24Hour(ctx context.Context, mint string) ([]*serializers.DataChartResp, error) {
	resps := []*serializers.DataChartResp{}
	mintInfo, err := s.cgc.GetSolanaTokenInfo(mint)
	if err != nil {
		return nil, errs.NewError(err)
	}
	charts, err := s.cgc.GetCoinMarketChart(mintInfo.ID, "USD")
	if err != nil {
		return nil, errs.NewError(err)
	}
	for _, v := range charts {
		resps = append(resps, &serializers.DataChartResp{
			Time:  v.Timestamp,
			Price: v.Price,
		})
	}
	return resps, nil
}

func (s *Service) GetAgentWalletSolanaTrades(ctx context.Context, networkID uint64, agentContractID string, mint string, page int, limit int) ([]*serializers.WalletActionTradeResp, error) {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agent == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	filters := map[string][]interface{}{
		"agent_info_id = ?": {agent.ID},
		"status = ?":        {models.AgentWalletActionStatusDone},
		"action_type = ?":   {"trade_raydium"},
	}
	if mint != "" {
		filters["mint = ?"] = []interface{}{mint}
	}
	waMs, _, err := s.dao.FindAgentWalletAction4Page(
		daos.GetDBMainCtx(ctx),
		filters,
		make(map[string][]interface{}),
		[]string{
			"id desc",
		},
		page,
		limit,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	resps := []*serializers.WalletActionTradeResp{}
	for _, waM := range waMs {
		resps = append(resps, &serializers.WalletActionTradeResp{
			CreatedAt: waM.CreatedAt,
			Mint:      waM.Mint,
			Side:      waM.Side,
			AmountIn:  waM.AmountIn,
			AmountOut: waM.AmountOut,
			TxHash:    waM.TxHash,
		})
	}
	return resps, nil
}

func (s *Service) JobMigrateWalletActions(ctx context.Context) error {
	waMs, err := s.dao.FindAgentWalletAction(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"status = ?":                     {models.AgentWalletActionStatusDone},
			"tx_hash is null or tx_hash = ?": {""},
		},
		make(map[string][]interface{}),
		[]string{
			"id desc",
		},
		0,
		999999,
	)
	if err != nil {
		return errs.NewError(err)
	}
	for _, waM := range waMs {
		err = s.MigrateWalletActionForID(ctx, waM.ID)
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) MigrateWalletActionForID(ctx context.Context, walletActionID uint) error {
	waM, err := s.dao.FirstAgentWalletActionByID(
		daos.GetDBMainCtx(ctx),
		walletActionID,
		make(map[string][]interface{}),
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if waM.Status == models.AgentWalletActionStatusDone && waM.TxHash == "" {
		var actionInput struct {
			ActionInput struct {
				RefID  string           `json:"ref_id"`
				Action string           `json:"action"`
				Mint   string           `json:"mint"`
				Amount numeric.BigFloat `json:"amount"`
			} `json:"action_input"`
		}
		var actionOutput struct {
			ActionOutput struct {
				OutputAmount numeric.BigInt `json:"output_amount"`
				Signatures   []string       `json:"signatures"`
			} `json:"action_output"`
		}
		err = helpers.ConvertJsonObject(waM.ActionInput, &actionInput)
		if err != nil {
			return errs.NewError(err)
		}
		err = helpers.ConvertJsonObject(waM.ActionOutput, &actionOutput)
		if err != nil {
			return errs.NewError(err)
		}
		decimalsMint, err := s.GetSolanaTokenDecimals(actionInput.ActionInput.Mint)
		if err != nil {
			return errs.NewError(err)
		}
		var amountIn, amountOut numeric.BigFloat
		amountIn = actionInput.ActionInput.Amount
		switch actionInput.ActionInput.Action {
		case "buy":
			{
				amountOut = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(&actionOutput.ActionOutput.OutputAmount.Int, uint(decimalsMint)))
			}
		case "sell":
			{
				amountOut = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(&actionOutput.ActionOutput.OutputAmount.Int, 9))
			}
		}
		err = daos.GetDBMainCtx(ctx).
			Model(waM).
			Updates(
				map[string]interface{}{
					"mint":       actionInput.ActionInput.Mint,
					"side":       actionInput.ActionInput.Action,
					"amount_in":  amountIn,
					"amount_out": amountOut,
					"tx_hash":    actionOutput.ActionOutput.Signatures[0],
				},
			).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) AgentWalletGetSolanaTokenPnls(ctx context.Context, networkID uint64, agentContractID string) (any, error) {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {networkID},
			"agent_contract_id = ?": {agentContractID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agent == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	wallet, err := s.GetAgentWallet(daos.GetDBMainCtx(ctx), models.SOLANA_CHAIN_ID, agent.ID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	balances, err := s.blockchainUtils.SolanaGetTokenBalances(wallet.Address)
	if err != nil {
		return nil, errs.NewError(err)
	}
	resps := map[string]*serializers.SolanaTokenBalanceResp{}
	for _, v := range balances {
		if v.TokenAmount.UIAmount > 0 && v.Mint != "" {
			resps[v.Mint] = &serializers.SolanaTokenBalanceResp{
				IsNative: v.IsNative,
				Mint:     v.Mint,
				Amount:   v.TokenAmount.UIAmount,
			}
		}
	}
	tradeSums, err := s.dao.GetWalletActionTradeSum(daos.GetDBMainCtx(ctx), agent.ID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	resMap := map[string]interface{}{}
	for _, tradeSum := range tradeSums {
		mintPrice, _ := s.GetPumpFunTokenPrice(ctx, tradeSum.Mint)
		pnl := models.SubBigFloats(&tradeSum.SellAmount.Float, &tradeSum.BuyAmount.Float)
		if resps[tradeSum.Mint] != nil {
			pnl = models.AddBigFloats(
				pnl,
				big.NewFloat(resps[tradeSum.Mint].Amount*mintPrice),
			)
		}
		resMap[tradeSum.Mint] = map[string]interface{}{
			"mint_price": mintPrice,
			"pnl_amount": numeric.BigFloat2TextDecimals(pnl, 9),
		}
		time.Sleep(50 * time.Millisecond)
	}
	return resMap, nil
}

func (s *Service) GetTradeAnalytic(ctx context.Context, token string) (interface{}, error) {
	var resp struct {
		TokenSymbol        string  `json:"token_symbol"`
		TokenName          string  `json:"token_name"`
		LastPrice          string  `json:"last_price_usd"`
		PriceChange        string  `json:"price_change_usd"`
		PriceChangePercent string  `json:"price_change_percent"`
		WeightedAvgPrice   string  `json:"weighted_avg_price_usd"`
		PrevClosePrice     string  `json:"prev_close_price_usd"`
		LastQty            string  `json:"last_qty"`
		BidPrice           string  `json:"bid_price_usd"`
		BidQty             string  `json:"bid_qty"`
		AskQty             string  `json:"ask_qty"`
		OpenPrice          string  `json:"open_price_usd"`
		HighPrice          string  `json:"high_price_usd"`
		LowPrice           string  `json:"low_price_usd"`
		Volume             string  `json:"volume"`
		QuoteVolume        string  `json:"volume_usd"`
		OpenTime           int64   `json:"open_time"`
		CloseTime          int64   `json:"close_time"`
		Sma20              float64 `json:"sma_20"`
		Ema12              float64 `json:"ema_12"`
		Rsi                float64 `json:"rsi"`
		Fibonacci          float64 `json:"fibonacci"`
	}

	cacheKey := fmt.Sprintf(`GetTradeAnalytic_%s`, token)
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		mapToken, _ := s.GetMapAgentSnapshotMissionTokens(ctx)
		for range 3 {
			binanceInfo, err := helpers.GetBinancePrice24h(fmt.Sprintf(`%sUSDT`, token))
			if err == nil {
				resp.TokenSymbol = token
				if v, ok := mapToken[token]; ok {
					resp.TokenName = v.Name
				}
				resp.LastPrice = binanceInfo.LastPrice
				resp.PriceChange = binanceInfo.PriceChange
				resp.PriceChangePercent = binanceInfo.PriceChangePercent
				resp.WeightedAvgPrice = binanceInfo.WeightedAvgPrice
				resp.PrevClosePrice = binanceInfo.PrevClosePrice
				resp.LastQty = binanceInfo.LastQty
				resp.BidPrice = binanceInfo.BidPrice
				resp.BidQty = binanceInfo.BidQty
				resp.AskQty = binanceInfo.AskQty
				resp.OpenPrice = binanceInfo.OpenPrice
				resp.HighPrice = binanceInfo.HighPrice
				resp.LowPrice = binanceInfo.LowPrice
				resp.Volume = binanceInfo.Volume
				resp.QuoteVolume = binanceInfo.QuoteVolume
				resp.OpenTime = binanceInfo.OpenTime
				resp.CloseTime = binanceInfo.CloseTime

				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}

		for range 3 {
			analyticInfo, err := s.taapi.BulkRequest(token)
			if analyticInfo != nil && err == nil {
				for _, item := range analyticInfo.Data {
					if item.ID == "fibonacciretracement" {
						resp.Fibonacci = item.Result.Value
					}

					if item.ID == "rsi" {
						resp.Rsi = item.Result.Value
					}

					if item.ID == "sma" {
						resp.Sma20 = item.Result.Value
					}

					if item.ID == "ema" {
						resp.Ema12 = item.Result.Value
					}
				}
				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}
		_ = s.SetRedisCachedWithKey(cacheKey, resp, 5*time.Minute)
	}

	return resp, nil
}
