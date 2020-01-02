package types

//go:generate ffjson $GOFILE

type NHAsset struct {
	ID             NHAssetID          `json:"id"`
	NhAssetCreator AccountID          `json:"nh_asset_creator"`
	NhAssetOwner   AccountID          `json:"nh_asset_owner"`
	NhAssetActive  AccountID          `json:"nh_asset_active"`
	Dealership     AccountID          `json:"dealership"`
	AssetQualifier string             `json:"asset_qualifier"`
	WorldView      string             `json:"world_view"`
	BaseDescribe   string             `json:"base_describe"`
	Parent         NHAssetMapType     `json:"parent"`
	Child          NHAssetMapType     `json:"child"`
	CreateTime     Time               `json:"create_time"`
	LimitList      ContractIDListType `json:"limit_list"`
}
