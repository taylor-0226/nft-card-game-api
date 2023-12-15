package nft

import (
	"context"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"math/rand"
)

func GenerateCardPackCards(ctx context.Context, card_series_id int64) (models.CardPackCards, error) {
	out := models.CardPackCards{}

	rarity := int64(rand.Intn(2))
	is_crafted := false

	celebrities, err := store.ListCelebrity(ctx)
	if err != nil {
		return out, err
	}

	i := rand.Intn(len(celebrities) - 1)

	celebrity := celebrities[i]

	card_year := models.NftCardYear{
		NftCardYearData: models.NftCardYearData{
			// OwnerId:      &oid,
			Year:         celebrity.BirthYear,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	card_day_month := models.NftCardDayMonth{
		NftCardDayMonthData: models.NftCardDayMonthData{
			// OwnerId:      &oid,
			Day:          celebrity.BirthDay,
			Month:        celebrity.BirthMonth,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	card_category := models.NftCardCategory{
		NftCardCategoryData: models.NftCardCategoryData{
			// OwnerId:      &oid,
			Category:     celebrity.Category,
			IsCrafted:    &is_crafted,
			Rarity:       &rarity,
			CardSeriesId: &card_series_id,
		},
	}

	card_crafting := models.NftCardCrafting{
		NftCardCraftingData: models.NftCardCraftingData{
			Rarity: &rarity,
			// OwnerId:      &oid,
			IsCrafted:    &is_crafted,
			CardSeriesId: &card_series_id,
		},
	}

	out.Crafting = append(out.Crafting, card_crafting)
	out.Category = append(out.Category, card_category)
	out.DayMonth = append(out.DayMonth, card_day_month)
	out.Year = append(out.Year, card_year)

	return out, nil
}
