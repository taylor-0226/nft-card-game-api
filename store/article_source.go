package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewArticleSource(ctx context.Context, data *models.ArticleSource) error {
	err := graphql.NewArticleSource(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetArticleSource(ctx context.Context, id int64) (models.ArticleSource, error) {
	data, err := graphql.GetArticleSource(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteArticleSource(ctx context.Context, id int64) error {
	err := graphql.DeleteArticleSource(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArticleSource(ctx context.Context, data models.ArticleSource) error {
	err := graphql.UpdateArticleSource(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
