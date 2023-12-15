package nft

import "gameon-twotwentyk-api/models"

const WALLET_KEY_PRIVATE = "wallet_key_private"

type AggPack struct {
	AmountNftCard          int64
	AmountNftCardCelebrity int64
	AmountNftCardCategory  int64
	AmountNftCardDayMonth  int64
	AmountNftCardTrigger   int64
}

type AggPackRaw struct {
	Standard []map[int64]int64
	Premium  []map[int64]int64
	Elite    []map[int64]int64
}

type CollectionConfig struct {
}

type CollectionConfigTier struct {
	CardAmount   int64
	CardsPerPack []CollectionConfigCardsPerPack
	Price        int64
	Name         string
}

type CollectionConfigCardsPerPack struct {
	NftTypeId int64
	Amount    int64
}

type CollectionConfigPack struct {
	Tiers []CollectionConfigTier
}

func CreateCardCollection(data models.CardCollection, agg_pack AggPack) error {
	//load data into alex's engine
	//wait for alex's engine to return
	//store metadata in ipfs
	//mint nft card packs

	//dayMonth, year, category, trigger

	//core, uncommon, rare

	// rare_count := int64(0)
	// uncommon_count := int64(0)
	// core_count := int64(0)

	// cards_per_pack = int64(0)
	// pack_quantity = int64(0)
	// price = int64(0)

	// guaranteed_core_count = 0
	// guaranteed_uncommon_count = 0
	// guaranteed_rare_count = 0

	// for _, series := range data.CardSeries {
	// 	if *series.Rarity == models.NFT_RARITY_COMMON {
	// 		core_count = *series.CardAmount
	// 	}

	// 	if *series.Rarity == models.NFT_RARITY_UNCOMMON {
	// 		uncommon_count = *series.CardAmount
	// 	}

	// 	if *series.Rarity == models.NFT_RARITY_RARE {
	// 		rare_count = *series.CardAmount
	// 	}

	// }

	return nil
}
