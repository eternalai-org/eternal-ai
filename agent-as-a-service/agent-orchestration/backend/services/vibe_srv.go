package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

func (s *Service) AddVibeWhiteList(ctx context.Context, email string) error {
	whiteList := &models.VibeWhiteList{
		Email: email,
	}
	return s.dao.Create(daos.GetDBMainCtx(ctx), whiteList)
}

func (s *Service) ValidateReferralCode(ctx context.Context, refCode, userAddress string) (bool, error) {
	if strings.EqualFold(refCode, "eaieai") {
		return true, nil
	}

	//check if refCode is valid
	vibeRefcode, err := s.dao.FirstVibeReferralCode(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{"ref_code = ?": {refCode}},
		nil, nil)
	if err != nil {
		return false, errs.NewError(err)
	}

	if vibeRefcode == nil {
		return false, errs.NewError(errs.ErrBadRequest)
	}

	if vibeRefcode.Used && !strings.EqualFold(vibeRefcode.UserAddress, userAddress) {
		return false, errs.NewError(errs.ErrReferralCodeUsed)
	}

	vibeRefcode, _ = s.dao.FirstVibeReferralCodeByID(daos.GetDBMainCtx(ctx), vibeRefcode.ID, nil, true)
	vibeRefcode.Used = true
	vibeRefcode.UserAddress = strings.ToLower(userAddress)
	err = s.dao.Save(daos.GetDBMainCtx(ctx), vibeRefcode)
	if err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (s *Service) GetVibeDashboard(ctx context.Context,
	includeHidden *bool,
	contractAddresses []string,
	userAddress string, networkID uint64, agentTypes []int,
	search string,
	ids, exludeIds []uint,
	categoryIds []string, sortListStr []string, page, limit int) ([]*models.AgentInfo, uint, error) {
	selected := []string{
		`ifnull(agent_infos.reply_latest_time, agent_infos.updated_at) reply_latest_time`,
		"agent_infos.*",
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), ifnull(agent_token_infos.price_change,0)) meme_percent`,
		`cast(case when ifnull(agent_token_infos.usd_market_cap, 0) then ifnull(agent_token_infos.usd_market_cap, 0)
		when ifnull(memes.price_usd*memes.total_suply, 0) > 0 then ifnull(memes.price_usd*memes.total_suply, 0) end as decimal(36, 18)) meme_market_cap`,
		`ifnull(memes.price_usd, agent_token_infos.price_usd) meme_price`,
		`ifnull(memes.volume_last24h*memes.price_usd, agent_token_infos.volume_last24h) meme_volume_last24h`,
	}
	joinFilters := map[string][]any{
		`
			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
		`: {},
	}

	filters := map[string][]any{
		`	
			agent_infos.agent_contract_address is not null and agent_infos.agent_contract_address != ""
		`: {},
	}

	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(agent_infos.token_name) like ? 
			or LOWER(agent_infos.token_symbol) like ? 
			or LOWER(agent_infos.token_address) like ?
			or LOWER(agent_infos.twitter_username) like ?
			or LOWER(agent_infos.agent_name) like ?
			or LOWER(agent_infos.display_name) like ? 
		`] = []any{search, search, search, search, search, search}
	}

	//filter agent type
	if len(agentTypes) > 0 {
		filters["agent_infos.agent_type in (?)"] = []any{agentTypes}
	} else {
		filters["agent_infos.agent_type in (?)"] = []any{[]models.AgentInfoAgentType{
			models.AgentInfoAgentTypeModel,
			models.AgentInfoAgentTypeModelOnline,
			models.AgentInfoAgentTypeJs,
			models.AgentInfoAgentTypePython,
			models.AgentInfoAgentTypeInfa,
			models.AgentInfoAgentTypeCustomUi,
			models.AgentInfoAgentTypeCustomPrompt,
		}}
	}

	//filter contract address
	if len(contractAddresses) > 0 {
		filters["agent_infos.agent_contract_address in (?)"] = []any{contractAddresses}
	}

	if networkID == models.SOLANA_CHAIN_ID_OLD {
		networkID = models.SOLANA_CHAIN_ID
	}
	if networkID > 0 {
		if networkID == models.SHARDAI_CHAIN_ID || networkID == models.SOLANA_CHAIN_ID {
			filters["agent_infos.network_id = ? or agent_infos.token_network_id = ? or agent_infos.id = 763"] = []any{networkID, networkID}
		} else if networkID == models.HERMES_CHAIN_ID {
			filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
			filters["agent_infos.id != 763"] = []any{}
		} else if networkID == models.BASE_CHAIN_ID {
			filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
			filters["agent_infos.id != 763"] = []any{}
		} else if networkID == models.ETHEREUM_CHAIN_ID {
			listEtherModels := []uint64{models.ETHEREUM_CHAIN_ID, models.BASE_CHAIN_ID, models.APE_CHAIN_ID, models.ABSTRACT_TESTNET_CHAIN_ID, models.ARBITRUM_CHAIN_ID}
			filters["agent_infos.network_id in (?) or agent_infos.token_network_id in (?)"] = []any{listEtherModels, listEtherModels}
		} else {
			filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
		}
	}

	if len(ids) > 0 {
		filters["agent_infos.id in (?)"] = []any{ids}
	} else {
		if includeHidden == nil || !(*includeHidden) {
			filters["agent_infos.is_public = 1"] = []any{}
		}
	}

	if len(exludeIds) > 0 {
		filters["agent_infos.id not in (?)"] = []any{exludeIds}
	}

	if len(categoryIds) > 0 {
		filters["agent_infos.agent_category_id in (?)"] = []any{categoryIds}
	}

	if userAddress != "" {
		selected = append(selected, "ifnull(agent_utility_recent_chats.updated_at, now() - interval 100 day) recent_chat_time")
		joinFilters = map[string][]any{
			`
			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
			left join agent_utility_recent_chats on agent_utility_recent_chats.agent_info_id = agent_infos.id
					and agent_utility_recent_chats.address = ?
		`: {strings.ToLower(userAddress)},
		}
	} else {
		selected = append(selected, "now() recent_chat_time")
	}

	sortDefault := "installed_count desc"
	if len(sortListStr) > 0 {
		sortDefault = strings.Join(sortListStr, ", ")
	}

	agents, err := s.dao.FindAgentInfoJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"Meme":          {`deleted_at IS NULL and status not in ("created", "pending")`},
			"TokenInfo":     {},
			"AgentCategory": {},
		},
		[]string{sortDefault},
		page, limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return agents, 0, nil
}

func (s *Service) GetVibeDashboardsDetail(ctx context.Context, agentID string) (*models.AgentInfo, error) {
	selected := []string{
		`ifnull(agent_infos.reply_latest_time, agent_infos.updated_at) reply_latest_time`,
		"agent_infos.*",
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), ifnull(agent_token_infos.price_change,0)) meme_percent`,
		`cast(case when ifnull(agent_token_infos.usd_market_cap, 0) then ifnull(agent_token_infos.usd_market_cap, 0)
		when ifnull(memes.price_usd*memes.total_suply, 0) > 0 then ifnull(memes.price_usd*memes.total_suply, 0) end as decimal(36, 18)) meme_market_cap`,
		`ifnull(memes.price_usd, agent_token_infos.price_usd) meme_price`,
		`ifnull(memes.volume_last24h*memes.price_usd, agent_token_infos.volume_last24h) meme_volume_last24h`,
	}

	joinFilters := map[string][]any{
		`
			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
		`: {},
	}

	filters := map[string][]any{
		`	
			agent_infos.agent_id = ? or agent_infos.id = ? or (agent_infos.token_address = ? and agent_infos.token_address != "")
		`: {agentID, agentID, agentID},
	}

	agents, err := s.dao.FindAgentInfoJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"Meme":          {`deleted_at IS NULL and status not in ("created", "pending")`},
			"TokenInfo":     {},
			"AgentCategory": {},
		},
		[]string{},
		1, 1,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if len(agents) == 0 {
		return nil, errs.NewError(errs.ErrAgentNotFound)
	}
	return agents[0], nil
}
