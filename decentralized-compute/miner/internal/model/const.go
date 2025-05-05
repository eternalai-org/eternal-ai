package model

const (
	BatchInferHistoryStatusPending       string = "pending"
	BatchInferHistoryStatusAgentInferred string = "agent-inferred"
	BatchInferHistoryStatusQueueHandled  string = "queue-handled"
	BatchInferHistoryStatusCompleted     string = "completed"

	BatchInferHistoryStatusFailed string = "failed"
)

const (
	ToolsetTypeDefault          string = "default"
	ToolsetTypeReplyMentions    string = "reply_mentions"
	ToolsetTypeReplyNonMentions string = "reply_non_mentions"
	ToolsetTypeFollow           string = "follow"
	ToolsetTypePost             string = "post"
	ToolsetTypeCreateToken      string = "create_token"
)

const (
	LightHouseStorageType StorageType = "lighthouse-filecoin"
	EaiChainStorageType   StorageType = "eai-chain"
)
