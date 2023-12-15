package handlers

import (
	"net/http"

	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"

	"github.com/pandoratoolbox/json"
)

const PINATA_ACCESS_TOKEN = "DmLxLXKWM6V3doP4C9YfMD0u1mYJvTmsmfHYJfvj9ShHRHVOqWud0w6irfCDGDyo"

func CreateNftCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := struct {
		Name       string
		CardSeries []models.CardSeriesData
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	card_collection := models.CardCollection{
		CardCollectionData: models.CardCollectionData{
			Name: &input.Name,
		},
	}

	err = store.NewCardCollection(ctx, &card_collection)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	var card_series []*models.CardSeries

	for _, v := range input.CardSeries {
		v.CardCollectionId = card_collection.Id

		o := models.CardSeries{
			CardSeriesData: v,
		}

		err := store.NewCardSeries(ctx, &o)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		card_series = append(card_series, &o)

		//send config to engine
		//await engine to produce agg_pack
		//generate nft card packs from agg_pack json
	}

	card_collection.CardSeries = card_series

	ServeJSON(w, card_collection)
}
