package handlers

import (
	"context"
	"encoding/json"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func CraftIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		NftCardCraftingId int64  `json:"nft_card_crafting_id"`
		NftCardDayMonthId int64  `json:"nft_card_day_month_id"`
		NftCardYearId     int64  `json:"nft_card_year_id"`
		NftCardCategoryId int64  `json:"nft_card_category_id"`
		CelebrityId       *int64 `json:"celebrity_id"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	spew.Dump(input)

	// var out models.NftCardIdentity

	nft_card_day_month, err := store.GetNftCardDayMonth(ctx, input.NftCardDayMonthId)
	if err != nil {
		ServeError(w, "nft_card_day_month: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_year, err := store.GetNftCardYear(ctx, input.NftCardYearId)
	if err != nil {
		ServeError(w, "nft_card_year: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_crafting, err := store.GetNftCardCrafting(ctx, input.NftCardCraftingId)
	if err != nil {
		ServeError(w, "nft_card_crafting: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_category, err := store.GetNftCardCategory(ctx, input.NftCardCategoryId)
	if err != nil {
		ServeError(w, "nft_card_category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//assign date
	day := nft_card_day_month.Day
	month := nft_card_day_month.Month
	year := nft_card_year.Year

	//check rarities
	rarities := make(map[int64]int)

	_, ok := rarities[*nft_card_day_month.Rarity]
	if !ok {
		rarities[*nft_card_day_month.Rarity] = 1
	} else {
		rarities[*nft_card_day_month.Rarity]++
	}

	_, ok = rarities[*nft_card_year.Rarity]
	if !ok {
		rarities[*nft_card_year.Rarity] = 1
	} else {
		rarities[*nft_card_year.Rarity]++
	}

	_, ok = rarities[*nft_card_category.Rarity]
	if !ok {
		rarities[*nft_card_category.Rarity] = 1
	} else {
		rarities[*nft_card_category.Rarity]++
	}

	rarity := int64(0)
	var rar_arr []int
	for k, v := range rarities {
		rar_arr = append(rar_arr, int(k))
		if k == *nft_card_year.Rarity {
			if v >= 2 {
				rarity = k
			}
		}
	}

	sort.Ints(rar_arr)
	lowest_rarity := rar_arr[0]

	if int(rarity)-lowest_rarity > 1 {
		rarity = 0
	}

	category := nft_card_category.Category
	is_crafted := false
	card_series_id := int64(1)
	out := models.NftCardIdentity{
		NftCardIdentityData: models.NftCardIdentityData{
			OwnerId:      &mid,
			Year:         year,
			Month:        month,
			Day:          day,
			Rarity:       &rarity,
			Category:     category,
			IsCrafted:    &is_crafted,
			CardSeriesId: &card_series_id,
		},
	}

	//get celebrity by id
	if input.CelebrityId != nil {
		celebrity, err := store.GetCelebrity(ctx, *input.CelebrityId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *nft_card_year.Year != *celebrity.BirthYear {
			ServeError(w, "Nft year doesn't match celebrity", 400)
			return
		}

		if *nft_card_day_month.Day != *celebrity.BirthDay {
			ServeError(w, "Nft day doesn't match celebrity", 400)
			return
		}

		if *nft_card_day_month.Month != *celebrity.BirthMonth {
			ServeError(w, "Nft month doesn't match celebrity", 400)
			return
		}

		out.CelebrityName = celebrity.Name
	}

	err = store.NewNftCardIdentity(context.TODO(), &out)
	if err != nil {
		ServeError(w, "new nft_card_identity: "+err.Error(), http.StatusInternalServerError)
		return
	}

	burn := true

	nft_card_category.IsCrafted = &burn
	store.UpdateNftCardCategory(ctx, nft_card_category)

	nft_card_crafting.IsCrafted = &burn
	store.UpdateNftCardCrafting(ctx, nft_card_crafting)

	nft_card_day_month.IsCrafted = &burn
	store.UpdateNftCardDayMonth(ctx, nft_card_day_month)

	nft_card_year.IsCrafted = &burn
	store.UpdateNftCardYear(ctx, nft_card_year)

	ServeJSON(w, out)
}

// func CraftDate(w http.ResponseWriter, r *http.Request) {
// 	var out models.NftCardDate

// 	input := struct {
// 		NftCardDayId   int64
// 		NftCardMonthId int64
// 		NftCardYearId  int64
// 	}{}

// 	ServeJSON(out)
// }

func CraftPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	var out models.NftCardPrediction

	input := struct {
		NftCardCraftingId int64   `json:"nft_card_crafting_id"`
		NftCardIdentityId int64   `json:"nft_card_identity_id"`
		NftCardTriggerIds []int64 `json:"nft_card_trigger_ids"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_crafting, err := store.GetNftCardCrafting(ctx, input.NftCardCraftingId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if *nft_card_crafting.OwnerId != mid {
		ServeError(w, "User does not own crafting card", 400)
		return
	}

	nft_card_identity, err := store.GetNftCardIdentity(ctx, input.NftCardIdentityId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var triggers []string
	var nft_card_triggers []models.NftCardTrigger
	trigger_highest_rarity := int64(0)
	for _, t := range input.NftCardTriggerIds {
		nft_card_trigger, err := store.GetNftCardTrigger(ctx, t)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nft_card_triggers = append(nft_card_triggers, nft_card_trigger)

		if *nft_card_trigger.OwnerId != mid {
			ServeError(w, "User does not own trigger card", 400)
			return
		}

		if *nft_card_trigger.Rarity > trigger_highest_rarity {
			trigger_highest_rarity = *nft_card_trigger.Rarity
		}

		triggers = append(triggers, *nft_card_trigger.Trigger)
	}

	rarity := trigger_highest_rarity

	is_claimed := false

	ttriggers := models.Strings(triggers)

	card_series_id := int64(1)

	out = models.NftCardPrediction{
		NftCardPredictionData: models.NftCardPredictionData{
			OwnerId:       &mid,
			IsClaimed:     &is_claimed,
			Rarity:        &rarity,
			Triggers:      &ttriggers,
			CelebrityName: nft_card_identity.CelebrityName,
			CardSeriesId:  &card_series_id,
		},
	}

	err = store.NewNftCardPrediction(ctx, &out)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	burn := true

	ni := models.NftCardIdentity{
		NftCardIdentityData: models.NftCardIdentityData{
			Id:        &input.NftCardIdentityId,
			IsCrafted: &burn,
		},
	}

	store.UpdateNftCardIdentity(ctx, ni)

	nc := models.NftCardCrafting{
		NftCardCraftingData: models.NftCardCraftingData{
			Id:        &input.NftCardCraftingId,
			IsCrafted: &burn,
		},
	}
	store.UpdateNftCardCrafting(ctx, nc)

	for _, t := range nft_card_triggers {
		n := models.NftCardTrigger{
			NftCardTriggerData: models.NftCardTriggerData{
				Id:        t.Id,
				IsCrafted: &burn,
			},
		}
		store.UpdateNftCardTrigger(ctx, n)
	}

	ServeJSON(w, out)
}
