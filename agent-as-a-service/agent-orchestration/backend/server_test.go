package main_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/databases"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services"
)

var ts *services.Service

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conf := configs.GetConfig()
	logger.NewLogger("agents-ai-api", conf.Env, "", true)
	defer logger.Sync()
	dbMain, err := databases.Init(
		conf.DbURL,
		nil,
		1,
		20,
		conf.Debug,
	)
	if err != nil {
		panic(err)
	}
	daos.InitDBConn(
		dbMain,
	)
	var (
		s = services.NewService(
			conf,
		)
	)
	ts = s
}

func Test_JOB(t *testing.T) {

	fmt.Println(
		ts.JobCreateTokenInfo(context.Background()),
	)

	// data, err := ts.ValidateTweetContentGenerateVideoWithLLM2(context.Background(), "man opens shirt and shows his fat belly")
	// fmt.Println(data, err)
	// ts.MemeEventsByTransaction(context.Background(), 56, "")
	// ts.AgentSnapshotPostCreate(context.Background(), 59166, "", "")
	// ts.JobScanAgentTwitterPostForTA(context.Background())
	// ts.RetryAgentDeployToken(context.Background(), 51265)
	// ts.JobUpdateOffchainAutoOutputForMission(context.Background())
	// ts.JobAgentTwitterPostTA(context.Background())
	// ts.JobLuckyMoneyProcessUserReward(context.Background())

	// fmt.Println(
	// // ts.DeployDAOTreasuryLogic(context.Background(), models.BASE_CHAIN_ID),
	// // ts.DeployDAOTreasuryAddress(context.Background(), models.BASE_CHAIN_ID),
	// // ts.AgentAddLiquidityDAOToken(context.Background(), 1),
	// // ts.CreateSOLAddress(context.Background()),
	// // ts.CreateETHAddress(context.Background()),
	// // ts.JobAgentTgeTransferDAOToken(context.Background()),
	// // ts.JobAgentAddLiquidityDAOToken(context.Background()),
	// // ts.DeployDAOTreasuryLogic(context.Background(), models.BASE_CHAIN_ID),
	// // ts.JobAgentTgeTransferDAOToken(context.Background()),
	// )
	// select {}
}

func Test_UTIL(t *testing.T) {
	// ts.JobCreateTokenInfo(context.Background())
	// resp := map[string]interface{}{
	// 	"method": "getUserByUsername",
	// 	"params": map[string]interface{}{
	// 		"username": "Uniswap",
	// 	},
	// }
	// jsonString, _ := json.Marshal(resp)
	// fmt.Println(jsonString)
	ts.JobUpdateTrendingTokens(context.Background())
}

func Test_OpenAI(t *testing.T) {
	ts.GetTelegramMessage(context.Background())
}

func Test_SRV(t *testing.T) {
	// ts.TestCrawlDexScreener(context.Background())
	// ts.JobRobotScanBalanceSOL(context.Background())
	ts.CreateClankerTokenForVideoByPostID(context.Background(), 35944)
}

func Test_UpdateTokenPrice(t *testing.T) {
	ts.GenerateTokenInfoFromVideoPrompt(context.Background(), "Dancing characters in green frog hoodie in a grid", false)
	// ts.GenerateTokenInfoWithLLMV2(context.Background(), "Dancing characters in green frog hoodie in a grid", false)
	// ts.GetGifImageUrlFromTokenInfo("ABC", "ABC", "ABC")
	// ts.CreateTokenInfo(context.Background(), 14754)
}

func Test_IPFS(t *testing.T) {
	resp := map[string]any{
		"name":          "XXX",
		"description":   "XXX Description ...",
		"symbol":        "XXX",
		"image":         "https://xxx.jpg",
		"animation_url": "https://xxx.mp4",
		"content": map[string]any{
			"uri":  "https://xxx.mp4",
			"mime": "video/mp4",
		},
	}
	jsonString, _ := json.Marshal(resp)
	fmt.Println(
		ts.IpfsUploadDataForName(context.Background(), "data", jsonString),
	)
}
