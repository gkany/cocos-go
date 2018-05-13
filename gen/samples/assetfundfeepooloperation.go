//this file is generated by gen [DO NOT EDIT]
//operation sample data for OperationTypeAssetFundFeePool

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeAssetFundFeePool] = sampleDataAssetFundFeePoolOperation
}

var sampleDataAssetFundFeePoolOperation = `{
  "amount": 20000000,
  "asset_id": "1.3.113",
  "extensions": [],
  "fee": {
    "amount": 100000,
    "asset_id": "1.3.0"
  },
  "from_account": "1.2.35298"
}`

//end of file