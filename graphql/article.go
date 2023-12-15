package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_article = ReflectToFragment(models.ArticleData{})
)

func NewArticle(ctx context.Context, data *models.Article) error {
	q := `
		mutation CreateArticle {
			article(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Article
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
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Article) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Article[0].Id

	data.Id = &id

	return nil
}

func DeleteArticle(ctx context.Context, id int64) error {
	q := `
		mutation DeleteArticle {
			article(where: { id: { eq: $id } }) {
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
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Article) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateArticle(ctx context.Context, data models.Article) error {
	q := `
		mutation UpdateArticle {
			article(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Article
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
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Article) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetArticle(ctx context.Context, id int64) (models.Article, error) {
	var data models.Article

	q := fragment_article + `
			query GetArticle {
			article(where: { id: { eq: $id } }) {
				...Article
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
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Article) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Article[0]

	return data, nil
}

func SearchArticles(ctx context.Context, q string) ([]models.Article, error) {
	var data []models.Article
	gq := fragment_article + `
	query SearchArticles {
	article(where: { tags: { contains: $q } }) {
		...Article
	}
}
`

	input := struct {
		Q string `json:"q"`
	}{
		Q: q,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, gq, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	data = out.Article

	return data, nil
}

func ListArticles(ctx context.Context) ([]models.Article, error) {
	var data []models.Article
	gq := fragment_article + `
	query ListArticles {
	article() {
		...Article
	}
}
`

	res, err := Graph.GraphQL(ctx, gq, nil, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Article []models.Article
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	data = out.Article

	return data, nil
}
