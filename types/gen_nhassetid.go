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

type NHAssetID struct {
	ObjectID
}

func (p NHAssetID) Marshal(enc *util.TypeEncoder) error {
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

func (p *NHAssetID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeNHAsset) << 48) | instance)
	return nil
}

type NHAssetIDs []NHAssetID

func (p NHAssetIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode NHAssetID")
		}
	}

	return nil
}

func NHAssetIDFromObject(ob GrapheneObject) NHAssetID {
	id, ok := ob.(*NHAssetID)
	if ok {
		return *id
	}

	p := NHAssetID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeNHAsset {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeNHAsset'", p.ID()))
	}

	return p
}

//NewNHAssetID creates an new NHAssetID object
func NewNHAssetID(id string) GrapheneObject {
	gid := new(NHAssetID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"NHAssetID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeNHAsset {
		logging.Errorf(
			"NHAssetID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeNHAsset'", id),
		)
		return nil
	}

	return gid
}