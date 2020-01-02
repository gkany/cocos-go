package operations

//go:generate ffjson $GOFILE

import (
	"fmt"

	"github.com/gkany/graphSDK/types"
	"github.com/gkany/graphSDK/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeTransfer] = func() types.Operation {
		op := &TransferOperation{}
		return op
	}
}

type TransferOperation struct {
	types.OperationFee
	From       types.AccountID   `json:"from"`
	To         types.AccountID   `json:"to"`
	Amount     types.AssetAmount `json:"amount"`
	Memo       types.MemoPair    `json:"memo,omitempty"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p TransferOperation) Type() types.OperationType {
	return types.OperationTypeTransfer
}

func (p TransferOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p TransferOperation) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("TransferOperation Marshal")
	fmt.Println("  ->type: ", p.Type())
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	fmt.Println("  ->Fee: ", p.Fee)
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	fmt.Println("  ->From: ", p.From)
	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}

	fmt.Println("  ->To: ", p.To)
	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}

	fmt.Println("  ->Amount: ", p.Amount)
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}

	fmt.Println("  ->p.Memo != nil: ", p.Memo != nil)
	if err := enc.Encode(p.Memo != nil); err != nil {
		return errors.Annotate(err, "encode have Memo")
	}

	fmt.Println("  ->Memo: ", p.Memo)
	if err := enc.Encode(p.Memo); err != nil {
		return errors.Annotate(err, "encode memo")
	}

	fmt.Println("  ->Extensions[transfer op]: ", p.Extensions)
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}

func (p *TransferOperation) SetFee(fee types.AssetAmount) {
	p.OperationFee.SetFee(fee)
}

func (p TransferOperation) GetFee() types.AssetAmount {
	return p.OperationFee.GetFee()
}
