package model

type DeviceInformation struct {
	Device    string `json:"device"`
	OS        string `json:"os"`
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	GPU       string `json:"gpu"`
	GPUCore   string `json:"gpu_core"`
}

type Network struct {
	Name    string `json:"name"`
	ChainID string `json:"chain_id"`
	Rpc     string `json:"rpc"`
}

type OllamaModelInformation struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

type OnChainData struct {
	Address          string                 `json:"address"`
	ID               string                 `json:"id"`
	ProcessingTasks  uint64                 `json:"processing_tasks"`
	Network          Network                `json:"network"`
	ModelInformation OllamaModelInformation `json:"model_information"`
}

type Node struct {
	Name string      `json:"name"`
	Data OnChainData `json:"data"`
}

type CreateNode struct {
	NodeID           string                 `json:"-"`
	ModelInformation OllamaModelInformation `json:"model_information"`
	Network          Network                `json:"network"`
	PrivateKey       string                 `json:"private_key"`

	RunPodAPI    string `json:"run_pod_api"`
	RunPodAPIKey string `json:"run_pod_api_key"`
}

type CreateInference struct {
	Request LLMInferRequest `json:"request"`
	PrvKey  string          `json:"prv_key"`
}
