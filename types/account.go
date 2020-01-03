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
}

type LockWithVotePairType struct {
	Number UInt32
	Asset  AssetAmount
}
