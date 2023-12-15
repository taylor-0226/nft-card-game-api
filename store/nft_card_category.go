package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardCategory(ctx context.Context, data *models.NftCardCategory) error {
	err := graphql.NewNftCardCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardCategory(ctx context.Context, id int64) (models.NftCardCategory, error) {
	data, err := graphql.GetNftCardCategory(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardCategory(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardCategory(ctx context.Context, data models.NftCardCategory) error {
	err := graphql.UpdateNftCardCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardCategoryByOwnerId(ctx context.Context, id int64, filters models.NftCardCategoryFilter) ([]models.NftCardCategory, error) {
	data, err := graphql.ListNftCardCategoryByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
