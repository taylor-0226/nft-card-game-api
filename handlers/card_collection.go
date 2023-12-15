package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetCardCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_collection_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetCardCollection(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

// func NewCardCollection(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	input := models.CardCollection{}

// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&input)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = store.NewCardCollection(ctx, &input)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	ServeJSON(w, input)
// }

func UpdateCardCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.CardCollection{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateCardCollection(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteCardCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_collection_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteCardCollection(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListCardCollection(w http.ResponseWriter, r *http.Request) {
	data, err := store.ListCardCollection(r.Context())
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, data)
}
