package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardIdentity(ctx context.Context, data *models.NftCardIdentity) error {
	err := graphql.NewNftCardIdentity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardIdentity(ctx context.Context, id int64) (models.NftCardIdentity, error) {
	data, err := graphql.GetNftCardIdentity(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardIdentity(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardIdentity(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardIdentity(ctx context.Context, data models.NftCardIdentity) error {
	err := graphql.UpdateNftCardIdentity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardIdentityByOwnerId(ctx context.Context, id int64, filters models.NftCardIdentityFilter) ([]models.NftCardIdentity, error) {
	data, err := graphql.ListNftCardIdentityByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
