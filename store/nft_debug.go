package store

import (
	"context"
	"gameon-twotwentyk-api/models"
	"math/rand"
)

func AddNftsToUser(ctx context.Context, user_id int64) error {

	oid := user_id

	rarity := int64(rand.Intn(2))
	is_crafted := false
	card_series_id := int64(1)

	celebrities, err := ListCelebrity(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		i := rand.Intn(len(celebrities) - 1)

		celebrity := celebrities[i]

		card_year := models.NftCardYear{
			NftCardYearData: models.NftCardYearData{
				OwnerId:      &oid,
				Year:         celebrity.BirthYear,
				IsCrafted:    &is_crafted,
				Rarity:       &rarity,
				CardSeriesId: &card_series_id,
			},
		}

		err = NewNftCardYear(ctx, &card_year)
		if err != nil {
			return err
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

		err = NewNftCardDayMonth(ctx, &card_day_month)
		if err != nil {
			return err
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

		err = NewNftCardCategory(ctx, &card_category)
		if err != nil {
			return err
		}

		card_crafting := models.NftCardCrafting{
			NftCardCraftingData: models.NftCardCraftingData{
				Rarity:       &rarity,
				OwnerId:      &oid,
				IsCrafted:    &is_crafted,
				CardSeriesId: &card_series_id,
			},
		}

		err = NewNftCardCrafting(ctx, &card_crafting)
		if err != nil {
			return err
		}
	}

	triggers, err := ListTrigger(ctx)
	if err != nil {
		return err
	}

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
		err := NewNftCardTrigger(ctx, &new)
		if err != nil {
			return err
		}
		out = append(out, new)
	}

	return nil
}
