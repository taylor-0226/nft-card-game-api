package handlers

import (
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetNftCardPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_prediction_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardPrediction(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.NftCardPrediction{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewNftCardPrediction(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateNftCardPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.NftCardPrediction{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateNftCardPrediction(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_prediction_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardPrediction(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardPredictionForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	filters := models.NftCardPredictionFilter{}

	q_card_series_id := r.URL.Query().Get("card_series_id")
	if q_card_series_id != "" {
		i, err := strconv.ParseInt(q_card_series_id, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		filters.CardSeriesId = &i
	}

	q_rarities := r.URL.Query().Get("rarities")
	if q_rarities != "" {
		var rarities models.Ints
		for _, v := range strings.Split(q_rarities, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			rarities = append(rarities, i)
		}

		filters.Rarities = &rarities
	}

	q_triggers := r.URL.Query().Get("triggers")
	if q_triggers != "" {
		var triggers models.Strings
		for _, v := range strings.Split(q_triggers, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			t, ok := store.TriggerMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid trigger id: %d", i), http.StatusBadRequest)
				return
			}

			triggers = append(triggers, *t.Name)
		}

		filters.Triggers = &triggers
	}

	q_celebrities := r.URL.Query().Get("celebrities")
	if q_celebrities != "" {
		var celebrities models.Strings
		for _, v := range strings.Split(q_celebrities, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			c, ok := store.CelebrityMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid celebrity id: %d", i), http.StatusBadRequest)
				return
			}

			celebrities = append(celebrities, *c.Name)
		}

		filters.Celebrities = &celebrities
	}

	q_status := r.URL.Query().Get("status")
	if q_status != "" {
		var status models.Ints
		for _, v := range strings.Split(q_status, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			status = append(status, i)
		}

		filters.Status = &status
	}

	data, err := store.ListNftCardPredictionByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
