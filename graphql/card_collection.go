package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_card_collection = ReflectToFragment(models.CardCollectionData{})
)

func NewCardCollection(ctx context.Context, data *models.CardCollection) error {
	q := `
		mutation CreateCardCollection {
			card_collection(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.CardCollection
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
		CardCollection []models.CardCollection
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardCollection) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.CardCollection[0].Id

	data.Id = &id

	return nil
}

func DeleteCardCollection(ctx context.Context, id int64) error {
	q := `
		mutation DeleteCardCollection {
			card_collection(where: { id: { eq: $id } }) {
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
		CardCollection []models.CardCollection
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardCollection) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateCardCollection(ctx context.Context, data models.CardCollection) error {
	q := `
		mutation UpdateCardCollection {
			card_collection(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.CardCollection
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
		CardCollection []models.CardCollection
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardCollection) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetCardCollection(ctx context.Context, id int64) (models.CardCollection, error) {
	var data models.CardCollection

	q := fragment_card_collection + `
			query GetCardCollection {
			card_collection(where: { id: { eq: $id } }) {
				...CardCollection
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
		CardCollection []models.CardCollection
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.CardCollection) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.CardCollection[0]

	return data, nil
}

func ListCardCollection(ctx context.Context) ([]models.CardCollection, error) {
	var data []models.CardCollection

	q := fragment_card_collection + `
			query ListCardCollection {
			card_collection {
				...CardCollection
			}
		}
		`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return nil, err
	}

	var out struct {
		CardCollection []models.CardCollection
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return nil, err
	}

	data = out.CardCollection

	return data, nil
}
