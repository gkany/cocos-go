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

// LuaMapItem ... [LuaKey, LuaType]
type LuaMapItem []interface{}

// Marshal ...
func (o LuaMapItem) Marshal(enc *util.TypeEncoder) error {
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
			kv := k2v.(float64)
			// fmt.Println("kv is int", kv)
			tempKey = kv
		case string:
			// kv := k2v.(string)
			// fmt.Println("kv is too large int --> string", kv)
			kv, err := strconv.ParseFloat(k2v.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", k2v)
				return tempKey, err
			}
			tempKey = kv
		}
	// case LuaTypeNumber:
	// 	kv := k2v.(float64)
	// 	// fmt.Println("kv is number", kv)
	// 	tempKey = kv
	case LuaTypeString:
		kv := k2v.(string)
		// fmt.Println("kv is string", kv)
		tempKey = kv
	case LuaTypeBool:
		kv := k2v.(bool)
		// fmt.Println("kv is bool", kv)
		tempKey = kv
	default:
		fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", k1))
		// return nil, errors.Annotate("Unknow type")
		return nil, errors.Errorf("UnKnow LuaKey type, type index: %v", k1)
	}
	return tempKey, nil
}

// LuaType ... basic type: int number string bool, complex: LuaMap, LuaTable, LuaFunction, maybe include basic and complex type
func parseLuaType(value []interface{}, output *[]interface{}) error {
	var tempValue interface{}
	// fmt.Printf("[start]sss [parseLuaType] value: %v \n outpt: %v\n", value, output)

	v1 := value[0].(float64)
	// fmt.Printf(">>> [parseLuaType]v1: %v, type: %v; v2:%v, type: %v\n", v1, v1, value[1], value[1])

	v2 := value[1].(map[string]interface{})
	v2v := v2["v"] // basic type

	switch v1 {
	case LuaTypeInt, LuaTypeNumber:
		switch v2v.(type) {
		case float64:
			vv := v2v.(float64)
			// fmt.Println("vv is string", vv)
			tempValue = vv
		case string:
			// vv := v2v.(string) // 强转成int
			vv, err := strconv.ParseFloat(v2v.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", v2v)
				return err
			}
			// fmt.Println("vv is too large int --> string", vv)
			tempValue = vv
		}
	// case LuaTypeNumber:
	// 	vv := v2v.(float64)
	// 	// fmt.Println("vv is number", vv)
	// 	tempValue = vv
	case LuaTypeString:
		vv := v2v.(string)
		// fmt.Println("vv is string", vv)
		tempValue = vv
	case LuaTypeBool:
		vv := v2v.(bool)
		// fmt.Println("vv is bool", vv)
		tempValue = vv
	case LuaTypeFunction:
		// fmt.Printf("[LuaTypeFunction] value[1]: %v, type: %T\n", value[1], value[1])
		isVarArg := v2["is_var_arg"].(bool)
		argList := v2["arglist"].([]interface{})
		temp := []string{}
		for _, val := range argList {
			temp = append(temp, val.(string))
		}
		vv := LuaFunction{IsVarArg: isVarArg, ArgList: temp}
		// fmt.Println("vv is function", vv)
		tempValue = vv
	case LuaTypeTable:
		// fmt.Printf("-- [LuaTypeTable] value[1]: %v, type: %T\n", value[1], value[1])
		tmap := value[1].(map[string]interface{})
		for _, tv := range tmap {
			// temp = append(temp, val.(string))
			// fmt.Printf("---- tk:%v, type:%T;\n", tk, tk)
			// fmt.Printf("---- tv:%v, type:%T\n", tv, tv)
			tvList := tv.([]interface{})
			for _, tvv := range tvList {
				// temp = append(temp, val.(string))
				// fmt.Printf("------ tvv:%v, type:%T\n", tvv, tvv)
				for _, tvvv := range tvv.([]interface{}) {
					// fmt.Printf("======= tvvv:%v, type:%T\n", tvvv, tvvv)
					switch tvvv.(type) {
					case []interface{}:
						// fmt.Println("1111  start")
						// fmt.Printf("data: %v, type: %v\n", tvvv, tvvv)
						err := parseLuaType(tvvv.([]interface{}), output)
						if err != nil {
							return err
						}
					case map[string]interface{}:
						// fmt.Println("\n2222  start")
						// tvvvMap := tvv.(map[string]interface{})
						tvvvMap := tvv.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
						tm0 := tvvvMap[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
						tm1 := tvvvMap[1]

						tm0Map := tm0.(map[string]interface{})
						value := tm0Map["key"] // 二元数组： 根据类型解析
						// fmt.Printf("****** [tm0['key']]value:%v, type:%T\n", value, value)
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
						// fmt.Printf("** tm1Result: %v\n", tm1Result)
						// fmt.Println("2222  end\n")
					}
				}
			}
		}
	default:
		fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", v1))
		return errors.Errorf("UnKnow LuaType type, type index: %v", v1)
	}
	// fmt.Printf("<<< [parseLuaType]parse result: %v ## end\n\n", tempValue)
	*output = append(*output, tempValue)
	// fmt.Printf("[end]sss [parseLuaType]outpt: %v\n", output)
	return nil
}

// UnmarshalJSON ...
func (p *LuaMapItem) UnmarshalJSON(data []byte) error {
	// fmt.Printf("[start]# [LuaMapItem]data: %s  ## start\n", string(data))
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

	keyResult, err := parseLuaKey(key)
	if err != nil {
		return err
	}
	// fmt.Printf("[LueKey] %v --> %v\n", key, keyResult)

	// fmt.Println("2.1 [LuaMapItem] start parse value")
	var result []interface{}
	err = parseLuaType(value, &result)
	if err != nil {
		return err
	}
	// fmt.Printf("[LueKey] %v \n----> %v\n", value, valueResult)
	// fmt.Println("2.2 [LuaMapItem] end parse value")

	itemResult := make(LuaMapItem, 2)
	// p = &LuaMapItem{keyResult, result}
	itemResult[0] = keyResult
	switch len(result) {
	case 1:
		itemResult[1] = result[0]
	default:
		itemResult[1] = result
	}
	*p = itemResult
	// fmt.Printf("# [LuaMapItem] p: %v ## end\n", p)
	// fmt.Println("\n")
	return nil
}

// LuaKey ... eg: {"key": [2, {"v": "lua string test"}]}
type LuaKey struct {
	Key LuaKeyObjectType `json:"key"`
}

type LuaMap []LuaMapItem

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

func (o *LuaTable) Unmarshal(data []byte) error {
	fmt.Printf("## [LuaTable]data: %s  ## start\n", string(data))

	if err := ffjson.Unmarshal(data, o); err != nil {
		return errors.Annotate(err, "unmarshal LuaTable object")
	}

	fmt.Printf("## [LuaTable] o: %v  ## end\n", o)
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

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{}

// LuaKeyObjectType ...
type LuaKeyObjectType []interface{}
