//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeProposalCreate

package samples

func init() {

	sampleDataProposalCreateOperation[37] = `{
  "expiration_time": "2016-06-06T22:12:49",
  "extensions": [],
  "fee": {
    "amount": 2341192,
    "asset_id": "1.3.0"
  },
  "fee_paying_account": "1.2.101860",
  "proposed_ops": [
    {
      "op": [
        6,
        {
          "account": "1.2.100429",
          "active": {
            "account_auths": [
              [
                "1.2.100429",
                4
              ]
            ],
            "address_auths": [],
            "key_auths": [],
            "weight_threshold": 1
          },
          "extensions": {},
          "fee": {
            "amount": 32233,
            "asset_id": "1.3.0"
          },
          "new_options": {
            "extensions": [],
            "memo_key": "BTS5VRaCZGCVQrPWsFAutV5fDVu8cGePg2cRowvHNdGQywhaQTyM5",
            "num_committee": 6,
            "num_witness": 9,
            "votes": [
              "1:22",
              "1:24",
              "1:27",
              "1:28",
              "1:36",
              "1:37",
              "1:44",
              "1:45",
              "1:56",
              "0:76",
              "0:84",
              "0:85",
              "0:87",
              "0:88",
              "0:147"
            ],
            "voting_account": "1.2.5"
          },
          "owner": {
            "account_auths": [
              [
                "1.2.100586",
                1
              ]
            ],
            "address_auths": [],
            "key_auths": [
              [
                "BTS7MCYAvFSdzNb5GNh5xpTcp6QqBW6Bvg6wCoE2GaJXA5rj3Abnu",
                1
              ]
            ],
            "weight_threshold": 1
          }
        }
      ]
    }
  ]
}`

}

//end of file
