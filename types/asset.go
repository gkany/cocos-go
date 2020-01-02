package types

//go:generate ffjson $GOFILE

type Assets []Asset

type Asset struct {
	ID                 AssetID             `json:"id"`
	Symbol             String              `json:"symbol"`
	Precision          int                 `json:"precision"`
	Issuer             AccountID           `json:"issuer"`
	Options            AssetOptions        `json:"options"`
	DynamicAssetDataID AssetDynamicDataID  `json:"dynamic_asset_data_id"`
	BitassetDataID     AssetBitAssetDataID `json:"bitasset_data_id"`
	BuybackAccount     AccountID           `json:"buyback_account"`
}
