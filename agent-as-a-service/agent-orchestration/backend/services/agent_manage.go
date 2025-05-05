package services

import (
	"context"
	"encoding/hex"
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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/delegate_cash"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/hiro"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/magiceden"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) GetModelDefaultByChainID(chainID uint64) string {
	var baseModel string
	s.RedisCached(
		fmt.Sprintf("GetModelDefaultByChainID_%d", chainID),
		true,
		10*time.Minute,
		&baseModel,
		func() (interface{}, error) {
			baseModel = "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
			listConfig, _ := s.dojoAPI.GetChainConfigs()
			for _, chain := range listConfig {
				if strings.EqualFold(chain.ChainId, fmt.Sprintf("%d", chainID)) {
					for modelName := range chain.SupportModelNames {
						baseModel = modelName
					}
					break
				}
			}
			return baseModel, nil
		},
	)
	fmt.Println(baseModel)
	return baseModel
}

func (s *Service) AgentCreateAgentAssistant(ctx context.Context, address string, req *serializers.AssistantsReq) (*models.AgentInfo, error) {
	if req.SystemContent == "" {
		req.SystemContent = "default"
	}
	agent := &models.AgentInfo{
		Version:          "2",
		AgentCategoryID:  req.CategoryID,
		AgentType:        models.AgentInfoAgentTypeReasoning,
		AgentID:          helpers.RandomBigInt(12).Text(16),
		Status:           models.AssistantStatusPending,
		NetworkID:        req.ChainID,
		NetworkName:      models.GetChainName(req.ChainID),
		AgentName:        req.AgentName,
		Creator:          strings.ToLower(address),
		MetaData:         req.SystemContent,
		SystemPrompt:     req.SystemContent,
		AgentBaseModel:   req.AgentBaseModel,
		Bio:              req.GetAssistantCharacter(req.Bio),
		Lore:             req.GetAssistantCharacter(req.Lore),
		Knowledge:        req.GetAssistantCharacter(req.Knowledge),
		MessageExamples:  req.GetAssistantCharacter(req.MessageExamples),
		PostExamples:     req.GetAssistantCharacter(req.PostExamples),
		Topics:           req.GetAssistantCharacter(req.Topics),
		Style:            req.GetAssistantCharacter(req.Style),
		Adjectives:       req.GetAssistantCharacter(req.Adjectives),
		SocialInfo:       req.GetAssistantCharacter(req.SocialInfo),
		ScanEnabled:      false,
		VerifiedNftOwner: req.VerifiedNFTOwner,
		NftAddress:       req.NFTAddress,
		NftTokenID:       req.NFTTokenID,
		NftOwnerAddress:  req.NFTOwnerAddress,
		Thumbnail:        req.Thumbnail,
		NftTokenImage:    req.NFTTokenImage,
		TokenImageUrl:    req.TokenImageUrl,
		MissionTopics:    req.MissionTopics,
		ConfigData:       req.ConfigData,
		SourceUrl:        req.SourceUrl,
		AuthenUrl:        req.AuthenUrl,
		DependAgents:     req.DependAgents,
		EnvExample:       req.EnvExample,
		CodeVersion:      1,
		Author:           req.Author,
	}
	if req.RequiredEnv != nil {
		agent.RequiredEnv = *req.RequiredEnv
	}

	if req.RequiredWallet != nil {
		agent.RequiredWallet = *req.RequiredWallet
	}
	if req.IsOnchain != nil {
		agent.IsOnchain = *req.IsOnchain
	}
	if req.IsStreaming != nil {
		agent.IsStreaming = *req.IsStreaming
	}
	if req.IsCustomUi != nil {
		agent.IsCustomUi = *req.IsCustomUi
	}
	if req.RequiredInfo != nil {
		agent.RequiredInfo = *req.RequiredInfo
	}
	if req.InferFee != nil {
		inferFee := numeric.NewBigFloatFromString(*req.InferFee)
		agent.InferFee = inferFee
	}
	if req.DisplayName != "" {
		agent.DisplayName = req.DisplayName
	}
	if req.ShortDescription != "" {
		agent.ShortDescription = req.ShortDescription
	}
	agent.MinFeeToUse = req.MinFeeToUse
	agent.Worker = req.Worker
	if req.IsForceUpdate != nil {
		agent.IsForceUpdate = *req.IsForceUpdate
	}

	tokenInfo, _ := s.GenerateTokenInfoFromSystemPrompt(ctx, req.AgentName, req.SystemContent)
	if tokenInfo != nil && tokenInfo.TokenSymbol != "" {
		agent.TokenSymbol = tokenInfo.TokenSymbol
		agent.TokenName = req.AgentName
		agent.TokenDesc = tokenInfo.TokenDesc
		if req.TokenImageUrl == "" {
			agent.TokenImageUrl = tokenInfo.TokenImageUrl
		}
	} else {
		agent.TokenName = req.AgentName
		agent.TokenSymbol = helpers.GenerateTokenSymbol(req.AgentName)
	}

	if req.CreateTokenMode == models.CreateTokenModeTypeLinkExisting {
		agent.TokenMode = string(models.CreateTokenModeTypeLinkExisting)
		agent.TokenAddress = req.TokenAddress
		agent.TokenSymbol = req.Ticker
		agent.TokenName = req.TokenName
		tokenChainId, _ := strconv.ParseUint(req.TokenChainId, 0, 64)
		agent.TokenNetworkID = tokenChainId
	} else if req.TokenChainId != "" {
		tokenChainId, _ := strconv.ParseUint(req.TokenChainId, 0, 64)
		if !(tokenChainId == models.POLYGON_CHAIN_ID || tokenChainId == models.ZKSYNC_CHAIN_ID) {
			agent.TokenNetworkID = tokenChainId
			if req.CreateTokenMode != "" {
				agent.TokenMode = string(req.CreateTokenMode)
			}
		} else {
			agent.TokenNetworkID = models.GENERTAL_NETWORK_ID
			agent.TokenMode = string(models.TokenSetupEnumNoToken)
		}
	}

	if agent.AgentBaseModel == "" {
		agent.AgentBaseModel = s.GetModelDefaultByChainID(req.ChainID)
	}

	if req.VerifiedNFTOwner {
		if req.NFTAddress == "" || req.NFTTokenID == "" || req.NFTOwnerAddress == "" {
			req.VerifiedNFTOwner = false
		} else {
			checked := false
			if strings.Contains(req.NFTTokenID, "i0") {
				// inscription
				checked = s.CheckOwnerInscription(req.NFTDelegateAddress, req.NFTOwnerAddress, req.NFTPublicKey, req.NFTTokenID, req.NFTSignature, req.NFTSignMessage)
			} else {
				checked = s.CheckOwnerNFT721(req.NFTDelegateAddress, req.NFTOwnerAddress, req.NFTAddress, req.NFTTokenID, req.NFTSignature, req.NFTSignMessage)
			}
			agent.VerifiedNftOwner = checked
		}
	}

	if req.TwinTwitterUsernames != "" {
		agent.TwinTwitterUsernames = req.TwinTwitterUsernames
		agent.TwinStatus = models.TwinStatusPending
		listPendingAgents, err := s.dao.FindAgentInfo(
			daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				`twin_twitter_usernames != '' and twin_status = ?`: {models.TwinStatusPending},
				"scan_enabled = ?": {true},
			},
			map[string][]interface{}{},
			[]string{
				"id asc",
			},
			0,
			999999,
		)
		if err != nil {
			return nil, err
		}
		estimateDoneTime := time.Now().Add(time.Duration(len(listPendingAgents)) * 30 * time.Minute)
		agent.EstimateTwinDoneTimestamp = &estimateDoneTime
	}

	// generate address
	switch agent.NetworkID {
	case models.LOCAL_CHAIN_ID:
		{
			// nothing for local
		}
	default:
		{
			ethAddress, err := s.CreateETHAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			agent.ETHAddress = strings.ToLower(ethAddress)
			agent.TronAddress = trxapi.AddrEvmToTron(ethAddress)
			solAddress, err := s.CreateSOLAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			agent.SOLAddress = solAddress
			addressBtc, err := s.CreateBTCAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			agent.TipBtcAddress = addressBtc
			addressEth, err := s.CreateETHAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			agent.TipEthAddress = addressEth
			addressSol, err := s.CreateSOLAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			agent.TipSolAddress = addressSol
		}
	}

	if req.AgentType > 0 {
		agent.AgentType = req.AgentType
		if req.AgentType == models.AgentInfoAgentTypeModel ||
			req.AgentType == models.AgentInfoAgentTypeModelOnline ||
			req.AgentType == models.AgentInfoAgentTypeJs ||
			req.AgentType == models.AgentInfoAgentTypePython ||
			req.AgentType == models.AgentInfoAgentTypeCustomUi ||
			req.AgentType == models.AgentInfoAgentTypeCustomPrompt {
			agent.IsPublic = false
		}
	}

	if req.CreateKnowledgeRequest != nil {
		agent.AgentType = models.AgentInfoAgentTypeKnowledgeBase
	}

	if agent.AgentName == "" && req.CreateKnowledgeRequest != nil {
		agent.AgentName = req.CreateKnowledgeRequest.Name
	}

	if err := s.dao.Create(daos.GetDBMainCtx(ctx), agent); err != nil {
		return nil, errs.NewError(err)
	}

	agentTokenInfo := &models.AgentTokenInfo{}
	agentTokenInfo.AgentInfoID = agent.ID
	agentTokenInfo.NetworkID = agent.TokenNetworkID
	agentTokenInfo.NetworkName = models.GetChainName(agent.TokenNetworkID)

	if err := s.dao.Create(daos.GetDBMainCtx(ctx), agentTokenInfo); err != nil {
		return nil, errs.NewError(err)
	}

	agent.TokenInfoID = agentTokenInfo.ID

	if err := s.dao.Save(daos.GetDBMainCtx(ctx), agent); err != nil {
		return nil, errs.NewError(err)
	}

	if agent.ID > 0 {
		if req.CreateTokenMode == models.CreateTokenModeTypeLinkExisting {
			go s.UpdateTokenPriceInfo(context.Background(), agent.ID)
		}
		go s.AgentCreateMissionDefault(context.Background(), agent.ID)
	}

	if req.CreateKnowledgeRequest != nil {
		ctx := context.Background()
		kbReq := req.CreateKnowledgeRequest
		kbReq.UserAddress = strings.ToLower(address)
		kbReq.DepositAddress = agent.ETHAddress
		kbReq.SolanaDepositAddress = agent.SOLAddress
		kbReq.NetworkID = req.ChainID
		kbReq.AgentInfoId = agent.ID
		kb, err := s.KnowledgeUsecase.CreateKnowledgeBase(ctx, req.CreateKnowledgeRequest)
		if err != nil {
			s.KnowledgeUsecase.SendMessage(ctx, fmt.Sprintf("CreateKnowledgeBase for agent %s (%d) - has error: %s ", agent.AgentName, agent.ID, err.Error()), -1)
			return nil, err
		}

		agent.AgentKBId = kb.ID
		if err := s.dao.Save(daos.GetDBMainCtx(ctx), agent); err != nil {
			return nil, errs.NewError(err)
		}
		oKb, _ := s.KnowledgeUsecase.GetKnowledgeBaseById(ctx, kb.ID)
		agent.KnowledgeBase = oKb
		s.KnowledgeUsecase.SendMessage(ctx, fmt.Sprintf("Create KB Agent DONE  %s (%d)", agent.AgentName, agent.ID), 0)
	}

	if agent.IsVibeAgent() {
		go s.CreateTokenInfoVibe(context.Background(), agent.ID)
	}

	return agent, nil
}

func (s *Service) CheckOwnerInscription(delegate, vault, publicKey, tokenId, signature, signMessage string) bool {
	if !strings.Contains(tokenId, "i0") {
		return false
	}
	if signature != "" {
		temp, err := s.verifyInscription(signature, delegate, publicKey, signMessage)
		if err != nil || !temp {
			return false
		}
	}
	owner := ""
	magicEdenService := magiceden.NewMagicedenService()
	item, err := magicEdenService.GetInscriptionInfo(tokenId)
	if err == nil {
		owner = item.Owner
	} else {
		service := hiro.NewHiroService(s.conf.HiroUrl)
		info, err := service.GetInscriptionInfo(tokenId)
		if err == nil {
			owner = info.Address
		}
	}
	if owner == "" {
		return false
	}
	if strings.ToLower(delegate) == strings.ToLower(vault) && strings.ToLower(owner) == strings.ToLower(delegate) {
		return true
	}
	return false
}

func (s *Service) CheckOwnerNFT721(delegate, vault, contract, tokenId, signature, signMessage string) bool {
	if strings.Contains(tokenId, "i0") {
		return false
	}
	if signature != "" {
		temp, _ := s.verify(signature, delegate, signMessage)
		if !temp {
			return false
		}
	}

	if strings.ToLower(delegate) == strings.ToLower(vault) {
		// return true
	}

	chainID := 1 // default ETH
	hardcodeCollection := s.openseaService.FindHardCodeCollectionByAddress(contract)
	if hardcodeCollection != nil {
		chainID = hardcodeCollection.ChainID
	}

	delegateCash := delegate_cash.NewDelegateCashAPIService(s.conf.DelegateCash.Url, s.conf.DelegateCash.ApiKey)
	checked, err := delegateCash.CheckDelegateForTokenERC721V1(delegate, vault, contract, tokenId, chainID)
	if err != nil {
		fmt.Println(err)
	}
	if !checked {
		checked, err = delegateCash.CheckDelegateForTokenERC721V2(delegate, vault, contract, tokenId, chainID)
		if err != nil {
			fmt.Println(err)
		}
	}
	return checked
}

func (s *Service) verify(signatureHex string, signer string, msgStr string) (bool, error) {
	_, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}
	sig := hexutil.MustDecode(signatureHex)

	msgBytes := []byte(msgStr)
	msgHash := accounts.TextHash(msgBytes)

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	signerHex := recoveredAddr.Hex()
	isVerified := strings.ToLower(signer) == strings.ToLower(signerHex)

	return isVerified, nil
}

func (s *Service) verifyInscription(signatureHex string, signer, publicKey string, msgStr string) (bool, error) {
	fullUrl := "https://api-verify-sig.eternalai.org/api/unisat/verify-sig"
	req := map[string]string{
		"address":   signer,
		"pubKey":    publicKey,
		"message":   msgStr,
		"signature": signatureHex,
	}
	resp, _, i, err := helpers.HttpRequest(fullUrl, "POST",
		map[string]string{
			"Content-Type": "application/json",
		},
		req)
	if err != nil {
		return false, err
	}
	if i != http.StatusOK && i != http.StatusCreated {
		return false, nil
	}
	return string(resp) == "true", nil
}

func (s *Service) AgentUpdateAgentAssistant(ctx context.Context, address string, req *serializers.AssistantsReq) (*models.AgentInfo, error) {
	var agent *models.AgentInfo
	updateMap := make(map[string]interface{})
	updateKb := false
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			var err error
			if req.AgentID != "" {
				agent, err = s.dao.FirstAgentInfo(tx,
					map[string][]interface{}{
						"agent_id = ?": {req.AgentID},
					},
					map[string][]interface{}{
						"KnowledgeBase": {},
					},
					[]string{})
				if err != nil {
					return errs.NewError(err)
				}

				if agent != nil {
					agent, err = s.dao.FirstAgentInfoByID(tx, agent.ID,
						map[string][]interface{}{
							"KnowledgeBase": {},
						}, true,
					)
					if err != nil {
						return errs.NewError(err)
					}
				}
			} else {
				agent, err = s.dao.FirstAgentInfoByID(tx, req.ID, map[string][]interface{}{
					"KnowledgeBase": {},
				}, true)
				if err != nil {
					return errs.NewError(err)
				}
			}

			if agent != nil {
				if !strings.EqualFold(agent.Creator, address) {
					return errs.NewError(errs.ErrInvalidOwner)
				}

				if !(agent.AgentContractID != "" || agent.AgentNftMinted == true) {
					agent.NetworkID = req.ChainID
					agent.NetworkName = models.GetChainName(req.ChainID)
				}

				if req.AgentName != "" {
					agent.AgentName = req.AgentName
				}

				if req.SystemContent != "" {
					agent.MetaData = req.SystemContent
					agent.SystemPrompt = req.SystemContent
				}

				agent.Bio = req.GetAssistantCharacter(req.Bio)
				agent.Lore = req.GetAssistantCharacter(req.Lore)
				agent.Knowledge = req.GetAssistantCharacter(req.Knowledge)
				agent.MessageExamples = req.GetAssistantCharacter(req.MessageExamples)
				agent.PostExamples = req.GetAssistantCharacter(req.PostExamples)
				agent.Topics = req.GetAssistantCharacter(req.Topics)
				agent.Style = req.GetAssistantCharacter(req.Style)
				agent.Adjectives = req.GetAssistantCharacter(req.Adjectives)
				agent.SocialInfo = req.GetAssistantCharacter(req.SocialInfo)

				if req.EnvExample != "" {
					agent.EnvExample = req.EnvExample
				}

				if req.RequiredEnv != nil {
					agent.RequiredEnv = *req.RequiredEnv
				}

				if req.SourceUrl != "" {
					agent.SourceUrl = req.SourceUrl
				}
				if req.AuthenUrl != "" {
					agent.AuthenUrl = req.AuthenUrl
				}
				if req.DependAgents != "" {
					agent.DependAgents = req.DependAgents
				}
				if req.RequiredWallet != nil {
					agent.RequiredWallet = *req.RequiredWallet
				}
				if req.IsOnchain != nil {
					agent.IsOnchain = *req.IsOnchain
				}
				if req.IsStreaming != nil {
					agent.IsStreaming = *req.IsStreaming
				}
				if req.IsCustomUi != nil {
					agent.IsCustomUi = *req.IsCustomUi
				}
				if req.RequiredInfo != nil {
					agent.RequiredInfo = *req.RequiredInfo
				}
				if req.InferFee != nil {
					inferFee := numeric.NewBigFloatFromString(*req.InferFee)
					agent.InferFee = inferFee
				}
				if req.MinFeeToUse.Cmp(big.NewFloat(0)) > 0 {
					agent.MinFeeToUse = req.MinFeeToUse
				}
				if req.Worker != "" {
					agent.Worker = req.Worker
				}
				if req.TokenImageUrl != "" {
					agent.TokenImageUrl = req.TokenImageUrl
				}
				if req.DisplayName != "" {
					agent.DisplayName = req.DisplayName
				}
				if req.ShortDescription != "" {
					agent.ShortDescription = req.ShortDescription
				}
				if req.CategoryID != 0 {
					agent.AgentCategoryID = req.CategoryID
				}
				if req.IsForceUpdate != nil {
					agent.IsForceUpdate = *req.IsForceUpdate
				}

				if req.Author != "" {
					agent.Author = req.Author
				}

				if agent.TokenStatus == "" && agent.TokenAddress == "" {
					if req.TokenChainId != "" {
						tokenChainId, _ := strconv.ParseUint(req.TokenChainId, 0, 64)
						if !(tokenChainId == models.POLYGON_CHAIN_ID || tokenChainId == models.ZKSYNC_CHAIN_ID) {
							agent.TokenNetworkID = tokenChainId
							if req.CreateTokenMode != "" {
								agent.TokenMode = string(req.CreateTokenMode)
							}
							tokenName := req.TokenName
							if req.TokenName == "" {
								tokenName = req.Ticker
							}
							agent.TokenSymbol = req.Ticker
							agent.TokenName = tokenName
							agent.TokenDesc = req.TokenDesc

							if agent.TokenMode == string(models.TokenSetupEnumAutoCreate) && (agent.AgentNftMinted || (agent.AgentType == models.AgentInfoAgentTypeKnowledgeBase && agent.Status == models.AssistantStatusReady)) {
								agent.TokenStatus = "pending"
							}
						} else {
							agent.TokenNetworkID = models.GENERTAL_NETWORK_ID
							agent.TokenMode = string(models.TokenSetupEnumNoToken)
						}
					}
				}

				err := s.dao.Save(tx, agent)
				if err != nil {
					return errs.NewError(err)
				}

				if req.CreateKnowledgeRequest != nil && agent.AgentType == models.AgentInfoAgentTypeKnowledgeBase {
					updateKb = true
					kbReq := req.CreateKnowledgeRequest
					if kbReq.Name != "" {
						updateMap["name"] = kbReq.Name
						agent.AgentName = kbReq.Name
						if err := s.dao.Save(tx, agent); err != nil {
							return errs.NewError(err)
						}
					}

					if kbReq.Description != "" {
						updateMap["description"] = kbReq.Description
					}

					if kbReq.NetworkID != 0 {
						updateMap["network_id"] = kbReq.NetworkID
					}

					if kbReq.ThumbnailUrl != "" {
						updateMap["thumbnail_url"] = kbReq.ThumbnailUrl
					}

					if kbReq.DomainUrl != "" {
						updateMap["domain_url"] = kbReq.DomainUrl
					}
				}

				go s.AgentCreateMissionDefault(context.Background(), agent.ID)
			}
			return nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if updateKb {
		i, err := s.KnowledgeUsecase.GetKnowledgeBaseById(ctx, agent.AgentKBId)
		if err != nil {
			return nil, err
		}
		if len(req.CreateKnowledgeRequest.Files) > 0 {
			updatedListFile, err := s.KnowledgeUsecase.UpdateListKnowledgeBaseFile(ctx, agent.AgentKBId, req.CreateKnowledgeRequest.Files)
			if err != nil {
				return nil, errs.NewError(err)
			}
			if updatedListFile {
				i.ChargeMore = i.CalcChargeMore()
				if i.ChargeMore != 0 {
					i.Fee = i.Fee + i.ChargeMore
					updateMap["fee"] = i.Fee
					updateMap["charge_more"] = i.ChargeMore
					updateMap["status"] = models.KnowledgeBaseStatusWaitingPayment
				}
			}
		}
		if err := s.KnowledgeUsecase.UpdateKnowledgeBaseById(ctx, agent.AgentKBId, updateMap); err != nil {
			return nil, err
		}
		agent.KnowledgeBase = i
	}

	agentInfoKbs := []*models.AgentInfoKnowledgeBase{}
	for _, id := range req.KbIds {
		i := &models.AgentInfoKnowledgeBase{
			AgentInfoId:     agent.ID,
			KnowledgeBaseId: id,
		}
		agentInfoKbs = append(agentInfoKbs, i)
	}
	_, err = s.KnowledgeUsecase.CreateAgentInfoKnowledgeBase(ctx, agentInfoKbs, agent.ID)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (s *Service) UpdateAgentInfoInContract(ctx context.Context, address string, req *serializers.UpdateAgentAssistantInContractRequest) (*models.AgentInfo, error) {
	var agent *models.AgentInfo
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			var err error
			if req.AgentID != "" {
				agent, err = s.dao.FirstAgentInfo(tx,
					map[string][]interface{}{
						"agent_id = ?": {req.AgentID},
					},
					map[string][]interface{}{}, []string{})
				if err != nil {
					return errs.NewError(err)
				}

				if agent != nil {
					agent, err = s.dao.FirstAgentInfoByID(tx, agent.ID, map[string][]interface{}{}, true)
					if err != nil {
						return errs.NewError(err)
					}
				}
			} else {
				agent, err = s.dao.FirstAgentInfoByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
			}

			if agent != nil {
				if !strings.EqualFold(agent.Creator, address) {
					return errs.NewError(errs.ErrInvalidOwner)
				}

				systemPromptBytes, _, err := lighthouse.DownloadDataSimple(req.HashSystemPrompt)
				if err != nil {
					return errs.NewError(err)
				}
				dataBytes, _, err := lighthouse.DownloadDataSimple(req.HashName)
				if err != nil {
					return errs.NewError(err)
				}
				agentUriData := models.AgentUriData{}
				err = json.Unmarshal(dataBytes, &agentUriData)
				if err != nil {
					return errs.NewError(err)
				}

				agent.SystemPrompt = string(systemPromptBytes)
				agent.AgentName = agentUriData.Name
				agent.Uri = req.HashName
				err = s.dao.Save(tx, agent)
				if err != nil {
					return errs.NewError(err)
				}

				// _, err = s.ExecuteUpdateAgentInfoInContract(ctx, agent, req)
				// if err != nil {
				// 	return errs.NewError(err)
				// }
			}
			return nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	return agent, nil
}

func (s *Service) UploadDataToLightHouse(ctx context.Context, address string, req *serializers.DataUploadToLightHouse) (string, error) {
	hash, err := lighthouse.UploadData(s.conf.Lighthouse.Apikey, address, []byte(req.Content))
	if err != nil {
		return "", errs.NewError(err)
	}
	return fmt.Sprintf("ipfs://%v", hash), nil
}

func (s *Service) GenerateTokenInfoFromSystemPrompt(ctx context.Context, tokenName, sysPrompt string) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{}
	sysPrompt = strings.ReplaceAll(sysPrompt, "@CryptoEternalAI", "")
	promptGenerateToken := fmt.Sprintf(`
						I want to generate my token base on this info
						'%s'

						token-name (generate if not provided, make sure it not empty)
						token-symbol (generate if not provided, make sure it not empty)
						token-story (generate if not provided, make sure it not empty)

						Please return in string in json format including token-name, token-symbol, token-story, just only json without explanation  and token name limit with 15 characters
					`, sysPrompt)
	// aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
	// if err != nil {
	// 	return nil, errs.NewError(err)
	// }
	message := []openai.ChatCompletionMessage{
		openai.ChatCompletionMessage{
			Role:    "system",
			Content: "You are a helpful assistant",
		},
		openai.ChatCompletionMessage{
			Role:    "user",
			Content: promptGenerateToken,
		},
	}
	aiStr, err := s.openais["Agent"].CallStreamDirectlyEternalLLMV2(ctx, message, "Qwen/QwQ-32B", s.conf.KnowledgeBaseConfig.DirectServiceUrl, nil)
	if err != nil {
		return nil, errs.NewError(err)
	}
	fmt.Println(aiStr)
	if aiStr != "" {
		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
		tokenSymbol := ""
		tokenDesc := ""
		imageUrl := ""
		if mapInfo != nil {

			if v, ok := mapInfo["token-symbol"]; ok {
				tokenSymbol = fmt.Sprintf(`%v`, v)
			}

			if v, ok := mapInfo["token-story"]; ok {
				tokenDesc = fmt.Sprintf(`%v`, v)
			}
			if tokenName == "" {
				if v, ok := mapInfo["token-name"]; ok {
					tokenName = fmt.Sprintf(`%v`, v)
				}
				if tokenName == "" {
					tokenName = tokenSymbol
				}
			}
			imageUrl, _ = s.GetGifImageUrlFromTokenInfo(tokenSymbol, tokenName, tokenDesc)
		}
		info = &models.TweetParseInfo{
			TokenSymbol:   tokenSymbol,
			TokenDesc:     tokenDesc,
			TokenImageUrl: imageUrl,
			TokenName:     tokenName,
		}
	}

	return info, nil
}

func (s *Service) JobMigrateTronAddress(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobMigrateTronAddress",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"tron_address = '' or tron_address is null": {},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				1000,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = daos.GetDBMainCtx(ctx).
					Model(agent).
					Updates(
						map[string]interface{}{
							"tron_address": trxapi.AddrEvmToTron(agent.ETHAddress),
						},
					).
					Error
				if err != nil {
					return errs.NewError(err)
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

// //
func (s *Service) AgentCreateAgentStudio(ctx context.Context, address, graphData string) ([]*models.AgentInfo, error) {
	var reqs *models.AgentStudioGraphData
	json.Unmarshal([]byte(graphData), &reqs)

	if reqs != nil && len(reqs.Data) <= 0 {
		return nil, errs.NewError(errs.ErrBadRequest)
	}

	listAgent := []*models.AgentInfo{}
	for _, req := range reqs.Data {
		if req.Idx == "agent_new" {
			agent := &models.AgentInfo{
				Version:     "2",
				AgentType:   models.AgentInfoAgentTypeReasoning,
				AgentID:     helpers.RandomBigInt(12).Text(16),
				Status:      models.AssistantStatusPending,
				AgentName:   fmt.Sprintf("%v", req.Data["agentName"]),
				Creator:     strings.ToLower(address),
				ScanEnabled: false,
				GraphData:   graphData,
			}

			listMission := []*serializers.AgentSnapshotMissionInfo{}
			for _, item := range req.Children {
				switch item.CategoryIdx {
				case "personality":
					{
						switch item.Idx {
						case "personality_customize":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
							}
						case "personality_nft":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
								var nftInfo map[string]interface{}
								somebytes, _ := json.Marshal(item.Data["selectedNFT"])

								err1 := json.Unmarshal(somebytes, &nftInfo)
								if err1 == nil {
									agent.VerifiedNftOwner = false
									agent.NftAddress = helpers.GetStringValueFromMap(nftInfo, "token_address")
									agent.NftTokenID = helpers.GetStringValueFromMap(nftInfo, "token_id")
									agent.NftTokenImage = helpers.GetStringValueFromMap(nftInfo, "token_uri")
									agent.NftOwnerAddress = helpers.GetStringValueFromMap(nftInfo, "owner_of")
								}
							}
						case "personality_ordinals":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")

								var nftInfo map[string]interface{}
								somebytes, _ := json.Marshal(item.Data["selectedNFT"])

								err1 := json.Unmarshal(somebytes, &nftInfo)
								if err1 == nil {
									agent.VerifiedNftOwner = false
									agent.NftAddress = helpers.GetStringValueFromMap(nftInfo, "token_address")
									agent.NftTokenID = helpers.GetStringValueFromMap(nftInfo, "token_id")
									agent.NftTokenImage = helpers.GetStringValueFromMap(nftInfo, "token_uri")
									agent.NftOwnerAddress = helpers.GetStringValueFromMap(nftInfo, "owner_of")
								}
							}
						case "personality_token":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
							}
						case "personality_genomics":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
								var twitterInfo []map[string]interface{}
								somebytes, _ := json.Marshal(item.Data["twitterInfos"])

								json.Unmarshal(somebytes, &twitterInfo)
								if len(twitterInfo) > 0 {
									twinTwitterUsernames := []string{}
									for _, tw := range twitterInfo {
										username := helpers.GetStringValueFromMap(tw, "username")
										if username != "" {
											twinTwitterUsernames = append(twinTwitterUsernames, username)
										}
									}
									if len(twinTwitterUsernames) > 0 {
										agent.TwinTwitterUsernames = strings.Join(twinTwitterUsernames, ",")
										agent.TwinStatus = models.TwinStatusPending
										listPendingAgents, err := s.dao.FindAgentInfo(
											daos.GetDBMainCtx(ctx),
											map[string][]interface{}{
												`twin_twitter_usernames != '' and twin_status = ?`: {models.TwinStatusPending},
												"scan_enabled = ?": {true},
											},
											map[string][]interface{}{},
											[]string{
												"id asc",
											},
											0,
											999999,
										)
										if err != nil {
											return nil, errs.NewError(err)
										}
										estimateDoneTime := time.Now().Add(time.Duration(len(listPendingAgents)) * 30 * time.Minute)
										agent.EstimateTwinDoneTimestamp = &estimateDoneTime
									}
								}
							}
						case "personality_knowledge":
							{
								agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
								agent.AgentType = models.AgentInfoAgentTypeKnowledgeBase
							}
						}
					}
				case "blockchain":
					{
						chainName := helpers.GetStringValueFromMap(item.Data, "chainId")
						agent.NetworkID = models.GetChainID(chainName)
						agent.NetworkName = models.GetChainName(agent.NetworkID)
					}
				case "decentralized_inference":
					{
						agent.AgentBaseModel = fmt.Sprintf("%v", item.Data["decentralizeId"])
					}
				case "ai_framework":
					{
						switch item.Idx {
						case "ai_framework_eternal_ai":
							{
								agent.AgentType = models.AgentInfoAgentTypeReasoning
							}
						case "ai_framework_eliza":
							{
								agent.AgentType = models.AgentInfoAgentTypeEliza
								agent.ConfigData = helpers.GetStringValueFromMap(item.Data, "config")
							}
						case "ai_framework_zerepy":
							{
								agent.AgentType = models.AgentInfoAgentTypeZerepy
								agent.ConfigData = helpers.GetStringValueFromMap(item.Data, "config")
							}
						}
					}
				case "token":
					{
						agent.TokenNetworkID = helpers.GetTokenIDFromMap(item.Data)
						if agent.TokenNetworkID > 0 {
							agent.TokenMode = string(models.CreateTokenModeTypeAutoCreate)
						} else {
							agent.TokenMode = string(models.CreateTokenModeTypeNoToken)
						}
					}
				case "mission_on_x":
					{
						frequency := helpers.GetFrequencyFromMap(item.Data)

						userPrompt := helpers.GetStringValueFromMap(item.Data, "details")
						agentModels := helpers.GetStringValueFromMap(item.Data, "modelName")

						switch item.Idx {
						case "mission_on_x_post":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypePost,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						case "mission_on_x_post_news":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypePostSearchV2,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						case "mission_on_x_reply":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypeReplyMentions,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						case "mission_on_x_engage":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypeReplyNonMentions,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						case "mission_on_x_follow":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypeFollow,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						}
					}
				case "mission_on_farcaster":
					{
						frequency := helpers.GetFrequencyFromMap(item.Data)

						userPrompt := helpers.GetStringValueFromMap(item.Data, "details")
						agentModels := helpers.GetStringValueFromMap(item.Data, "modelName")
						switch item.Idx {
						case "mission_on_farcaster_post":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypePostFarcaster,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						case "mission_on_farcaster_reply":
							{
								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:        models.ToolsetTypeReplyMentionsFarcaster,
									UserPrompt:     userPrompt,
									Interval:       frequency,
									AgentBaseModel: agentModels,
								})
							}
						}
					}
				case "mission_on_defi":
					{
						frequency := helpers.GetFrequencyFromMap(item.Data)

						switch item.Idx {
						case "mission_on_defi_trade_analytics":
							{
								tokens := helpers.GetStringValueFromMap(item.Data, "token")
								toolList := s.conf.ToolLists.TradeAnalytic
								if tokens == "" {
									return nil, errs.NewError(errs.ErrBadRequest)
								}
								userPrompt := fmt.Sprintf(`Conduct a technical analysis of $%s price data. Based on your findings, provide a recommended buy price and sell price to maximize potential returns.`, tokens)
								toolList = strings.ReplaceAll(toolList, "{api_key}", s.conf.InternalApiKey)
								toolList = strings.ReplaceAll(toolList, "{token_symbol}", tokens)

								listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
									ToolSet:    models.ToolsetTypeTradeAnalyticsOnTwitter,
									Tokens:     tokens,
									Interval:   frequency,
									UserPrompt: userPrompt,
									ToolList:   toolList,
								})
							}
						}
					}
				default:
					{
					}
				}
			}

			if agent.NetworkID == 0 || agent.AgentBaseModel == "" {
				return nil, errs.NewError(errs.ErrBadRequest)
			}

			// gen token
			tokenInfo, _ := s.GenerateTokenInfoFromSystemPrompt(ctx, agent.AgentName, agent.SystemPrompt)
			if tokenInfo != nil && tokenInfo.TokenSymbol != "" {
				agent.TokenSymbol = tokenInfo.TokenSymbol
				agent.TokenName = agent.AgentName
				agent.TokenDesc = tokenInfo.TokenDesc
				agent.TokenImageUrl = tokenInfo.TokenImageUrl
			}

			// // generate address
			{
				ethAddress, err := s.CreateETHAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				agent.ETHAddress = strings.ToLower(ethAddress)
				agent.TronAddress = trxapi.AddrEvmToTron(ethAddress)

				solAddress, err := s.CreateSOLAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				agent.SOLAddress = solAddress

				addressBtc, err := s.CreateBTCAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				agent.TipBtcAddress = addressBtc

				addressEth, err := s.CreateETHAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				agent.TipEthAddress = addressEth

				addressSol, err := s.CreateSOLAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				agent.TipSolAddress = addressSol
			}

			if err := s.dao.Create(daos.GetDBMainCtx(ctx), agent); err != nil {
				return nil, errs.NewError(err)
			}
			agentTokenInfo := &models.AgentTokenInfo{}
			agentTokenInfo.AgentInfoID = agent.ID
			agentTokenInfo.NetworkID = agent.TokenNetworkID
			agentTokenInfo.NetworkName = models.GetChainName(agent.TokenNetworkID)

			if err := s.dao.Create(daos.GetDBMainCtx(ctx), agentTokenInfo); err != nil {
				return nil, errs.NewError(err)
			}

			agent.TokenInfoID = agentTokenInfo.ID

			if err := s.dao.Save(daos.GetDBMainCtx(ctx), agent); err != nil {
				return nil, errs.NewError(err)
			}

			if len(listMission) > 0 {
				var err error
				agent, err = s.CreateUpdateAgentSnapshotMission(ctx, agent.AgentID, "", listMission)
				if err != nil {
					return nil, errs.NewError(errs.ErrBadRequest)
				}
			}
			listAgent = append(listAgent, agent)
		}
	}

	return listAgent, nil
}

func (s *Service) AgentUpdateAgentStudio(ctx context.Context, address, agentID, graphData string) (*models.AgentInfo, error) {
	var agent *models.AgentInfo
	listMission := []*serializers.AgentSnapshotMissionInfo{}
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			var err error
			agent, err = s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{})
			if err != nil {
				return errs.NewError(err)
			}

			if agent != nil {
				if !strings.EqualFold(agent.Creator, address) {
					return errs.NewError(errs.ErrInvalidOwner)
				}
				agent, _ = s.dao.FirstAgentInfoByID(tx, agent.ID, map[string][]interface{}{}, true)

				var reqs *models.AgentStudioGraphData
				json.Unmarshal([]byte(graphData), &reqs)

				if reqs != nil && len(reqs.Data) <= 0 {
					return errs.NewError(errs.ErrBadRequest)
				}
				req := reqs.Data[0]

				if req.Idx == "agent_new" {
					for _, item := range req.Children {
						switch item.CategoryIdx {
						case "personality":
							{
								switch item.Idx {
								case "personality_customize":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								case "personality_nft":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								case "personality_ordinals":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								case "personality_token":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								case "personality_genomics":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								case "personality_knowledge":
									{
										agent.SystemPrompt = helpers.GetStringValueFromMap(item.Data, "personality")
									}
								}
							}
						case "blockchain":
							{
								if !(agent.AgentContractID != "" || agent.AgentNftMinted == true) {
									chainName := helpers.GetStringValueFromMap(item.Data, "chainId")
									agent.NetworkID = models.GetChainID(chainName)
									agent.NetworkName = models.GetChainName(agent.NetworkID)
								}
							}
						case "decentralized_inference":
							{
								agent.AgentBaseModel = fmt.Sprintf("%v", item.Data["decentralizeId"])
							}
						case "token":
							{
								if agent.TokenStatus == "" && agent.TokenAddress == "" {
									agent.TokenNetworkID = helpers.GetTokenIDFromMap(item.Data)
									if agent.TokenNetworkID > 0 {
										agent.TokenMode = string(models.CreateTokenModeTypeAutoCreate)
									} else {
										agent.TokenMode = string(models.CreateTokenModeTypeNoToken)
									}

									if agent.TokenMode == string(models.CreateTokenModeTypeAutoCreate) && (agent.AgentNftMinted || (agent.AgentType == models.AgentInfoAgentTypeKnowledgeBase && agent.Status == models.AssistantStatusReady)) {
										agent.TokenStatus = "pending"
									}
								}
							}
						case "ai_framework":
							{
								switch item.Idx {
								case "ai_framework_eternal_ai":
									{
										agent.AgentType = models.AgentInfoAgentTypeReasoning
									}
								case "ai_framework_eliza":
									{
										agent.AgentType = models.AgentInfoAgentTypeEliza
										agent.ConfigData = helpers.GetStringValueFromMap(item.Data, "config")
									}
								case "ai_framework_zerepy":
									{
										agent.AgentType = models.AgentInfoAgentTypeZerepy
										agent.ConfigData = helpers.GetStringValueFromMap(item.Data, "config")
									}
								}
							}
						case "mission_on_x":
							{
								frequency := helpers.GetFrequencyFromMap(item.Data)
								userPrompt := helpers.GetStringValueFromMap(item.Data, "details")
								agentModels := helpers.GetStringValueFromMap(item.Data, "modelName")
								switch item.Idx {
								case "mission_on_x_post":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypePost,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								case "mission_on_x_post_news":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypePostSearchV2,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								case "mission_on_x_reply":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypeReplyMentions,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								case "mission_on_x_engage":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypeReplyNonMentions,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								case "mission_on_x_follow":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypeFollow,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								}
							}
						case "mission_on_farcaster":
							{
								frequency := helpers.GetFrequencyFromMap(item.Data)
								userPrompt := helpers.GetStringValueFromMap(item.Data, "details")
								agentModels := helpers.GetStringValueFromMap(item.Data, "modelName")

								switch item.Idx {
								case "mission_on_farcaster_post":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypePostFarcaster,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								case "mission_on_farcaster_reply":
									{
										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:        models.ToolsetTypeReplyMentionsFarcaster,
											UserPrompt:     userPrompt,
											Interval:       frequency,
											AgentBaseModel: agentModels,
										})
									}
								}
							}
						case "mission_on_defi":
							{
								frequency := helpers.GetFrequencyFromMap(item.Data)
								switch item.Idx {
								case "mission_on_defi_trade_analytics":
									{
										tokens := helpers.GetStringValueFromMap(item.Data, "token")
										toolList := s.conf.ToolLists.TradeAnalytic
										if tokens == "" {
											return errs.NewError(errs.ErrBadRequest)
										}
										userPrompt := fmt.Sprintf(`Conduct a technical analysis of $%s price data. Based on your findings, provide a recommended buy price and sell price to maximize potential returns.`, tokens)
										toolList = strings.ReplaceAll(toolList, "{api_key}", s.conf.InternalApiKey)
										toolList = strings.ReplaceAll(toolList, "{token_symbol}", tokens)

										listMission = append(listMission, &serializers.AgentSnapshotMissionInfo{
											ToolSet:    models.ToolsetTypeTradeAnalyticsOnTwitter,
											Tokens:     tokens,
											Interval:   frequency,
											UserPrompt: userPrompt,
											ToolList:   toolList,
										})
									}
								}
							}
						default:
							{
							}
						}
					}
					agent.GraphData = graphData
					err := s.dao.Save(tx, agent)
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if len(listMission) > 0 && agent != nil {
		var err error
		agent, err = s.CreateUpdateAgentSnapshotMission(ctx, agent.AgentID, "", listMission)
		if err != nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}

	return agent, nil
}

func (s *Service) AgentCreateAgentAssistantForLocal(ctx context.Context, req *serializers.AssistantsReq) (*models.AgentInfo, error) {
	if !s.conf.ExistsedConfigKey(models.LOCAL_CHAIN_ID, "network_id") {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if req.SystemContent == "" {
		req.SystemContent = "default"
	}
	agent := &models.AgentInfo{
		Version:              "2",
		AgentType:            models.AgentInfoAgentTypeReasoning,
		AgentID:              helpers.RandomBigInt(12).Text(16),
		Status:               models.AssistantStatusReady,
		NetworkID:            models.LOCAL_CHAIN_ID,
		NetworkName:          models.GetChainName(req.ChainID),
		AgentName:            req.AgentName,
		Creator:              req.Creator,
		AgentContractAddress: req.AgentContractAddress,
		AgentContractID:      req.AgentContractID,
		SystemPrompt:         req.SystemContent,
		AgentBaseModel:       req.AgentBaseModel,
		ScanEnabled:          true,
	}
	if err := s.dao.Create(daos.GetDBMainCtx(ctx), agent); err != nil {
		return nil, errs.NewError(err)
	}
	return agent, nil
}

func (s *Service) GetAgentChainFees(ctx context.Context) (map[string]interface{}, error) {
	res, err := s.dao.FindAgentChainFee(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{},
		map[string][]interface{}{},
		[]string{"id desc"},
		0,
		999999,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	chainFeeMap := map[string]interface{}{}
	for _, v := range res {
		chainFeeMap[strconv.Itoa(int(v.NetworkID))] = map[string]interface{}{
			"network_id":       v.NetworkID,
			"mint_fee":         numeric.BigFloat2Text(&v.MintFee.Float),
			"post_fee":         numeric.BigFloat2Text(&v.InferFee.Float),
			"token_fee":        numeric.BigFloat2Text(&v.TokenFee.Float),
			"agent_deploy_fee": numeric.BigFloat2Text(&v.AgentDeployFee.Float),
		}
	}
	return chainFeeMap, nil
}

func (s *Service) MarkInstalledUtilityAgent(ctx context.Context, address string, req *serializers.AgentActionReq) ([]uint, error) {
	resp := []uint{}
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if len(req.ContractAddress) > 0 {
				for _, contractAddress := range req.ContractAddress {
					agentInfo, err := s.dao.FirstAgentInfo(
						tx,
						map[string][]any{
							"agent_contract_address = ?": []any{contractAddress},
						},
						map[string][]any{},
						[]string{"id desc"},
					)
					if err != nil {
						return errs.NewError(err)
					}
					if agentInfo == nil {
						return errs.NewError(errs.ErrAgentNotFound)
					}

					if req.Action == "uninstall" {
						err := tx.Where("agent_info_id = ? and address = ?", agentInfo.ID, strings.ToLower(address)).
							Unscoped().
							Delete(&models.AgentUtilityInstall{}).Error
						if err != nil {
							return errs.NewError(err)
						}
					} else {
						inst := &models.AgentUtilityInstall{
							Address:     strings.ToLower(address),
							AgentInfoID: agentInfo.ID,
						}
						err = s.dao.Create(tx, inst)
						if err == nil {
							err = tx.Model(agentInfo).Update("installed_count", gorm.Expr("installed_count + 1")).Error
							if err != nil {
								return errs.NewError(err)
							}
						}

					}
					resp = append(resp, agentInfo.ID)
				}
			} else {
				if req.Action == "uninstall" {
					for _, agentID := range req.Ids {
						err := tx.Where("agent_info_id = ? and address = ?", agentID, strings.ToLower(address)).
							Unscoped().
							Delete(&models.AgentUtilityInstall{}).Error
						if err != nil {
							return errs.NewError(err)
						}
						resp = append(resp, agentID)
					}
				} else {
					for _, agentID := range req.Ids {
						inst := &models.AgentUtilityInstall{
							Address:     strings.ToLower(address),
							AgentInfoID: agentID,
						}
						err := s.dao.Create(tx, inst)
						if err == nil {
							err = tx.Model(&models.AgentInfo{}).Where("id = ?", agentID).Update("installed_count", gorm.Expr("installed_count + 1")).Error
							if err != nil {
								return errs.NewError(err)
							}
						}
						resp = append(resp, agentID)

					}
				}
			}
			return nil
		},
	)

	if err != nil {
		return resp, errs.NewError(err)
	}

	return resp, nil
}

func (s *Service) MarkRecentChatUtilityAgent(ctx context.Context, address string, req *serializers.AgentActionReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			for _, agentID := range req.Ids {
				now := helpers.TimeNow()
				inst := &models.AgentUtilityRecentChat{
					Address:     strings.ToLower(address),
					AgentInfoID: agentID,
				}
				inst.UpdatedAt = *now
				err := s.dao.Save(tx, inst)
				if err != nil {
					_ = tx.Model(&models.AgentUtilityRecentChat{}).Where("address = ? and agent_info_id = ?", strings.ToLower(address), agentID).Update("updated_at", now).Error
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (s *Service) MarkPromptCountUtilityAgent(ctx context.Context, address string, agentID uint) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfoByID(
				tx, agentID,
				map[string][]any{},
				true,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo == nil {
				return errs.NewError(errs.ErrAgentNotFound)
			}

			if agentInfo.PromptCalls > 0 {
				err = tx.Model(agentInfo).
					UpdateColumn("prompt_calls", gorm.Expr("prompt_calls + 1")).Error
			} else {
				err = tx.Model(agentInfo).
					UpdateColumn("prompt_calls", 1).Error
			}
			if err != nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (s *Service) LikeAgent(ctx context.Context, address string, agentID uint) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentReactionHistory, err := s.dao.FirstAgentReactionHistory(
				tx,
				map[string][]any{
					"agent_info_id = ?": []any{agentID},
					"user_address = ?":  []any{strings.ToLower(address)},
				},
				map[string][]any{},
				[]string{"id desc"},
			)
			if err != nil {
				return errs.NewError(err)
			}

			agentInfo, err := s.dao.FirstAgentInfoByID(
				tx, agentID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo == nil {
				return errs.NewError(errs.ErrAgentNotFound)
			}

			if agentReactionHistory == nil {
				agentReactionHistory = &models.AgentReactionHistory{
					AgentInfoID: agentID,
					UserAddress: strings.ToLower(address),
					Reaction:    "like",
				}
				err = s.dao.Create(tx, agentReactionHistory)
				if err != nil {
					return errs.NewError(err)
				}

				err = tx.Model(agentInfo).Update("likes", gorm.Expr("likes + 1")).Error
				if err != nil {
					return errs.NewError(err)
				}
			} else {
				//unlike
				err = tx.Unscoped().Delete(agentReactionHistory).Error
				if err != nil {
					return errs.NewError(err)
				}

				err = tx.Model(agentInfo).Update("likes", gorm.Expr("likes - 1")).Error
				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (s *Service) CheckAgentLiked(ctx context.Context, address string, agentID uint) (bool, error) {
	agentReactionHistory, err := s.dao.FirstAgentReactionHistory(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"agent_info_id = ?": []any{agentID},
			"user_address = ?":  []any{strings.ToLower(address)},
		},
		map[string][]any{},
		[]string{"id desc"},
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	return agentReactionHistory != nil, nil
}

func (s *Service) PublicAgent(ctx context.Context, address string, agentID uint) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfoByID(
				tx, agentID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo == nil {
				return errs.NewError(errs.ErrAgentNotFound)
			}

			if strings.ToLower(address) != agentInfo.Creator {
				return errs.NewError(errs.ErrUnAuthorization)
			}

			err = tx.Model(agentInfo).Update("is_public", !agentInfo.IsPublic).Error
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (s *Service) AgentComment(ctx context.Context, address string, agentID uint, req *serializers.AgentCommentReq) (*models.AgentUserComment, error) {
	var agentUserComment *models.AgentUserComment
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfoByID(
				tx, agentID,
				map[string][]any{},
				true,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo == nil {
				return errs.NewError(errs.ErrAgentNotFound)
			}

			agentUserComment = &models.AgentUserComment{
				AgentInfoID: agentID,
				UserAddress: strings.ToLower(address),
				Comment:     req.Comment,
				Rating:      req.Rating,
			}

			err = s.dao.Create(tx, agentUserComment)
			if err != nil {
				return errs.NewError(err)
			}

			switch req.Rating {
			case 1:
				err = tx.Model(agentInfo).UpdateColumn("num_of_one_star", gorm.Expr("num_of_one_star + 1")).
					UpdateColumn("num_of_rating", gorm.Expr("num_of_rating + 1")).Error
			case 2:
				err = tx.Model(agentInfo).UpdateColumn("num_of_two_star", gorm.Expr("num_of_two_star + 1")).
					UpdateColumn("num_of_rating", gorm.Expr("num_of_rating + 1")).Error
			case 3:
				err = tx.Model(agentInfo).UpdateColumn("num_of_three_star", gorm.Expr("num_of_three_star + 1")).
					UpdateColumn("num_of_rating", gorm.Expr("num_of_rating + 1")).Error
			case 4:
				err = tx.Model(agentInfo).UpdateColumn("num_of_four_star", gorm.Expr("num_of_four_star + 1")).
					UpdateColumn("num_of_rating", gorm.Expr("num_of_rating + 1")).Error
			case 5:
				err = tx.Model(agentInfo).UpdateColumn("num_of_five_star", gorm.Expr("num_of_five_star + 1")).
					UpdateColumn("num_of_rating", gorm.Expr("num_of_rating + 1")).Error
			}
			if err != nil {
				return errs.NewError(err)
			}

			agentInfo.Rating = (agentInfo.Rating*float64(agentInfo.NumOfRating) + req.Rating) / float64(agentInfo.NumOfRating+1)
			err = tx.Model(agentInfo).Update("rating", agentInfo.Rating).Error
			if err != nil {
				return errs.NewError(err)
			}

			return nil
		},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return agentUserComment, nil
}

func (s *Service) GetListAgentComment(ctx context.Context, agentID uint, page, limit int) ([]*models.AgentUserComment, error) {
	agentUserComments, err := s.dao.AgentUserComment4Page(
		daos.GetDBMainCtx(ctx),
		map[string][]any{"agent_info_id = ?": {agentID}},
		map[string][]any{},
		[]string{"created_at desc"},
		page,
		limit,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return agentUserComments, nil
}
