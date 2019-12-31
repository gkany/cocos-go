package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "github.com/gkany/graphSDK"
	// "github.com/gkany/graphSDK/config"
	// "github.com/gkany/graphSDK/crypto"
	"github.com/gkany/graphSDK/types"

	"log"
)

// type MemoExtension interface{}

// type TransferOp struct {
// 	From       types.AccountID   `json:"from"`
// 	To         types.AccountID   `json:"to"`
// 	Amount     types.AssetAmount `json:"amount"`
// 	Memo       MemoExtension     `json:"memo,omitempty"`
// 	Extensions types.Extensions  `json:"extensions"`
// }

// func NewMemoBuilderTest(p graphSDK.WebsocketAPI, fromObj, toObj types.GrapheneObject, memo string) *graphSDK.MemoBuilder {
// 	builder := graphSDK.MemoBuilder{
// 		from: fromObj,
// 		to:   toObj,
// 		memo: memo,
// 		api:  p,
// 	}
// 	return &builder
// }

// func TransferTest(p graphSDK.WebsocketAPI, keyBag crypto.KeyBag, from, to, feeAsset types.GrapheneObject, amount types.AssetAmount, memo string, isEncrypt bool) {
// 	op := TransferOp{
// 		Amount:     amount,
// 		Extensions: types.Extensions{},
// 		From:       types.AccountIDFromObject(from),
// 		To:         types.AccountIDFromObject(to),
// 	}

// 	if memo != "" {
// 		if isEncrypt {
// 			builder := p.NewMemoBuilder(from, to, memo)
// 			// builder := graphSDK.MemoBuilder{
// 			// 	from: from,
// 			// 	to:   to,
// 			// 	memo: memo,
// 			// 	api:  p,
// 			// }
// 			m, err := builder.Encrypt(keyBag)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			// m := types.Memo{
// 			// 	From:  from.Options.MemoKey,
// 			// 	To:    to.Options.MemoKey,
// 			// 	Nonce: types.UInt64(100),
// 			// }
// 			op.Memo = m
// 		} else {
// 			op.Memo = memo
// 		}
// 	}
// 	fmt.Printf("transfer op: %v\n", op)
// }

// func test() {
// 	from := types.NewAccountID("1.2.16") // nicotest 1.2.16
// 	to := types.NewAccountID("1.2.16")   // nicotest 1.2.16
// 	coreAsset := types.NewAssetID("1.3.0")
// 	amount := types.AssetAmount{
// 		Amount: types.Int64(100),
// 		Asset:  types.AssetIDFromObject(coreAsset),
// 	}
// 	// p *websocketAPI) Transfer(keyBag *crypto.KeyBag
// 	// p WebsocketAPI, keyBay crypto.KeyBag

// 	// api
// 	config.SetCurrent(config.ChainIDBCXTest)
// 	wsURL := "ws://test.cocosbcx.net"
// 	api := graphSDK.NewWebsocketAPI(wsURL)
// 	if err := api.Connect(); err != nil {
// 		log.Println(err)
// 	}
// 	//keyBag
// 	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
// 	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
// 	localKeyBag := crypto.NewKeyBag()
// 	err := localKeyBag.Add(privateKey)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	TransferTest(from, to, coreAsset, amount, "this is a test", true)
// 	TransferTest(from, to, coreAsset, amount, "this is a test", false)
// }

//test
type People struct{}
type People2 struct{}

func (p *People) ShowA() {
	fmt.Println("### showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("### showB")
}

func (p *People) ShowC() {
	fmt.Println("### showC")
}

func (p *People) ShowD() {
	fmt.Println("### People:showD")
}

func (p *People2) ShowD() {
	fmt.Println(">>> People2:showD")
}

type Teacher struct {
	People  //组合People
	People2 //组合People2
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
func (t *Teacher) ShowC(arg string) {
	fmt.Println(arg)
}

func UnMarshWitness() {
	rawMessage := `
	{"id":"1.6.1","witness_account":"1.2.5","last_aslot":12148831,"signing_key":"COCOS8hxjGaAwkNHewgQqg6ERLA7L4J6wkzuhhLRVbKKuJZUYM3dfuS","pay_vb":"1.13.2","vote_id":"1:0","total_votes":"5000000000000","url":"","total_missed":0,"last_confirmed_block_num":52619,"work_status":true,"next_maintenance_time":"1970-01-01T00:00:00","supporters":[]}
	`

	var witness types.Witness
	err := json.Unmarshal([]byte(rawMessage), &witness)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(witness)
}

func testExtend() {
	t := Teacher{}

	//print showA
	//print showB
	t.ShowA()

	//print teacher showB
	t.ShowB()

	//print showB
	t.People.ShowB()

	//print test
	t.ShowC("test")

	//print showC
	t.People.ShowC()

	//因为组合方法中多次包含ShowD，所以调用时必须显示指定匿名方法
	//print People2:showD
	t.People2.ShowD()
}

func testInterfaceMap() {
	slice := make([]interface{}, 10)

	map1 := make(map[string]string)
	map2 := make(map[string]int)

	map2["TaskID"] = 1
	map1["Command"] = "ping"

	map3 := make(map[string]map[string]string)
	map3["mapvalue"] = map1

	slice[0] = map2
	slice[1] = map1
	slice[3] = map3

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[3])
}

// json Marshal
type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

func test_crypto_ecdsa() {
	// 生成公钥和私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	// 公钥是存在在私钥中的，从私钥中读取公钥
	publicKey := &privateKey.PublicKey
	message := []byte("hello,dsa签名")

	// 进入签名操作
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)
	// 进入验证
	flag := ecdsa.Verify(publicKey, message, r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
	flag = ecdsa.Verify(publicKey, []byte("hello"), r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
}

type Student struct {
	Name  string `json:"name"`
	Age   int
	High  bool
	sex   string
	Class *Class `json:"class"`
}
type Class struct {
	Name  string
	Grade int
}

func testJSON() {
	st1 := Student{
		Name: "zhang San",
		Age:  18,
		High: true,
		sex:  "男",
	}
	c1 := new(Class)
	c1.Name = "1班"
	c1.Grade = 3
	st1.Class = c1

	stJSON, err := json.Marshal(st1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stJSON))
}

func testByte() {
	str := string("USDT")
	bytes := []byte(str)
	fmt.Printf("%v, %v", str, bytes)
}

func testDecodeString() {
	chainID := "dd896d2d415224156f95e61c93687ccc3fb38a7ec16e02bc509b6510a952936d"
	rawChainID, err := hex.DecodeString(chainID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rawChainID)
	fmt.Printf("%T\n", rawChainID)
}

func testDecodeJson() {
	rawMessage := `
	{
		"id":"2.1.0",
		"head_block_number":1349,
		"head_block_id":"0000054575ed56b5e2f8547ad3d6cd9ab309a386",
		"time":"2019-12-25T07:58:54",
		"current_witness":"1.6.6",
		"current_transaction_count":0,
		"next_maintenance_time":"2019-12-25T08:00:00",
		"last_budget_time":"2019-12-25T07:09:12",
		"witness_budget":0,
		"accounts_registered_this_interval":0,
		"recently_missed_count":0,
		"current_aslot":8855924,
		"recent_slots_filled":"319014556689469546759347944519926545855",
		"dynamic_flags":0,
		"last_irreversible_block_num":1340
	}
	`
	var dgp types.DynamicGlobalProperties
	err := json.Unmarshal([]byte(rawMessage), &dgp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dgp)
}

// test nil
type Pair []interface{}
type Operation struct {
	ID   int  `json:"id"`
	Memo Pair `json:"memo,omitempty"`
}

func testNil() {
	op := Operation{
		ID: 100,
	}
	fmt.Println(op.Memo == nil)

	op1 := Operation{
		ID:   200,
		Memo: Pair{},
	}
	fmt.Println(op1.Memo == nil)
}

func main() {

	// var a = make(interface{})
	// var i int = 5
	// s := "Hello world"
	// // a可以存储任意类型的数值
	// a = i
	// a = s

	// test()
	// test_crypto_ecdsa()
	// testJSON()
	// testByte()

	// testDecodeString()

	// testDecodeJson()

	// testNil()
	UnMarshWitness()
}
