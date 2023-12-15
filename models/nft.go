package models

const (
	TRIGGER_DEATH    = "death"
	TRIGGER_MARRIAGE = "marriage"
)

const (
	NFT_TYPE_ID_CARD_PACK  = 0
	NFT_TYPE_ID_CRAFTING   = 1
	NFT_TYPE_ID_CATEGORY   = 2
	NFT_TYPE_ID_DAY_MONTH  = 3
	NFT_TYPE_ID_YEAR       = 4
	NFT_TYPE_ID_TRIGGER    = 5
	NFT_TYPE_ID_IDENTITY   = 6
	NFT_TYPE_ID_PREDICTION = 7
)

const (
	NFT_RARITY_COMMON   = 0
	NFT_RARITY_UNCOMMON = 1
	NFT_RARITY_RARE     = 2
)

type NftType struct {
	Id              int
	ContractAddress string
	Name            string
}
