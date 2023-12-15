package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var fragment_trigger = ReflectToFragment(models.TriggerData{})

func ListTrigger(ctx context.Context) ([]models.Trigger, error) {
	out := struct {
		Trigger []models.Trigger
	}{}

	q := fragment_trigger + `query ListTrigger {
		trigger() {
			...Trigger
		}
	}`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return nil, err
	}

	data := out.Trigger

	return data, nil
}

func NewTrigger(ctx context.Context, data *models.Trigger) error {
	q := `
		mutation CreateTrigger {
			trigger(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Trigger
	}{
		Data: *data,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Trigger []models.Trigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Trigger) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Trigger[0].Id

	data.Id = &id

	return nil
}

func DeleteTrigger(ctx context.Context, id int64) error {
	q := `
		mutation DeleteTrigger {
			trigger(where: { id: { eq: $id } }) {
				id
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Trigger []models.Trigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Trigger) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateTrigger(ctx context.Context, data models.Trigger) error {
	q := `
		mutation UpdateTrigger {
			trigger(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Trigger
	}{
		Id: *data.Id,
	}

	data.Id = nil
	input.Data = data

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Trigger []models.Trigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Trigger) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetTrigger(ctx context.Context, id int64) (models.Trigger, error) {
	var data models.Trigger

	q := fragment_trigger + `
			query GetTrigger {
			trigger(where: { id: { eq: $id } }) {
				...Trigger
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Trigger []models.Trigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Trigger) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Trigger[0]

	return data, nil
}
