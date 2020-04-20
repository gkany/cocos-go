package operations

//go:generate ffjson $GOFILE

import (
	"github.com/gkany/cocos-go/types"
	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCallContractFunction] = func() types.Operation {
		op := &CallContractFunction{}
		return op
	}
}

// CallContractFunction ...
type CallContractFunction struct {
	types.OperationFee
	Caller       types.AccountID  `json:"caller"`
	Creator      types.AccountID  `json:"creator"`
	ContractID   types.ContractID `json:"contract_id"`
	FunctionName string           `json:"function_name"`
	ValueList    []interface{}    `json:"value_list"`
	Extensions   types.Extensions `json:"extensions"`
	Amount       float64          `json:"amount"`
}

// Type ...
func (p CallContractFunction) Type() types.OperationType {
	return types.OperationTypeCallContractFunction
}

// MarshalFeeScheduleParams ...
func (p CallContractFunction) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
		}
	}

	return nil
}

// Marshal ...
func (p CallContractFunction) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Caller); err != nil {
		return errors.Annotate(err, "encode Caller")
	}

	if err := enc.Encode(p.Creator); err != nil {
		return errors.Annotate(err, "encode Creator")
	}

	if err := enc.Encode(p.ContractID); err != nil {
		return errors.Annotate(err, "encode ContractID")
	}

	if err := enc.Encode(p.FunctionName); err != nil {
		return errors.Annotate(err, "encode FunctionName")
	}

	if err := enc.Encode(p.ValueList); err != nil {
		return errors.Annotate(err, "encode ValueList")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode Amount")
	}

	return nil
}
