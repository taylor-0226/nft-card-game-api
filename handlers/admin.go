package handlers

import (
	"gameon-twotwentyk-api/models"
	"net/http"

	"github.com/pandoratoolbox/json"
)

func ReceiveAggPack(w http.ResponseWriter, r *http.Request) {
	input := struct {
		CardCollectionId int64
		AggPackPath      string
	}{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// err = store.UpdateCardCollectionAggPack(ctx, input.CardCollectionId, input.AggPackPath)
	// if err != nil {
	// 	ServeError(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = nft.DeployNewCollection(input)

	//get the agg pack from the request
	//save it to the database
	//save it to the factory contract

	//return the agg pack id
}

func CreateCardCollection(w http.ResponseWriter, r *http.Request) {
	input := models.CollectionCreation{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	input.AggPackPath = nil
	input.ID = nil

	// err = store.NewCardCollection(ctx, &input)
	// if err != nil {
	// 	ServeError(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// CreateCollectionAggPack(input)

	ServeJSON(w, input)

}
