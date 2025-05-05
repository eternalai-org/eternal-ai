package services

import (
	"context"
	"errors"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/evmapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/zkapi"
	"github.com/google/uuid"
)

func (s *Service) CreateETHAddress(ctx context.Context) (string, error) {
	addr, prk, err := ethapi.CreateETHAddress()
	if err != nil {
		return "", err
	}
	addr, _, err = s.coreClient.StoreAddress(
		addr,
		prk,
		0,
		"ethereum",
	)
	if err != nil {
		return "", err
	}
	return strings.ToLower(addr), nil
}

func (s *Service) CreateTRONAddress(ctx context.Context) (string, error) {
	addr, prk, err := trxapi.CreateTRONAddress()
	if err != nil {
		return "", err
	}
	addr, _, err = s.coreClient.StoreAddress(
		addr,
		prk,
		0,
		"tron",
	)
	if err != nil {
		return "", err
	}
	return addr, nil
}

func (s *Service) StoreAddress(ctx context.Context, address, prk string) (string, error) {
	address, _, err := s.coreClient.StoreAddress(
		address,
		prk,
		0,
		"ethereum",
	)
	if err != nil {
		return "", err
	}
	return strings.ToLower(address), nil
}

func (s *Service) CreateBTCAddress(ctx context.Context) (string, error) {
	addr, prk, err := s.btcAPI.Address()
	if err != nil {
		return "", err
	}
	addr, _, err = s.coreClient.StoreAddress(
		addr,
		prk,
		0,
		"bitcoin",
	)
	if err != nil {
		return "", err
	}
	return addr, nil
}

func (s *Service) CreateSOLAddress(ctx context.Context) (string, error) {
	addr, err := s.blockchainUtils.SolanaAddress()
	if err != nil {
		return "", err
	}
	return addr, nil
}

func (s *Service) CreateAPIKey(ctx context.Context) (string, string, error) {
	apiKey := uuid.NewString()
	secrectKey := uuid.NewString()
	apiKey, _, err := s.coreClient.StoreAddress(
		apiKey,
		secrectKey,
		0,
		"apiKey",
	)
	if err != nil {
		return "", "", err
	}
	return apiKey, secrectKey, nil
}

func (s *Service) GetEthereumClient(ctx context.Context, networkID uint64) *ethapi.Client {
	s.jobMutex.Lock()
	defer s.jobMutex.Unlock()
	ethApi, ok := s.ethApiMap[networkID]
	if !ok {
		rpcUrl := s.conf.GetConfigKeyString(
			networkID,
			"rpc_url",
		)
		minGasPrice := "0"
		if s.conf.ExistsedConfigKey(
			networkID,
			"min_gas_price",
		) {
			minGasPrice = s.conf.GetConfigKeyString(
				networkID,
				"min_gas_price",
			)
		}
		isBtcL1 := s.conf.GetConfigKeyBool(networkID, "is_btc_l1")
		if rpcUrl != "" {
			ethApi = &ethapi.Client{
				BaseURL:           rpcUrl,
				MinGasPrice:       minGasPrice,
				BTCL1:             isBtcL1,
				BlockTimeDisabled: true,
			}
			if isBtcL1 {
				ethApi.InscribeTxsLog = s.InscribeTxsLog
			}
			if s.conf.ExistsedConfigKey(
				networkID,
				"gas_rpc_url",
			) {
				ethApi.BaseGasURL = s.conf.GetConfigKeyString(
					networkID,
					"gas_rpc_url",
				)
			}
			s.ethApiMap[networkID] = ethApi
		} else {
			panic(errors.New("rpc url is not found"))
		}
	}
	return ethApi
}

func (s *Service) GetZkClient(ctx context.Context, networkID uint64) *zkapi.Client {
	s.jobMutex.Lock()
	defer s.jobMutex.Unlock()
	zkApi, ok := s.zkApiMap[networkID]
	if !ok {
		rpcUrl := s.conf.GetConfigKeyString(
			networkID,
			"rpc_url",
		)
		minGasPrice := "0"
		if s.conf.ExistsedConfigKey(
			networkID,
			"min_gas_price",
		) {
			minGasPrice = s.conf.GetConfigKeyString(
				networkID,
				"min_gas_price",
			)
		}
		if rpcUrl != "" {
			var paymasterAddress, paymasterToken string
			var paymasterFeeZero bool
			if s.conf.ExistsedConfigKey(networkID, "paymaster_address") &&
				s.conf.ExistsedConfigKey(networkID, "paymaster_token") {
				paymasterAddress = s.conf.GetConfigKeyString(networkID, "paymaster_address")
				paymasterToken = s.conf.GetConfigKeyString(networkID, "paymaster_token")
				paymasterFeeZero = s.conf.GetConfigKeyBool(networkID, "paymaster_fee_zero")
			}
			zkApi = &zkapi.Client{
				BaseURL:          rpcUrl,
				MinGasPrice:      minGasPrice,
				PaymasterAddress: paymasterAddress,
				PaymasterToken:   paymasterToken,
				PaymasterFeeZero: paymasterFeeZero,
			}
			s.zkApiMap[networkID] = zkApi
		} else {
			panic(errors.New("rpc url is not found"))
		}
	}
	return zkApi
}

func (s *Service) GetEVMClient(ctx context.Context, networkID uint64) evmapi.BaseClient {
	if !s.conf.GetConfigKeyBool(networkID, "zk_sync") {
		return s.GetEthereumClient(ctx, networkID)
	} else {
		return s.GetZkClient(ctx, networkID)
	}
}

func (s *Service) InscribeTxsLog(txHash string, inscribeTxHash string, logErr string) {
	go func() error {
		err := daos.GetDBMainCtx(context.Background()).
			Create(
				&models.BTCL1InscribeTx{
					TxHash:         txHash,
					InscribeTxHash: inscribeTxHash,
					Error:          logErr,
				},
			).Error
		if err != nil {
			return errs.NewError(err)
		}
		return nil
	}()
}

func (s *Service) ApproveMaxEAIForAddress(
	ctx context.Context,
	networkID uint64,
	address string,
	spender string,
) (string, error) {
	eaiTokenAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
	txHash, err := s.GetEthereumClient(ctx, networkID).
		Erc20ApproveMax(
			eaiTokenAddress,
			s.GetAddressPrk(address),
			spender,
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return txHash, nil
}

func (s *Service) DeployProxyAdminAddress(
	ctx context.Context,
	networkID uint64,
) (string, string, error) {
	memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "meme_pool_address"))
	contractAddress, txHash, err := s.GetEthereumClient(ctx, networkID).
		DeployProxyAdmin(
			s.GetAddressPrk(memePoolAddress),
		)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	return contractAddress, txHash, nil
}
