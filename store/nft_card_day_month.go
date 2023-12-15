package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardDayMonth(ctx context.Context, data *models.NftCardDayMonth) error {
	err := graphql.NewNftCardDayMonth(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardDayMonth(ctx context.Context, id int64) (models.NftCardDayMonth, error) {
	data, err := graphql.GetNftCardDayMonth(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardDayMonth(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardDayMonth(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardDayMonth(ctx context.Context, data models.NftCardDayMonth) error {
	err := graphql.UpdateNftCardDayMonth(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardDayMonthByOwnerId(ctx context.Context, id int64, filters models.NftCardDayMonthFilter) ([]models.NftCardDayMonth, error) {
	data, err := graphql.ListNftCardDayMonthByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
