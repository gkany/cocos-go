package types

//go:generate ffjson $GOFILE

type Accounts []Account

func (p Accounts) Lookup(ID GrapheneObject) *Account {
	for _, acct := range p {
		if acct.ID.Equals(ID) {
			return &acct
		}
	}

	return nil
}

type Account struct {
	ID                       AccountID              `json:"id"`
	MembershipExpirationDate Time                   `json:"membership_expiration_date"`
	Registrar                AccountID              `json:"registrar"`
	Name                     String                 `json:"name"`
	Owner                    Authority              `json:"owner"`
	Active                   Authority              `json:"active"`
	Options                  AccountOptions         `json:"options"`
	Statistics               ObjectID               `json:"statistics"`
	CashbackGAS              VestingBalanceID       `json:"cashback_gas"`
	CashbackVB               VestingBalanceID       `json:"cashback_vb"`
	CashbackVote             VestingBalanceID       `json:"cashback_vote"`
	OwnerSpecialAuthority    OwnerSpecialAuthority  `json:"owner_special_authority"`
	ActiveSpecialAuthority   ActiveSpecialAuthority `json:"active_special_authority"`
	TopNControlFlags         UInt64                 `json:"top_n_control_flags"`

	//TODO
	// witness_status
	// committee_status
	// asse_locked

	// NetworkFeePercentage          UInt64                 `json:"network_fee_percentage"`
	// LifetimeReferrerFeePercentage UInt64                 `json:"lifetime_referrer_fee_percentage"`
	// ReferrerRewardsPercentage     UInt64                 `json:"referrer_rewards_percentage"`
	// WhitelistingAccounts          AccountIDs             `json:"whitelisting_accounts"`
	// BlacklistingAccounts          AccountIDs             `json:"blacklisting_accounts"`
	// WhitelistedAccounts           AccountIDs             `json:"whitelisted_accounts"`
	// BlacklistedAccounts           AccountIDs             `json:"blacklisted_accounts"`
	// Referrer                      AccountID              `json:"referrer"`
	// LifetimeReferrer              AccountID              `json:"lifetime_referrer"`
}
