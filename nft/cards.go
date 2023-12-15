package nft

import (
	"context"
	"gameon-twotwentyk-api/models"
)

// func CraftCardPack(tier, amount_trigger, amount_date_day, amount_date_month, amount_date_year int64) error {

// 	return nil
// }

func CraftIdentity(nft_card_day_month_id, nft_card_year_id, celebrity_id int64) (models.NftCardIdentity, error) {
	var out models.NftCardIdentity
	return out, nil
}

func CraftPrediction(nft_card_prediction_id, nft_card_trigger_id int64) (models.NftCardPrediction, error) {
	var out models.NftCardPrediction
	return out, nil
}

func GetNftCardDayMonth(ctx context.Context, id int64) (models.NftCardDayMonth, error) {
	var out models.NftCardDayMonth
	return out, nil
}

func GetNftCardYear(ctx context.Context, id int64) (models.NftCardYear, error) {
	var out models.NftCardYear
	return out, nil
}

func GetNftCardTrigger(ctx context.Context, id int64) (models.NftCardTrigger, error) {
	var out models.NftCardTrigger
	return out, nil
}

func GetNftCardIdentity(ctx context.Context, id int64) (models.NftCardIdentity, error) {
	var out models.NftCardIdentity
	return out, nil
}

func GetNftCardCrafting(ctx context.Context, id int64) (models.NftCardCrafting, error) {
	var out models.NftCardCrafting
	return out, nil
}
