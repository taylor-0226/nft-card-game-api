package nft

import "math/rand"

func SimulateNewNftIdentityCard(user_id int64, celebrity_id int64) {

}

func SimulateNewNftPredictionCard(user_id int64, celebrity_id int64, trigger string) {

}

func SimulateNewNftTriggerCard(user_id int64, trigger string) {

}

func SimulateOpenCardPack(user_id int64, rarity string) {
	//add fake nft cards to user wallet
}

func RandomRarity() int64 {
	return int64(rand.Intn(2))
}

func RarityName(i int64) string {
	switch i {
	case 0:
		return "common"
	case 1:
		return "uncommon"
	case 2:
		return "rare"
	default:
		return "common"
	}
}
