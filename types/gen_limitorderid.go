// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/denkhaus/logging"
	"github.com/gkany/graphSDK/util"
	"github.com/juju/errors"
)

type LimitOrderID struct {
	ObjectID
}

func (p LimitOrderID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *LimitOrderID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeLimitOrder) << 48) | instance)
	return nil
}

type LimitOrderIDs []LimitOrderID

func (p LimitOrderIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode LimitOrderID")
		}
	}

	return nil
}

func LimitOrderIDFromObject(ob GrapheneObject) LimitOrderID {
	id, ok := ob.(*LimitOrderID)
	if ok {
		return *id
	}

	p := LimitOrderID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeLimitOrder {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeLimitOrder'", p.ID()))
	}

	return p
}

//NewLimitOrderID creates an new LimitOrderID object
func NewLimitOrderID(id string) GrapheneObject {
	gid := new(LimitOrderID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"LimitOrderID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeLimitOrder {
		logging.Errorf(
			"LimitOrderID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeLimitOrder'", id),
		)
		return nil
	}

	return gid
}
