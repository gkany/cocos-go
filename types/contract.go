package types

//go:generate ffjson $GOFILE

type Contract struct {
	ID                     ContractID `json:"id"`
	CreationDate           Time       `json:"creation_date"`
	Owner                  AccountID  `json:"owner"`
	Name                   string     `json:"name"`
	IsRelease              bool       `json:"is_release"`
	CurrentVersion         string     `json:"current_version"`
	CheckContractAuthority bool       `json:"check_contract_authority"`
	ContractAuthority      PublicKey  `json:"contract_authority"`
}
