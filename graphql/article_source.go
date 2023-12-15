package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_article_source = ReflectToFragment(models.ArticleSourceData{})
)

func NewArticleSource(ctx context.Context, data *models.ArticleSource) error {
	q := `
		mutation CreateArticleSource {
			article_source(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.ArticleSource
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
		ArticleSource []models.ArticleSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.ArticleSource) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.ArticleSource[0].Id

	data.Id = &id

	return nil
}

func DeleteArticleSource(ctx context.Context, id int64) error {
	q := `
		mutation DeleteArticleSource {
			article_source(where: { id: { eq: $id } }) {
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
		ArticleSource []models.ArticleSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.ArticleSource) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateArticleSource(ctx context.Context, data models.ArticleSource) error {
	q := `
		mutation UpdateArticleSource {
			article_source(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.ArticleSource
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
		ArticleSource []models.ArticleSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.ArticleSource) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetArticleSource(ctx context.Context, id int64) (models.ArticleSource, error) {
	var data models.ArticleSource

	q := fragment_article_source + `
			query GetArticleSource {
			article_source(where: { id: { eq: $id } }) {
				...ArticleSource
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
		ArticleSource []models.ArticleSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.ArticleSource) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.ArticleSource[0]

	return data, nil
}
