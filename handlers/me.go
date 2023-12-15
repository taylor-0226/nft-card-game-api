package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
)

func GetMyNfts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	nft_card_crafting, err := store.ListNftCardCraftingByOwnerId(ctx, mid, models.NftCardCraftingFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_category, err := store.ListNftCardCategoryByOwnerId(ctx, mid, models.NftCardCategoryFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_day_month, err := store.ListNftCardDayMonthByOwnerId(ctx, mid, models.NftCardDayMonthFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_identity, err := store.ListNftCardIdentityByOwnerId(ctx, mid, models.NftCardIdentityFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_prediction, err := store.ListNftCardPredictionByOwnerId(ctx, mid, models.NftCardPredictionFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_trigger, err := store.ListNftCardTriggerByOwnerId(ctx, mid, models.NftCardTriggerFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_year, err := store.ListNftCardYearByOwnerId(ctx, mid, models.NftCardYearFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out := struct {
		NftCardCrafting   []models.NftCardCrafting   `json:"nft_card_crafting"`
		NftCardPrediction []models.NftCardPrediction `json:"nft_card_prediction"`
		NftCardTrigger    []models.NftCardTrigger    `json:"nft_card_trigger"`
		NftCardDayMonth   []models.NftCardDayMonth   `json:"nft_card_day_month"`
		NftCardYear       []models.NftCardYear       `json:"nft_card_year"`
		NftCardIdentity   []models.NftCardIdentity   `json:"nft_card_identity"`
		NftCardCategory   []models.NftCardCategory   `json:"nft_card_category"`
	}{
		NftCardCrafting:   nft_card_crafting,
		NftCardPrediction: nft_card_prediction,
		NftCardTrigger:    nft_card_trigger,
		NftCardDayMonth:   nft_card_day_month,
		NftCardYear:       nft_card_year,
		NftCardIdentity:   nft_card_identity,
		NftCardCategory:   nft_card_category,
	}

	ServeJSON(w, out)
}
