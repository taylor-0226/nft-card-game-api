package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNftCardPrediction(ctx context.Context, data *models.NftCardPrediction) error {
	err := graphql.NewNftCardPrediction(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNftCardPrediction(ctx context.Context, id int64) (models.NftCardPrediction, error) {
	data, err := graphql.GetNftCardPrediction(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNftCardPrediction(ctx context.Context, id int64) error {
	err := graphql.DeleteNftCardPrediction(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNftCardPrediction(ctx context.Context, data models.NftCardPrediction) error {
	err := graphql.UpdateNftCardPrediction(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNftCardPredictionByOwnerId(ctx context.Context, id int64, filters models.NftCardPredictionFilter) ([]models.NftCardPrediction, error) {
	data, err := graphql.ListNftCardPredictionByOwnerId(ctx, id, filters)
	if err != nil {
		return data, err
	}

	return data, nil
}
