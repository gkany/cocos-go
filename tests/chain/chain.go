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

	// memo := string("memo test false")
	// err := api.Transfer(localKeyBag, from, to, coreAsset, amount, memo, false)
	// fmt.Println(err)

	// memo := string("memo test false")
	err := api.Transfer(localKeyBag, from, to, coreAsset, amount, "", false)
	fmt.Println(err)
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

func main() {
	config.SetCurrent(config.ChainIDBCXDev)
	// wsURL := "ws://test.cocosbcx.net"
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
	test_createAsset(api)
}
