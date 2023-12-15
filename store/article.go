package store

import (
	"context"
	"gameon-twotwentyk-api/connections"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewArticle(ctx context.Context, data *models.Article) error {
	err := graphql.NewArticle(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetArticle(ctx context.Context, id int64) (models.Article, error) {
	data, err := graphql.GetArticle(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteArticle(ctx context.Context, id int64) error {
	err := graphql.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArticle(ctx context.Context, data models.Article) error {
	err := graphql.UpdateArticle(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetArticlesByTagsAndExcerpt(ctx context.Context, q string) ([]models.Article, error) {
	pgq := "SELECT id, article_source_id, excerpt, url, created_at, thumbnail_src, title, tags FROM public.article WHERE $1 = ANY(tags) OR excerpt ILIKE $1"
	res, err := connections.Postgres.Query(pgq, q)
	if err != nil {
		return nil, err
	}

	var data []models.Article
	for res.Next() {
		var new models.Article
		err = res.Scan(&new.Id, &new.ArticleSourceId, &new.Excerpt, &new.Url, &new.CreatedAt, &new.ThumbnailSrc, &new.Title, *new.Tags)
		if err != nil {
			return nil, err
		}
		data = append(data, new)
	}

	return data, nil
}

func GetArticlesNewest(ctx context.Context) ([]models.Article, error) {
	pgq := "SELECT id, article_source_id, excerpt, url, created_at, thumbnail_src, title FROM public.article"
	res, err := connections.Postgres.Query(pgq)
	if err != nil {
		return nil, err
	}

	var data []models.Article
	for res.Next() {
		var new models.Article
		err = res.Scan(&new.Id, &new.ArticleSourceId, &new.Excerpt, &new.Url, &new.CreatedAt, &new.ThumbnailSrc, &new.Title)
		if err != nil {
			return nil, err
		}
		data = append(data, new)
	}

	return data, nil
}

func SearchArticles(ctx context.Context, q string) ([]models.Article, error) {
	data, err := graphql.SearchArticles(ctx, q)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ListArticles(ctx context.Context) ([]models.Article, error) {
	data, err := graphql.ListArticles(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
