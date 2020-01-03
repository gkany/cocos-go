package types

type Info struct {
	ActiveCommitteeMembers CommitteeMemberIDs `json:"active_committee_members"`
	ActiveWitnesses        WitnessIDs         `json:"active_witnesses"`
	ChainID                string             `json:"chain_id"`
	HeadBlockAge           Time               `json:"head_block_age"`
	HeadBlockID            string             `json:"head_block_id"`
	HeadBlockNum           UInt32             `json:"head_block_num"`
	NextMaintenanceTime    Time               `json:"next_maintenance_time"`
	Participation          string             `json:"participation"`
}
