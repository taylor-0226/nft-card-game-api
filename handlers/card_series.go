package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/nft"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetCardSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_series_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetCardSeries(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewCardSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.CardSeries{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewCardSeries(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateCardSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.CardSeries{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateCardSeries(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteCardSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_series_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteCardSeries(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

// func InitiateCardSeriesOrder(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	q := chi.URLParam(r, "card_series_id")
// 	id, err := strconv.ParseInt(q, 10, 64)
// 	if err != nil {
// 		ServeError(w, err.Error(), 500)
// 		return
// 	}

// 	data, err := store.GetCardSeries(ctx, id)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	cost := data.CostUSD

// 	packs, err := store.ListCardPackByCardSeriesId(ctx, id)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if len(packs)+1 > int(*data.Quantity) {
// 		ServeError(w, "There are no card packs left for this card series", http.StatusNotAcceptable)
// 		return
// 	}

// 	//initiate moonpay payment with cost from card series

// 	//ServeJSON(res.payment_token)

// }

func CompleteCardSeriesOrder(w http.ResponseWriter, r *http.Request) {

}

func BuyCardSeriesPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	q := chi.URLParam(r, "card_series_id")
	card_series_id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	input := struct {
		PaymentMethodId int64
		Quantity        int
	}{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	if input.PaymentMethodId == 0 {
		ServeError(w, "Invalid payment_method_id", 400)
		return
	}

	series, err := store.GetCardSeries(ctx, card_series_id)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	// count, err := store.GetCardPackCountByCardSeriesId(ctx, input.CardSeriesId)
	// if err != nil {
	// 	ServeError(w, err.Error(), 400)
	// 	return
	// }

	// if series.Quantity >= count {
	// 	ServeError(w, "There are no more card packs left in this card series", 400)
	// 	return
	// }
	var card_packs []models.CardPack
	for i := 0; i < input.Quantity; i++ {
		cards, err := nft.GenerateCardPackCards(ctx, 1)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		card_pack := models.CardPack{
			CardPackData: models.CardPackData{
				OwnerId:      &mid,
				CardSeriesId: &card_series_id,
				Cards:        &cards,
			},
		}

		switch input.PaymentMethodId {
		case 1:
			err := store.NewCardPack(ctx, &card_pack)

			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			err = store.AddBalanceToUser(ctx, mid, 0-*series.CostUsd)
			if err != nil {
				ServeError(w, err.Error(), 400)
				return
			}

		}

		card_pack.CardSeries = &series

		card_packs = append(card_packs, card_pack)
	}

	ServeJSON(w, card_packs)
}

func ListCardSeries(w http.ResponseWriter, r *http.Request) {
	data, err := store.ListCardSeries(r.Context())
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}
