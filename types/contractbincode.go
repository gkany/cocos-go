package types

//go:generate ffjson $GOFILE

type ContractBinCode struct {
	ID         ContractBinCodeID `json:"id"`
	ContractID ContractID        `json:"contract_id"`
	LuaCodeB   CharListType      `json:"lua_code_b"`
}
