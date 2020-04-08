package types

import (
	"encoding/json"
	"fmt"

	"github.com/gkany/cocos-go/logging"
	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

// LuaType ... lua type index number
const (
	LuaTypeInt = iota
	LuaTypeNumber
	LuaTypeString
	LuaTypeBool
	LuaTypeTable
	LuaTypeFunction
	// LuaTypeUserData
)

// LuaKey ... lua key index number
const (
	LuaKeyInt = iota
	LuaKeyNumber
	LuaKeyString
	LuaKeyBool
)

// LuaInt ...
type LuaInt uint64
// LuaNumber ...
type LuaNumber float64
// LuaString ...
type LuaString string
// LuaBool ...
type LuaBool bool

// LuaMap ... map[LuaKey]LuaType
type LuaMap map[uint64]LuaType

// LuaTable ...
type LuaTable struct {
	Value LuaMap `json:"v"`
}

// Marshal ...
func (o LuaTable) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(o.Value); err != nil {
		return errors.Annotate(err, "encode lua map value")
	}
	return nil
}

// LuaFunction ... 
type LuaFunction struct {
	IsVarArg bool     `json:"is_var_arg"`
	ArgList  []string `json:"arglist"`
}

// Marshal ...
func (o LuaFunction) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(o.IsVarArg); err != nil {
		return errors.Annotate(err, "encode IsVarArg")
	}

	if err := enc.Encode(o.ArgList); err != nil {
		return errors.Annotate(err, "encode ArgList")
	}
	return nil
}

// LuaObject ... LuaInt, LuaNumber, LuaBool, LuaFunction, LuaString, LuaMap, LuaUserData
type LuaObject struct {
	Value interface{} `json:"v"` 
}

// Marshal ...
func (o LuaObject) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(o.Value); err != nil {
		return errors.Annotate(err, "encode lua type obj")
	}
	return nil
}

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{} 

// Marshal ... 
func (o LuaType) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("[test] lua type object Marshal")
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua type object length error. %v", len(v))
	}

	if err := enc.Encode(o[0]); err != nil {
		return errors.Annotate(err, "encode type index")
	}

	if err := enc.Encode(o[1]); err != nil {
		return errors.Annotate(err, "encode type object")
	}

	return nil
}

// UnmarshalJSON ... 
func (o *LuaType) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	var typeIndex uint64
	if err := ffjson.Unmarshal(raw[0], &typeIndex); err != nil {
		return errors.Annotate(err, "unmarshal lua type index")
	}

	desc := fmt.Sprintf("encrypt type %s", typeIndex)
	if typeIndex == LuaTypeInt {
		var value LuaInt
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua int type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else if typeIndex == LuaTypeNumber {
		var value LuaNumber
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua number type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else if typeIndex == LuaTypeString {
		var value LuaString
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua string type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else if typeIndex == LuaTypeBool {
		var value LuaBool
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua bool type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else if typeIndex == LuaTypeFunction {
		var value LuaFunction
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua function type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else if typeIndex == LuaTypeTable {
		var value LuaTable
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua table type %v", typeIndex)
		}
		o = &LuaType{typeIndex, value}
	} else {
		return errors.Errorf("unmarshal lua type error, typeIndex(%v) unknown", typeIndex)
	}

	return nil
}
