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

func GetNftCardDayMonth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_day_month_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardDayMonth(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardDayMonth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.NftCardDayMonth{}

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

	new := models.NftCardDayMonth{
		NftCardDayMonthData: models.NftCardDayMonthData{
			OwnerId:   &oid,
			Month:     celebrity.BirthMonth,
			Day:       celebrity.BirthDay,
			IsCrafted: &is_crafted,
			Rarity:    &rarity,
		},
	}

	err = store.NewNftCardDayMonth(ctx, &new)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, new)
}

func UpdateNftCardDayMonth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.NftCardDayMonth{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateNftCardDayMonth(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardDayMonth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_day_month_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardDayMonth(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardDayMonthForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	filters := models.NftCardDayMonthFilter{}

	q_card_series_id := r.URL.Query().Get("card_series_id")
	if q_card_series_id != "" {
		i, err := strconv.ParseInt(q_card_series_id, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		filters.CardSeriesId = &i
	}

	q_day := r.URL.Query().Get("day")
	if q_day != "" {
		i, err := strconv.ParseInt(q_day, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		filters.Day = &i
	}

	q_month := r.URL.Query().Get("month")
	if q_month != "" {
		i, err := strconv.ParseInt(q_month, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		filters.Month = &i
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

		filters.Rarities = &status
	}

	data, err := store.ListNftCardDayMonthByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
