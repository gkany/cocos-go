package main

import (
	"fmt"
	"log"

	"github.com/gkany/graphSDK"
	"github.com/gkany/graphSDK/config"
	"github.com/gkany/graphSDK/crypto"
	"github.com/gkany/graphSDK/types"
)

func getData(api graphSDK.WebsocketAPI) {
	accountID := types.NewAccountID("1.2.16") // nicotest 1.2.16
	coreAsset := types.NewAssetID("1.3.0")

	balances, err := api.GetAccountBalances(accountID, coreAsset)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("balances: %v", balances)
	}

	block, err := api.GetBlock(uint64(100))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("block: %v", block)
	}

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
	// localKeyBag.Add("5Jdvatdk3qpZ8Ek9tQyqh3QwQ5mWNZ7kfnwSVwMUsLLmdUAfUwo")
	// if err != nil {
	// 	log.Println(err)
	// }

	from := types.NewAccountID("1.2.17") // init6 1.2.11
	to := types.NewAccountID("1.2.18")   // nicotest 1.2.16
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
	// memo2 := string("memo test false")
	// err2 := api.Transfer(localKeyBag, from, to, coreAsset, amount, memo2, false)
	// fmt.Println(err2)
	// {"ref_block_num":53627,"ref_block_prefix":3375976042,"expiration":"2019-12-19T11:11:40","operations":[[0,{"from":"1.2.17","to":"1.2.18","amount":{"amount":100,"asset_id":"1.3.0"},"memo":[0,"memo test false"],"extensions":[]}]],"extensions":[]}

	// 3. no memo
	// err3 := api.Transfer(localKeyBag, from, to, coreAsset, amount, "", false)
	// {"ref_block_num":53539,"ref_block_prefix":2200014947,"expiration":"2019-12-19T11:08:29","operations":[[0,{"from":"1.2.17","to":"1.2.18","amount":{"amount":100,"asset_id":"1.3.0"},"extensions":[]}]],"extensions":[]}
	// fmt.Println(err3)
}

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

	issuer := types.NewAccountID("1.2.17") // nicotest
	coreAsset := types.NewAssetID("1.3.0")

	rate := types.Price{
		Base: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(types.NewAssetID("1.3.2")),
		},
		Quote: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(coreAsset),
		},
	}

	common := types.AssetOptions{
		MaxSupply:        100000000,
		CoreExchangeRate: rate,
	}
	err := api.CreateAsset(keyBag, issuer, coreAsset, "USDT", 5, common, nil)
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
	}
}

func main() {
	// config.SetCurrent(config.ChainIDBCXTest)
	// wsURL := "ws://test.cocosbcx.net"
	config.SetCurrent(config.ChainIDBCXDev)
	wsURL := "ws://127.0.0.1:8049"

	// chain api 测试
	log.Println("------- chain api test ----------")
	api := graphSDK.NewWebsocketAPI(wsURL)
	if err := api.Connect(); err != nil {
		log.Println(err)
	}
	// getData(api)
	// transfer(api)

	// test_ListAssets(api)  // success
	// test_createAsset(api)

	testGetBlock(api, 4495, 4499)
}
