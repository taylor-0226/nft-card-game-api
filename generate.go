package main

import (
	"context"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"log"

	"github.com/XANi/loremipsum"
)

func GenerateArticles(limit int) {
	gen := loremipsum.New()

	for i := 0; i < limit; i++ {
		title := gen.Sentence()
		except := gen.Paragraph()
		url := "http://" + gen.Word() + ".com"
		as := int64(1)
		tags := models.Strings{gen.Word()}

		new := models.Article{
			ArticleData: models.ArticleData{
				Title:           &title,
				Excerpt:         &except,
				Url:             &url,
				ArticleSourceId: &as,
				Tags:            &tags,
			},
		}

		err := store.NewArticle(context.TODO(), &new)
		if err != nil {
			log.Fatalf("Unable to insert article: %s", err)
		}
	}

}
