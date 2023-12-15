package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetNftCardYear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_year_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardYear(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardYear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.NftCardYear{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rarity := int64(rand.Intn(2))
	is_crafted := false

	celebrities, err := store.ListCelebrity(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	i := rand.Intn(len(celebrities) - 1)

	celebrity := celebrities[i]

	oid := ctx.Value(models.CTX_user_id).(int64)
	if input.OwnerId != nil {
		oid = *input.OwnerId
	}

	new := models.NftCardYear{
		NftCardYearData: models.NftCardYearData{
			OwnerId:   &oid,
			Year:      celebrity.BirthYear,
			IsCrafted: &is_crafted,
			Rarity:    &rarity,
		},
	}

	err = store.NewNftCardYear(ctx, &new)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, new)
}

func UpdateNftCardYear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.NftCardYear{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateNftCardYear(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardYear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_year_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardYear(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardYearForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	filters := models.NftCardYearFilter{}

	q_card_series_id := r.URL.Query().Get("card_series_id")
	if q_card_series_id != "" {
		i, err := strconv.ParseInt(q_card_series_id, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		filters.CardSeriesId = &i
	}

	q_year := r.URL.Query().Get("year")
	if q_year != "" {
		i, err := strconv.ParseInt(q_year, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		filters.Year = &i
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

	data, err := store.ListNftCardYearByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
