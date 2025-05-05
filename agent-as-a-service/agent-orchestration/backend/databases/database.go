package databases

import (
	"sync"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

var (
	dbMutex sync.Mutex
	dbMap   map[string]*gorm.DB
)

func init() {
	dbMap = map[string]*gorm.DB{}
}

func Init(dbURL string, migrateFunc func(db *gorm.DB) error, idleNum int, openNum int, debug bool) (*gorm.DB, error) {
	dbConn, err := gorm.Open("mysql", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	dbConn.LogMode(debug)
	dbConn = dbConn.Set("gorm:save_associations", false)
	dbConn = dbConn.Set("gorm:association_save_reference", false)
	dbConn.DB().SetMaxIdleConns(idleNum)
	dbConn.DB().SetMaxOpenConns(openNum)
	if migrateFunc != nil {
		func() {
			err = migrateFunc(dbConn)
			if err != nil {
				panic(err)
			}
		}()
	}
	return dbConn, nil
}

func MigrateDBMain(db *gorm.DB) error {
	allTables := []any{
		(*models.BlockScanInfo)(nil),
		(*models.AppConfig)(nil),
		(*models.BTCL1InscribeTx)(nil),
		(*models.User)(nil),
		(*models.UserTransaction)(nil),
		(*models.Erc20Holder)(nil),
		(*models.Erc721Holder)(nil),
		(*models.Erc1155Holder)(nil),
		(*models.TwitterPost)(nil),
		(*models.TwitterInfo)(nil),
		(*models.AgentCategory)(nil),
		(*models.AgentInfo)(nil),
		(*models.AgentTwitterPost)(nil),
		(*models.UserTwitterPost)(nil),
		(*models.AgentTokenInfo)(nil),
		(*models.AgentTradeHistory)(nil),
		(*models.TokenPrice)(nil),
		(*models.TwitterTweet)(nil),
		(*models.AgentEaiTopup)(nil),
		(*models.AgentSnapshotPost)(nil),
		(*models.AgentSnapshotMission)(nil),
		(*models.AgentSnapshotPostAction)(nil),
		(*models.AuthCode)(nil),
		(*models.AgentWallet)(nil),
		(*models.AgentWalletAction)(nil),
		(*models.TwitterTweetLiked)(nil),
		(*models.ExternalWallet)(nil),
		(*models.ExternalWalletOrder)(nil),
		(*models.ExternalWalletToken)(nil),
		(*models.AgentChainFee)(nil),
		// meme
		(*models.Meme)(nil),
		(*models.MemeFeesCollected)(nil),

		(*models.ApiSubscriptionPackage)(nil),
		(*models.ApiSubscriptionKey)(nil),
		(*models.ApiSubscriptionHistory)(nil),
		(*models.ApiSubscriptionUsageLog)(nil),

		(*models.AgentSnapshotMissionConfigs)(nil),
		(*models.AgentTradeToken)(nil),
		(*models.AgentExternalInfo)(nil),

		(*models.AgentTeleMsg)(nil),

		(*models.BatchInferHistory)(nil),
		(*models.ChainConfig)(nil),
		(*models.TrainingRequestERC20Info)(nil),
		(*models.ZkSyncNetwork)(nil),
		(*models.TrainingRequest)(nil),
		(*models.ModelPredictHistory)(nil),
		(*models.ModelMarket)(nil),
		(*models.JobConfig)(nil),
		(*models.KnowledgeBase)(nil),
		(*models.KnowledgeBaseFile)(nil),
		(*models.AgentInfoKnowledgeBase)(nil),
		// missionstore
		(*models.MissionStore)(nil),
		(*models.MissionStoreRating)(nil),
		(*models.MissionStoreHistory)(nil),
		// launchpad
		(*models.Launchpad)(nil),
		(*models.LaunchpadMember)(nil),
		(*models.LaunchpadTransaction)(nil),
		(*models.AbilityLuckyMoney)(nil),

		//
		(*models.AgentStore)(nil),
		(*models.AgentStoreMission)(nil),
		(*models.AgentStoreInstall)(nil),

		(*models.AgentStoreLog)(nil),

		(*models.SampleTwitterApp)(nil),

		(*models.InfraTwitterApp)(nil),
		(*models.InfraTwitterTopupTx)(nil),

		(*models.StoreDefiApp)(nil),
		(*models.AgentStoreTry)(nil),
		(*models.AgentStoreTryDetail)(nil),

		(*models.AgentInfoInstall)(nil),

		(*models.AgentLibrary)(nil),
		(*models.AgentUtilityInstall)(nil),
		(*models.AgentUtilityRecentChat)(nil),
		(*models.ClankerVideoToken)(nil),
		(*models.AgentReactionHistory)(nil),
		(*models.PrivyWallet)(nil),
		(*models.RobotSaleWallet)(nil),
		(*models.RobotProject)(nil),
		(*models.RobotTokenTransfer)(nil),
		(*models.TrendingToken)(nil),
		(*models.VibeWhiteList)(nil),
		(*models.VibeReferralCode)(nil),
		(*models.AgentUserComment)(nil),
	}

	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return err
	}
	return nil
}
