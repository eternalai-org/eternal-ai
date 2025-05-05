package model

type LocalChain struct {
	Contracts            map[string]string `json:"contracts"`
	Rpc                  string            `json:"rpc"`
	ModelName            string            `json:"model_name"`
	UseExternalRunPod    bool              `json:"use_external_run_pod"`
	RunPodInternal       string            `json:"run_pod_internal"` // if we use custom runpod url (not local).  RunPodInternal=RunPodExternal=url
	RunPodExternal       string            `json:"run_pod_external"` //for heath check
	RunPodAPIKEY         string            `json:"run_pod_api_key"`
	ModelID              string            `json:"model_id"`
	ChainID              string            `json:"chain_id"`
	PrivateKey           string            `json:"private_key"`
	Miners               map[string]Miners `json:"miners"`
	Platform             string            `json:"platform"`
	ApiUrl               string            `json:"api_url"`
	AgentContractAddress string            `json:"agent_contract_address"`
	PubSubURL            string            `json:"pub_sub_url"`
}

type Miners struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}
