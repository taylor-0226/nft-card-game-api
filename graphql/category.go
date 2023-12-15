package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_category = ReflectToFragment(models.CategoryData{})
)

func NewCategory(ctx context.Context, data *models.Category) error {
	q := `
		mutation CreateCategory {
			category(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Category
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
		Category []models.Category
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Category) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Category[0].Id

	data.Id = &id

	return nil
}

func DeleteCategory(ctx context.Context, id int64) error {
	q := `
		mutation DeleteCategory {
			category(where: { id: { eq: $id } }) {
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
		Category []models.Category
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Category) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateCategory(ctx context.Context, data models.Category) error {
	q := `
		mutation UpdateCategory {
			category(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Category
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
		Category []models.Category
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Category) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetCategory(ctx context.Context, id int64) (models.Category, error) {
	var data models.Category

	q := fragment_category + `
			query GetCategory {
			category(where: { id: { eq: $id } }) {
				...Category
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
		Category []models.Category
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Category) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Category[0]

	return data, nil
}

func ListCategory(ctx context.Context) ([]models.Category, error) {
	var data []models.Category

	out := struct {
		Category []models.Category
	}{}

	q := fragment_category + `query ListCategory {
		category() {
			...Category
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

	data = out.Category

	return data, nil
}
