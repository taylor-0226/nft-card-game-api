package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardTrigger(ctx context.Context, data *models.NftCardTrigger) error {
	err := graphql.NewNftCardTrigger(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardTrigger(ctx context.Context, id int64) (models.NftCardTrigger, error) {
	data, err := graphql.GetNftCardTrigger(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardTrigger(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardTrigger(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardTrigger(ctx context.Context, data models.NftCardTrigger) error {
	err := graphql.UpdateNftCardTrigger(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardTriggerByOwnerId(ctx context.Context, id int64, filters models.NftCardTriggerFilter) ([]models.NftCardTrigger, error) {
	data, err := graphql.ListNftCardTriggerByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
