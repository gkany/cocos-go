// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/gkany/graphSDK/logging"
	"github.com/gkany/graphSDK/util"
	"github.com/juju/errors"
)

type ContractID struct {
	ObjectID
}

func (p ContractID) Marshal(enc *util.TypeEncoder) error {
	n, err := enc.EncodeUVarintByByte(uint64(p.Instance()))
	if err != nil {
		return errors.Annotate(err, "encode instance")
	}

	for i := 0; i < 8-n; i++ {
		if err := enc.EncodeUVarint(uint64(0)); err != nil {
			return errors.Annotate(err, "encode zero")
		}
	}

	return nil
}

func (p *ContractID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeContract) << 48) | instance)
	return nil
}

type ContractIDs []ContractID

func (p ContractIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode ContractID")
		}
	}

	return nil
}

func ContractIDFromObject(ob GrapheneObject) ContractID {
	id, ok := ob.(*ContractID)
	if ok {
		return *id
	}

	p := ContractID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeContract {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeContract'", p.ID()))
	}

	return p
}

//NewContractID creates an new ContractID object
func NewContractID(id string) GrapheneObject {
	gid := new(ContractID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"ContractID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeContract {
		logging.Errorf(
			"ContractID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeContract'", id),
		)
		return nil
	}

	return gid
}
