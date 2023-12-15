package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewCardSeries(ctx context.Context, data *models.CardSeries) error {
	err := graphql.NewCardSeries(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCardSeries(ctx context.Context, id int64) (models.CardSeries, error) {
	data, err := graphql.GetCardSeries(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCardSeries(ctx context.Context, id int64) error {
	err := graphql.DeleteCardSeries(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCardSeries(ctx context.Context, data models.CardSeries) error {
	err := graphql.UpdateCardSeries(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCardSeries(ctx context.Context) ([]models.CardSeries, error) {
	return graphql.ListCardSeries(ctx)
}
