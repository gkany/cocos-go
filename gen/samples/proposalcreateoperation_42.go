//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeProposalCreate

package samples

func init() {

	sampleDataProposalCreateOperation[42] = `{
  "expiration_time": "2016-07-30T15:13:15",
  "extensions": [],
  "fee": {
    "amount": 2507447,
    "asset_id": "1.3.0"
  },
  "fee_paying_account": "1.2.116522",
  "proposed_ops": [
    {
      "op": [
        11,
        {
          "asset_to_update": "1.3.1074",
          "extensions": [],
          "fee": {
            "amount": 29393527,
            "asset_id": "1.3.0"
          },
          "issuer": "1.2.114105",
          "new_issuer": "1.2.116524",
          "new_options": {
            "blacklist_authorities": [],
            "blacklist_markets": [],
            "core_exchange_rate": {
              "base": {
                "amount": 100000,
                "asset_id": "1.3.0"
              },
              "quote": {
                "amount": 100,
                "asset_id": "1.3.1074"
              }
            },
            "description": "{\"main\":\"bbbingji.com （BB评级）是中国首家、影响力最大的专门针对区块链资产（股权）公布网站之一，BB评级网致力于公布去中心化资产的众筹公告、评级、财报、资料、历史分红等，秉承为用户着想的区块链资产查询网站。\",\"short_name\":\"BB评级网股权\",\"market\":\"\"}",
            "extensions": [],
            "flags": 4,
            "issuer_permissions": 79,
            "market_fee_percent": 0,
            "max_market_fee": 0,
            "max_supply": 1000000000,
            "whitelist_authorities": [],
            "whitelist_markets": []
          }
        }
      ]
    }
  ]
}`

}

//end of file
