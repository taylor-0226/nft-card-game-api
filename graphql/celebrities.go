package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_celebrity = ReflectToFragment(models.CelebrityData{})
)

func NewCelebrity(ctx context.Context, data *models.Celebrity) error {
	q := `
		mutation CreateCelebrity {
			celebrity(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Celebrity
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
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Celebrity) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Celebrity[0].Id

	data.Id = &id

	return nil
}

func DeleteCelebrity(ctx context.Context, id int64) error {
	q := `
		mutation DeleteCelebrity {
			celebrity(where: { id: { eq: $id } }) {
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
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Celebrity) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateCelebrity(ctx context.Context, data models.Celebrity) error {
	q := `
		mutation UpdateCelebrity {
			celebrity(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Celebrity
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
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Celebrity) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetCelebrity(ctx context.Context, id int64) (models.Celebrity, error) {
	var data models.Celebrity

	q := fragment_celebrity + `
			query GetCelebrity {
			celebrity(where: { id: { eq: $id } }) {
				...Celebrity
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
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Celebrity) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Celebrity[0]

	return data, nil
}

func ListCelebrity(ctx context.Context) ([]models.Celebrity, error) {
	var data []models.Celebrity

	q := fragment_celebrity + `
			query ListCelebrity {
			celebrity() {
				...Celebrity
			}
		}
		`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	data = out.Celebrity

	return data, nil
}

func ListCelebrityByArrays(ctx context.Context, day []int, month []int, year []int, category []string) ([]models.Celebrity, error) {
	var data []models.Celebrity

	q := fragment_celebrity + `query ListCelebrityByArrays {
		celebrity(where: { and: { birth_day: { in: $days }, birth_month: { in: $months }, birth_year: { in: $years }, categories: { in: $categories } }  }) {
			...Celebrity
		}
	}`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Celebrity []models.Celebrity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	data = out.Celebrity

	return data, nil
}
