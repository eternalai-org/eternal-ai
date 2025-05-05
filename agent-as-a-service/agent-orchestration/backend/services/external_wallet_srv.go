package services

import (
	"context"
	"math"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

func (s *Service) ExternalWalletCreateSOL(ctx context.Context) (*serializers.ExternalWalletResp, error) {
	apiKey, secrectKey, err := s.CreateAPIKey(ctx)
	if err != nil {
		return nil, errs.NewError(err)
	}
	solAddress, err := s.CreateSOLAddress(ctx)
	if err != nil {
		return nil, errs.NewError(err)
	}
	err = s.dao.Create(
		daos.GetDBMainCtx(ctx),
		&models.ExternalWallet{
			APIKey:  secrectKey,
			Type:    models.ExternalWalletTypeSOL,
			Address: solAddress,
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return &serializers.ExternalWalletResp{
		ApiKey:  apiKey,
		Address: solAddress,
	}, nil
}

func (s *Service) ExternalWalletGet(ctx context.Context, apiKey string) (*serializers.ExternalWalletResp, error) {
	m, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if m == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	return &serializers.ExternalWalletResp{
		Address: m.Address,
	}, nil
}

func (s *Service) ExternalWalletBalances(ctx context.Context, apiKey string) ([]*serializers.SolanaTokenBalanceResp, error) {
	m, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if m == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	balances, err := s.blockchainUtils.SolanaGetTokenBalances(m.Address)
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

func (s *Service) ExternalWalletComputeOrder(ctx context.Context, apiKey string, req *serializers.ExternalWalletOrderReq) (float64, error) {
	wallet, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return 0, errs.NewError(err)
	}
	if wallet == nil {
		return 0, errs.NewError(errs.ErrBadRequest)
	}
	switch req.Action {
	case models.ExternalWalletOrderTypeBuy, models.ExternalWalletOrderTypeSell:
		{
			var inputMint, outputMint string
			var decimalsIn, decimalsOut, decimalsMint int
			decimalsMint, err = s.GetSolanaTokenDecimals(req.Mint)
			if err != nil {
				return 0, errs.NewError(err)
			}
			switch req.Action {
			case models.ExternalWalletOrderTypeBuy:
				{
					inputMint = "So11111111111111111111111111111111111111112"
					outputMint = req.Mint
					decimalsIn = 9
					decimalsOut = decimalsMint
				}
			case models.ExternalWalletOrderTypeSell:
				{
					inputMint = req.Mint
					outputMint = "So11111111111111111111111111111111111111112"
					decimalsIn = decimalsMint
					decimalsOut = 9
				}
			default:
				{
					return 0, errs.NewError(errs.ErrBadContent)
				}
			}
			resp, err := s.blockchainUtils.SolanaComputeRaydiumToken(
				&blockchainutils.SolanaTradeRaydiumTokenReq{
					Address:    wallet.Address,
					InputMint:  inputMint,
					OutputMint: outputMint,
					Slippage:   0.5,
					Amount:     uint64(req.Amount * math.Pow10(decimalsIn)),
				},
			)
			if err != nil {
				return 0, errs.NewError(err)
			}
			if resp == nil {
				return 0, nil
			}
			return float64(resp.OutputAmount.Int64()) / math.Pow10(decimalsOut), nil
		}
	default:
		{
			return 0, errs.NewError(errs.ErrBadContent)
		}
	}
}

func (s *Service) ExternalWalletCreateOrder(ctx context.Context, apiKey string, req *serializers.ExternalWalletOrderReq) (*models.ExternalWalletOrder, error) {
	wallet, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if wallet == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	order := &models.ExternalWalletOrder{
		ExternalWalletID: wallet.ID,
		Type:             models.ExternalWalletOrderType(req.Action),
		TokenAddress:     req.Mint,
		Destination:      req.Destination,
		AmountIn:         numeric.NewBigFloatFromFloat(big.NewFloat(req.Amount)),
		Status:           models.ExternalWalletOrderStatusNew,
	}
	err = s.dao.Create(
		daos.GetDBMainCtx(ctx),
		order,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	switch req.Action {
	case models.ExternalWalletOrderTypeBuy, models.ExternalWalletOrderTypeSell:
		{
			var inputMint, outputMint string
			var decimalsIn, decimalsOut, decimalsMint int
			decimalsMint, err = s.GetSolanaTokenDecimals(req.Mint)
			if err != nil {
				return nil, errs.NewError(err)
			}
			switch req.Action {
			case models.ExternalWalletOrderTypeBuy:
				{
					inputMint = "So11111111111111111111111111111111111111112"
					outputMint = req.Mint
					decimalsIn = 9
					decimalsOut = decimalsMint
				}
			case models.ExternalWalletOrderTypeSell:
				{
					inputMint = req.Mint
					outputMint = "So11111111111111111111111111111111111111112"
					decimalsIn = decimalsMint
					decimalsOut = 9
				}
			default:
				{
					return nil, errs.NewError(errs.ErrBadContent)
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
				order.Error = err.Error()
				order.Status = models.ExternalWalletOrderStatusError
			} else {
				order.Status = models.ExternalWalletOrderStatusDone
				order.TxHash = strings.Join(resp.Signatures, ",")
				order.AmountOut = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(&resp.OutputAmount.Int, uint(decimalsOut)))
			}
			err = s.dao.Save(
				daos.GetDBMainCtx(ctx),
				order,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}
	case models.ExternalWalletOrderTypeWithdraw:
		{
			decimalsMint := 9
			if req.Mint != "" {
				decimalsMint, err = s.GetSolanaTokenDecimals(req.Mint)
				if err != nil {
					return nil, errs.NewError(err)
				}
			}
			resp, err := s.blockchainUtils.SolanaTransfer(
				wallet.Address,
				&blockchainutils.SolanaTransferReq{
					Mint:      req.Mint,
					Amount:    req.Amount * math.Pow10(decimalsMint),
					ToAddress: req.Destination,
				},
			)
			if err != nil {
				order.Error = err.Error()
				order.Status = models.ExternalWalletOrderStatusError
			} else {
				order.Status = models.ExternalWalletOrderStatusDone
				order.TxHash = resp
			}
			err = s.dao.Save(
				daos.GetDBMainCtx(ctx),
				order,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}
	default:
		{
			return nil, errs.NewError(errs.ErrBadContent)
		}
	}
	return order, nil
}

func (s *Service) ExternalWalletGetOrders(ctx context.Context, apiKey string, page int, limit int) ([]*models.ExternalWalletOrder, error) {
	m, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if m == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	orders, _, err := s.dao.FindExternalWalletOrder4Page(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"external_wallet_id = ?": {m.ID},
		},
		map[string][]interface{}{},
		[]string{
			"id desc",
		},
		page,
		limit,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return orders, nil
}

func (s *Service) ExternalWalletGetTokens(ctx context.Context, apiKey string) ([]*models.ExternalWalletToken, error) {
	m, err := s.dao.FirstExternalWallet(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {s.GetAddressPrk(apiKey)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if m == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	tokens, err := s.dao.FindExternalWalletToken(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"enabled = ?": {true},
		},
		map[string][]interface{}{},
		[]string{},
		0,
		999999,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tokens, nil
}
