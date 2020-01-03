package types

//go:generate ffjson $GOFILE

// type FeeSchedule struct {
// 	Parameters         FeeScheduleParameters `json:"parameters"`
// 	Scale              UInt32                `json:"scale"`
// 	MaximunHandlingFee Int64                 `json:"maximun_handling_fee"`
// }

type ChainParametersType struct {
	CurrentFees                          *FeeSchedule    `json:"current_fees"`
	BlockInterval                        UInt8           `json:"block_interval"`
	maintenance_interval                 uint32_t        `json:"maintenance_interval             "`
	maintenance_skip_slots               uint8_t         `json:"maintenance_skip_slots           "`
	committee_proposal_review_period     uint32_t        `json:"committee_proposal_review_period "`
	maximum_block_size                   uint32_t        `json:"maximum_block_size               "`
	maximum_time_until_expiration        uint32_t        `json:"maximum_time_until_expiration    "`
	maximum_proposal_lifetime            uint32_t        `json:"maximum_proposal_lifetime        "`
	maximum_asset_feed_publishers        uint8_t         `json:"maximum_asset_feed_publishers    "`
	witness_number_of_election           uint16_t        `json:"witness_number_of_election       "`
	committee_number_of_election         uint16_t        `json:"committee_number_of_election     "`
	maximum_authority_membership         uint16_t        `json:"maximum_authority_membership     "`
	cashback_gas_period_seconds          uint32_t        `json:"cashback_gas_period_seconds      "`
	cashback_vb_period_seconds           uint32_t        `json:"cashback_vb_period_seconds       "`
	cashback_vote_period_seconds         uint32_t        `json:"cashback_vote_period_seconds     "`
	witness_pay_per_block                share_type      `json:"witness_pay_per_block            "`
	witness_pay_vesting_seconds          uint32_t        `json:"witness_pay_vesting_seconds      "`
	worker_budget_per_day                share_type      `json:"worker_budget_per_day            "`
	accounts_per_fee_scale               uint16_t        `json:"accounts_per_fee_scale           "`
	account_fee_scale_bitshifts          uint8_t         `json:"account_fee_scale_bitshifts      "`
	max_authority_depth                  uint8_t         `json:"max_authority_depth              "`
	maximum_run_time_ratio               uint16_t        `json:"maximum_run_time_ratio           "`
	maximum_nh_asset_order_expiration    uint32_t        `json:"maximum_nh_asset_order_expiration"`
	assigned_task_life_cycle             uint32_t        `json:"assigned_task_life_cycle         "`
	crontab_suspend_threshold            uint32_t        `json:"crontab_suspend_threshold        "`
	crontab_suspend_expiration           uint32_t        `json:"crontab_suspend_expiration       "`
	witness_candidate_freeze             share_type      `json:"witness_candidate_freeze         "`
	committee_candidate_freeze           share_type      `json:"committee_candidate_freeze       "`
	candidate_award_budget               share_type      `json:"candidate_award_budget           "`
	committee_percent_of_candidate_award uint16_t        `json:"committee_percent_of_candidate_award"`
	unsuccessful_candidates_percent      uint16_t        `json:"unsuccessful_candidates_percent  "`
	extensions                           extensions_type `json:"extensions"`
}

type GlobalProperty struct {
	ID GlobalPropertyID `json:"id"`
	Parameters

	ChainID             string                       `json:"chain_id"`
	ImmutableParameters ImmutableChainParametersType `json:"immutable_parameters"`
	BaseContract        string                       `json:"base_contract"`
}
