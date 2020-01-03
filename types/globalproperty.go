package types

//go:generate ffjson $GOFILE

type ChainParametersType struct {
	CurrentFees                      *FeeSchedule `json:"current_fees"`
	BlockInterval                    UInt8        `json:"block_interval"`
	MaintenanceInterval              UInt32       `json:"maintenance_interval"`
	MaintenanceSkipSlots             UInt8        `json:"maintenance_skip_slots"`
	CommitteeProposalReviewPeriod    UInt32       `json:"committee_proposal_review_period "`
	MaximumBlockSize                 UInt32       `json:"maximum_block_size"`
	MaximumTimeUntilExpiration       UInt32       `json:"maximum_time_until_expiration"`
	MaximumProposalLifetime          UInt32       `json:"maximum_proposal_lifetime"`
	MaximumAssetFeedPublishers       UInt8        `json:"maximum_asset_feed_publishers"`
	WitnessNumberOfElection          UInt16       `json:"witness_number_of_election"`
	CommitteeNumberOfElection        UInt16       `json:"committee_number_of_election"`
	MaximumAuthorityMembership       UInt16       `json:"maximum_authority_membership"`
	CashbackGasPeriodSeconds         UInt32       `json:"cashback_gas_period_seconds"`
	CashbackVbPeriodSeconds          UInt32       `json:"cashback_vb_period_seconds"`
	CashbackVotePeriodSeconds        UInt32       `json:"cashback_vote_period_seconds"`
	WitnessPayPerBlock               Int64        `json:"witness_pay_per_block"`
	WitnessPayVestingSeconds         UInt32       `json:"witness_pay_vesting_seconds"`
	WorkerBudgetPerDay               Int64        `json:"worker_budget_per_day"`
	AccountsPerFeeScale              UInt16       `json:"accounts_per_fee_scale"`
	AccountFeeScaleBitshifts         UInt8        `json:"account_fee_scale_bitshifts"`
	MaxAuthorityDepth                UInt8        `json:"max_authority_depth"`
	MaximumRunTimeRatio              UInt16       `json:"maximum_run_time_ratio"`
	MaximumNhAssetOrderExpiration    UInt32       `json:"maximum_nh_asset_order_expiration"`
	AssignedTaskLifeCycle            UInt32       `json:"assigned_task_life_cycle"`
	CrontabSuspendThreshold          UInt32       `json:"crontab_suspend_threshold"`
	CrontabSuspendExpiration         UInt32       `json:"crontab_suspend_expiration"`
	WitnessCandidateFreeze           Int64        `json:"witness_candidate_freeze"`
	CommitteeCandidateFreeze         Int64        `json:"committee_candidate_freeze"`
	CandidateAwardBudget             Int64        `json:"candidate_award_budget"`
	CommitteePercentOfCandidateAward UInt16       `json:"committee_percent_of_candidate_award"`
	UnsuccessfulCandidatesPercent    UInt16       `json:"unsuccessful_candidates_percent"`
	Extensions                       Extensions   `json:"extensions"`
}

type GlobalProperty struct {
	ID                     GlobalPropertyID     `json:"id"`
	Parameters             ChainParametersType  `json:"parameters"`
	PendingParameters      *ChainParametersType `json:"pending_parameters"`
	NextAvailableVoteID    UInt32               `json:"next_available_vote_id"`
	ActiveCommitteeMembers CommitteeMemberIDs   `json:"active_committee_members"`
	ActiveWitnesses        WitnessIDs           `json:"active_witnesses"`
}
