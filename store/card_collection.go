package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewCardCollection(ctx context.Context, data *models.CardCollection) error {
	err := graphql.NewCardCollection(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCardCollection(ctx context.Context, id int64) (models.CardCollection, error) {
	data, err := graphql.GetCardCollection(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCardCollection(ctx context.Context, id int64) error {
	err := graphql.DeleteCardCollection(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCardCollection(ctx context.Context, data models.CardCollection) error {
	err := graphql.UpdateCardCollection(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCardCollection(ctx context.Context) ([]models.CardCollection, error) {
	data, err := graphql.ListCardCollection(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
