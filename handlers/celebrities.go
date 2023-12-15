package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetCelebrity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "celebrity_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetCelebrity(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewCelebrity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.Celebrity{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewCelebrity(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateCelebrity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.Celebrity{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateCelebrity(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteCelebrity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "celebrity_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteCelebrity(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListCelebrity(w http.ResponseWriter, r *http.Request) {
	data, err := store.ListCelebrity(r.Context())
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, data)
}

func GetAvailableCelebrityRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	nft_card_day_month, err := store.ListNftCardDayMonthByOwnerId(ctx, mid, models.NftCardDayMonthFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_year, err := store.ListNftCardYearByOwnerId(ctx, mid, models.NftCardYearFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_category, err := store.ListNftCardCategoryByOwnerId(ctx, mid, models.NftCardCategoryFilter{})
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var day_arr []int
	var month_arr []int
	var year_arr []int
	var category_arr []string

	for _, n := range nft_card_day_month {
		day_arr = append(day_arr, int(*n.Day))
		month_arr = append(month_arr, int(*n.Month))
	}

	for _, n := range nft_card_year {
		year_arr = append(year_arr, int(*n.Year))
	}

	for _, n := range nft_card_category {
		category_arr = append(category_arr, *n.Category)
	}

	celebrities, err := store.ListCelebrityByArrays(ctx, day_arr, month_arr, year_arr, category_arr)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, celebrities)

}

// func GetCelebrityMatch(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	input := struct {
// 		Category string
// 		Month    int
// 		Year     int
// 		Day      int
// 	}{}

// 	decoder :=

// 	matches, err := store.GetCelebrityMatch(ctx, category, day, month, year)
// }
