package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_card_series = ReflectToFragment(models.CardSeriesData{})
)

func NewCardSeries(ctx context.Context, data *models.CardSeries) error {
	q := `
		mutation CreateCardSeries {
			card_series(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.CardSeries
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
		CardSeries []models.CardSeries
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardSeries) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.CardSeries[0].Id

	data.Id = &id

	return nil
}

func DeleteCardSeries(ctx context.Context, id int64) error {
	q := `
		mutation DeleteCardSeries {
			card_series(where: { id: { eq: $id } }) {
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
		CardSeries []models.CardSeries
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardSeries) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateCardSeries(ctx context.Context, data models.CardSeries) error {
	q := `
		mutation UpdateCardSeries {
			card_series(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.CardSeries
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
		CardSeries []models.CardSeries
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardSeries) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetCardSeries(ctx context.Context, id int64) (models.CardSeries, error) {
	var data models.CardSeries

	q := fragment_card_series + fragment_card_collection + `
			query GetCardSeries {
			card_series(where: { id: { eq: $id } }) {
				...CardSeries
				card_collection {
					...CardCollection
				}
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
		CardSeries []models.CardSeries
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.CardSeries) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.CardSeries[0]

	return data, nil
}

func ListCardSeries(ctx context.Context) ([]models.CardSeries, error) {
	var data []models.CardSeries

	q := fragment_card_series + fragment_card_collection + `
			query ListCardSeries {
			card_series {
				...CardSeries
				card_collection {
					...CardCollection
				}
			}
		}
		`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		CardSeries []models.CardSeries
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	data = out.CardSeries

	return data, nil
}
