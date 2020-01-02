package types

//go:generate ffjson $GOFILE

type BlockHeader struct {
	Previous              Buffer     `json:"previous"`
	TimeStamp             Time       `json:"timestamp"`
	Witness               WitnessID  `json:"witness"`
	TransactionMerkleRoot Buffer     `json:"transaction_merkle_root"`
	Extensions            Extensions `json:"extensions"`
}

type Block struct {
	Previous              Buffer                      `json:"previous"`
	TimeStamp             Time                        `json:"timestamp"`
	Witness               WitnessID                   `json:"witness"`
	TransactionMerkleRoot Buffer                      `json:"transaction_merkle_root"`
	WitnessSignature      Buffer                      `json:"witness_signature"`
	BlockID               Buffer                      `json:"block_id"`
	Transactions          SignedTransactionsWithTrxID `json:"transactions"`
	Extensions            Extensions                  `json:"extensions"`
}
