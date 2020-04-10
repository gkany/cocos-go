package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gkany/cocos-go/logging"
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

// printJSON ...  map[string]interface{}
func printJSON(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			printJSON(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}

// test
// type Key struct {
// 	Path, Country string
// }

// type TestMapKey map[Key]int

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
	case LuaTypeInt:
		switch k2v.(type) {
		case float64:
			kv := k2v.(float64)
			// fmt.Println("kv is int", kv)
			tempKey = kv
		case string:
			kv := k2v.(string)
			// fmt.Println("kv is too large int --> string", kv)
			tempKey = kv
		}
	case LuaTypeNumber:
		kv := k2v.(float64)
		// fmt.Println("kv is number", kv)
		tempKey = kv
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

// func parseLuaType2(value interface{}) (interface{}, error) {
// 	var tempValue interface{}

// 	switch value.(type) {
// 	case map[string]interface{}:
// 		temp := value.(map[string]interface{})

// 	case []interface{}:

// 	default:
// 		fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", value.(type)）)
// 		// return nil, errors.Annotate("Unknow type")
// 		return nil, errors.Errorf("UnKnow LuaKey type, type index: %v", k1)
// 	}
// 	return tempValue, nil
// }

// LuaType ... basic type: int number string bool, complex: LuaMap, LuaTable, LuaFunction, maybe include basic and complex type
func parseLuaType(value []interface{}, output *[]interface{}) error {
	var tempValue interface{}
	fmt.Printf("[start]sss [parseLuaType] value: %v \n outpt: %v\n", value, output)

	v1 := value[0].(float64)
	// fmt.Printf(">>> [parseLuaType]v1: %v, type: %v; v2:%v, type: %v\n", v1, v1, value[1], value[1])

	// fmt.Printf("v2:%v, type: %v\n", value[1], value[1])
	// v2vTemp := value[1].(map[string]interface{})
	// switch value[1].(type) {
	// case map[string]interface{}:
	// 	v2 := value[1].(map[string]interface{})
	// 	v2v := v2["v"] // basic type
	// case []interface{}:
	// 	v2 := value[1].([]interface{})
	// 	// v2v := v2["v"] // basic type

	// }

	// if value[1].(type) != map[string]interface{} {
	// 	fmt.Printf("v2:%v, type: %v\n", value[1], value[1])
	// 	for kt, vt := range value[1] {
	// 		fmt.Printf("k: %v, v: %v\n", k, v)
	// 	}
	// 	return tempValue, nil
	// }

	// switch value[1].(type) {
	// case map[string]interface{}:
	// case []interface{}:
	// 	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")
	// 	return tempValue, nil
	// default:
	// 	return tempValue, nil
	// }

	v2 := value[1].(map[string]interface{})
	v2v := v2["v"] // basic type

	switch v1 {
	case LuaTypeInt:
		switch v2v.(type) {
		case float64:
			vv := v2v.(float64)
			// fmt.Println("vv is string", vv)
			tempValue = vv
		case string:
			vv := v2v.(string)
			fmt.Println("vv is too large int --> string", vv)
			tempValue = vv
		}
	case LuaTypeNumber:
		vv := v2v.(float64)
		// fmt.Println("vv is number", vv)
		tempValue = vv
	case LuaTypeString:
		vv := v2v.(string)
		// fmt.Println("vv is string", vv)
		tempValue = vv
	case LuaTypeBool:
		vv := v2v.(bool)
		// fmt.Println("vv is bool", vv)
		tempValue = vv
	case LuaTypeFunction:
		fmt.Printf("[LuaTypeFunction] value[1]: %v, type: %T\n", value[1], value[1])
		isVarArg := v2["is_var_arg"].(bool)
		argList := v2["arglist"].([]interface{})
		temp := []string{}
		for _, val := range argList {
			temp = append(temp, val.(string))
		}
		vv := LuaFunction{IsVarArg: isVarArg, ArgList: temp}
		fmt.Println("vv is function", vv)
		tempValue = vv
	case LuaTypeTable:
		fmt.Printf("-- [LuaTypeTable] value[1]: %v, type: %T\n", value[1], value[1])
		tmap := value[1].(map[string]interface{})
		for tk, tv := range tmap {
			// temp = append(temp, val.(string))
			fmt.Printf("---- tk:%v, type:%T;\n", tk, tk)
			fmt.Printf("---- tv:%v, type:%T\n", tv, tv)
			tvList := tv.([]interface{})
			for _, tvv := range tvList {
				// temp = append(temp, val.(string))
				fmt.Printf("------ tvv:%v, type:%T\n", tvv, tvv)
				for _, tvvv := range tvv.([]interface{}) {
					fmt.Printf("======= tvvv:%v, type:%T\n", tvvv, tvvv)
					switch tvvv.(type) {
					case []interface{}:
						fmt.Println("1111  start")
						fmt.Printf("data: %v, type: %v\n", tvvv, tvvv)
						err := parseLuaType(tvvv.([]interface{}), output)
						if err != nil {
							return err
						}
						// tvvvMap := tvvv.([]interface{})
						// var tvvvMapResult []interface{}
						// for _, t3vMapv := range tvvvMap {
						// 	// fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
						// 	fmt.Printf("----++++---- [t3vMapv]value:%v, type:%T\n", t3vMapv, t3vMapv)
						// 	t3vMapvKey := t3vMapv.(map[string]interface{})["key"]
						// 	fmt.Printf("[][] %v, type: %v\n", t3vMapvKey, t3vMapvKey)
						// 	// t3vResult, err := parseLuaType(t3vMapvKey.([]interface{}))
						// 	// if err != nil {
						// 	// 	return nil, err
						// 	// }
						// 	// tvvvMapResult = append(tvvvMapResult, t3vResult)
						// }

						// fmt.Printf("** t3vResult: %v\n", t3vResult)
						fmt.Println("1111  end\n\n")
						// tempValue = t3vResult
					case map[string]interface{}:
						fmt.Println("\n2222  start")
						// tvvvMap := tvv.(map[string]interface{})
						tvvvMap := tvv.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
						tm0 := tvvvMap[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
						tm1 := tvvvMap[1]
						fmt.Printf(">>>>>>>> tm0:%v, type:%T\n", tm0, tm0)
						fmt.Printf(">>>>>>>> tm1:%v, type:%T\n", tm1, tm1)

						// for key, value := range tm0.(map[string]interface{}) {
						// 	fmt.Printf("-------- [tm0]key:%v, type:%T\n", key, key)
						// 	fmt.Printf("-------- [tm0]value:%v, type:%T\n", value, value)
						// }
						tm0Map := tm0.(map[string]interface{})
						value := tm0Map["key"] // 二元数组： 根据类型解析
						fmt.Printf("****** [tm0['key']]value:%v, type:%T\n", value, value)
						err := parseLuaType(value.([]interface{}), output)
						if err != nil {
							fmt.Printf("** [error] tmoResult: %v\n", err)
							return err
						}
						// fmt.Printf("** tm0Result: %v\n", tm0Result)
						// for _, tm1v := range value.([]interface{}) {
						// 	// fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
						// 	fmt.Printf("-------- [tm0['key']]item:%v, type:%T\n", tm1v, tm1v)
						// }

						////////////////////////////// value
						// for _, tm1v := range tm1.([]interface{}) {
						// 	// fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
						// 	fmt.Printf("-------- tm1v:%v, type:%T\n", tm1v, tm1v)
						// }
						err = parseLuaType(tm1.([]interface{}), output)
						if err != nil {
							fmt.Printf("** [error] tm1Result: %v\n", err)
							return err
						}
						// fmt.Printf("** tm1Result: %v\n", tm1Result)
						fmt.Println("2222  end\n")
					}
				}
			}
		}

		// vv := v2v.(LuaTable)
		// fmt.Println("vv is table", vv)
		// tempValue = vv
	default:
		fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", v1))
		return errors.Errorf("UnKnow LuaType type, type index: %v", v1)
	}
	fmt.Printf("<<< [parseLuaType]parse result: %v ## end\n\n", tempValue)
	*output = append(*output, tempValue)
	fmt.Printf("[end]sss [parseLuaType]outpt: %v\n", output)
	return nil
}

// UnmarshalJSON ...
func (o *LuaMapItem) UnmarshalJSON(data []byte) error {
	fmt.Printf("[start]# [LuaMapItem]data: %s  ## start\n", string(data))
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

	// fmt.Printf("# [LuaMapItem] key: %v, value: %v #\n", key, value)
	// fmt.Printf("value type: %T\n", value)
	// LuaKey
	// fmt.Printf("LuaKey ... %v\n", key)
	/*
		k1 := key.Key[0].(float64)
		k2 := key.Key[1].(map[string]interface{})
		k2v := k2["v"]

		var tempKey interface{}
		switch k1 {
		case LuaTypeInt:
			kv := k2v.(uint64)
			fmt.Println("kv is string", kv)
			tempKey = kv
		case LuaTypeNumber:
			kv := k2v.(float64)
			fmt.Println("kv is number", kv)
			tempKey = kv
		case LuaTypeString:
			kv := k2v.(string)
			fmt.Println("kv is string", kv)
			tempKey = kv
		case LuaTypeBool:
			kv := k2v.(bool)
			fmt.Println("kv is bool", kv)
			tempKey = kv
		default:
			fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", k2v))
		}
	*/
	keyResult, err := parseLuaKey(key)
	if err != nil {
		return err
	}
	fmt.Printf("[LueKey] %v --> %v\n", key, keyResult)

	// test := make([]interface, 2)
	// LuaType
	// fmt.Printf("LuaType ... %v\n", value)
	// fmt.Printf("[0] %v, type: %T;[1] %v, type: %T\n", value[0], value[0], value[1], value[1])

	/*
		v1 := value[0].(float64)
		v2 := value[1].(map[string]interface{})
		v2v := v2["v"]
		var tempValue interface{}

		switch v1 {
		case LuaTypeInt:
			switch v2v.(type) {
			case float64:
				vv := v2v.(float64)
				fmt.Println("vv is string", vv)
				tempValue = vv
			case string:
				vv := v2v.(string)
				fmt.Println("vv is string", vv)
				tempValue = vv
			}

		case LuaTypeNumber:
			vv := v2v.(float64)
			fmt.Println("vv is number", vv)
			tempValue = vv
		case LuaTypeString:
			vv := v2v.(string)
			fmt.Println("vv is string", vv)
			tempValue = vv
		case LuaTypeBool:
			vv := v2v.(bool)
			fmt.Println("vv is bool", vv)
			tempValue = vv
		case LuaTypeFunction:
			// fmt.Printf("[LuaTypeFunction] value[1]: %v, type: %T\n", value[1], value[1])
			isVarArg := v2["is_var_arg"].(bool)
			argList := v2["arglist"].([]interface{})
			temp := []string{}
			for _, val := range argList {
				temp = append(temp, val.(string))
			}
			// vv := LuaFunction{IsVarArg: isVarArg, ArgList: temp}
			// fmt.Println("vv is function", vv)
			tempValue = LuaFunction{IsVarArg: isVarArg, ArgList: temp}
		case LuaTypeTable:
			fmt.Printf("-- [LuaTypeTable] value[1]: %v, type: %T\n", value[1], value[1])
			tmap := value[1].(map[string]interface{})
			for tk, tv := range tmap {
				// temp = append(temp, val.(string))
				fmt.Printf("---- tk:%v, type:%T;\n", tk, tk)
				fmt.Printf("---- tv:%v, type:%T\n", tv, tv)
				tvList := tv.([]interface{})
				for _, tvv := range tvList {
					// temp = append(temp, val.(string))
					fmt.Printf("------ tvv:%v, type:%T\n", tvv, tvv)
					for _, tvvv := range tvv.([]interface{}) {
						fmt.Printf("======= tvvv:%v, type:%T\n", tvvv, tvvv)
						switch tvvv.(type) {
						case []interface{}:
							fmt.Println("1111  start")
							tvvvMap := tvv.([]interface{})
							for t3vMapk, t3vMapv := range tvvvMap {
								fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
								fmt.Printf("-------- value:%v, type:%T\n", t3vMapv, t3vMapv)
							}
							fmt.Println("1111  end")
						case map[string]interface{}:
							fmt.Println("2222  start")
							// tvvvMap := tvv.(map[string]interface{})
							tvvvMap := tvv.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
							tm0 := tvvvMap[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
							tm1 := tvvvMap[1]
							fmt.Printf(">>>>>>>> tm0:%v, type:%T\n", tm0, tm0)
							fmt.Printf(">>>>>>>> tm1:%v, type:%T\n", tm1, tm1)

							// for key, value := range tm0.(map[string]interface{}) {
							// 	fmt.Printf("-------- [tm0]key:%v, type:%T\n", key, key)
							// 	fmt.Printf("-------- [tm0]value:%v, type:%T\n", value, value)
							// }
							tm0Map := tm0.(map[string]interface{})
							value := tm0Map["key"] // 二元数组： 根据类型解析
							fmt.Printf("-------- [tm0['key']]value:%v, type:%T\n", value, value)
							// for _, tm1v := range value.([]interface{}) {
							// 	// fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
							// 	fmt.Printf("-------- [tm0['key']]item:%v, type:%T\n", tm1v, tm1v)
							// }

							////////////////////////////// value
							for _, tm1v := range tm1.([]interface{}) {
								// fmt.Printf("-------- key:%v, type:%T\n", t3vMapk, t3vMapk)
								fmt.Printf("-------- tm1v:%v, type:%T\n", tm1v, tm1v)
							}
							fmt.Println("2222  end")
						}
					}
				}
			}

			// vv := v2v.(LuaTable)
			// fmt.Println("vv is table", vv)
			// tempValue = vv
		default:
			fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", v1))
		}
		fmt.Printf("[LueType]%v --> %v\n", value, tempValue)

	*/
	fmt.Println("2.1 [LuaMapItem] start parse value")
	var result []interface{}
	err = parseLuaType(value, &result)
	if err != nil {
		return err
	}
	// fmt.Printf("[LueKey] %v \n----> %v\n", value, valueResult)
	fmt.Println("2.2 [LuaMapItem] end parse value")

	/*
		// // test
		// k1 := key.Key[0].(float64)
		// k2 := key.Key[1]
		// fmt.Printf("k1: %v, type: %T\n", k1, k1)
		// fmt.Printf("k2: %v, type: %T\n", k2, k2)
		// if k1 == LuaTypeString {
		// 	fmt.Println("----- is LuaTypeString")
		// 	str := k2.(map[string]interface{})
		// 	printJSON(str)
		// 	fmt.Printf("str: %v, type: %T\n", str, str)
		// 	fmt.Println("<<<<<<<<<<<< ")
		// }
		// v1 := value[0]
		// v2 := value[1]
		// fmt.Printf("# [LuaMapItem] key: %v, value: %v #end\n\n", key.Key[1], value[1])
		// # [LuaMapItem] k2: map[v:max_bet], v2: map[v:1e+06] #end
		// fmt.Printf("# [LuaMapItem] k2: %v, type: %T; v2: %v, type: %T #end\n", k2, k2, v2, v2)
		// printJSON(k2)

		// test
		// switch vv := k2.(type) {
		// 	case string:
		// 		fmt.Println("is string", vv)
		// 	case float64:
		// 		fmt.Println("is float", int64(vv))
		// 	case int:
		// 		fmt.Println("is int", vv)
		// 	case []interface{}:
		// 		fmt.Println("is an array:")
		// 		for i, u := range vv {
		// 			fmt.Println(i, u)
		// 		}
		// 	case nil:
		// 		fmt.Println("is nil", "null")
		// 	case map[string]interface{}:
		// 		fmt.Println("is an map:")
		// 		// printJSON(vv)
		// 	default:
		// 		fmt.Println("is of a type I don't know how to handle ", fmt.Sprintf("%T", vv))
		// }
		// test end
	*/
	o = &LuaMapItem{keyResult, result}
	fmt.Printf("# [LuaMapItem] o: %v ## end\n", o)
	fmt.Println("\n")
	return nil
}

// LuaKey ... eg: {"key": [2, {"v": "lua string test"}]}
type LuaKey struct {
	Key LuaKeyObjectType `json:"key"`
}

// type LuaKeyObjectType []interface {}  # eg：[0, {"v": 100}] [2, {"v": "string_type"}]

// LuaMap ... map[LuaKey]LuaType
// type LuaMap map[LuaKey]LuaType

type LuaMap []LuaMapItem

// Marshal ...
func (o LuaMap) Marshal(enc *util.TypeEncoder) error {
	for _, item := range o {
		if err := enc.Encode(item); err != nil {
			return errors.Annotate(err, "encode LuaMap item")
		}
	}
	return nil
}

// Unmarshal ...
func (o *LuaMap) Unmarshal(data []byte) error {
	fmt.Printf("## [LuaMap]data: %s  ## start\n", string(data))

	if err := ffjson.Unmarshal(data, o); err != nil {
		return errors.Annotate(err, "unmarshal LuaMap object")
	}

	fmt.Printf("## [LuaMap] o: %v  ## end\n", o)
	return nil
}

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

// LuaInt, LuaNumber ... replace LuaObject
// LuaObject ... LuaInt, LuaNumber, LuaBool, LuaFunction, LuaString, LuaMap, LuaUserData
// type LuaObject struct {
// 	Value interface{} `json:"v"`
// }

// // Marshal ...
// func (o LuaObject) Marshal(enc *util.TypeEncoder) error {
// 	if err := enc.Encode(o.Value); err != nil {
// 		return errors.Annotate(err, "encode lua type obj")
// 	}
// 	return nil
// }

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{}

// type LuaType struct {
// 	Index  uint64
// 	Value interface{}
// }

// Marshal ...
func (o LuaType) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("[test] lua type object Marshal")
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua type object length error. %v", len(o))
	}

	if err := enc.Encode(o[0]); err != nil {
		return errors.Annotate(err, "encode type index")
	}

	if err := enc.Encode(o[1]); err != nil {
		return errors.Annotate(err, "encode type object")
	}

	// if err := enc.Encode(o.Index); err != nil {
	// 	return errors.Annotate(err, "encode type index")
	// }

	// if err := enc.Encode(o.Value); err != nil {
	// 	return errors.Annotate(err, "encode type object")
	// }

	return nil
}

// UnmarshalJSON ...
func (o *LuaType) Unmarshal(data []byte) error {
	fmt.Printf("## [LuaType]data: %s  ## start\n", string(data))
	dataStr := string(data)

	if strings.HasPrefix(dataStr, "[") {
		raw := make([]json.RawMessage, 2)
		if err := ffjson.Unmarshal(data, &raw); err != nil {
			return errors.Annotate(err, "unmarshal raw object")
		}

		var typeIndex uint64
		if err := ffjson.Unmarshal(raw[0], &typeIndex); err != nil {
			return errors.Annotate(err, "unmarshal lua type index")
		}
		// fmt.Printf("type index: %d, ", typeIndex)
		desc := fmt.Sprintf("encrypt type %s", typeIndex)

		if typeIndex == LuaTypeInt {
			var value LuaInt
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua int type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value.V}
			o = &LuaType{value.V}
		} else if typeIndex == LuaTypeNumber {
			var value LuaNumber
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua number type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value.V}
			o = &LuaType{value.V}
		} else if typeIndex == LuaTypeString {
			var value LuaString
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua string type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value.V}
			o = &LuaType{value.V}
		} else if typeIndex == LuaTypeBool {
			var value LuaBool
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua bool type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value.V}
			o = &LuaType{value.V}
		} else if typeIndex == LuaTypeFunction {
			var value LuaFunction
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua function type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value}
			o = &LuaType{value}
			fmt.Printf("[LuaType]LuaFunction value: %v\n", value)
		} else if typeIndex == LuaTypeTable {
			var value LuaTable
			if err := ffjson.Unmarshal(raw[1], &value); err != nil {
				logging.DDumpUnmarshaled(desc, raw[1])
				return errors.Annotatef(err, "unmarshal lua table type %v", typeIndex)
			}
			// o = &LuaType{typeIndex, value}
			o = &LuaType{value}
		} else {
			return errors.Errorf("unmarshal lua type error, typeIndex(%v) unknown", typeIndex)
		}
	} else if strings.HasPrefix(dataStr, "{") {
		// lua key
		var value LuaKey
		fmt.Printf("### [LuaType]parse LuaKey -- start\n")
		if err := ffjson.Unmarshal(data, &value); err != nil {
			return errors.Annotate(err, "unmarshal raw object")
		}
		// o = &LuaType{LuaTypeFunction, value}
		o = &LuaType{value}
		fmt.Printf("###  [LuaType]parse LuaKey: %v\n ### end", o)
	}
	fmt.Printf("## [LuaType] o: %v  ## end\n", o)
	return nil
}

// LuaKeyObjectType ...
type LuaKeyObjectType []interface{}

// Marshal ...
func (o LuaKeyObjectType) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("[test] lua type object Marshal")
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua type object length error. %v", len(o))
	}

	if err := enc.Encode(o[0]); err != nil {
		return errors.Annotate(err, "encode type index")
	}

	if err := enc.Encode(o[1]); err != nil {
		return errors.Annotate(err, "encode type object")
	}

	// if err := enc.Encode(o.Index); err != nil {
	// 	return errors.Annotate(err, "encode type index")
	// }

	// if err := enc.Encode(o.Value); err != nil {
	// 	return errors.Annotate(err, "encode type object")
	// }

	return nil
}

// UnmarshalJSON ...
func (o *LuaKeyObjectType) Unmarshal(data []byte) error {
	fmt.Printf("## [LuaKeyObjectType]data: %s  ## start\n", string(data))

	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	var typeIndex uint64 // 0 1 2 3
	if err := ffjson.Unmarshal(raw[0], &typeIndex); err != nil {
		return errors.Annotate(err, "unmarshal lua type index")
	}
	// fmt.Printf("type index: %d, ", typeIndex)
	desc := fmt.Sprintf("encrypt type %s", typeIndex)

	if typeIndex == LuaTypeInt {
		var value LuaInt
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua int type %v", typeIndex)
		}
		// o = &LuaKeyObjectType{typeIndex, value.V}
		o = &LuaKeyObjectType{value.V}
	} else if typeIndex == LuaTypeNumber {
		var value LuaNumber
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua number type %v", typeIndex)
		}
		// o = &LuaType{typeIndex, value.V}
		// o = &LuaKeyObjectType{typeIndex, value.V}
		o = &LuaKeyObjectType{value.V}
	} else if typeIndex == LuaTypeString {
		var value LuaString
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua string type %v", typeIndex)
		}
		// fmt.Printf("value: %v\n", value.V)
		// o = &LuaKeyObjectType{typeIndex, value.V}
		o = &LuaKeyObjectType{value.V}
	} else if typeIndex == LuaTypeBool {
		var value LuaBool
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua bool type %v", typeIndex)
		}
		// o = &LuaKeyObjectType{typeIndex, value.V}
		o = &LuaKeyObjectType{value.V}
	} else {
		return errors.Errorf("unmarshal lua type error, typeIndex(%v) unknown", typeIndex)
	}
	fmt.Printf("## [LuaKeyObjectType] o: %v  ## end\n", o)
	return nil
}

/*
// LuaKeyObjectType ...
type LuaKeyObjectType struct {
	Index   uint64 		// 0 1 2 3
	VInt    LuaInt  	// 0
	VNumber LuaNumber  	// 1
	VString LuaString 	// 2
	VBool   LuaBool  	// 3
}

// Marshal ...
func (o LuaKeyObjectType) Marshal(enc *util.TypeEncoder) error {
	fmt.Println("[test] lua type object Marshal")

	if o.Index > 3 {
		return errors.Errorf("lua type object Index error. %v", o.Index)
	}

	if err := enc.Encode(o.Index); err != nil {
		return errors.Annotate(err, "encode type index")
	}
	if o.Index == 0 {
		if err := enc.Encode(o.VInt); err != nil {
			return errors.Annotate(err, "encode type object")
		}
	} else if o.Index == 1 {
		if err := enc.Encode(o.VNumber); err != nil {
			return errors.Annotate(err, "encode type object")
		}
	} else if o.Index == 2 {
		if err := enc.Encode(o.VString); err != nil {
			return errors.Annotate(err, "encode type object")
		}
	} else if o.Index == 3 {
		if err := enc.Encode(o.VBool); err != nil {
			return errors.Annotate(err, "encode type object")
		}
	} else {
		return errors.Errorf("unmarshal lua type error, typeIndex(%v) unknown", o.Index)
	}

	return nil
}

// UnmarshalJSON ...
func (o *LuaKeyObjectType) UnmarshalJSON(data []byte) error {
	fmt.Println("--------------- LuaKeyObjectType UnmarshalJson")
	fmt.Println(string(data))

	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	var typeIndex uint64  // 0 1 2 3
	if err := ffjson.Unmarshal(raw[0], &typeIndex); err != nil {
		return errors.Annotate(err, "unmarshal lua type index")
	}
	fmt.Printf("type index: %d, ", typeIndex)
	desc := fmt.Sprintf("encrypt type %s", typeIndex)

	obj := LuaKeyObjectType{Index:typeIndex}
	if typeIndex == LuaTypeInt {
		var value LuaInt
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua int type %v", typeIndex)
		}
		obj.VInt = value
	} else if typeIndex == LuaTypeNumber {
		var value LuaNumber
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua number type %v", typeIndex)
		}
		obj.VNumber = value
	} else if typeIndex == LuaTypeString {
		var value LuaString
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua string type %v", typeIndex)
		}
		obj.VString = value
	} else if typeIndex == LuaTypeBool {
		var value LuaBool
		if err := ffjson.Unmarshal(raw[1], &value); err != nil {
			logging.DDumpUnmarshaled(desc, raw[1])
			return errors.Annotatef(err, "unmarshal lua bool type %v", typeIndex)
		}
		obj.VBool = value
		} else {
		return errors.Errorf("unmarshal lua type error, typeIndex(%v) unknown", typeIndex)
	}
	o = &obj
	fmt.Println(o)
	fmt.Println("/////////////////////")
	return nil
}
*/
