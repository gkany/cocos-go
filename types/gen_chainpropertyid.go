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

type ChainPropertyID struct {
	ObjectID
}

func (p ChainPropertyID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *ChainPropertyID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeChainProperty) << 48) | instance)
	return nil
}

type ChainPropertyIDs []ChainPropertyID

func (p ChainPropertyIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode ChainPropertyID")
		}
	}

	return nil
}

func ChainPropertyIDFromObject(ob GrapheneObject) ChainPropertyID {
	id, ok := ob.(*ChainPropertyID)
	if ok {
		return *id
	}

	p := ChainPropertyID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeChainProperty {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeChainProperty'", p.ID()))
	}

	return p
}

//NewChainPropertyID creates an new ChainPropertyID object
func NewChainPropertyID(id string) GrapheneObject {
	gid := new(ChainPropertyID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"ChainPropertyID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeChainProperty {
		logging.Errorf(
			"ChainPropertyID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeChainProperty'", id),
		)
		return nil
	}

	return gid
}
