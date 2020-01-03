package types

//go:generate ffjson $GOFILE

type ImmutableChainParametersType struct {
	MinCommitteeMemberCount UInt16 `json:"min_committee_member_count"`
	MinWitnessCount         UInt16 `json:"min_witness_count"`
}

type ChainProperty struct {
	ID                  ChainPropertyID              `json:"id"`
	ChainID             string                       `json:"chain_id"`
	ImmutableParameters ImmutableChainParametersType `json:"immutable_parameters"`
	BaseContract        string                       `json:"base_contract"`
}
