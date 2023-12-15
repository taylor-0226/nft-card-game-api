package store

// var NftCollectionMap = map[string]models.NftCollection{"year": {}, "day_month": {}, "trigger": {
// 	Name:            "trigger",
// 	ContractAddress: "",
// }, "crafting": {Name: "crafting", ContractAddress: ""}, "category": {Name: "category", ContractAddress: "", Id: 2}, "prediction": {Name: "prediction", ContractAddress: ""}, "identity": {Name: "identity", ContractAddress: ""}}

// var NftTypeMap = map[int64]models.NftType{
// 	models.NFT_TYPE_ID_CATEGORY: {
// 		Id: models.NFT_TYPE_ID_CATEGORY,
// 		Name: "category",
// 	},
// 	models.NFT_TYPE_ID_CRAFTING: {
// 		Id: models.NFT_TYPE_ID_CRAFTING,
// 		Name: "crafting",

// 	},
// }

var NftRarityMap = map[int64]string{
	0: "Common",
	1: "Uncommon",
	2: "Rare",
}

// func GetNftCollectionByContractAddress(ctx context.Context, contract_address string) (models.NftCollection, bool) {
// 	return nil, false
// }
