package services

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentfactory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentupgradeable"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20utilityagent"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) ERC20UtilityAgentFetchCode(
	ctx context.Context,
	networkID uint64,
	contractAddress string,
) (string, error) {
	resp, err := s.GetEthereumClient(ctx, networkID).
		ERC20UtilityAgentFetchCode(
			contractAddress,
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return resp, nil
}

func (s *Service) ERC20UtilityAgentGetStorageInfo(
	ctx context.Context,
	networkID uint64,
	contractAddress string,
) (*erc20utilityagent.IUtilityAgentStorageInfo, error) {
	resp, err := s.GetEthereumClient(ctx, networkID).
		ERC20UtilityAgentGetStorageInfo(
			contractAddress,
		)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return resp, nil
}

func (s *Service) DeployAgentUpgradeableAddress(
	ctx context.Context,
	networkID uint64,
	agentID string,
	agentName string,
	agentVersion string,
	codeLanguage string,
	pointers []agentupgradeable.IAgentCodePointer,
	depsAgents []common.Address,
	agentOwner common.Address,
) (string, error) {
	memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "meme_pool_address"))
	agentFactoryAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "agent_factory_address"))
	txHash, err := s.GetEthereumClient(ctx, networkID).
		AgentFactoryCreateAgent(
			agentFactoryAddress,
			s.GetAddressPrk(memePoolAddress),
			common.HexToHash(agentID),
			agentName,
			agentVersion,
			codeLanguage,
			pointers,
			depsAgents,
			agentOwner,
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return txHash, nil
}

func (s *Service) DeployAgentUpgradeable(ctx context.Context, agentInfoID uint) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]any{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agentInfo != nil {
		if agentInfo.AgentType != models.AgentInfoAgentTypeModel &&
			agentInfo.AgentType != models.AgentInfoAgentTypeModelOnline &&
			agentInfo.AgentType != models.AgentInfoAgentTypeInfa &&
			agentInfo.AgentType != models.AgentInfoAgentTypeJs &&
			agentInfo.AgentType != models.AgentInfoAgentTypePython &&
			agentInfo.AgentType != models.AgentInfoAgentTypeCustomUi &&
			agentInfo.AgentType != models.AgentInfoAgentTypeCustomPrompt {
			return errs.NewError(errs.ErrBadRequest)
		}
		if agentInfo.MintHash == "" {
			switch agentInfo.NetworkID {
			case models.ETHEREUM_CHAIN_ID,
				models.BASE_CHAIN_ID,
				models.BASE_SEPOLIA_CHAIN_ID:
				{
					if agentInfo.SourceUrl == "" {
						return errs.NewError(errs.ErrBadRequest)
					}
					var fileNames []string
					err = json.Unmarshal([]byte(agentInfo.SourceUrl), &fileNames)
					if err != nil {
						return errs.NewError(err)
					}
					if len(fileNames) == 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					storageInfos := []agentupgradeable.IAgentCodePointer{}
					for _, fileName := range fileNames {
						if strings.HasPrefix(fileName, "ethfs_") {
							storageInfos = append(storageInfos, agentupgradeable.IAgentCodePointer{
								RetrieveAddress: helpers.HexToAddress(strings.ToLower(s.conf.GetConfigKeyString(agentInfo.NetworkID, "ethfs_address"))),
								FileType:        0,
								FileName:        strings.TrimPrefix(fileName, "ethfs_"),
							})
						} else if strings.HasPrefix(fileName, "ipfs_") {
							storageInfos = append(storageInfos, agentupgradeable.IAgentCodePointer{
								RetrieveAddress: helpers.HexToAddress(models.ETH_ZERO_ADDRESS),
								FileType:        0,
								FileName:        strings.TrimPrefix(fileName, "ipfs_"),
							})
						} else {
							storageInfos = append(storageInfos, agentupgradeable.IAgentCodePointer{
								RetrieveAddress: helpers.HexToAddress(models.ETH_ZERO_ADDRESS),
								FileType:        0,
								FileName:        fileName,
							})
						}
					}
					dependAgentAddrs := []common.Address{}
					if agentInfo.DependAgents != "" {
						var dependAgents []string
						err = json.Unmarshal([]byte(agentInfo.DependAgents), &dependAgents)
						if err != nil {
							return errs.NewError(err)
						}
						for _, v := range dependAgents {
							dependAgentAddrs = append(dependAgentAddrs, helpers.HexToAddress(v))
						}
					}
					codeLanguage := agentInfo.GetCodeLanguage()
					if codeLanguage == "" {
						return errs.NewError(errs.ErrBadRequest)
					}
					txHash, err := s.DeployAgentUpgradeableAddress(
						ctx,
						agentInfo.NetworkID,
						agentInfo.AgentID,
						agentInfo.AgentName,
						"1",
						codeLanguage,
						storageInfos,
						dependAgentAddrs,
						helpers.HexToAddress(agentInfo.Creator),
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).
						Model(agentInfo).
						Updates(
							map[string]any{
								"mint_hash": txHash,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			default:
				{
					return errs.NewError(errs.ErrBadRequest)
				}
			}
		}
	}
	return nil
}

func (s *Service) JobUpdateAgentUpgradeableCodeVersion(ctx context.Context) error {
	agents, err := s.dao.FindAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"agent_contract_address != ?": {""},
			"agent_type in (?)": {
				[]models.AgentInfoAgentType{
					models.AgentInfoAgentTypeModel,
					models.AgentInfoAgentTypeModelOnline,
					models.AgentInfoAgentTypeJs,
					models.AgentInfoAgentTypePython,
					models.AgentInfoAgentTypeInfa,
					models.AgentInfoAgentTypeCustomUi,
					models.AgentInfoAgentTypeCustomPrompt,
				},
			},
			"network_id in (?)": {
				[]uint64{
					models.SHARDAI_CHAIN_ID,
					models.ETHEREUM_CHAIN_ID,
					models.BITTENSOR_CHAIN_ID,
					models.BASE_CHAIN_ID,
					models.HERMES_CHAIN_ID,
					models.ARBITRUM_CHAIN_ID,
					models.ZKSYNC_CHAIN_ID,
					models.POLYGON_CHAIN_ID,
					models.BSC_CHAIN_ID,
					models.APE_CHAIN_ID,
					models.AVALANCHE_C_CHAIN_ID,
					models.ABSTRACT_TESTNET_CHAIN_ID,
					models.DUCK_CHAIN_ID,
					models.TRON_CHAIN_ID,
					models.MODE_CHAIN_ID,
					models.ZETA_CHAIN_ID,
					models.STORY_CHAIN_ID,
					models.HYPE_CHAIN_ID,
					models.MONAD_TESTNET_CHAIN_ID,
					models.MEGAETH_TESTNET_CHAIN_ID,
					models.CELO_CHAIN_ID,
					models.BASE_SEPOLIA_CHAIN_ID,
				},
			},
		},
		map[string][]any{},
		[]string{},
		0,
		999999,
	)
	if err != nil {
		return errs.NewError(err)
	}
	for _, agent := range agents {
		err = s.UpdateAgentUpgradeableCodeVersion(ctx, agent.ID)
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) HandleAgentUpgradeableCodePointerCreated(ctx context.Context, event *agentupgradeable.AgentUpgradeableCodePointerCreated) error {
	agentInfo, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"agent_contract_address = ?": {event.Raw.Address.Hex()},
		},
		map[string][]any{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agentInfo != nil {
		err = s.UpdateAgentUpgradeableCodeVersion(ctx, agentInfo.ID)
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) UpdateAgentUpgradeableCodeVersion(ctx context.Context, agentInfoID uint) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]any{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agentInfo.AgentContractAddress == "" ||
		agentInfo.AgentContractID != "0" {
		return errs.NewError(errs.ErrBadRequest)
	}
	codeVersion, err := s.GetEthereumClient(ctx, agentInfo.NetworkID).AgentUpgradeableCodeVersion(agentInfo.AgentContractAddress)
	if err != nil {
		return errs.NewError(err)
	}
	depsAgents, err := s.GetEthereumClient(ctx, agentInfo.NetworkID).AgentUpgradeableDepsAgents(agentInfo.AgentContractAddress, codeVersion)
	if err != nil {
		return errs.NewError(err)
	}
	depsAgentsBytes, err := json.Marshal(depsAgents)
	if err != nil {
		return errs.NewError(err)
	}
	depsAgentsJson := strings.ToLower(string(depsAgentsBytes))
	if codeVersion != agentInfo.CodeVersion || !strings.EqualFold(depsAgentsJson, agentInfo.DependAgents) {
		err = daos.GetDBMainCtx(ctx).
			Model(agentInfo).
			Updates(
				map[string]any{
					"code_version":  codeVersion,
					"depend_agents": depsAgentsJson,
				},
			).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) AgentFactoryAgentCreatedEvent(ctx context.Context, networkID uint64, event *agentfactory.AgentFactoryAgentCreated) error {
	if !s.conf.ExistsedConfigKey(networkID, "agent_factory_address") {
		return nil
	}
	agentFactoryAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "agent_factory_address"))
	if strings.EqualFold(agentFactoryAddress, event.Raw.Address.Hex()) {
		agentID := big.NewInt(0).SetBytes(event.AgentId[:]).Text(16)
		agentInfo, err := s.dao.FirstAgentInfo(
			daos.GetDBMainCtx(ctx),
			map[string][]any{
				"agent_id = ?": {agentID},
			},
			map[string][]any{},
			[]string{},
		)
		if err != nil {
			return errs.NewError(err)
		}
		agentAddress := strings.ToLower(event.Agent.Hex())
		if agentInfo != nil && !strings.EqualFold(agentAddress, agentInfo.AgentContractAddress) {
			err = daos.GetDBMainCtx(ctx).
				Model(agentInfo).
				Updates(
					map[string]any{
						"agent_contract_address": agentAddress,
						"agent_contract_id":      "0",
						"agent_logic_address":    "",
						"mint_hash":              event.Raw.TxHash.Hex(),
						"status":                 models.AssistantStatusReady,
						"reply_enabled":          true,
						"agent_nft_minted":       true,
						"factory_address":        strings.ToLower(event.Raw.Address.Hex()),
					},
				).Error
			if err != nil {
				return errs.NewError(err)
			}
			s.DeleteFilterAddrs(ctx, agentInfo.NetworkID)
			go s.UpdateAgentUpgradeableCodeVersion(ctx, agentInfo.ID)
		}
	}
	return nil
}
