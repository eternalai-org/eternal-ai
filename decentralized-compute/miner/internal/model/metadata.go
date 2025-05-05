package model

type ClusterMetaData struct {
	Version          int    `json:"version"`
	ModelName        string `json:"model_name"`
	ModelType        string `json:"model_type"`
	ModelUrl         string `json:"model_url"`
	ModelFileHash    string `json:"model_file_hash"`
	MinHardware      int    `json:"min_hardware"`
	VerifierUrl      string `json:"verifier_url"`
	VerifierFileHash string `json:"verifier_file_hash"`
}
