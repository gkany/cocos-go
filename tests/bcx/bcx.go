package main

import (
	"fmt"
	"log"

	"github.com/gkany/graph-sdk"
	"github.com/gkany/graph-sdk/config"
	"github.com/gkany/graph-sdk/crypto"
	"github.com/gkany/graph-sdk/types"
)

func main() {
	// config.SetCurrent(config.ChainIDBCXDev)
	// wsURL := "ws://127.0.0.1:8049"
	// // wsURL := "ws://test.cocosbcx.net"
	// walletURL := "http://127.0.0.1:8048"

	config.SetCurrent(config.ChainIDBCXTest)
	wsURL := "ws://test.cocosbcx.net"
	walletURL := "http://127.0.0.1:8048"

	// testAccount := "nicotest"
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"

	// 1. key bag 测试
	log.Println("------- key bag test ----------")
	localKeyBag := crypto.NewKeyBag()
	err := localKeyBag.Add(privateKey)
	if err != nil {
		log.Println(err)
	} else {
		publics := localKeyBag.Publics()
		fmt.Printf("publics: %v %v\n", publics, publicKey)
	}

	// 2. chain api 测试
	log.Println("------- chain api test ----------")
	api := graph-sdk.NewWebsocketAPI(wsURL)
	if err := api.Connect(); err != nil {
		log.Println(err)
	}

	accountID := types.NewAccountID("1.2.16") // nicotest 1.2.16
	coreAsset := types.NewAssetID("1.3.0")

	balances, err := api.GetAccountBalances(accountID, coreAsset)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("balances: %v", balances)

	// 3. wallet api 测试
	log.Println("------- wallet api test ----------")
	walletAPI := graph-sdk.NewWalletAPI(walletURL)
	if err := walletAPI.Connect(); err != nil {
		log.Println(err)
	}

	info, err := walletAPI.Info()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("info: %v", info)
}

/*
test result:
1. local develop chain
2019/12/12 18:06:04 ------- key bag test ----------
publics: [COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx] COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx
2019/12/12 18:06:04 ------- chain api test ----------
2019/12/12 18:06:04 balances: [{9677387512479481 1.3.0}]
2019/12/12 18:06:04 ------- wallet api test ----------
2019/12/12 18:06:04 info: &{[1.5.0 1.5.1 1.5.2 1.5.3 1.5.4 1.5.5 1.5.6 1.5.7 1.5.8 1.5.9 1.5.10] [1.6.1 1.6.2 1.6.3 1.6.4 1.6.5 1.6.6 1.6.7 1.6.8 1.6.9 1.6.10 1.6.11] 179db3c6a2e08d610f718f05e9cc2aabad62aff80305b9621b162b8b6f2fd79f 23 hours ago 00008c89e07b78f0e0fd23ebf17674c9c258c267 35977 22 hours ago 91.40625000000000000}

2. testnet
2019/12/12 18:07:30 ------- key bag test ----------
publics: [COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx] COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx
2019/12/12 18:07:30 ------- chain api test ----------
2019/12/12 18:07:30 balances: [{6883739559325716 1.3.0}]
2019/12/12 18:07:30 ------- wallet api test ----------
2019/12/12 18:07:30 info: &{[1.5.0 1.5.1 1.5.2 1.5.3 1.5.4 1.5.5 1.5.6 1.5.7 1.5.8 1.5.9 1.5.10] [1.6.1 1.6.2 1.6.3 1.6.4 1.6.5 1.6.6 1.6.7 1.6.8 1.6.9 1.6.10 1.6.11] 179db3c6a2e08d610f718f05e9cc2aabad62aff80305b9621b162b8b6f2fd79f 23 hours ago 00008c89e07b78f0e0fd23ebf17674c9c258c267 35977 22 hours ago 91.40625000000000000}
*/
