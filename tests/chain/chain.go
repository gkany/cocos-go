package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gkany/graphSDK"
	"github.com/gkany/graphSDK/config"
	"github.com/gkany/graphSDK/crypto"
	"github.com/gkany/graphSDK/operations"
	"github.com/gkany/graphSDK/types"
)

func getData(api graphSDK.WebsocketAPI) {
	accountID := types.NewAccountID("1.2.16") // nicotest 1.2.16
	coreAsset := types.NewAssetID("1.3.0")

	fmt.Println("\nGetAccountBalances: ")
	balances, err := api.GetAccountBalances(accountID, coreAsset)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("balances: %v", balances)
	}

	fmt.Println("\nGetBlock: ")
	block, err := api.GetBlock(uint64(100))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("block: %v", block)
	}

	fmt.Println("\nGetDynamicGlobalProperties: ")
	gdp, err := api.GetDynamicGlobalProperties()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("gdb: %v", gdp)
	}
}

func transfer(api graphSDK.WebsocketAPI) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	localKeyBag := crypto.NewKeyBag()
	localKeyBag.Add(privateKey)

	from := types.NewAccountID("1.2.16") // nicotest 1.2.16
	to := types.NewAccountID("1.2.5")
	coreAsset := types.NewAssetID("1.3.0")
	amount := types.AssetAmount{
		Amount: types.Int64(100),
		Asset:  types.AssetIDFromObject(coreAsset),
	}
	// memo := string("memo test true")
	// err := api.Transfer(localKeyBag, from, to, coreAsset, amount, memo, true)
	// fmt.Println(err)

	memo1 := string("memo test false")
	err1 := api.Transfer(localKeyBag, from, to, coreAsset, amount, memo1, true)
	// {"ref_block_num":53683,"ref_block_prefix":4187846010,"expiration":"2019-12-19T11:13:43","operations":[[0,{"from":"1.2.17","to":"1.2.18","amount":{"amount":100,"asset_id":"1.3.0"},"memo":[1,{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","nonce":5577006791947779410,"message":"8ea26a387aa8c4c4aa2ab6175177a5c70eb3637da62d0748e17d29a286d65d9f"}],"extensions":[]}]],"extensions":[]}
	fmt.Println(err1)
	// unlocked >>> read_memo {"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","nonce":"5577006791947779410","message":"8ea26a387aa8c4c4aa2ab6175177a5c70eb3637da62d0748e17d29a286d65d9f"}
	// read_memo {"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","nonce":"5577006791947779410","message":"8ea26a387aa8c4c4aa2ab6175177a5c70eb3637da62d0748e17d29a286d65d9f"}
	// "memo test false"
	// unlocked >>>

	// 2. false
	fmt.Println("\n-------> 2. transfer memo false")
	memo2 := string("memo test false")
	err2 := api.Transfer(localKeyBag, from, to, coreAsset, amount, memo2, false)
	fmt.Println(err2)
	// {"ref_block_num":53627,"ref_block_prefix":3375976042,"expiration":"2019-12-19T11:11:40","operations":[[0,{"from":"1.2.17","to":"1.2.18","amount":{"amount":100,"asset_id":"1.3.0"},"memo":[0,"memo test false"],"extensions":[]}]],"extensions":[]}

	// 3. no memo
	// err3 := api.Transfer(localKeyBag, from, to, coreAsset, amount, "", false)
	// {"ref_block_num":53539,"ref_block_prefix":2200014947,"expiration":"2019-12-19T11:08:29","operations":[[0,{"from":"1.2.17","to":"1.2.18","amount":{"amount":100,"asset_id":"1.3.0"},"extensions":[]}]],"extensions":[]}
	// fmt.Println(err3)
}

/*
func transfer2(api graphSDK.WebsocketAPI) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	localKeyBag := crypto.NewKeyBag()
	localKeyBag.Add(privateKey)

	from := types.NewAccountID("1.2.15") // nicotest 1.2.15
	to := types.NewAccountID("1.2.5")
	coreAsset := types.NewAssetID("1.3.0")
	amount := types.AssetAmount{
		Amount: types.Int64(100),
		Asset:  types.AssetIDFromObject(coreAsset),
	}

	fmt.Println("\n-------> 1. transfer ")
	memo1 := string("memo test false")
	err1 := api.Transfer2(localKeyBag, from, to, coreAsset, amount, memo1)
	fmt.Println(err1)

	// 2. memo empty
	fmt.Println("\n-------> 2. transfer memo empty")
	err2 := api.Transfer2(localKeyBag, from, to, coreAsset, amount, "")
	fmt.Println(err2)
}
*/

func test_broadcast(api graphSDK.WebsocketAPI) {

	// if err := api.BroadcastTransaction(trx); err != nil {
	// 	return errors.Annotate(err, "BroadcastTransaction")
	// }
}

func test_GetAccountByName(api graphSDK.WebsocketAPI) {

}

func test_ListAssets(api graphSDK.WebsocketAPI) {
	assets, error := api.ListAssets("", 5)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(assets)
}

// CreateAsset(keyBag *crypto.KeyBag, issuer, feeAsset types.GrapheneObject, symbol string, precision uint8) error
func test_createAsset(api graphSDK.WebsocketAPI) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(privateKey)
	// localKeyBag.Add("5Jdvatdk3qpZ8Ek9tQyqh3QwQ5mWNZ7kfnwSVwMUsLLmdUAfUwo")
	// if err != nil {
	// 	log.Println(err)
	// }

	issuer := types.NewAccountID("1.2.16") // nicotest
	coreAsset := types.NewAssetID("1.3.0")

	rate := types.Price{
		Base: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(types.NewAssetID("1.3.3")),
		},
		Quote: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(coreAsset),
		},
	}

	common := types.AssetOptions{
		MaxSupply:        210000000000,
		CoreExchangeRate: &rate,
	}
	err := api.CreateAsset(keyBag, issuer, coreAsset, "USDC", 5, common, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetBlock(api graphSDK.WebsocketAPI, from, to uint64) {
	for i := from; i <= to; i++ {
		log.Printf("block: %v ", i)
		block, err := api.GetBlock(i)
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Printf("block %v: %v", i, block)
		log.Printf("block operations: %v\n", block.Transactions)
		for _, trx := range block.Transactions {
			// TransactionId     string
			// SignedTransaction ProcessedTransaction
			fmt.Printf("trx id: %v, SignedTransaction: %v\n", trx.TransactionId, trx.SignedTransaction)
		}
	}
}

func testHeadBlockPrefix(api graphSDK.WebsocketAPI) {
	fmt.Println("\nGetDynamicGlobalProperties: ")
	gdp, err := api.GetDynamicGlobalProperties()
	if err != nil {
		fmt.Println(err)
		return
	}

	rawBlockID, err := hex.DecodeString(gdp.HeadBlockID.String())
	if err != nil {
		return
	}
	fmt.Println("HeadBlockID: ", gdp.HeadBlockID.String(), ", rawBlockID: ", rawBlockID)
	if len(rawBlockID) < 8 {
		return
	}

	rawPrefix := rawBlockID[4:8]
	fmt.Println("rawPrefix: ", rawPrefix)

	var prefix uint32
	//binary.LittleEndian 小端序
	if err := binary.Read(bytes.NewReader(rawPrefix), binary.LittleEndian, &prefix); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("prefix: ", types.UInt32(prefix))
	fmt.Println("prefix: ", prefix)
}

//test
func NewTransferOperation() operations.TransferOperation {
	from := types.NewAccountID("1.2.16") // nicotest 1.2.16
	to := types.NewAccountID("1.2.10")
	coreAsset := types.NewAssetID("1.3.0")
	amount := types.AssetAmount{
		Amount: types.Int64(100),
		Asset:  types.AssetIDFromObject(coreAsset),
	}

	op := operations.TransferOperation{
		Amount:     amount,
		Extensions: types.Extensions{},
		From:       types.AccountIDFromObject(from),
		To:         types.AccountIDFromObject(to),
		// Memo:       []interface{}{0, "test trx"},
	}
	return op
}

func getOperations(ops ...types.Operation) types.Operations {
	return types.Operations(ops)
}

func testBroadcastTrx(api graphSDK.WebsocketAPI) {
	fmt.Println("Unmarshal transfer memo")
	memoJSON := `{
		"from": "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx",
		"to": "COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7",
		"nonce": "1370671550721035852",
		"message": "06f0785ded07956401b53c34b02a4f38eee1374984ff892e8d456b8263c1ef57"
	  }`
	var memo types.Memo
	err := json.Unmarshal([]byte(memoJSON), &memo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(memo)

	fmt.Println("\nUnmarshal transfer operation")
	transferOpJSON := `
	{
		"fee": {
		  "amount": 21054,
		  "asset_id": "1.3.0"
		},
		"from": "1.2.15",
		"to": "1.2.5",
		"amount": {
		  "amount": 1000000,
		  "asset_id": "1.3.0"
		},
		"extensions": []
	  }
	`
	var transferOp operations.TransferOperation
	err = json.Unmarshal([]byte(transferOpJSON), &transferOp)
	if err != nil {
		fmt.Println(err)
		return
	}
	// transferOp.Memo = &memo
	fmt.Println(transferOp)

	fmt.Println("\nUnmarshal trx")
	trxJSON := `
	{
		"ref_block_num": 2819,
		"ref_block_prefix": 3987120764,
		"expiration": "2019-12-25T09:13:50",
		"extensions": [],
		"signatures": [
		  "1f34a2401e10df67413c5c7573a706b1c5d8e1d6e8a7e9c407314213aeda5eda625155b813aabf9545eb347a650e1d4d75fb62ba3ac6b36f32b88392171471eca1"
		]
	  }
	`
	var signedTrx types.SignedTransaction
	err = json.Unmarshal([]byte(trxJSON), &signedTrx)
	if err != nil {
		fmt.Println(err)
		return
	}
	// var op types.Operation
	// op = transferOp
	signedTrx.Operations = getOperations(&transferOp)
	fmt.Println(signedTrx)

	if err := api.BroadcastTransaction(&signedTrx); err != nil {
		fmt.Println(err)
	}
	/*
		371114ms th_a       websocket_api.cpp:148         on_message           ] websocket api exception :{"code":10,"name":"assert_exception","message":"Assert Exception","stack":[{"context":{"level":"error","file":"db_block.cpp","line":736,"method":"_apply_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:06:11"},"format":"now <= trx.expiration: ","data":{"now":"2019-12-25T10:06:10","trx.exp":"2019-12-25T09:13:50"}},{"context":{"level":"warn","file":"db_block.cpp","line":813,"method":"_apply_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:06:11"},"format":"","data":{"trx":{"ref_block_num":2819,"ref_block_prefix":3987120764,"expiration":"2019-12-25T09:13:50","operations":[[0,{"fee":{"amount":21054,"asset_id":"1.3.0"},"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"1370671550721035852","message":"06f0785ded07956401b53c34b02a4f38eee1374984ff892e8d456b8263c1ef57"},"extensions":[]}]],"extensions":[],"signatures":["1f34a2401e10df67413c5c7573a706b1c5d8e1d6e8a7e9c407314213aeda5eda625155b813aabf9545eb347a650e1d4d75fb62ba3ac6b36f32b88392171471eca1"]}}},{"context":{"level":"warn","file":"db_block.cpp","line":258,"method":"push_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:06:11"},"format":"","data":{"trx":{"ref_block_num":2819,"ref_block_prefix":3987120764,"expiration":"2019-12-25T09:13:50","operations":[[0,{"fee":{"amount":21054,"asset_id":"1.3.0"},"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"1370671550721035852","message":"06f0785ded07956401b53c34b02a4f38eee1374984ff892e8d456b8263c1ef57"},"extensions":[]}]],"extensions":[],"signatures":["1f34a2401e10df67413c5c7573a706b1c5d8e1d6e8a7e9c407314213aeda5eda625155b813aabf9545eb347a650e1d4d75fb62ba3ac6b36f32b88392171471eca1"]}}},{"context":{"level":"warn","file":"websocket_api.cpp","line":144,"method":"on_message","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:06:11"},"format":"","data":{"call.method":"call","call.params":[4,"broadcast_transaction",[{"signatures":["1f34a2401e10df67413c5c7573a706b1c5d8e1d6e8a7e9c407314213aeda5eda625155b813aabf9545eb347a650e1d4d75fb62ba3ac6b36f32b88392171471eca1"],"ref_block_num":2819,"ref_block_prefix":3987120764,"expiration":"2019-12-25T09:13:50","operations":[[0,{"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"1370671550721035852","message":"06f0785ded07956401b53c34b02a4f38eee1374984ff892e8d456b8263c1ef57"},"extensions":[],"fee":{"amount":21054,"asset_id":"1.3.0"}}]],"extensions":[]}]]}}]}
	*/
}

func testBroadcastTrx2(api graphSDK.WebsocketAPI) {
	fmt.Println("\nUnmarshal trx")
	trxJSON := `
	{
		"ref_block_num": 3019,
		"ref_block_prefix": 1345571086,
		"expiration": "2019-12-25T09:21:20",
		"operations": [[
			0,{
			  "fee": {
				"amount": 21054,
				"asset_id": "1.3.0"
			  },
			  "from": "1.2.15",
			  "to": "1.2.5",
			  "amount": {
				"amount": 1000000,
				"asset_id": "1.3.0"
			  },
			  "memo": {
				"from": "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx",
				"to": "COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7",
				"nonce": "6558818321903381236",
				"message": "1a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed"
			  },
			  "extensions": []
			}
		  ]
		],
		"extensions": [],
		"signatures": [
		  "207f701184579cdc559cc782ad12df2f7ff345d874afca790c7b79cf8fb6faa9ab3e06a8cfc12b4af1bc582d6acb02655e6347f764653cb5c33db54e511eb9d26f"
		]
	  }
	`
	var signedTrx types.SignedTransaction
	err := json.Unmarshal([]byte(trxJSON), &signedTrx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(signedTrx)

	if err := api.BroadcastTransaction(&signedTrx); err != nil {
		fmt.Println(err)
	}

	/*
		434610ms th_a       websocket_api.cpp:148         on_message           ] websocket api exception :{"code":10,"name":"assert_exception","message":"Assert Exception","stack":[{"context":{"level":"error","file":"db_block.cpp","line":736,"method":"_apply_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:07:14"},"format":"now <= trx.expiration: ","data":{"now":"2019-12-25T10:07:14","trx.exp":"2019-12-25T09:21:20"}},{"context":{"level":"warn","file":"db_block.cpp","line":813,"method":"_apply_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:07:14"},"format":"","data":{"trx":{"ref_block_num":3019,"ref_block_prefix":1345571086,"expiration":"2019-12-25T09:21:20","operations":[[0,{"fee":{"amount":21054,"asset_id":"1.3.0"},"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"6558818321903381236","message":"1a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed"},"extensions":[]}]],"extensions":[],"signatures":["207f701184579cdc559cc782ad12df2f7ff345d874afca790c7b79cf8fb6faa9ab3e06a8cfc12b4af1bc582d6acb02655e6347f764653cb5c33db54e511eb9d26f"]}}},{"context":{"level":"warn","file":"db_block.cpp","line":258,"method":"push_transaction","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:07:14"},"format":"","data":{"trx":{"ref_block_num":3019,"ref_block_prefix":1345571086,"expiration":"2019-12-25T09:21:20","operations":[[0,{"fee":{"amount":21054,"asset_id":"1.3.0"},"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"6558818321903381236","message":"1a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed"},"extensions":[]}]],"extensions":[],"signatures":["207f701184579cdc559cc782ad12df2f7ff345d874afca790c7b79cf8fb6faa9ab3e06a8cfc12b4af1bc582d6acb02655e6347f764653cb5c33db54e511eb9d26f"]}}},{"context":{"level":"warn","file":"websocket_api.cpp","line":144,"method":"on_message","hostname":"","thread_name":"th_a","timestamp":"2019-12-25T10:07:14"},"format":"","data":{"call.method":"call","call.params":[4,"broadcast_transaction",[{"signatures":["207f701184579cdc559cc782ad12df2f7ff345d874afca790c7b79cf8fb6faa9ab3e06a8cfc12b4af1bc582d6acb02655e6347f764653cb5c33db54e511eb9d26f"],"ref_block_num":3019,"ref_block_prefix":1345571086,"expiration":"2019-12-25T09:21:20","operations":[[0,{"from":"1.2.15","to":"1.2.5","amount":{"amount":1000000,"asset_id":"1.3.0"},"memo":{"from":"COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx","to":"COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7","nonce":"6558818321903381236","message":"1a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed"},"extensions":[],"fee":{"amount":21054,"asset_id":"1.3.0"}}]],"extensions":[]}]]}}]}
	*/
}

/*
func testSign(api graphSDK.WebsocketAPI) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	localKeyBag := crypto.NewKeyBag()
	localKeyBag.Add(privateKey)

	serialize_transaction := "35180a05475e5e45035e01003e52000000000000000f0540420f00000000000001021b6eb9d35c7f7ed23b64cb19fd96513cb50a6acd768bc35cfccabc1ad2d5c5c6024bc0412e636f8b2801764c94f734391bc05d0f5ef528f4bdb6b01257737cc443f4ea5f7f839a055b201a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed00000000"

	error := api.SignTest(localKeyBag, serialize_transaction)
	if error != nil {
		fmt.Println(error)
	}
}
*/

func main() {
	// config.SetCurrent(config.ChainIDTestnet)
	// wsURL := "ws://test.cocosbcx.net"

	chainID := config.ChainIDLocal
	wsURL := "ws://127.0.0.1:8049"
	config.SetCurrent(chainID)

	// chain api 测试
	log.Println("------- chain api test ----------")
	api := graphSDK.NewWebsocketAPI(wsURL)
	if err := api.Connect(); err != nil {
		log.Println(err)
	}
	// getData(api)
	// transfer(api)
	// transfer2(api)  // old version

	// test_ListAssets(api)  // success
	test_createAsset(api)

	// testGetBlock(api, 7997, 7998)

	// testBroadcastTrx(api)
	// testBroadcastTrx2(api)

	// testSign(api)

	// testHeadBlockPrefix(api)
}
