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
	LuaKeyInt    = iota // 0
	LuaKeyNumber        // 1
	LuaKeyString        // 2
	LuaKeyBool          // 3
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
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua map item type length error. size: %v", len(o))
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
func parseLuaKey(keyObject LuaKey) (interface{}, error) {
	valueType := keyObject.Key[0].(float64)
	value := keyObject.Key[1].(map[string]interface{})
	valuev := value["v"] // basic type

	var result interface{}
	switch valueType {
	case LuaTypeInt, LuaTypeNumber:
		switch valuev.(type) {
		case float64:
			result = valuev.(float64)
		case string:
			strf, err := strconv.ParseFloat(valuev.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", valuev)
				return result, err
			}
			result = strf
		}
	case LuaTypeString:
		result = valuev.(string)
	case LuaTypeBool:
		result = valuev.(bool)
	default:
		fmt.Println("UnKnow LuaKey type. ", fmt.Sprintf("%T", valueType))
		return nil, errors.Errorf("UnKnow LuaKey type, type index: %v", valueType)
	}
	return result, nil
}

// LuaType ... basic type: int number string bool, complex: LuaMap, LuaTable, LuaFunction, maybe include basic and complex type
func parseLuaType(pair []interface{}, dst *[]interface{}) error {
	var result interface{}

	valueType := pair[0].(float64)
	value := pair[1].(map[string]interface{})
	valuev := value["v"] // basic type

	switch valueType {
	case LuaTypeInt, LuaTypeNumber:
		switch valuev.(type) {
		case float64:
			result = valuev.(float64)
		case string:
			// result = valuev.(string) // 强转成float64
			strf, err := strconv.ParseFloat(valuev.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", valuev)
				return err
			}
			result = strf
		}
	case LuaTypeString:
		result = valuev.(string)
	case LuaTypeBool:
		result = valuev.(bool)
	case LuaTypeFunction:
		isVarArg := value["is_var_arg"].(bool)
		argList := value["arglist"].([]interface{})
		args := []string{}
		for _, arg := range argList {
			args = append(args, arg.(string))
		}
		result = LuaFunction{IsVarArg: isVarArg, ArgList: args}
	case LuaTypeTable:
		vMap := pair[1].(map[string]interface{})
		for _, vMapList := range vMap {
			for _, eList := range vMapList.([]interface{}) {
				for _, element := range eList.([]interface{}) {
					switch element.(type) {
					case []interface{}:
						err := parseLuaType(element.([]interface{}), dst)
						if err != nil {
							return err
						}
					case map[string]interface{}:
						ePair := eList.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
						first := ePair[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
						second := ePair[1]

						firstMap := first.(map[string]interface{})
						value := firstMap["key"] // pair type, parse by type
						err := parseLuaType(value.([]interface{}), dst)
						if err != nil {
							fmt.Printf("parseLuaType error. data: %v, error: %v\n", value, err)
							return err
						}

						err = parseLuaType(second.([]interface{}), dst)
						if err != nil {
							fmt.Printf("parseLuaType error. data: %v, error: %v\n", second, err)
							return err
						}
					}
				}
			}
		}
	default:
		fmt.Println("UnKnow LuaType type. ", fmt.Sprintf("%T", valueType))
		return errors.Errorf("UnKnow LuaType type, type index: %v", valueType)
	}
	*dst = append(*dst, result)
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
	Key LuaType `json:"key"`
}

type LuaMap []LuaTypeItem

// LuaTable ...
type LuaTable struct {
	Value LuaMap `json:"v"`
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

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{}
