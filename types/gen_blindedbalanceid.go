// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/gkany/graphSDK/util"
	"github.com/gkany/graphSDK/util/logging"
	"github.com/juju/errors"
)

type BlindedBalanceID struct {
	ObjectID
}

func (p BlindedBalanceID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *BlindedBalanceID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeBlindedBalance) << 48) | instance)
	return nil
}

type BlindedBalanceIDs []BlindedBalanceID

func (p BlindedBalanceIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode BlindedBalanceID")
		}
	}

	return nil
}

func BlindedBalanceIDFromObject(ob GrapheneObject) BlindedBalanceID {
	id, ok := ob.(*BlindedBalanceID)
	if ok {
		return *id
	}

	p := BlindedBalanceID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeBlindedBalance {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeBlindedBalance'", p.ID()))
	}

	return p
}

//NewBlindedBalanceID creates an new BlindedBalanceID object
func NewBlindedBalanceID(id string) GrapheneObject {
	gid := new(BlindedBalanceID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"BlindedBalanceID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeBlindedBalance {
		logging.Errorf(
			"BlindedBalanceID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeBlindedBalance'", id),
		)
		return nil
	}

	return gid
}
