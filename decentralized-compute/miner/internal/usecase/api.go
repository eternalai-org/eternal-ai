package usecase

import (
	"context"
	"core/contracts/prompt_scheduler"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"solo/config"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/db/sqlite"
	"solo/pkg/eth"
	"solo/pkg/logger"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/common"
)

type api struct {
	port int
}

func NewAPI(port int) port.IApi {
	return &api{port: port}
}

func (a *api) HealthCheck(ctx context.Context) (bool, error) {
	return true, nil
}

func (a *api) Information(ctx context.Context) (*model.DeviceInformation, error) {
	swInfo, err := pkg.ExecCommand("system_profiler", "SPSoftwareDataType")
	if err != nil {
		return nil, err
	}
	_swinfo := a.parseInfo(*swInfo, ":")
	_ = _swinfo
	macOs, ok := _swinfo["System Version"]
	if !ok {
		macOs = "..."
	}

	info, err := pkg.ExecCommand("system_profiler", "SPHardwareDataType")
	if err != nil {
		return nil, err
	}
	_info := a.parseInfo(*info, ":")
	deviceName, ok := _info["Model Name"]
	if !ok {
		deviceName = "..."
	}

	ram, ok := _info["Memory"]
	if !ok {
		ram = "..."
	}

	processor, ok := _info["Processor Name"]
	if !ok {
		processor = "..."
	}

	//GPU
	gpu, err := pkg.ExecCommand("system_profiler", "SPDisplaysDataType")
	if err != nil {
		gpu = new(string)
		*gpu = "..."
	}

	_gpu := a.parseInfo(*gpu, ":")
	_gpuModel, ok := _gpu["Chipset Model"]
	_gpuCores := "1" //TODO - implement me
	if !ok {
		_gpuModel = "..."
	}

	resp := &model.DeviceInformation{
		Device:    deviceName,
		OS:        macOs,
		Processor: processor,
		GPU:       _gpuModel,
		GPUCore:   _gpuCores,
		Ram:       ram,
	}

	return resp, nil
}

func (a *api) OnChainData(ctx context.Context) (*model.OnChainData, error) {
	//TODO - implement me
	resp := &model.OnChainData{
		Address:         "0xae32",
		ID:              "12345",
		ProcessingTasks: 2345,
		Network: model.Network{
			Name:    "Arbitrum",
			ChainID: "123",
		},
		ModelInformation: model.OllamaModelInformation{
			Name: "Llama 3.2 - 3B",
			Size: "2.32 GB",
		},
	}

	return resp, nil
}

func (a *api) ListNodes(ctx context.Context) ([]*model.Node, error) {
	resp := []*model.Node{}
	dir := fmt.Sprintf(pkg.ENV_FOLDER, pkg.CurrentDir())

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		node, err := a.nodeInfo(file.Name())
		if err != nil {
			continue
		}
		resp = append(resp, node)
	}

	return resp, nil
}

func (a *api) CreateNodes(ctx context.Context, input model.CreateNode) (*model.Node, error) {
	//generate node_id {node_id}
	fname, err := a.createFileName(&input)
	if err != nil {
		return nil, err
	}

	_, err = a.CreateConfigENV(&input, *fname)
	if err != nil {
		return nil, err
	}

	node, err := a.nodeInfo(*fname)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (a *api) ListInferences(ctx context.Context, nodeID string) (interface{}, error) {
	resp := []*model.SuccessResponse{}

	sqldb, err := sqlite.NewSqliteDB(nodeID)
	if err != nil {
		return nil, err
	}
	sqldb.Connect()
	defer sqldb.Close()

	task := model.SuccessResponse{}
	sql := fmt.Sprintf("select task_id, created_by, created_at,assignment_id,model_contract,params,assignment_role,result_data,submit_tx, submitted_by from %s order by id desc", task.TableName())
	rows, err := sqldb.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			createdBy      string
			createdAt      string
			taskID         string
			assignmentID   string
			modelContract  string
			params         string
			assignmentRole string
			resultData     string
			submitTX       string
			submittedBy    string
		)

		if err := rows.Scan(&taskID, &createdBy, &createdAt, &assignmentID, &modelContract, &params, &assignmentRole, &resultData, &submitTX, &submittedBy); err != nil {
			continue
		}

		item := &model.SuccessResponse{
			InferenceID:  taskID,
			AssignmentID: assignmentID,
			CreatedBy:    createdBy,
			CreatedAt:    createdAt,
			SubmittedBy:  submittedBy,
			SubmitTx:     submitTX,
		}

		taskResult := new(model.TaskResult)
		_resultData := []byte(resultData)
		err = json.Unmarshal(_resultData, taskResult)
		if err == nil {
			data := taskResult.Data
			_resultDataData := new(model.LLMInferResponse)
			err = json.Unmarshal(data, _resultDataData)
			if err == nil {
				item.ResultData = _resultDataData.Choices
			}
		}

		_params := []byte(params)
		_paramObject := new(model.LLMInferRequest)
		err = json.Unmarshal(_params, _paramObject)
		if err == nil {
			item.Param = _paramObject
		}

		resp = append(resp, item)

	}
	return resp, nil
}

func (a *api) Chains(ctx context.Context) map[string]config.Chain {
	return config.ChainConfig()
}

func (a *api) CreateInferences(ctx context.Context, chainID string, input *model.CreateInference) (interface{}, error) {
	logkey := "CreateInferences"
	logs := new([]zap.Field)

	*logs = append(*logs, zap.Any("input", input))
	*logs = append(*logs, zap.Any("chainID", chainID))
	var err error

	defer func() {
		if err != nil {
			*logs = append(*logs, zap.Error(err))
			logger.AtLog.Logger.Error(logkey, *logs...)
		} else {
			logger.AtLog.Logger.Info(logkey, *logs...)
		}
	}()

	cnfs := config.ChainConfig()
	cnf, ok := cnfs[chainID]
	if !ok {
		err := errors.New("chain_id is incorrect")
		return nil, err
	}

	//TODO - remove hard code
	client, err := eth.NewEthClient(cnf.RPC)
	if err != nil {
		return nil, err
	}

	promptScheduleAddress := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER]

	promptScheduleContract, err := prompt_scheduler.NewPromptScheduler(common.HexToAddress(promptScheduleAddress), client)
	if err != nil {
		return nil, err
	}

	privKey := input.PrvKey
	auth, err := eth.CreateBindTransactionOpts(ctx, client, privKey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		return nil, err
	}

	modelID := cnf.ModelID
	_input, _ := json.Marshal(input.Request)
	_, address, err := eth.GetAccountInfo(privKey)

	*logs = append(*logs, zap.String("address", address.Hex()))

	tx, err := promptScheduleContract.Infer(auth, modelID, _input, *address, true)
	if err != nil {
		return nil, err
	}

	*logs = append(*logs, zap.String("tx", tx.Hash().Hex()))
	_, err = eth.WaitForTxReceipt(client, tx.Hash())
	if err != nil {
		return nil, err
	}

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}

	txLogs := receipt.Logs
	var inferId uint64
	for _, item := range txLogs {
		if item == nil {
			continue
		}

		inferData, err := promptScheduleContract.ParseNewInference(*item)
		if err == nil {
			inferId = inferData.InferenceId
		}
	}

	*logs = append(*logs, zap.Uint64("inferId", inferId))
	out := []byte{}
	//wait for success data

	processedMiner := make(map[string]common.Address)
	for i := 1; i <= 1000; i++ {
		time.Sleep(time.Second * 2)
		infer, err1 := promptScheduleContract.GetInferenceInfo(nil, inferId)
		if err1 != nil {
			continue
		}
		out = infer.Output
		if len(out) > 0 {
			break
		}

		processedMiner[infer.ProcessedMiner.Hex()] = infer.ProcessedMiner
	}

	taskResult := new(model.TaskResult)
	err = json.Unmarshal(out, taskResult)
	if err != nil {
		return nil, err
	}

	_resultDataData := new(model.LLMInferResponse)
	err = json.Unmarshal(taskResult.Data, _resultDataData)
	if err != nil {
		return nil, err
	}

	_resultDataData.OnchainData.InferTx = tx.Hash().Hex()
	_resultDataData.OnchainData.InferId = inferId
	_resultDataData.OnchainData.Proposer = address.Hex()

	for _, v := range processedMiner {
		_resultDataData.OnchainData.PbftCommittee = append(_resultDataData.OnchainData.PbftCommittee, v.Hex())
	}

	*logs = append(*logs, zap.Any("_resultDataData", _resultDataData))
	return _resultDataData, nil
}

// Private methods
func (a *api) createFileName(input *model.CreateNode) (*string, error) {
	envDir := fmt.Sprintf(pkg.ENV_FOLDER, pkg.CurrentDir())
	if _, err := os.Stat(envDir); os.IsNotExist(err) {
		err = os.MkdirAll(envDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create env directory: %v", err)
		}
	}

	fnames := []string{}
	dir := fmt.Sprintf(pkg.ENV_FOLDER, pkg.CurrentDir())
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fnames = append(fnames, file.Name())
	}

goto_here:
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(9999) + 1
	formattedInt := fmt.Sprintf("%04d", randomInt)

	nodeID := formattedInt
	input.NodeID = nodeID
	//create node's config - (config_{node_id}.env)
	fname := fmt.Sprintf("config_%s.env", nodeID)

	for _, v := range fnames {
		if v == fname {
			// generate again
			goto goto_here
		}
	}

	return &fname, nil
}

func (a *api) parseInfo(input string, sep string) map[string]string {
	resp := make(map[string]string)
	arr := strings.Split(input, "\n")
	for _, i := range arr {
		arr1 := strings.Split(i, sep)
		if len(arr1) <= 1 {
			continue
		}

		if arr1[1] == "" {
			continue
		}

		k := strings.TrimSpace(arr1[0])
		v := strings.TrimSpace(arr1[1])
		resp[k] = v
	}
	return resp
}

func (a *api) nodeInfo(fname string) (*model.Node, error) {
	node := new(model.Node)

	_dir := pkg.ENV_FOLDER + "/" + fname
	_b, err := pkg.ReadFile(fmt.Sprintf(_dir, pkg.CurrentDir()))
	if err != nil {
		return nil, err
	}

	parsed := a.parseInfo(string(_b), "=")
	_ = parsed

	nodeID, ok := parsed["NODE_ID"]
	if !ok {
		nodeID = "..."
	}

	chainID, ok := parsed["CHAIN_ID"]
	if !ok {
		err = errors.New("CHAIN_ID is missing")
		return nil, err
	}

	prvKey, ok := parsed["ACCOUNT_PRIV"]
	if !ok {
		err = errors.New("ACCOUNT_PRIV is missing")
		return nil, err
	}

	_, address, err := eth.GetAccountInfo(prvKey)
	if err != nil {
		return nil, err
	}

	chainName, ok := parsed["CHAIN_NAME"]
	if !ok {
		chainName = "..."
	}

	modelName, ok := parsed["MODEL_NAME"]
	if !ok {
		err = errors.New("MODEL_NAME is missing")
		return nil, err
	}

	rpc, ok := parsed["CHAIN_RPC"]
	if !ok {
		err = errors.New("CHAIN_RPC is missing")
		return nil, err
	}

	node.Name = fname
	node.Data = model.OnChainData{
		Address:         strings.ToLower(address.Hex()),
		ProcessingTasks: 0,
		ID:              nodeID,
		Network: model.Network{
			Name:    chainName,
			ChainID: chainID,
			Rpc:     rpc,
		},
		ModelInformation: model.OllamaModelInformation{
			Name: modelName,
			Size: "",
		},
	}

	return node, nil
}

func (a *api) CreateConfigENV(input *model.CreateNode, fname string) (string, error) {
	//contract
	cnfs := config.ChainConfig()
	cnf, ok := cnfs[input.Network.ChainID]
	if !ok {
		err := errors.New("chain_id is incorrect")
		return "", err
	}

	env := ""
	//env += fmt.Sprintf("PLATFORM=%v\n", cnf.Platform)
	//env += fmt.Sprintf("API_URL=%v\n", apiURL)
	//env += fmt.Sprintf("API_KEY=%v\n", cnf.RunPodAPIKEY)
	//env += fmt.Sprintf("LIGHT_HOUSE_API_KEY=%v\n", os.Getenv("LIGHT_HOUSE_API_KEY"))
	//env += fmt.Sprintf("CLUSTER_ID=%v\n", cnf.ModelID)
	//env += fmt.Sprintf("MODEL_ID=%v\n", cnf.ModelID)

	//Node information
	env += fmt.Sprintf("NODE_ID=%v\n", input.NodeID)
	env += fmt.Sprintf("CHAIN_ID=%v\n", input.Network.ChainID)
	env += fmt.Sprintf("CHAIN_NAME='%s'\n", input.Network.Name)
	env += fmt.Sprintf("CHAIN_RPC=%v\n", input.Network.Rpc)
	env += fmt.Sprintf("ACCOUNT_PRIV=%v\n", input.PrivateKey)
	env += fmt.Sprintf("MODEL_NAME='%s'\n", input.ModelInformation.Name)

	//RUN POD
	env += fmt.Sprintf("API_KEY='%s'\n", input.RunPodAPIKey)
	env += fmt.Sprintf("API_URL='%s'\n", input.RunPodAPI)

	//contracts
	env += fmt.Sprintf("STAKING_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER])
	env += fmt.Sprintf("MODEL_LOAD_BALANCER_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_LOAD_BALANCER])
	env += fmt.Sprintf("WORKER_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER])
	env += fmt.Sprintf("ERC20_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI])
	env += fmt.Sprintf("COLLECTION_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION])

	//save config to a file (./env/config_{node_id}.env).
	p := pkg.ENV_FOLDER + "/" + fname
	path := fmt.Sprintf(p, pkg.CurrentDir())
	err := pkg.CreateFile(path, []byte(env))
	if err != nil {
		return "", err
	}

	return env, nil
}
