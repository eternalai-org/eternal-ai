package knowledge

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/repository"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/eth"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/zkclient"
	"github.com/ethereum/go-ethereum/common"
	resty "github.com/go-resty/resty/v2"
	"github.com/mymmrac/telego"

	"go.uber.org/zap"
)

var categoryNameTracer string = "knowledge_usecase_tracer"

type options func(*knowledgeUsecase)

type knowledgeUsecase struct {
	knowledgeBaseRepo          repository.KnowledgeBaseRepo
	knowledgeBaseFileRepo      repository.KnowledgeBaseFileRepo
	agentInfoKnowledgeBaseRepo repository.IAgentInfoKnowledgeBaseRepo
	agentInfoRepo              repository.IAgentInfoRepo

	secretKey       string
	networks        map[string]map[string]string
	ethApiMap       map[uint64]*ethapi.Client
	trxApi          *trxapi.Client
	ragApi          string
	lighthouseKey   string
	webhookUrl      string
	notiBot         *telego.Bot
	notiActChanId   int64
	notiErrorChanId int64
	conf            *configs.Config
}

func WithRepos(
	knowledgeBaseRepo repository.KnowledgeBaseRepo,
	knowledgeBaseFileRepo repository.KnowledgeBaseFileRepo,
	agentInfoKnowledgeBaseRepo repository.IAgentInfoKnowledgeBaseRepo,
	agentInfoRepo repository.IAgentInfoRepo,
) options {
	return func(uc *knowledgeUsecase) {
		uc.knowledgeBaseRepo = knowledgeBaseRepo
		uc.knowledgeBaseFileRepo = knowledgeBaseFileRepo
		uc.agentInfoKnowledgeBaseRepo = agentInfoKnowledgeBaseRepo
		uc.agentInfoRepo = agentInfoRepo
	}
}

func WithSecretKey(secretKey string) options {
	return func(uc *knowledgeUsecase) {
		uc.secretKey = secretKey
	}
}

func WithConfig(conf *configs.Config) options {
	return func(uc *knowledgeUsecase) {
		uc.conf = conf
	}
}

func WithEthApiMap(ethApiMap map[uint64]*ethapi.Client) options {
	return func(uc *knowledgeUsecase) {
		uc.ethApiMap = ethApiMap
	}
}

func WithNetworks(networks map[string]map[string]string) options {
	return func(uc *knowledgeUsecase) {
		uc.networks = networks
	}
}

func WithTrxApi(trxApi *trxapi.Client) options {
	return func(uc *knowledgeUsecase) {
		uc.trxApi = trxApi
	}
}

func WithRagApi(ragApi string) options {
	return func(uc *knowledgeUsecase) {
		uc.ragApi = ragApi
	}
}

func WithLighthousekey(lighthousekey string) options {
	return func(uc *knowledgeUsecase) {
		uc.lighthouseKey = lighthousekey
	}
}

func WithWebhookUrl(webhookUrl string) options {
	return func(uc *knowledgeUsecase) {
		uc.webhookUrl = webhookUrl
	}
}

func WithNotiBot(teleKey, notiActChanId, notiErrorChanId string) options {
	return func(uc *knowledgeUsecase) {
		if teleKey != "" {
			bot, err := telego.NewBot(teleKey, telego.WithDefaultDebugLogger())
			if err != nil {
				logger.Error(categoryNameTracer, "with_noti_bot", zap.Error(err))
			}
			uc.notiBot = bot
			i, _ := strconv.ParseInt(notiActChanId, 10, 64)
			uc.notiActChanId = i

			ei, _ := strconv.ParseInt(notiErrorChanId, 10, 64)
			uc.notiErrorChanId = ei
		}
	}
}

func NewKnowledgeUsecase(options ...options) ports.IKnowledgeUsecase {
	uc := &knowledgeUsecase{}
	for _, opt := range options {
		opt(uc)
	}
	return uc
}

func (uc *knowledgeUsecase) CalcFeeByKnowledgeBaseId(ctx context.Context, kbId uint) (float64, error) {
	return uc.knowledgeBaseFileRepo.CalcTotalFee(ctx, kbId)
}

func (uc *knowledgeUsecase) SendMessage(_ context.Context, content string, chanId int64) (int, error) {
	if chanId == 0 {
		chanId = uc.notiActChanId
	} else if chanId == -1 {
		chanId = uc.notiErrorChanId
	}

	msg := &telego.SendMessageParams{
		ChatID: telego.ChatID{ID: chanId},
		Text:   strings.TrimSpace(content),
	}

	resp, err := uc.notiBot.SendMessage(msg)
	if err != nil {
		return 0, err
	}
	return resp.MessageID, nil
}

func (uc *knowledgeUsecase) CreateAgentInfoKnowledgeBase(ctx context.Context, models []*models.AgentInfoKnowledgeBase, agentInfoId uint) ([]*models.AgentInfoKnowledgeBase, error) {
	return uc.agentInfoKnowledgeBaseRepo.CreateList(ctx, models, agentInfoId)
}

func (uc *knowledgeUsecase) GetKBAgentsUsedOfSocialAgent(ctx context.Context, socialAgentId uint) ([]*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetKBAgentsUsedOfSocialAgent(ctx, socialAgentId)
}

func (uc *knowledgeUsecase) WebhookFile(ctx context.Context, filename string, bytes []byte, id uint) (*models.KnowledgeBase, error) {
	return nil, nil
	kn, err := uc.knowledgeBaseRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	logger.Info(categoryNameTracer, "start_webhook_file", zap.Any("knowledge_base_id", id), zap.Any("filename", filename))
	updatedFields := make(map[string]interface{})
	hash, err := lighthouse.UploadDataWithRetry(uc.lighthouseKey, fmt.Sprintf("%d_%s", time.Now().Unix(), filename), bytes)
	if err != nil {
		logger.Error(categoryNameTracer, "upload_data_with_retry", zap.Error(err))
		updatedFields["status"] = models.KnowledgeBaseStatusProcessingFailed
		updatedFields["last_error_message"] = err.Error()
		uc.SendMessage(ctx, fmt.Sprintf("webhook_file error upload lighthouse to agent: %s (%d) - error %s", kn.Name, kn.ID, updatedFields["last_error_message"]), uc.notiErrorChanId)
		_ = uc.knowledgeBaseRepo.UpdateById(ctx, id, updatedFields)
		return nil, err
	}

	if kn.KbId != "" {
		updatedFields["status"] = models.KnowledgeBaseStatusDone
	}

	updatedFields["filecoin_hash"] = fmt.Sprintf("ipfs://%s", hash)
	if err := uc.knowledgeBaseRepo.UpdateById(ctx, id, updatedFields); err != nil {
		return nil, err
	}
	_, _ = uc.SendMessage(ctx, fmt.Sprintf("start_webhook_file agent: %s (%d): %s - filecoin hash: %s", kn.Name, kn.ID, updatedFields["filecoin_hash"], filename), uc.notiActChanId)
	return kn, nil
}

func (uc *knowledgeUsecase) Webhook(ctx context.Context, req *models.RagHookResponse) (*models.KnowledgeBase, error) {
	logger.Info(categoryNameTracer, "webhook_update_kb", zap.Any("data", req))
	if req.Result == nil {
		return nil, nil
	}

	id, err := strconv.ParseUint(req.Result.Ref, 10, 64)
	if err != nil {
		return nil, err
	}

	kn, err := uc.knowledgeBaseRepo.GetById(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	if len(req.Result.Identifier) > 0 {
		// update status a file in kb list file
		identifies := strings.Split(req.Result.Identifier, "/")
		if len(identifies) < 2 {
			return nil, fmt.Errorf("invalid identifier %s", req.Result.Identifier)
		}
		hash := identifies[0]
		indexFile, err := strconv.Atoi(identifies[1])
		if err != nil {
			return nil, fmt.Errorf("invalid identifier index %s", identifies[1])
		}
		body, _, err := lighthouse.DownloadDataSimple(hash)
		if err != nil {
			return nil, fmt.Errorf("failed to download lighthouse %s", hash)
		}
		var listFileInfo []*lighthouse.FileInLightHouse
		if err := json.Unmarshal(body, &listFileInfo); err != nil {
			return nil, fmt.Errorf("failed to unmarshal lighthouse %s", hash)
		}
		if len(listFileInfo) < indexFile {
			return nil, fmt.Errorf("invalid file index len(listFile) %v , index file %v", len(listFileInfo), indexFile)
		}
		for _, file := range kn.KnowledgeBaseFiles {
			var lighthouseFile lighthouse.FileInLightHouse
			if err := json.Unmarshal([]byte(file.FilecoinHashRawData), &lighthouseFile); err != nil {
				continue
			}
			if lighthouseFile.Name == listFileInfo[indexFile].Name {
				if req.Status == "ok" {
					file.Status = models.KnowledgeBaseFileStatusDone
					err = uc.knowledgeBaseFileRepo.UpdateByKnowledgeBaseId(
						ctx, file.ID,
						map[string]interface{}{"status": file.Status},
					)
					if err != nil {
						return nil, fmt.Errorf("failed to update knowledge base file status %d", file.Status)
					}

					uc.SendMessage(ctx, fmt.Sprintf("process kb file  for kb agent via webhook DONE (kb_id: %d: %s : %v)", kn.ID, req.Result.Kb, req.Result.Identifier), uc.notiActChanId)

				} else if req.Status == "error" {
					file.Status = models.KnowledgeBaseFileStatusFail
					file.LastErrorMessage = req.Result.Message
					err = uc.knowledgeBaseFileRepo.UpdateByKnowledgeBaseId(
						ctx, file.ID,
						map[string]interface{}{
							"status":             file.Status,
							"last_error_message": file.LastErrorMessage,
						},
					)
					if err != nil {
						return nil, fmt.Errorf("failed to update knowledge base file status %d", file.Status)
					}
					uc.SendMessage(ctx, fmt.Sprintf("webhook update kb file status failed: %s (%d) %s - error %s", kn.Name, kn.ID, req.Result.Identifier, file.LastErrorMessage), uc.notiActChanId)
				}
				break
			}
		}
	} else {
		// update kb status
		updatedFields := make(map[string]interface{})
		if req.Status == "ok" && req.Result.Kb == "" {
			msg := "the kb_id is missing from the webhook API response."
			uc.SendMessage(ctx, fmt.Sprintf("webhook update agent status failed: %s (%d) - error %s", kn.Name, kn.ID, msg), uc.notiActChanId)

			updatedFields["status"] = models.KnowledgeBaseStatusProcessingFailed
			updatedFields["last_error_message"] = msg
			if err := uc.knowledgeBaseRepo.UpdateById(ctx, kn.ID, updatedFields); err != nil {
				return nil, err
			}
			return nil, errors.New(msg)
		}

		if req.Status != "ok" {
			updatedFields["status"] = models.KnowledgeBaseStatusProcessingFailed
			updatedFields["last_error_message"] = req.Result.Message
			uc.SendMessage(ctx, fmt.Sprintf("webhook update agent status failed: %s (%d) - error %s", kn.Name, kn.ID, req.Result.Message), uc.notiActChanId)
		} else if kn.KbId == "" {
			updatedFields["kb_id"] = req.Result.Kb
			updatedFields["status"] = models.KnowledgeBaseStatusDone
			uc.SendMessage(ctx, fmt.Sprintf("Process update kb_id for agent via webhook DONE (kb_id: %d: %s)", kn.ID, req.Result.Kb), uc.notiActChanId)
		} else {
			updatedFields["status"] = models.KnowledgeBaseStatusProcessUpdate
			uc.SendMessage(ctx, fmt.Sprintf("Process update kb_id for agent via webhook DONE (kb_id: %d: %s)", kn.ID, req.Result.Kb), uc.notiActChanId)
		}

		if err := uc.knowledgeBaseRepo.UpdateById(ctx, kn.ID, updatedFields); err != nil {
			return nil, err
		}
	}

	return kn, nil
}

func (uc *knowledgeUsecase) CreateKnowledgeBase(ctx context.Context, req *serializers.CreateKnowledgeRequest) (*serializers.KnowledgeBase, error) {
	model := &models.KnowledgeBase{}
	if err := utils.Copy(model, req); err != nil {
		return nil, err
	}

	// encryptedTempKey, tempAddr, err := utils.GenerateAddress(uc.secretKey)
	// if err != nil {
	// 	return nil, err
	// }
	// model.DepositPrivKey = encryptedTempKey
	model.DepositAddress = strings.ToLower(req.DepositAddress)

	// encryptedTempKey, tempAddr, err = utils.GenerateSolanaAddress(uc.secretKey)
	// if err != nil {
	// 	return nil, err
	// }
	// model.SolanaDepositPrivKey = encryptedTempKey
	model.SolanaDepositAddress = strings.ToLower(req.SolanaDepositAddress)

	model.Status = models.KnowledgeBaseStatusWaitingPayment

	resp, err := uc.knowledgeBaseRepo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	grFileId := time.Now().Unix()
	files := []*models.KnowledgeBaseFile{}
	for _, f := range req.Files {
		file := &models.KnowledgeBaseFile{
			FileUrl:         f.FileUrl,
			FileName:        f.FileName,
			FileSize:        f.FileSize,
			KnowledgeBaseId: resp.ID,
			GroupFileId:     grFileId,
			Status:          models.KnowledgeBaseFileStatusPending,
		}
		_, err := uc.knowledgeBaseFileRepo.Create(ctx, file)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	model.KnowledgeBaseFiles = files
	model.Fee, _ = uc.knowledgeBaseFileRepo.CalcTotalFee(ctx, model.ID)
	model.ChargeMore = model.Fee

	updatedFields := make(map[string]interface{})
	updatedFields["fee"] = model.Fee
	updatedFields["charge_more"] = model.ChargeMore
	if err := uc.UpdateKnowledgeBaseById(ctx, resp.ID, updatedFields); err != nil {
		return nil, err
	}

	result := &serializers.KnowledgeBase{}
	if err := utils.Copy(result, model); err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *knowledgeUsecase) ListKnowledgeBase(ctx context.Context, req *models.ListKnowledgeBaseRequest) ([]*serializers.KnowledgeBase, error) {
	resp, err := uc.knowledgeBaseRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := []*serializers.KnowledgeBase{}
	if err := utils.Copy(&result, resp); err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *knowledgeUsecase) MapKnowledgeBaseByAgentIds(ctx context.Context, ids []uint) (map[uint][]*models.KnowledgeBase, error) {
	resp, err := uc.agentInfoKnowledgeBaseRepo.ListByAgentIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	data := make(map[uint][]*models.KnowledgeBase)
	for _, r := range resp {
		if _, ok := data[r.AgentInfoId]; !ok {
			data[r.AgentInfoId] = make([]*models.KnowledgeBase, 0)
		}
		data[r.AgentInfoId] = append(data[r.AgentInfoId], r.KnowledgeBase)
	}

	return data, nil
}

func (uc *knowledgeUsecase) GetKnowledgeBaseById(ctx context.Context, id uint) (*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetById(ctx, id)
}

func (uc *knowledgeUsecase) GetAgentInfoKnowledgeBaseByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error) {
	return uc.agentInfoKnowledgeBaseRepo.GetByAgentId(ctx, id)
}

func (uc *knowledgeUsecase) DeleteKnowledgeBaseById(ctx context.Context, id uint) error {
	return uc.knowledgeBaseRepo.DeleteById(ctx, id)
}

func (uc *knowledgeUsecase) GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetByStatus(ctx, status, offset, limit)
}

func (uc *knowledgeUsecase) UpdateListKnowledgeBaseFile(ctx context.Context, kbId uint, files []*serializers.File) (bool, error) {
	fileIds := []uint{}
	grFileId := time.Now().Unix()
	updated := false
	for _, f := range files {
		if f.KbFileId != 0 {
			fileIds = append(fileIds, f.KbFileId)
			continue
		}

		file := &models.KnowledgeBaseFile{
			FileUrl:         f.FileUrl,
			FileName:        f.FileName,
			FileSize:        f.FileSize,
			KnowledgeBaseId: kbId,
			GroupFileId:     grFileId,
			Status:          models.KnowledgeBaseFileStatusPending,
		}
		_, err := uc.knowledgeBaseFileRepo.Create(ctx, file)
		if err != nil {
			return false, err
		}
		updated = true
	}

	// still not support delete
	/*
		currentFiles, err := uc.knowledgeBaseFileRepo.ListByKbId(ctx, kbId)
		if err != nil {
			return err
		}
		mapFiles := make(map[uint]*models.KnowledgeBaseFile)
		for _, f := range currentFiles {
			mapFiles[f.ID] = f
		}

		deletedIds := []uint{}
		for k, v := range mapFiles {
			if v.Status == models.KnowledgeBaseFileStatusDone {
				continue
			}
			if !slices.Contains(fileIds, k) {
				deletedIds = append(deletedIds, v.ID)
			}
		}

		return uc.knowledgeBaseFileRepo.DeleteByIds(ctx, deletedIds)
	*/
	return updated, nil
}

func (uc *knowledgeUsecase) UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	return uc.knowledgeBaseRepo.UpdateById(ctx, id, updatedFields)
}

func (uc *knowledgeUsecase) WatchWalletChange(ctx context.Context) error {
	start := time.Now()
	defer logger.Info(categoryNameTracer, "watch_wallet_change", zap.Any("start", start), zap.Any("end", time.Now()))
	offset := 0
	limit := 30
	for {
		resp, err := uc.knowledgeBaseRepo.GetByStatus(
			ctx, models.KnowledgeBaseStatusWaitingPayment, offset, limit,
		)
		if err != nil {
			return err
		}

		if len(resp) == 0 {
			break
		}

		for _, k := range resp {
			time.Sleep(500 * time.Microsecond)
			if err := uc.CheckBalance(ctx, k); err != nil {
				continue
			}
		}

		offset += len(resp)
	}
	return nil
}

func (uc *knowledgeUsecase) ScanKnowledgeBaseStatusPaymentReceipt(ctx context.Context) {
	start := time.Now()
	defer logger.Info(categoryNameTracer, "scan_knowledge_base_payment_receipt", zap.Any("start", start), zap.Any("end", time.Now()))
	offset := 0
	limit := 30
	for {
		resp, err := uc.knowledgeBaseRepo.GetByStatus(
			ctx, models.KnowledgeBaseStatusPaymentReceipt, offset, limit,
		)
		if err != nil {
			return
		}

		if len(resp) == 0 {
			break
		}

		for _, k := range resp {
			// chargeMore := k.CalcChargeMore() // Ensure charge_more is executed before the file's status is changed.
			_, _, err := uc.insertFilesToRAG(ctx, k)
			if err != nil {
				continue
			}

			// TODO transfer fee to backend wallet
			// i := 0
			// var transferErr error
			// var hash string
			// for i < 10 {
			// 	amount := new(big.Int).SetInt64(int64(chargeMore))
			// 	hash, transferErr = uc.transferFund(k.DepositPrivKey, uc.conf.KnowledgeBaseConfig.BackendWallet, amount, k.NetworkID)
			// 	if transferErr != nil {
			// 		i += 1
			// 		time.Sleep(3 * time.Second)
			// 		continue
			// 	}

			// 	if err := uc.knowledgeBaseFileRepo.UpdateTransferHash(ctx, kbFileIds, hash); err != nil {
			// 		return err
			// 	}
			// 	break
			// }

			// if transferErr != nil && hash == "" {
			// 	_, _ = uc.SendMessage(ctx, fmt.Sprintf("transferFund for agent %s (%d) - has error: %s ", k.Name, k.ID, transferErr.Error()), uc.notiErrorChanId)
			// }

		}
		offset += len(resp)
	}
}

func (uc *knowledgeUsecase) CheckBalance(ctx context.Context, kn *models.KnowledgeBase) error {
	knPrice := new(big.Float).SetFloat64(kn.Fee)
	knPrice = knPrice.Mul(knPrice, big.NewFloat(1e18))
	_knPrice := new(big.Int)
	_knPrice, _ = knPrice.Int(_knPrice)

	logger.Info(categoryNameTracer, "check_balance_and_process",
		zap.Any("knowledge_base", kn),
		zap.Any("knPrice", knPrice),
		zap.Any("_knPrice", _knPrice),
	)

	for networkId, net := range uc.networks {
		nId, err := strconv.ParseUint(networkId, 10, 64)
		if err != nil {
			continue
		}

		ethClient, ok := uc.ethApiMap[nId]
		if !ok {
			if net["rpc_url"] == "" {
				continue
			}

			uc.ethApiMap[nId] = &ethapi.Client{
				BaseURL:           net["rpc_url"],
				MinGasPrice:       net["min_gas_price"],
				BTCL1:             strings.ToUpper(net["is_btc_l1"]) == "TRUE",
				BlockTimeDisabled: true,
			}
			ethClient = uc.ethApiMap[nId]
		}

		balance, err := uc.balanceOfAddress(ctx, kn.DepositAddress, ethClient, net)
		if err != nil {
			continue
		}

		if balance.Cmp(_knPrice) >= 0 && _knPrice.Uint64() > 0 {
			kn1, err := uc.GetKnowledgeBaseById(ctx, kn.ID)
			if err != nil {
				return err
			}

			if int(kn1.Status) >= int(models.KnowledgeBaseStatusPaymentReceipt) {
				return nil
			}

			updatedFields := make(map[string]interface{})
			kn.Status = models.KnowledgeBaseStatusPaymentReceipt
			updatedFields["status"] = kn.Status
			updatedFields["deposit_tx_hash"] = fmt.Sprintf("%s/address/%s", net["explorer_url"], kn.DepositAddress)
			updatedFields["deposit_chain_id"] = nId
			if err := uc.knowledgeBaseRepo.UpdateById(ctx, kn.ID, updatedFields); err != nil {
				return err
			}
			content := fmt.Sprintf("Received amount for kb: %s (%d) on chain #%d", kn.Name, kn.ID, nId)
			uc.SendMessage(ctx, content, uc.notiActChanId)
		}
	}
	return nil
}

func (uc *knowledgeUsecase) balanceOfAddress(_ context.Context, address string, client *ethapi.Client, netInfo map[string]string) (*big.Int, error) {
	chainId, err := client.ChainID()
	if err != nil {
		return nil, err
	}
	conAddress := netInfo["eai_contract_address"]
	if address == "0x0000000000000000000000000000000000000000" || chainId == 0 || conAddress == "" {
		return big.NewInt(0), nil
	}
	if strings.EqualFold(conAddress, "0x000000000000000000000000000000000000800A") {
		balance, err := client.Balance(address)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}

	if chainId == models.TRON_CHAIN_ID {
		balance, err := uc.trxApi.Trc20Balance(conAddress, address)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}

	balanace, err := client.Erc20Balance(conAddress, address)
	if err != nil {
		return nil, err
	}

	return balanace, nil
}

func (uc *knowledgeUsecase) insertFilesToRAG(ctx context.Context, kn *models.KnowledgeBase) (*models.InsertRagResponse, []uint, error) {
	resp := &models.InsertRagResponse{}
	hash, kbFileIds, err := uc.uploadKBFileToLighthouseAndProcess(ctx, kn)
	if err != nil {
		uc.SendMessage(ctx, fmt.Sprintf("uploadKBFileToLighthouseAndProcess for agent %s (%d) - has error: %s ", kn.Name, kn.ID, err.Error()), uc.notiErrorChanId)
		return nil, nil, err
	}

	body := struct {
		FilecoinMetadataUrl string `json:"filecoin_metadata_url"`
		Ref                 string `json:"ref"`
		Hook                string `json:"hook"`
		Kb                  string `json:"kb"`
	}{
		FilecoinMetadataUrl: fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash),
		Ref:                 fmt.Sprintf("%d", kn.ID),
		Hook:                uc.webhookUrl,
		Kb:                  kn.KbId,
	}
	logger.Info(categoryNameTracer, "insert_file_to_rag", zap.Any("body", body))

	// if kn.KbId != "" {
	// 	kn.Status = models.KnowledgeBaseStatusProcessUpdate
	// } else {
	// 	kn.Status = models.KnowledgeBaseStatusProcessing
	// }
	kn.FilecoinHash = fmt.Sprintf("ipfs://%s", hash)
	_, err = resty.New().R().SetContext(ctx).SetDebug(true).
		SetBody(body).
		SetResult(resp).
		Post(fmt.Sprintf("%s/api/insert", uc.ragApi))
	if err != nil {
		_, _ = uc.SendMessage(ctx, fmt.Sprintf("insertFilesToRAG for agent %s (%d) - has error: %s ", kn.Name, kn.ID, err.Error()), uc.notiErrorChanId)
		return nil, nil, err
	}

	bBody, _ := json.Marshal(body)
	kn.RagInsertFileRequest = string(bBody)
	kn.Status = models.KnowledgeBaseStatusProcessing
	updatedFields := make(map[string]interface{})
	updatedFields["status"] = kn.Status
	updatedFields["rag_insert_file_request"] = kn.RagInsertFileRequest
	updatedFields["filecoin_hash"] = kn.FilecoinHash
	if err = uc.knowledgeBaseRepo.UpdateById(ctx, kn.ID, updatedFields); err != nil {
		_, _ = uc.SendMessage(ctx, fmt.Sprintf(" uc.knowledgeBaseRepo.UpdateById for agent %s (%d) - has error: %s ", kn.Name, kn.ID, err.Error()), uc.notiErrorChanId)
		return nil, nil, err
	}
	_, _ = uc.SendMessage(ctx, fmt.Sprintf("insertFilesToRAG for agent_id %s (%d): %s", kn.Name, kn.ID, string(bBody)), uc.notiActChanId)
	return resp, kbFileIds, nil
}

func (uc *knowledgeUsecase) GetKnowledgeBaseByKBId(ctx context.Context, kbId string) (*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetByKBId(ctx, kbId)
}

func (uc *knowledgeUsecase) GetKnowledgeBaseByKBTokenId(ctx context.Context, kbId string) (*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetByKBTokenId(ctx, kbId)
}

func (uc *knowledgeUsecase) GetManyKnowledgeBaseByQuery(ctx context.Context, query string, orderOption string, offset int, limit int) ([]*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetManyByQuery(ctx, query, orderOption, offset, limit)
}

func (uc *knowledgeUsecase) uploadKBFileToLighthouseAndProcess(ctx context.Context, kn *models.KnowledgeBase) (string, []uint, error) {
	result := []*lighthouse.FileInLightHouse{}
	kbFileIds := []uint{}
	for _, f := range kn.KnowledgeBaseFiles {
		if f.FilecoinHashRawData != "" {
			r := &lighthouse.FileInLightHouse{}
			if err := json.Unmarshal([]byte(f.FilecoinHashRawData), r); err == nil {
				if f.Status == models.KnowledgeBaseFileStatusDone {
					r.IsInserted = true
				}
				result = append(result, r)
				continue
			}
		}

		r, err := lighthouse.ZipAndUploadFileInMultiplePartsToLightHouseByUrl(f.FileUrl, "/tmp/data", uc.lighthouseKey)
		if err != nil {
			uc.SendMessage(ctx, fmt.Sprintf("uploadKBFileToLighthouseAndProcess for agent %s (%d) - has error: %s ", kn.Name, kn.ID, err.Error()), uc.notiErrorChanId)
			return "", nil, err
		}

		rw, _ := json.Marshal(r)
		if len(r.Files) > 0 {
			f.FilecoinHash = r.Files[0].Hash
		}
		f.FilecoinHashRawData = string(rw)
		err = uc.knowledgeBaseFileRepo.UpdateByKnowledgeBaseId(
			ctx, f.ID,
			map[string]interface{}{
				"filecoin_hash":          f.FilecoinHash,
				"filecoin_hash_raw_data": f.FilecoinHashRawData,
			},
		)
		if err != nil {
			return "", nil, err
		}
		kbFileIds = append(kbFileIds, f.ID)
		r.IsInserted = false
		result = append(result, r)
	}

	data, _ := json.Marshal(result)
	hash, err := lighthouse.UploadData(uc.lighthouseKey, kn.Name, data)
	if err != nil {
		uc.SendMessage(ctx, fmt.Sprintf("uploadKBFileToLighthouseAndProcess for agent %s (%d) - has error: %s ", kn.Name, kn.ID, err.Error()), uc.notiErrorChanId)
		return "", nil, err
	}
	return hash, kbFileIds, nil
}

func (uc *knowledgeUsecase) transferFund(priKeyFrom string, toAddress string, fund *big.Int, networkId uint64) (string, error) {
	_, pubKey, err := eth.GetAccountInfo(priKeyFrom)
	if err != nil {
		return "", fmt.Errorf("get account info: %v", err)
	}

	rpc := uc.conf.GetConfigKeyString(networkId, "rpc_url")
	var paymasterAddress, paymasterToken string
	var paymasterFeeZero bool
	if uc.conf.ExistsedConfigKey(networkId, "paymaster_address") &&
		uc.conf.ExistsedConfigKey(networkId, "paymaster_token") {
		paymasterAddress = uc.conf.GetConfigKeyString(networkId, "paymaster_address")
		paymasterToken = uc.conf.GetConfigKeyString(networkId, "paymaster_token")
		paymasterFeeZero = uc.conf.GetConfigKeyBool(networkId, "paymaster_fee_zero")
	}
	aiZkClient := zkclient.NewZkClient(
		rpc,
		paymasterFeeZero,
		paymasterAddress,
		paymasterToken,
	)
	tx, err := aiZkClient.Transact(priKeyFrom, *pubKey, common.HexToAddress(toAddress), fund, nil)
	if err != nil {
		return "", fmt.Errorf("failed to transact: %v", err)
	}
	return tx.TxHash.Hex(), nil
}
