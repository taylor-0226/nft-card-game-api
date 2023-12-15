package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardCrafting(ctx context.Context, data *models.NftCardCrafting) error {
	err := graphql.NewNftCardCrafting(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardCrafting(ctx context.Context, id int64) (models.NftCardCrafting, error) {
	data, err := graphql.GetNftCardCrafting(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardCrafting(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardCrafting(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardCrafting(ctx context.Context, data models.NftCardCrafting) error {
	err := graphql.UpdateNftCardCrafting(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardCraftingByOwnerId(ctx context.Context, id int64, filters models.NftCardCraftingFilter) ([]models.NftCardCrafting, error) {
	data, err := graphql.ListNftCardCraftingByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
