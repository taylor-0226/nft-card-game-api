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

func GetNftCardTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_trigger_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardTrigger(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := struct {
		models.NftCardTrigger
		Generate bool
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	is_crafted := false
	oid := int64(1)

	if input.OwnerId != nil {
		oid = *input.OwnerId
	} else {
		oid = ctx.Value(models.CTX_user_id).(int64)
	}

	triggers, err := store.ListTrigger(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
	}

	i := rand.Intn(len(triggers) - 1)

	trigger := triggers[i]

	rarity := int64(rand.Intn(2))

	new := models.NftCardTrigger{
		NftCardTriggerData: models.NftCardTriggerData{
			IsCrafted: &is_crafted,
			Tier:      trigger.Tier,
			Rarity:    &rarity,
			OwnerId:   &oid,
			Trigger:   trigger.Name,
		},
	}

	err = store.NewNftCardTrigger(ctx, &new)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, new)
}

func UpdateNftCardTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.NftCardTrigger{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateNftCardTrigger(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_trigger_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardTrigger(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardTriggerForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	filters := models.NftCardTriggerFilter{}

	q_card_series_id := r.URL.Query().Get("card_series_id")
	if q_card_series_id != "" {
		i, err := strconv.ParseInt(q_card_series_id, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		filters.CardSeriesId = &i
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

			c, ok := store.TriggerMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid trigger id: %d", i), http.StatusBadRequest)
				return
			}
			triggers = append(triggers, *c.Name)

		}

		filters.Triggers = &triggers
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

	q_tiers := r.URL.Query().Get("tiers")
	if q_tiers != "" {
		tiers := models.Strings(strings.Split(q_tiers, ","))
		filters.Tiers = &tiers
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

	data, err := store.ListNftCardTriggerByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
