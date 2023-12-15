package handlers

import (
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetNftCardCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_category_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardCategory(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.NftCardCategory{}

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

	new := models.NftCardCategory{
		NftCardCategoryData: models.NftCardCategoryData{
			OwnerId:   &oid,
			Category:  celebrity.Category,
			IsCrafted: &is_crafted,
			Rarity:    &rarity,
		},
	}

	err = store.NewNftCardCategory(ctx, &new)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, new)
}

func UpdateNftCardCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.NftCardCategory{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateNftCardCategory(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_category_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardCategory(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardCategoryForUserById(w http.ResponseWriter, r *http.Request) {
	var err error

	var filters models.NftCardCategoryFilter

	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	var categories models.Strings
	categories_raw := r.URL.Query().Get("categories")
	if categories_raw != "" {
		categories_str := strings.Split(categories_raw, ",")
		for _, v := range categories_str {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c, ok := store.CategoryMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid category id: %d", i), http.StatusBadRequest)
				return
			}
			categories = append(categories, *c.Name)
		}
		filters.Categories = &categories
	}

	var card_series_id int64
	card_series_id_raw := r.URL.Query().Get("card_series_id")
	if card_series_id_raw != "" {
		card_series_id, err = strconv.ParseInt(card_series_id_raw, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filters.CardSeriesId = &card_series_id
	}

	var rarities models.Ints
	rarities_raw := r.URL.Query().Get("rarities")
	if rarities_raw != "" {
		rarities_str := strings.Split(rarities_raw, ",")
		for _, v := range rarities_str {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			rarities = append(rarities, i)
		}
		filters.Rarities = &rarities
	}

	var status models.Ints
	status_raw := r.URL.Query().Get("status")
	if status_raw != "" {
		status_str := strings.Split(status_raw, ",")
		for _, v := range status_str {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			status = append(status, i)
		}
		filters.Status = &status
	}

	data, err := store.ListNftCardCategoryByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
