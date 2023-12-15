package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardYear(ctx context.Context, data *models.NftCardYear) error {
	err := graphql.NewNftCardYear(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardYear(ctx context.Context, id int64) (models.NftCardYear, error) {
	data, err := graphql.GetNftCardYear(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardYear(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardYear(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardYear(ctx context.Context, data models.NftCardYear) error {
	err := graphql.UpdateNftCardYear(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardYearByOwnerId(ctx context.Context, id int64, filters models.NftCardYearFilter) ([]models.NftCardYear, error) {
	data, err := graphql.ListNftCardYearByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
