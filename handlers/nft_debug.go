package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"math/rand"
	"net/http"

	"github.com/pandoratoolbox/json"
)

func GenerateAllTriggers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	oid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		OwnerId int64
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if input.OwnerId != 0 {
		oid = input.OwnerId
	}

	triggers, err := store.ListTrigger(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rarity := int64(rand.Intn(2))

	card_series_id := int64(1)

	var out []models.NftCardTrigger

	for _, t := range triggers {
		new := models.NftCardTrigger{
			NftCardTriggerData: models.NftCardTriggerData{
				CardSeriesId: &card_series_id,
				Trigger:      t.Name,
				Tier:         t.Tier,
				Rarity:       &rarity,
				OwnerId:      &oid,
			},
		}
		err := store.NewNftCardTrigger(ctx, &new)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		out = append(out, new)
	}

	ServeJSON(w, out)
}

func GenerateNFTsForRecipeIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := struct {
		OwnerId int64
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	oid := ctx.Value(models.CTX_user_id).(int64)

	if input.OwnerId != 0 {
		oid = input.OwnerId
	}

	rarity := int64(rand.Intn(2))
	is_crafted := false

	celebrities, err := store.ListCelebrity(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	i := rand.Intn(len(celebrities) - 1)

	celebrity := celebrities[i]

	card_series_id := int64(1)
	card_year := models.NftCardYear{
		NftCardYearData: models.NftCardYearData{
			OwnerId:      &oid,
			Year:         celebrity.BirthYear,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	err = store.NewNftCardYear(ctx, &card_year)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	card_day_month := models.NftCardDayMonth{
		NftCardDayMonthData: models.NftCardDayMonthData{
			OwnerId:      &oid,
			Day:          celebrity.BirthDay,
			Month:        celebrity.BirthMonth,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	err = store.NewNftCardDayMonth(ctx, &card_day_month)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	card_category := models.NftCardCategory{
		NftCardCategoryData: models.NftCardCategoryData{
			OwnerId:      &oid,
			Category:     celebrity.Category,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	err = store.NewNftCardCategory(ctx, &card_category)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	card_crafting := models.NftCardCrafting{
		NftCardCraftingData: models.NftCardCraftingData{
			Rarity:       &rarity,
			OwnerId:      &oid,
			IsCrafted:    &is_crafted,
			CardSeriesId: &card_series_id,
		},
	}

	err = store.NewNftCardCrafting(ctx, &card_crafting)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, struct {
		NftCardCrafting models.NftCardCrafting
		NftCardYear     models.NftCardYear
		NftCardDayMonth models.NftCardDayMonth
		NftCardCategory models.NftCardCategory
		CelebrityId     int64
	}{
		NftCardCrafting: card_crafting,
		NftCardYear:     card_year,
		NftCardDayMonth: card_day_month,
		NftCardCategory: card_category,
		CelebrityId:     *celebrity.Id,
	})
}
