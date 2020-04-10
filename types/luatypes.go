package types

//go:generate ffjson $GOFILE

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

// LuaType ... lua type index number
const (
	LuaTypeInt      = iota // 0
	LuaTypeNumber          // 1
	LuaTypeString          // 2
	LuaTypeBool            // 3
	LuaTypeTable           // 4
	LuaTypeFunction        // 5
	// LuaTypeUserData
)

// LuaKey ... lua key index number
const (
	LuaKeyInt = iota
	LuaKeyNumber
	LuaKeyString
	LuaKeyBool
)

// LuaInt ... type LuaInt uint64
type LuaInt struct {
	V uint64 `json:"v"`
}

// LuaNumber ... float64
type LuaNumber struct {
	V float64 `json:"v"`
}

// LuaString ... string
type LuaString struct {
	V string `json:"v"`
}

// LuaBool ... bool
type LuaBool struct {
	V bool `json:"v"`
}

// LuaTypeItem ... [LuaKey, LuaType]
type LuaTypeItem []interface{}

// Marshal ...
func (o LuaTypeItem) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("[test] lua map item Marshal")
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua map item type length error. %v", len(o))
	}

	if err := enc.Encode(o[0]); err != nil {
		return errors.Annotate(err, "encode item key")
	}

	if err := enc.Encode(o[1]); err != nil {
		return errors.Annotate(err, "encode item value")
	}

	return nil
}

// LuaKey ... int number string bool
func parseLuaKey(key LuaKey) (interface{}, error) {
	k1 := key.Key[0].(float64)
	k2 := key.Key[1].(map[string]interface{})
	k2v := k2["v"] // basic type

	var tempKey interface{}
	switch k1 {
	case LuaTypeInt, LuaTypeNumber:
		switch k2v.(type) {
		case float64:
			tempKey = k2v.(float64)
		case string:
			kv, err := strconv.ParseFloat(k2v.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", k2v)
				return tempKey, err
			}
			tempKey = kv
		}
	case LuaTypeString:
		tempKey = k2v.(string)
	case LuaTypeBool:
		tempKey = k2v.(bool)
	default:
		fmt.Println("UnKnow LuaKey type. ", fmt.Sprintf("%T", k1))
		return nil, errors.Errorf("UnKnow LuaKey type, type index: %v", k1)
	}
	return tempKey, nil
}

// LuaType ... basic type: int number string bool, complex: LuaMap, LuaTable, LuaFunction, maybe include basic and complex type
func parseLuaType(value []interface{}, output *[]interface{}) error {
	var tempValue interface{}

	v1 := value[0].(float64)
	v2 := value[1].(map[string]interface{})
	v2v := v2["v"] // basic type

	switch v1 {
	case LuaTypeInt, LuaTypeNumber:
		switch v2v.(type) {
		case float64:
			tempValue = v2v.(float64)
		case string:
			// tempValue = v2v.(string) // 强转成float64
			vv, err := strconv.ParseFloat(v2v.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", v2v)
				return err
			}
			tempValue = vv
		}
	case LuaTypeString:
		tempValue = v2v.(string)
	case LuaTypeBool:
		tempValue = v2v.(bool)
	case LuaTypeFunction:
		isVarArg := v2["is_var_arg"].(bool)
		argList := v2["arglist"].([]interface{})
		temp := []string{}
		for _, val := range argList {
			temp = append(temp, val.(string))
		}
		tempValue = LuaFunction{IsVarArg: isVarArg, ArgList: temp}
	case LuaTypeTable:
		tmap := value[1].(map[string]interface{})
		for _, tv := range tmap {
			tvList := tv.([]interface{})
			for _, tvv := range tvList {
				for _, tvvv := range tvv.([]interface{}) {
					switch tvvv.(type) {
					case []interface{}:
						err := parseLuaType(tvvv.([]interface{}), output)
						if err != nil {
							return err
						}
					case map[string]interface{}:
						tvvvMap := tvv.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
						tm0 := tvvvMap[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
						tm1 := tvvvMap[1]

						tm0Map := tm0.(map[string]interface{})
						value := tm0Map["key"] // 二元数组： 根据类型解析
						err := parseLuaType(value.([]interface{}), output)
						if err != nil {
							fmt.Printf("** [error] tmoResult: %v\n", err)
							return err
						}

						err = parseLuaType(tm1.([]interface{}), output)
						if err != nil {
							fmt.Printf("** [error] tm1Result: %v\n", err)
							return err
						}
					}
				}
			}
		}
	default:
		fmt.Println("UnKnow LuaType type. ", fmt.Sprintf("%T", v1))
		return errors.Errorf("UnKnow LuaType type, type index: %v", v1)
	}
	*output = append(*output, tempValue)
	return nil
}

// UnmarshalJSON ...
func (p *LuaTypeItem) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	var key LuaKey
	if err := ffjson.Unmarshal(raw[0], &key); err != nil {
		return errors.Annotate(err, "unmarshal key")
	}

	var value LuaType
	if err := ffjson.Unmarshal(raw[1], &value); err != nil {
		return errors.Annotate(err, "unmarshal value")
	}

	// keyResult, err := parseLuaKey(key)
	// if err != nil {
	// 	return err
	// }
	var keyResult []interface{}
	err := parseLuaType(key.Key, &keyResult)
	if err != nil {
		return err
	}

	var result []interface{}
	err = parseLuaType(value, &result)
	if err != nil {
		return err
	}

	itemResult := make(LuaTypeItem, 2)
	itemResult[0] = keyResult
	switch len(result) {
	case 1:
		itemResult[1] = result[0]
	default:
		itemResult[1] = result
	}
	*p = itemResult

	return nil
}

// LuaKey ... eg: {"key": [2, {"v": "lua string test"}]}
type LuaKey struct {
	Key LuaKeyObjectType `json:"key"`
}

type LuaMap []LuaTypeItem

// LuaTable ...
type LuaTable struct {
	Value LuaMap `json:"v"`
}

// Marshal ...
// func (o LuaTable) Marshal(enc *util.TypeEncoder) error {
// 	if err := enc.Encode(o.Value); err != nil {
// 		return errors.Annotate(err, "encode lua map value")
// 	}
// 	return nil
// }

// func (o *LuaTable) Unmarshal(data []byte) error {
// 	if err := ffjson.Unmarshal(data, o); err != nil {
// 		return errors.Annotate(err, "unmarshal LuaTable object")
// 	}

// 	return nil
// }

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

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{}

// LuaKeyObjectType ...
type LuaKeyObjectType []interface{}
