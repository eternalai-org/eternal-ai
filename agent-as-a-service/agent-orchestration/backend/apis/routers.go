package apis

import (
	"net/http"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	g    *gin.Engine
	conf *configs.Config
	nls  *services.Service
}

func NewServer(
	g *gin.Engine,
	conf *configs.Config,
	nls *services.Service,
) *Server {
	return &Server{
		g:    g,
		conf: conf,
		nls:  nls,
	}
}

func (s *Server) Routers() {
	s.g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowWildcard:    true,
		MaxAge:           12 * time.Hour,
	}))
	s.g.Use(s.logApiMiddleware())
	s.g.Use(s.recoveryMiddleware(raven.DefaultClient, false))
	rootAPI := s.g.Group("/api")
	{
		rootAPI.GET("/", func(c *gin.Context) {
			ctxJSON(c, http.StatusOK, &serializers.Resp{Error: nil})
		})
		rootAPI.GET("/health", func(c *gin.Context) {
			ctxJSON(c, http.StatusOK, &serializers.Resp{Error: nil})
		})

		rootAPI.GET("/configs/explorer", s.GetAllConfigsExplorer)
		rootAPI.GET("/clear-cache", s.ClearCacheKey)
		rootAPI.GET("/eai/supply/total", s.GetEAISupplyTotal)
		rootAPI.GET("/eai/supply/circulating", s.GetEAISupplyCirculating)
		rootAPI.GET("/coin-prices", s.GetTokenPrice)

		rootAPI.POST("/vibe-white-list", s.AddVibeWhiteList)
		rootAPI.GET("/scan-transaction", s.MemeEventsByTransaction)

		webhookAPI := rootAPI.Group("/webhook")
		{
			webhookAPI.GET("/twitter-oauth", s.TwitterOauthCallback)
			webhookAPI.GET("/twitter-oauth/internal", s.TwitterOauthCallbackForInternalData)
		}

		pumpAPI := rootAPI.Group("/pump")
		{
			pumpAPI.GET("/orders", s.GetPumpOrderHistory)
			pumpAPI.GET("/balances", s.GetPumpBalance)
			rootAPI.GET("/notify-change-price-pump", s.NotifyChangePricePump)
			rootAPI.GET("/trending-tokens", s.GetListTrendingTokens)
		}

		// user
		authAPI := rootAPI.Group("/auth")
		{
			authAPI.POST("/verify", s.VerifyLoginUserByWeb3)
		}

		userAPI := rootAPI.Group("/users")
		{
			userAPI.POST("/upload", s.UserUploadFile)
			userAPI.GET("/profile", s.authCheckTK1TokenMiddleware(), s.GetUserProfileWithAuth)
			userAPI.GET("/transactions", s.authCheckTK1TokenMiddleware(), s.GetListUserTransactions)
		}

		adminAPI := rootAPI.Group("/admin")
		{
			adminAPI.GET("/agent/:id", s.AdminGetAgentDetailByAgentID)
			adminAPI.GET("/snapshot_post_acctions", s.AdminGetSnapshotPostActionByAgent)
		}

		// Agent management
		agentAPI := rootAPI.Group("/agent")
		{
			agentAPI.GET("/categories", s.GetListAgentCategory)
			agentAPI.GET("/report", s.GetAgentSummaryReport)
			agentAPI.GET("/list", s.GetListAgent)
			agentAPI.GET("/unclaimed", s.GetListAgentUnClaimed)
			agentAPI.GET("/:id", s.GetAgentDetailByAgentID)
			agentAPI.GET("/admin/:id", s.GetAgentDetailByAgentID)
			agentAPI.GET("/by-contract/:address/:id", s.GetAgentDetailByContract)
			agentAPI.GET("/detail/:id", s.GetAgentDetail)
			agentAPI.POST("/update-code-version/:id", s.UpdateAgentCodeVersion)

			agentAPI.GET("/post", s.GetListAgentTwitterPost)
			agentAPI.GET("/post/:id", s.GetAgentTwitterPostDetail)
			agentAPI.GET("/topup-history/:id", s.GetListAgentEaiTopup)

			agentAPI.POST("/update-farcaster/:id", s.UpdateAgentFarcaster)
			agentAPI.POST("/update-external-info/:id", s.UpdateAgentExternalInfo)

			agentAPI.POST("/unlink/:id", s.UnlinkAgentTwitterInfo)
			agentAPI.POST("/pause/:id", s.PauseAgent)

			agentAPI.POST("/:id/chat/completions", s.AgentChats)
			agentAPI.POST("/preview", s.PreviewAgentSystemPromp)
			agentAPI.POST("/preview/v1", s.PreviewAgentSystemPrompV1)
			agentAPI.POST("/chats", s.AgentChatSupport)

			agentAPI.GET("/network-fees", s.GetAgentChainFees)
			agentAPI.GET("/liquidity-networks", s.GetEaiLiquidityNetowrks)

			twitterAPI := agentAPI.Group("/twitter")
			{
				twitterAPI.GET("/user/by/username", s.GetTwitterUserByUserNameByQuery)
			}

			faucetAPI := agentAPI.Group("/faucet")
			{
				faucetAPI.GET("/request", s.AgentRequestTwitterShareCode)
				faucetAPI.POST("/verify", s.AgentVerifyShareTwitter)
			}

			brainAPI := agentAPI.Group("/brain")
			{
				brainAPI.GET("/detail-by-tweet/:id", s.GetBrainDetailByTweetID)
				brainAPI.GET("/:id", s.GetAgentBrainHistory)
			}

			// abilities
			abilitiesAPI := agentAPI.Group("/mission")
			{
				abilitiesAPI.POST("update/:id", s.CreateUpdateAgentSnapshotMission)
				abilitiesAPI.DELETE("/:id", s.DeleteAgentSnapshotMission)
				abilitiesAPI.GET("configs", s.GetAgentMissionConfigs)
				abilitiesAPI.GET("tokens", s.GetAgentMissionTokens)
			}

			agentAPI.GET("/dashboard", s.GetDashBoardAgent)
			agentAPI.GET("/dashboard/:token_address", s.GetDashBoardAgentDetail)

			agentAPI.GET("/token-info/:id", s.GetTokenInfoByContract)

			// dojo
			agentAPI.GET("/dojo/list", s.GetListAgentForDojo)
			agentAPI.GET("/dojo/:id", s.GetAgentDetailByAgentIDForDojo)
			agentAPI.GET("/dojo/:id/knowledge-base", s.listKnowledgeByAgent)
			agentAPI.POST("/create_agent_assistant", s.authCheckTK1TokenMiddleware(), s.AgentCreateAgentAssistant)
			agentAPI.POST("/update_agent_assistant", s.authCheckTK1TokenMiddleware(), s.AgentUpdateAgentAssistant)
			agentAPI.POST("/install", s.authCheckTK1TokenMiddleware(), s.MarkInstalledUtilityAgent)
			agentAPI.POST("/recent-chat", s.authCheckTK1TokenMiddleware(), s.MarkRecentChatUtilityAgent)
			agentAPI.POST("/prompt/:id", s.authCheckTK1TokenMiddleware(), s.MarkPromptCountUtilityAgent)

			agentAPI.POST("/comment/:id", s.authCheckTK1TokenMiddleware(), s.AgentComment)
			agentAPI.GET("/comment/:id", s.GetListAgentComment)

			agentAPI.POST("/public/:id", s.authCheckTK1TokenMiddleware(), s.PublicAgent)
			agentAPI.POST("/like/:id", s.authCheckTK1TokenMiddleware(), s.LikeAgent)
			agentAPI.GET("/like/:id", s.authCheckTK1TokenMiddleware(), s.CheckAgentLiked)

			agentAPI.POST("/create-local-agent", s.AgentCreateAgentAssistantForLocal)
			agentAPI.GET("/list-local-agent", s.GetListAgentForDojo)

			agentAPI.POST("/create_agent_studio", s.authCheckTK1TokenMiddleware(), s.AgentCreateAgentStudio)
			agentAPI.POST("/update_agent_studio/:id", s.authCheckTK1TokenMiddleware(), s.AgentUpdateAgentStudio)

			agentAPI.POST("/update_agent_assistant_in_contract", s.authCheckTK1TokenMiddleware(), s.AgentUpdateAgentAssistantInContract)
			agentAPI.POST("/update_twin_status", s.UpdateTwinStatus)

			// knowledge base
			agentAPI.POST("/use-knowledge-base", s.authCheckTK1TokenMiddleware(), s.AgentUseKnowledgeBase)

			// infer
			agentAPI.POST("/async-batch-prompt", s.AsyncBatchPrompt)
			agentAPI.GET("/get-async-prompt-output/:id", s.GetBatchItem)
			//
			agentAPI.GET("/:id/install-code", s.authCheckTK1TokenMiddleware(), s.GetAgentStoreInstallCode)
			agentAPI.GET("/install/info", s.GetAgentInfoInstallInfo)
			//
			agentAPI.GET("/library", s.GetAgentLibrary)
			agentAPI.POST("/add-library", s.AddAgentLibrary)
			agentAPI.GET("/check-exist", s.CheckNameExist)

			agentAPI.GET("/video", s.GetListUserVideo)
			agentAPI.GET("/video/user", s.GetVideoUserInfo)
			agentAPI.POST("/video/export", s.authCheckTK1TokenMiddleware(), s.ExportUserPrivateKeyForClaimVideoReward)

		}

		nftWithAuthAPI := rootAPI.Group("/nft", s.authCheckTK1TokenMiddleware())
		{
			nftWithAuthAPI.GET("/collection", s.GetNftOpenseaCollections)
			nftWithAuthAPI.GET("/collection/:address", s.GetNftCollectionsDetail)
			nftWithAuthAPI.GET("/collection/:address/:tokenId", s.GetNftCollectionsByTokenID)
		}

		serviceAPI := rootAPI.Group("/service")
		{
			serviceAPI.POST("/light-house/upload", s.authCheckTK1TokenMiddleware(), s.UploadDataToLightHouse)
		}

		internalAPI := rootAPI.Group("/internal", s.internalApiMiddleware())
		{
			internalAPI.GET("/webtext", s.GetWebpageText)
			internalAPI.POST("/trigger_job", s.AdminTriggerJob)
			internalAPI.GET("/handle_video", s.AdminHandleVideo)

			twitterAPI := internalAPI.Group("/twitter")
			{
				twitterAPI.GET("/user/recent-info", s.GetTwitterDataForLaunchpad)
				twitterAPI.GET("/user/:id", s.GetTwitterUserByID)
				twitterAPI.GET("/user", s.GetTwitterUserByIDByQuery)

				twitterAPI.GET("/user/by/username/:username", s.GetTwitterUserByUserName)
				twitterAPI.GET("/user/by/username", s.GetTwitterUserByUserNameByQuery)
				twitterAPI.GET("/user/by", s.GetTwitterUserByQuery)

				twitterAPI.GET("/user/:id/following", s.GetTwitterUserFollowing)
				twitterAPI.GET("/user/following", s.GetTwitterUserFollowingByQuery)

				twitterAPI.GET("/tweets/:id", s.GetListUserTweets)
				twitterAPI.GET("/tweets/:id/all", s.GetListUserTweetsAll)
				twitterAPI.GET("/tweets-by-id", s.GetListUserTweetsByQuery)

				twitterAPI.GET("/tweets", s.LookupUserTweets)
				twitterAPI.GET("/tweets/v1", s.LookupUserTweetsV1)
				twitterAPI.GET("/user/:id/mentions", s.GetListUserMentions)
				twitterAPI.GET("/user/mentions", s.GetListUserMentionsByQuery)

				twitterAPI.GET("/user/by/username/:username/following", s.GetTwitterUserFollowingByUsername)
				twitterAPI.GET("/user/by/username/following", s.GetTwitterUserFollowingByUsernameByQuery)

				twitterAPI.GET("/tweets/by/username/:username", s.GetListUserTweetsByUserName)
				twitterAPI.GET("/tweets/by/username", s.GetListUserTweetsByUserNameByQuery)

				// get tweets by list users for mission
				twitterAPI.GET("/tweets/by/users", s.GetListUserTweetsByUsersForTradeMission)
				twitterAPI.GET("/tweets/by/agent", s.GetListUserTweetsByAgentForTradeMission)

				twitterAPI.GET("/user/by/username/:username/mentions", s.GetListUserMentionsByUsername)
				twitterAPI.GET("/user/by/username/:username/mentions/all", s.GetAllUserMentionsByUsername)
				twitterAPI.GET("/user/by/username/mentions", s.GetListUserMentionsByUsernameByQuery)

				twitterAPI.GET("/tweets/search/recent", s.SearchRecentTweet)
				twitterAPI.GET("/tweets/search/token", s.SearchTokenTweet)
				twitterAPI.GET("/user/search", s.SearchUsers)

				twitterAPI.GET("/user/liked", s.GetUser3700Liked)
				//
				twitterAPI.POST("/user/action", s.CreateAgentInternalAction)
				//
				twitterAPI.POST("/wallet/pumfun/create-token/:chain_id/:agent_contract_id", s.AgentWalletCreatePumpFunMeme)
				twitterAPI.POST("/wallet/pumfun/trade-token/:chain_id/:agent_contract_id", s.AgentWalletTradePumpFunMeme)
				twitterAPI.POST("/wallet/raydium/trade-token/:chain_id/:agent_contract_id", s.AgentWalletTradeRaydiumToken)
				twitterAPI.GET("/wallet/solana/balances/:chain_id/:agent_contract_id", s.AgentWalletGetSolanaTokenBalances)
				twitterAPI.GET("/wallet/solana/trades/:chain_id/:agent_contract_id", s.GetAgentWalletSolanaTrades)
				twitterAPI.GET("/wallet/solana/pnls/:chain_id/:agent_contract_id", s.AgentWalletGetSolanaTokenPnls)
				twitterAPI.GET("/wallet/pumfun/trades/:mint", s.GetPumpFunTrades)
				twitterAPI.GET("/wallet/pumfun/price/:mint", s.GetPumpFunTokenPrice)
				twitterAPI.GET("/wallet/solana/quote/latest", s.GetTokenQuoteLatestForSolana)

				twitterAPI.POST("/user/tweet-by-token", s.TweetByToken)
				twitterAPI.POST("/user/action-by-ref", s.CreateAgentInternalActionByRefID)
			}

			// Token management
			tokenAPI := internalAPI.Group("/trade")
			{
				tokenAPI.GET("/tokens", s.GetAgentTradeTokens)
				tokenAPI.GET("/search", s.DexSearchPair)
				tokenAPI.GET("/trade-history/latest", s.DexPairsTradeLatest)
				tokenAPI.GET("/solana/chart-24h/:mint", s.GetSolanaDataChart24Hour)
				tokenAPI.GET("/dexscreen-info", s.DexScreenInfo)
				tokenAPI.GET("/analytic", s.GetTradeAnalytic)
			}
			// launchpad management
			launchpadAPI := internalAPI.Group("/launchpad")
			{
				launchpadAPI.POST("/:id/tier/:member_id", s.ExecuteLaunchpadTier)
			}

			robotAPI := internalAPI.Group("/robot")
			{
				robotAPI.POST("/wallet", s.GenerateRobotSaleWallet)
				robotAPI.GET("/wallet", s.GetRobotSaleWallet)
				robotAPI.GET("/project", s.GetRobotProject)
				robotAPI.POST("/token", s.RobotCreateToken)
				robotAPI.POST("/transfer", s.RobotTransferToken)
				robotAPI.GET("/leaderboards", s.GetRobotProjectLeaderBoards)
			}
		}
		// deprecated
		externalWalletAPI := rootAPI.Group("/external-wallet")
		{
			externalWalletAPI.POST("/address", s.ExternalWalletCreateSOL)
			externalWalletAPI.GET("/address", s.ExternalWalletGet)
			externalWalletAPI.GET("/balances", s.ExternalWalletBalances)
			externalWalletAPI.POST("/compute-order", s.ExternalWalletComputeOrder)
			externalWalletAPI.POST("/create-order", s.ExternalWalletCreateOrder)
			externalWalletAPI.GET("/list-orders", s.ExternalWalletGetOrders)
			externalWalletAPI.GET("/list-tokens", s.ExternalWalletGetTokens)
			externalWalletAPI.GET("/tweets/:username", s.GetListUserTweetsByUserName)
		}

		subcriptionAPI := rootAPI.Group("/subscription")
		{
			subcriptionAPI.POST("/use-api-token", s.internalApiMiddleware(), s.CreateApiTokenUsage)
			subcriptionAPI.POST("/refund-api-token", s.internalApiMiddleware(), s.CreateApiTokenUsage)
			subcriptionAPI.GET("/usages", s.GetApiUsage)
			subcriptionAPI.GET("/packages", s.GetApiPackages)
			subcriptionAPI.GET("/info", s.GetApiSubscriptionInfo)
		}

		bubbleAPI := rootAPI.Group("/bubble")
		{
			bubbleAPI.GET("/list", s.GetListBubbleCrypto)
		}

		memeAPI := rootAPI.Group("/meme")
		{
			// memeAPI.GET("/configs", s.GetMemeConfigs)
			memeAPI.GET("/memes-by-address", s.GetListMemeByAddress)

			memeAPI.GET("/list", s.GetListMemeReport)
			memeAPI.GET("/feed", s.GetFeedMemeReport)
			memeAPI.GET("/:id", s.GetMemeDetail)
			memeAPI.GET("/seen/:id", s.SeenMemeDetail)
			memeAPI.GET("/story", s.GenerateMemeStory)
			memeAPI.GET("/chart/:id", s.GetMemeCandleChart)
			memeAPI.GET("/trade-history", s.GetMemeTradeHistory)
			memeAPI.GET("/trade-history/:id/latest", s.GetMemeTradeHistoryLatest)
			memeAPI.GET("/holders", s.GetMemeListHolders)
			memeAPI.GET("/holding", s.GetMemeListHolding)
			memeAPI.GET("/whitelist-address", s.GetMemeWhiteListAddress)
			memeAPI.GET("/burn-history", s.GetMemeBurnHistory)
			memeAPI.POST("/share/:id", s.ShareMeme)

			threadAPI := memeAPI.Group("/thread")
			{

				threadAPI.GET("/list/:id/latest", s.GetListMemeThreadLatest)
				threadAPI.POST("/hide/:id", s.HideMemeThread)
			}

			userAPI := memeAPI.Group("/user")
			{
				userAPI.GET("/profile", s.GetMemeUserProfile)
				userAPI.GET("/validate-follow", s.MemeValidatedFollow)
				userAPI.POST("/follow", s.FollowUser)
				userAPI.POST("/unfollow", s.UnFollowUser)
				userAPI.GET("/followers", s.GetListFollowers)
				userAPI.GET("/following", s.GetListFollowings)

				userAPI.GET("/notification", s.GetMemeNotification)
				userAPI.GET("/notification/latest", s.GetMemeNotificationLatest)
				userAPI.POST("/notification/seen/:id", s.UserSeenMemeNotification)
			}
		}
		knowledgeHookApi := rootAPI.Group("/knowledge")
		{
			knowledgeHookApi.POST("/webhook", s.webhookKnowledge)
			knowledgeHookApi.POST("/webhook-file/:id", s.webhookKnowledgeFile)
		}

		knowledgeApi := rootAPI.Group("/knowledge", s.authCheckTK1TokenMiddleware())
		{
			knowledgeApi.POST("", s.createKnowledge)
			knowledgeApi.GET("", s.listKnowledge)
			knowledgeApi.PATCH("/:id", s.updateKnowledge)
			knowledgeApi.GET("/:id", s.detailKnowledge)
			knowledgeApi.DELETE("/:id", s.deleteKnowledge)
			knowledgeApi.POST("/:id/check_balance", s.checkBalance)
			knowledgeApi.POST("/update-with-signature", s.updateKnowledgeBaseInContractWithSignature)
		}

		knowledgeBasePublicApi := rootAPI.Group("/knowledge")
		{
			knowledgeBasePublicApi.POST("/retrieve", s.retrieveKnowledge)
		}

		infraTwitterApp := rootAPI.Group("/infra-twitter-app")
		{
			// infraTwitterApp.GET("/install", s.InfraTwitterAppAuthenInstall)
			infraTwitterApp.GET("/callback", s.InfraTwitterAppAuthenCallback)
			infraTwitterApp.GET("/tweets/search", s.middlewareApiLimit(3, 1*time.Second), s.InfraTwitterAppSearchRecentTweet)
		}

		utilityApi := rootAPI.Group("/utility", s.authCheckSignatureMiddleware())
		// utilityApi := rootAPI.Group("/utility")
		{
			utilityApi.GET("/infra-twitter/info", s.GetInfraTwitterAppInfo)
			utilityApi.POST("/twitter/post", s.UtilityPostTwitter)
			// utilityApi.POST("/twitter/verify-deposit", s.UtilityTwitterVerifyDeposit)
		}

		vibeApi := rootAPI.Group("/vibe")
		{
			vibeApi.POST("/validate-ref-code", s.VibeValidateReferralCode)
			vibeApi.GET("/dashboard", s.GetVibeDashBoards)
			vibeApi.GET("/dashboard/:agent_id", s.GetVibeDashBoardsDetail)
			vibeApi.POST("/deploy-token/:agent_id", s.authCheckTK1TokenMiddleware(), s.VibeTokenGetDeployInfo)
		}
	}
}
