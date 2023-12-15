package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

var TriggerMap = make(map[int64]models.Trigger)

func RefreshTriggerMap(ctx context.Context) error {
	list, err := ListTrigger(ctx)
	if err != nil {
		return err
	}

	for _, item := range list {
		TriggerMap[*item.Id] = item
	}

	return nil
}

func ListTrigger(ctx context.Context) ([]models.Trigger, error) {
	data, err := graphql.ListTrigger(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewTrigger(ctx context.Context, data *models.Trigger) error {
	err := graphql.NewTrigger(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetTrigger(ctx context.Context, id int64) (models.Trigger, error) {
	data, err := graphql.GetTrigger(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteTrigger(ctx context.Context, id int64) error {
	err := graphql.DeleteTrigger(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTrigger(ctx context.Context, data models.Trigger) error {
	err := graphql.UpdateTrigger(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
