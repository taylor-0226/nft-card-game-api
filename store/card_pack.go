package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewCardPack(ctx context.Context, data *models.CardPack) error {
	err := graphql.NewCardPack(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCardPack(ctx context.Context, id int64) (models.CardPack, error) {
	data, err := graphql.GetCardPack(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCardPack(ctx context.Context, id int64) error {
	err := graphql.DeleteCardPack(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCardPack(ctx context.Context, data models.CardPack) error {
	err := graphql.UpdateCardPack(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCardPackByCardSeriesId(ctx context.Context, id int64) ([]models.CardPack, error) {
	data, err := graphql.ListCardPackByCardSeriesId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func ListCardPackByOwnerId(ctx context.Context, id int64) ([]models.CardPack, error) {
	data, err := graphql.ListCardPackByOwnerId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}
