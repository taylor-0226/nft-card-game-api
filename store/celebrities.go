package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

var CelebrityMap = make(map[int64]models.Celebrity)

func RefreshCelebrityMap(ctx context.Context) error {
	list, err := ListCelebrity(ctx)
	if err != nil {
		return err
	}

	for _, item := range list {
		CelebrityMap[*item.Id] = item
	}

	return nil
}

func NewCelebrity(ctx context.Context, data *models.Celebrity) error {
	err := graphql.NewCelebrity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCelebrity(ctx context.Context, id int64) (models.Celebrity, error) {
	data, err := graphql.GetCelebrity(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCelebrity(ctx context.Context, id int64) error {
	err := graphql.DeleteCelebrity(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCelebrity(ctx context.Context, data models.Celebrity) error {
	err := graphql.UpdateCelebrity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCelebrity(ctx context.Context) ([]models.Celebrity, error) {
	out, err := graphql.ListCelebrity(ctx)
	if err != nil {
		return out, err
	}

	return out, nil
}

func ListCelebrityByArrays(ctx context.Context, day []int, month []int, year []int, category []string) ([]models.Celebrity, error) {
	data, err := graphql.ListCelebrityByArrays(ctx, day, month, year, category)
	if err != nil {
		return nil, err
	}

	return data, nil
}
