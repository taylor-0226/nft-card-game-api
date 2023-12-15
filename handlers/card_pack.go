package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetCardPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_pack_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetCardPack(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewCardPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.CardPack{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewCardPack(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateCardPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.CardPack{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateCardPack(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteCardPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "card_pack_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteCardPack(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListCardPackForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListCardPackByOwnerId(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}

func OpenCardPack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	q := chi.URLParam(r, "card_pack_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetCardPack(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	if *data.IsOpened {
		ServeError(w, "Card pack has already been opened", 400)
		return
	}

	if *data.OwnerId != mid {
		ServeError(w, "User does not own this card pack", 401)
		return
	}

	for _, c := range data.Cards.Category {
		c.OwnerId = &mid
		err := store.NewNftCardCategory(ctx, &c)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	for _, c := range data.Cards.Crafting {
		c.OwnerId = &mid
		err := store.NewNftCardCrafting(ctx, &c)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	for _, c := range data.Cards.Trigger {
		c.OwnerId = &mid
		err := store.NewNftCardTrigger(ctx, &c)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	for _, c := range data.Cards.DayMonth {
		c.OwnerId = &mid
		err := store.NewNftCardDayMonth(ctx, &c)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	for _, c := range data.Cards.Year {
		c.OwnerId = &mid
		err := store.NewNftCardYear(ctx, &c)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	opened := true
	err = store.UpdateCardPack(ctx, models.CardPack{
		CardPackData: models.CardPackData{
			Id:       data.Id,
			IsOpened: &opened,
		},
	})
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	js, err := json.Marshal(data.Cards)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	w.Write(js)

}
