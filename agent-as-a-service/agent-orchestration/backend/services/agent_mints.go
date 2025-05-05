package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/systempromptmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (s *Service) JobAgentMintNft(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobAgentMintNft",
		func() error {
			var retErr error
			{
				agents, err := s.dao.FindAgentInfoJoin(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						"join agent_chain_fees on agent_chain_fees.network_id = agent_infos.network_id": {},
					},
					map[string][]any{
						"agent_infos.agent_type in (?)": {
							[]models.AgentInfoAgentType{
								models.AgentInfoAgentTypeNormal,
								models.AgentInfoAgentTypeReasoning,
								models.AgentInfoAgentTypeEliza,
								models.AgentInfoAgentTypeZerepy,
							},
						},
						"agent_infos.agent_id != ''":        {},
						"agent_infos.agent_contract_id = ?": {""},
						"agent_infos.agent_nft_minted = ?":  {false},
						`agent_infos.twin_twitter_usernames is null 
						or agent_infos.twin_twitter_usernames = '' 
						or (agent_infos.twin_twitter_usernames != '' and agent_infos.twin_status = ?)
					`: {models.TwinStatusDoneSuccess},
						"agent_infos.scan_enabled = ?":    {true},
						"agent_infos.system_prompt != ''": {},
						"agent_infos.eai_balance > 0":     {},
						`(agent_infos.ref_tweet_id > 0 
						or (agent_infos.eai_balance >= (agent_chain_fees.mint_fee + 9.9 * agent_chain_fees.infer_fee))
						or (agent_infos.agent_type in (3,4) and agent_infos.eai_balance >= agent_chain_fees.mint_fee)
						)`: {},
						"agent_infos.network_id in (?)": {
							[]uint64{
								models.SHARDAI_CHAIN_ID,
								models.ETHEREUM_CHAIN_ID,
								models.BITTENSOR_CHAIN_ID,
								models.SOLANA_CHAIN_ID,
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
					[]string{
						"updated_at asc",
					},
					0,
					10,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agent := range agents {
					err = s.AgentMintNft(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
					err = s.AgentCreateMissionDefault(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
					time.Sleep(10 * time.Second)
				}
			}
			{
				agents, err := s.dao.FindAgentInfoJoin(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						"join agent_chain_fees on agent_chain_fees.network_id = agent_infos.network_id": {},
					},
					map[string][]any{
						"agent_infos.agent_contract_id = ?": {""},
						"agent_infos.agent_nft_minted = ?":  {false},
						"agent_infos.eai_balance > 0":       {},
						`(1 != 1
							or (agent_infos.agent_type in (?) and agent_infos.eai_balance >= agent_chain_fees.agent_deploy_fee)
						)`: {
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
						"agent_infos.network_id in (?)": {
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
					[]string{
						"updated_at asc",
					},
					0,
					10,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agent := range agents {
					err = s.AgentMintNft(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
					time.Sleep(10 * time.Second)
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentMintNft(ctx context.Context, agentInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentMintNft_%d", agentInfoID),
		func() error {
			agent, err := s.dao.FirstAgentInfoByID(
				daos.GetDBMainCtx(ctx),
				agentInfoID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agent.AgentContractID == "" &&
				!agent.AgentNftMinted {
				var isOk bool
				mintFee := numeric.NewFloatFromString("0.0")
				checkFee := numeric.NewFloatFromString("0.0")
				if agent.RefTweetID <= 0 {
					agentChainFee, err := s.GetAgentChainFee(
						daos.GetDBMainCtx(ctx),
						agent.NetworkID,
					)
					if err != nil {
						return errs.NewError(err)
					}
					switch agent.AgentType {
					case models.AgentInfoAgentTypeNormal,
						models.AgentInfoAgentTypeReasoning:
						{
							mintFee = &agentChainFee.MintFee.Float
							checkFee = models.AddBigFloats(&agentChainFee.MintFee.Float, models.MulBigFloats(&agentChainFee.InferFee.Float, big.NewFloat(9.9)))
						}
					case models.AgentInfoAgentTypeEliza,
						models.AgentInfoAgentTypeZerepy:
						{
							mintFee = &agentChainFee.MintFee.Float
							checkFee = &agentChainFee.MintFee.Float
						}
					case models.AgentInfoAgentTypeModel,
						models.AgentInfoAgentTypeModelOnline,
						models.AgentInfoAgentTypeJs,
						models.AgentInfoAgentTypePython,
						models.AgentInfoAgentTypeInfa,
						models.AgentInfoAgentTypeCustomUi,
						models.AgentInfoAgentTypeCustomPrompt:
						{
							mintFee = &agentChainFee.AgentDeployFee.Float
							checkFee = &agentChainFee.AgentDeployFee.Float
						}
					default:
						{
							return errs.NewError(errs.ErrBadRequest)
						}
					}
				}
				if agent.EaiBalance.Float.Cmp(checkFee) >= 0 {
					isOk = true
				}
				if isOk {
					updateAgentFields := map[string]any{
						"agent_nft_minted": true,
					}
					if agent.TokenMode == string(models.TokenSetupEnumAutoCreate) &&
						agent.TokenAddress == "" &&
						agent.TokenStatus == "" {
						updateAgentFields["token_status"] = "pending"
					}
					err = daos.GetDBMainCtx(ctx).
						Model(agent).
						Updates(updateAgentFields).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					switch agent.AgentType {
					case models.AgentInfoAgentTypeNormal,
						models.AgentInfoAgentTypeReasoning,
						models.AgentInfoAgentTypeKnowledgeBase,
						models.AgentInfoAgentTypeEliza,
						models.AgentInfoAgentTypeZerepy:
						{
							for range 5 {
								err = s.MintAgent(ctx, agent.ID)
								if err == nil {
									break
								}
							}
						}
					case models.AgentInfoAgentTypeModel,
						models.AgentInfoAgentTypeModelOnline,
						models.AgentInfoAgentTypeJs,
						models.AgentInfoAgentTypePython,
						models.AgentInfoAgentTypeInfa,
						models.AgentInfoAgentTypeCustomUi,
						models.AgentInfoAgentTypeCustomPrompt:
						{
							for range 2 {
								err = s.DeployAgentUpgradeable(ctx, agent.ID)
								if err == nil {
									break
								}
							}
						}
					default:
						{
							err = errs.NewError(errs.ErrBadRequest)
						}
					}
					if err != nil {
						_ = daos.GetDBMainCtx(ctx).
							Model(agent).
							Updates(
								map[string]any{
									"scan_error": "mint nft error " + err.Error(),
								},
							).
							Error
						return errs.NewError(err)
					} else {
						if mintFee.Cmp(big.NewFloat(0)) > 0 {
							err = daos.WithTransaction(
								daos.GetDBMainCtx(ctx),
								func(tx *gorm.DB) error {
									err = s.dao.Create(
										tx,
										&models.AgentEaiTopup{
											NetworkID:      agent.NetworkID,
											EventId:        fmt.Sprintf("agent_mint_fee_%d", agent.ID),
											AgentInfoID:    agent.ID,
											Type:           models.AgentEaiTopupTypeSpent,
											Amount:         numeric.NewBigFloatFromFloat(mintFee),
											Status:         models.AgentEaiTopupStatusDone,
											DepositAddress: agent.ETHAddress,
											ToAddress:      agent.ETHAddress,
											Toolset:        "mint_fee",
										},
									)
									if err != nil {
										return errs.NewError(err)
									}
									err = tx.
										Model(agent).
										Updates(
											map[string]any{
												"eai_balance": gorm.Expr("eai_balance - ?", numeric.NewBigFloatFromFloat(mintFee)),
												"mint_fee":    numeric.NewBigFloatFromFloat(mintFee),
											},
										).
										Error
									if err != nil {
										return errs.NewError(err)
									}
									return nil
								},
							)
							if err != nil {
								return errs.NewError(err)
							}
						}
					}
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

func (s *Service) JobRetryAgentMintNft(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobRetryAgentMintNft",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"agent_type in (?)": {
						[]models.AgentInfoAgentType{
							models.AgentInfoAgentTypeNormal,
							models.AgentInfoAgentTypeReasoning,
							models.AgentInfoAgentTypeEliza,
							models.AgentInfoAgentTypeZerepy,
						},
					},
					"updated_at <= ?":       {time.Now().Add(-60 * time.Minute)},
					"agent_contract_id = ?": {""},
					"agent_nft_minted = ?":  {true},
					"mint_hash != ?":        {""},
					"network_id in (?)": {
						[]uint64{
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
							models.MODE_CHAIN_ID,
							models.ZETA_CHAIN_ID,
							models.STORY_CHAIN_ID,
							models.HYPE_CHAIN_ID,
							models.MONAD_TESTNET_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				1000,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = s.GetEVMClient(ctx, agent.NetworkID).TransactionConfirmed(agent.MintHash)
				if err != nil {
					fmt.Println(err.Error())
					if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "transaction is not Successful") {
						err = daos.GetDBMainCtx(ctx).
							Model(agent).
							Updates(
								map[string]any{
									"mint_hash": "",
								},
							).
							Error
						if err != nil {
							return errs.NewError(err)
						}
						for i := 0; i < 5; i++ {
							err = s.MintAgent(ctx, agent.ID)
							if err == nil {
								break
							}
						}
					}
				} else {
					s.MemeEventsByTransaction(ctx, agent.NetworkID, agent.MintHash)
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) JobRetryAgentMintNftError(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobRetryAgentMintNftError",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"agent_infos.agent_type in (?)": {
						[]models.AgentInfoAgentType{
							models.AgentInfoAgentTypeNormal,
							models.AgentInfoAgentTypeReasoning,
							models.AgentInfoAgentTypeEliza,
							models.AgentInfoAgentTypeZerepy,
						},
					},
					"updated_at <= ?":                     {time.Now().Add(-60 * time.Minute)},
					"agent_contract_id = ?":               {""},
					"agent_nft_minted = ?":                {true},
					"mint_hash is null or  mint_hash = ?": {""},
					"scan_error != ?":                     {""},
					"network_id in (?)": {
						[]uint64{
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
				[]string{
					"rand()",
				},
				0,
				1000,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				if strings.Contains(agent.ScanError, "mint nft error") {
					err = daos.GetDBMainCtx(ctx).
						Model(agent).
						Updates(
							map[string]any{
								"agent_nft_minted": false,
								"scan_error":       "",
							},
						).
						Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MintAgent(ctx context.Context, agentInfoID uint) error {
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
		if agentInfo.MintHash == "" {
			switch agentInfo.NetworkID {
			case models.LOCAL_CHAIN_ID:
				{
					agentUriData := models.AgentUriData{
						Name: agentInfo.AgentName,
					}
					agentUriBytes, err := json.Marshal(agentUriData)
					if err != nil {
						return errs.NewError(err)
					}
					uriHash := string(agentUriBytes)
					systemContentHash := agentInfo.SystemPrompt
					modelID, err := s.GetEthereumClient(ctx, agentInfo.NetworkID).GPUManagerGetModelID(
						s.conf.GetConfigKeyString(agentInfo.NetworkID, "gpu_manager_contract_address"),
					)
					if err != nil {
						return errs.NewError(err)
					}
					agentContractAddress := s.conf.GetConfigKeyString(agentInfo.NetworkID, "dagent721_contract_address")
					txHash, err := s.GetEthereumClient(ctx, agentInfo.NetworkID).Dagent721Mint(
						agentContractAddress,
						s.GetAddressPrk(
							helpers.RandomInStrings(
								strings.Split(s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_admin_address"), ","),
							),
						),
						helpers.HexToAddress(agentInfo.Creator),
						uriHash,
						[]byte(systemContentHash),
						models.ConvertBigFloatToWei(&agentInfo.InferFee.Float, 18),
						"ai721",
						helpers.HexToAddress(s.conf.GetConfigKeyString(agentInfo.NetworkID, "prompt_scheduler_contract_address")),
						modelID,
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).
						Model(agentInfo).
						Updates(
							map[string]any{
								"agent_contract_address": agentContractAddress,
								"mint_hash":              txHash,
								"status":                 models.AssistantStatusMinting,
								"reply_enabled":          true,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			case models.SHARDAI_CHAIN_ID,
				models.HERMES_CHAIN_ID,
				models.BASE_CHAIN_ID,
				models.ETHEREUM_CHAIN_ID,
				models.ARBITRUM_CHAIN_ID,
				models.BSC_CHAIN_ID,
				models.POLYGON_CHAIN_ID,
				models.ZKSYNC_CHAIN_ID,
				models.APE_CHAIN_ID,
				models.AVALANCHE_C_CHAIN_ID,
				models.ABSTRACT_TESTNET_CHAIN_ID,
				models.BITTENSOR_CHAIN_ID,
				models.DUCK_CHAIN_ID,
				models.MODE_CHAIN_ID,
				models.ZETA_CHAIN_ID,
				models.STORY_CHAIN_ID,
				models.HYPE_CHAIN_ID,
				models.MONAD_TESTNET_CHAIN_ID,
				models.MEGAETH_TESTNET_CHAIN_ID,
				models.CELO_CHAIN_ID,
				models.BASE_SEPOLIA_CHAIN_ID:
				{
					agentUriData := models.AgentUriData{
						Name: agentInfo.AgentName,
					}
					agentUriBytes, err := json.Marshal(agentUriData)
					if err != nil {
						return errs.NewError(err)
					}
					uriHash, err := s.IpfsUploadDataForName(ctx, fmt.Sprintf("%v_%v", agentInfo.AgentID, "uri"), agentUriBytes)
					if err != nil {
						return errs.NewError(err)
					}
					systemContentHash, err := s.IpfsUploadDataForName(ctx, fmt.Sprintf("%v_%v", agentInfo.AgentID, "system_content"), []byte(agentInfo.SystemPrompt))
					if err != nil {
						return errs.NewError(err)
					}
					agentContractAddress := s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_contract_address")
					txHash, err := s.GetEVMClient(ctx, agentInfo.NetworkID).SystemPromptManagerMint(
						agentContractAddress,
						s.GetAddressPrk(
							helpers.RandomInStrings(
								strings.Split(s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_admin_address"), ","),
							),
						),
						helpers.HexToAddress(agentInfo.Creator),
						"ipfs://"+uriHash,
						[]byte("ipfs://"+systemContentHash),
						models.ConvertBigFloatToWei(&agentInfo.InferFee.Float, 18),
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).
						Model(agentInfo).
						Updates(
							map[string]any{
								"agent_contract_address": agentContractAddress,
								"mint_hash":              txHash,
								"status":                 models.AssistantStatusMinting,
								"reply_enabled":          true,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			case models.SOLANA_CHAIN_ID:
				{
					err = daos.GetDBMainCtx(ctx).
						Model(agentInfo).
						Updates(
							map[string]any{
								"agent_contract_address": s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_contract_address"),
								"agent_contract_id":      strconv.FormatUint(uint64(agentInfo.ID), 10),
								"status":                 models.AssistantStatusReady,
								"reply_enabled":          true,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			case models.TRON_CHAIN_ID:
				{
					agentUriData := models.AgentUriData{
						Name: agentInfo.AgentName,
					}
					agentUriBytes, err := json.Marshal(agentUriData)
					if err != nil {
						return errs.NewError(err)
					}
					uriHash, err := s.IpfsUploadDataForName(ctx, fmt.Sprintf("%v_%v", agentInfo.AgentID, "uri"), agentUriBytes)
					if err != nil {
						return errs.NewError(err)
					}
					systemContentHash, err := s.IpfsUploadDataForName(ctx, fmt.Sprintf("%v_%v", agentInfo.AgentID, "system_content"), []byte(agentInfo.SystemPrompt))
					if err != nil {
						return errs.NewError(err)
					}
					txHash, err := s.trxApi.SystemPromptManagerMint(
						s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_contract_address"),
						s.GetAddressPrk(
							helpers.RandomInStrings(
								strings.Split(s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_admin_address"), ","),
							),
						),
						helpers.HexToAddress(agentInfo.Creator),
						"ipfs://"+uriHash,
						[]byte("ipfs://"+systemContentHash),
						models.ConvertBigFloatToWei(&agentInfo.InferFee.Float, 18),
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).
						Model(agentInfo).
						Updates(
							map[string]any{
								"agent_contract_address": s.conf.GetConfigKeyString(agentInfo.NetworkID, "agent_contract_address"),
								"mint_hash":              txHash,
								"status":                 models.AssistantStatusMinting,
								"reply_enabled":          true,
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

func (s *Service) SystemPromptManagerNewTokenEvent(ctx context.Context, networkID uint64, event *systempromptmanager.SystemPromptManagerNewToken) error {
	agentInfo, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"network_id = ?": {networkID},
			"mint_hash = ?":  {event.Raw.TxHash.Hex()},
		},
		map[string][]any{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agentInfo != nil {
		var info models.AgentUriData
		var systemPrompt []byte
		switch agentInfo.NetworkID {
		case models.LOCAL_CHAIN_ID:
			{
				data := []byte(event.Uri)
				systemPrompt = event.SysPrompt
				err = json.Unmarshal(data, &info)
				if err != nil {
					return errs.NewError(err)
				}
			}
		default:
			{
				data, _, err := lighthouse.DownloadDataSimple(event.Uri)
				if err != nil {
					dataInfo := map[string]any{}
					err = json.Unmarshal([]byte(event.Uri), &dataInfo)
					if err != nil {
						return errs.NewError(err)
					}
					uri := dataInfo["agent_uri"].(string)
					data, _, err = lighthouse.DownloadDataSimple(uri)
					if err != nil {
						return errs.NewError(err)
					}
				}
				systemPrompt, _, err = lighthouse.DownloadDataSimple(string(event.SysPrompt))
				if err != nil {
					return errs.NewError(err)
				}
				err = json.Unmarshal(data, &info)
				if err != nil {
					return errs.NewError(err)
				}
			}
		}
		err = daos.GetDBMainCtx(ctx).
			Model(agentInfo).
			Updates(
				map[string]any{
					"agent_name":        info.Name,
					"creator":           strings.ToLower(event.Minter.Hex()),
					"agent_contract_id": event.TokenId.String(),
					"status":            models.AssistantStatusReady,
					"system_prompt":     string(systemPrompt),
					"reply_enabled":     true,
				},
			).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) SystemPromptManagerAgentDataUpdateEvent(ctx context.Context, networkID uint64, event *systempromptmanager.SystemPromptManagerAgentDataUpdate) error {
	contractAgentID := event.AgentId.String()
	lightHouseHash := string(event.NewSysPrompt)
	systemPromptBytes, _, err := lighthouse.DownloadDataSimple(lightHouseHash)
	if err != nil {
		if !strings.HasPrefix(lightHouseHash, "ipfs://") {
			systemPromptBytes = event.NewSysPrompt
		} else {
			return errs.NewError(err)
		}
	}
	err = daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where("network_id = ?", networkID).
		Where("agent_contract_address = ?", s.GetEVMClient(ctx, networkID).ConvertAddressForOut(strings.ToLower(event.Raw.Address.Hex()))).
		Where("agent_contract_id = ?", contractAgentID).
		Updates(
			map[string]any{
				"system_prompt": string(systemPromptBytes),
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) SystemPromptManagerAgentURIUpdateEvent(ctx context.Context, networkID uint64, event *systempromptmanager.SystemPromptManagerAgentURIUpdate) error {
	contractAgentID := event.AgentId.String()
	uri := event.Uri
	dataBytes, _, err := lighthouse.DownloadDataSimple(uri)
	if err != nil {
		return err
	}
	agentUriData := models.AgentUriData{}
	err = json.Unmarshal(dataBytes, &agentUriData)
	if err != nil {
		return err
	}
	err = daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where("network_id = ?", networkID).
		Where("agent_contract_address = ?", s.GetEVMClient(ctx, networkID).ConvertAddressForOut(strings.ToLower(event.Raw.Address.Hex()))).
		Where("agent_contract_id = ?", contractAgentID).
		Updates(
			map[string]any{
				"agent_name": agentUriData.Name,
				"uri":        uri,
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) ExecuteUpdateAgentInfoInContract(ctx context.Context, assistant *models.AgentInfo, request *serializers.UpdateAgentAssistantInContractRequest) (*serializers.UpdateAgentAssistantInContractResponse, error) {
	var txUpdateNameHash, txUpdateSystemPromptHash string
	switch assistant.NetworkID {
	case models.BASE_CHAIN_ID,
		models.ETHEREUM_CHAIN_ID,
		models.ARBITRUM_CHAIN_ID,
		models.BSC_CHAIN_ID,
		models.POLYGON_CHAIN_ID,
		models.ZKSYNC_CHAIN_ID:
		{
			ethClient := s.GetEVMClient(ctx, assistant.NetworkID)
			instanceABI, err := abi.JSON(strings.NewReader(systempromptmanager.SystemPromptManagerMetaData.ABI))
			if err != nil {
				return nil, errs.NewError(err)
			}
			agentId, ok := new(big.Int).SetString(assistant.AgentContractID, 10)
			if !ok {
				return nil, errs.NewError(fmt.Errorf("error while getting agent id :%v", assistant.AgentContractID))
			}
			{
				randomNonceName, ok := new(big.Int).SetString(request.RandomNonceName, 10)
				if !ok {
					return nil, fmt.Errorf("error while getting  random nonce name:%v", request.RandomNonceName)
				}
				input, err := instanceABI.Pack("updateAgentUriWithSignature", agentId, request.HashName, randomNonceName, common.Hex2Bytes(request.SignatureName[2:]))
				if err != nil {
					return nil, errs.NewError(err)
				}
				txUpdateNameHash, err = ethClient.Transact(
					assistant.AgentContractAddress,
					s.GetAddressPrk(
						helpers.RandomInStrings(
							strings.Split(s.conf.GetConfigKeyString(assistant.NetworkID, "agent_admin_address"), ","),
						),
					),
					input,
					common.Big0,
				)
				if err != nil {
					return nil, errs.NewError(err)
				}
			}
			{
				randomNonceSystemPrompt, ok := new(big.Int).SetString(request.RandomNonceSystemPrompt, 10)
				if !ok {
					return nil, fmt.Errorf("error while getting random nonce system prompt :%v", request.RandomNonceSystemPrompt)
				}
				input, err := instanceABI.Pack("updateAgentDataWithSignature", agentId, []byte(request.HashSystemPrompt),
					big.NewInt(0), randomNonceSystemPrompt, common.Hex2Bytes(request.SignatureSystemPrompt[2:]))
				if err != nil {
					return nil, errs.NewError(err)
				}
				txUpdateSystemPromptHash, err = ethClient.Transact(
					assistant.AgentContractAddress,
					s.GetAddressPrk(
						helpers.RandomInStrings(
							strings.Split(s.conf.GetConfigKeyString(assistant.NetworkID, "agent_admin_address"), ","),
						),
					),
					input,
					common.Big0,
				)
				if err != nil {
					return nil, errs.NewError(err)
				}
			}
		}
	default:
		{
			txUpdateNameHash = uuid.NewString()
			txUpdateSystemPromptHash = uuid.NewString()
		}
	}
	return &serializers.UpdateAgentAssistantInContractResponse{
		TxUpdateSystemPrompt: txUpdateSystemPromptHash,
		TxUpdateName:         txUpdateNameHash,
	}, nil
}

func (s *Service) JobAgentStart(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobAgentStart",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"agent_infos.agent_type in (3,4)":     {},
					"agent_infos.agent_contract_id != ''": {},
					"agent_infos.deployed_ref_id = ''":    {},
					"agent_infos.network_id in (?)": {
						[]uint64{
							models.SHARDAI_CHAIN_ID,
							models.ETHEREUM_CHAIN_ID,
							models.BITTENSOR_CHAIN_ID,
							models.SOLANA_CHAIN_ID,
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
							models.CELO_CHAIN_ID,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = func() error {
					startReq := map[string]any{}
					err = json.Unmarshal([]byte(agent.ConfigData), &startReq)
					if err != nil {
						return errs.NewError(err)
					}
					startReq["agentName"] = agent.AgentName
					startReq["model"] = agent.AgentBaseModel
					startReq["chainID"] = fmt.Sprintf("%d", agent.NetworkID)
					startReq["agentContractAddress"] = agent.AgentContractAddress
					startReq["agentID"] = agent.AgentContractID
					err = helpers.CurlURL(
						s.conf.AgentDeployer.Url+"/api/agent/start",
						http.MethodPost,
						map[string]string{},
						startReq,
						nil,
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).Model(agent).Updates(
						map[string]any{
							"deployed_ref_id": helpers.RandomBigInt(12).Text(16),
						},
					).Error
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				}()
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewError(err))
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
