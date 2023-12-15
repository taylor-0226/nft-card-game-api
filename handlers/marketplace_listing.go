package handlers

import (
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetMarketplaceListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "marketplace_listing_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetMarketplaceListing(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewMarketplaceListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := models.MarketplaceListing{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if input.NftTypeId == nil {
		ServeError(w, "nft_type_id must be specified", 400)
		return
	}

	if input.Price == nil {
		ServeError(w, "price must be specified", 400)
		return
	}

	switch *input.NftTypeId {
	case models.NFT_TYPE_ID_CATEGORY:
		if input.NftCardCategoryId == nil {
			ServeError(w, "nft_card_category_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardCategory(ctx, *input.NftCardCategoryId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_DAY_MONTH:
		if input.NftCardDayMonthId == nil {
			ServeError(w, "nft_card_day_month_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardDayMonth(ctx, *input.NftCardDayMonthId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_CRAFTING:
		if input.NftCardCraftingId == nil {
			ServeError(w, "nft_card_crafting must be specified", 400)
			return
		}

		data, err := store.GetNftCardCrafting(ctx, *input.NftCardCraftingId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_PREDICTION:
		if input.NftCardPredictionId == nil {
			ServeError(w, "nft_card_prediction_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardPrediction(ctx, *input.NftCardPredictionId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_IDENTITY:
		if input.NftCardIdentityId == nil {
			ServeError(w, "nft_card_identity_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardIdentity(ctx, *input.NftCardIdentityId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_TRIGGER:
		if input.NftCardTriggerId == nil {
			ServeError(w, "nft_card_trigger_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardTrigger(ctx, *input.NftCardTriggerId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_YEAR:
		if input.NftCardYearId == nil {
			ServeError(w, "nft_card_year_id must be specified", 400)
			return
		}

		data, err := store.GetNftCardYear(ctx, *input.NftCardYearId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	case models.NFT_TYPE_ID_CARD_PACK:
		if input.NftCardYearId == nil {
			ServeError(w, "card_pack_id must be specified", 400)
			return
		}

		data, err := store.GetCardPack(ctx, *input.CardPackId)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if *data.OwnerId != mid {
			ServeError(w, "user does not own specified nft", http.StatusForbidden)
			return
		}
	}

	input.CreatedAt = nil
	input.OwnerId = &mid

	err = store.NewMarketplaceListing(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateMarketplaceListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.MarketplaceListing{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateMarketplaceListing(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteMarketplaceListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "marketplace_listing_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteMarketplaceListing(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListMarketplaceListingForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListMarketplaceListingByOwnerId(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}

func SearchMarketplaceListings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := r.URL.Query().Get("q")

	nft_type_ids_q := r.URL.Query().Get("nft_type_ids")

	nft_collection_id, err := strconv.ParseInt(r.URL.Query().Get("nft_collection_id"), 10, 64)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var nft_type_ids []int64
	for _, v := range strings.Split(nft_type_ids_q, ",") {
		nft_type_id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nft_type_ids = append(nft_type_ids, nft_type_id)
	}

	limit_q := r.URL.Query().Get("limit")

	limit, err := strconv.ParseInt(limit_q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := store.SearchMarketplaceListings(ctx, q, nft_collection_id, nft_type_ids, int(limit), 0, false)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, out)
}

func BuyMarketplaceListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		PaymentMethodId int64
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	q := chi.URLParam(r, "marketplace_listing_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	gvars := struct {
		Id int64 `json:"id"`
	}{
		Id: id,
	}

	vars, err := json.Marshal(gvars)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	gql := graphql.Fragment_user + graphql.Fragment_marketplace_listing + `query GetMarketplaceListingWithOwner {
		marketplace_listing(where: {id: { eq: $id }}) {
			...MarketplaceListing
			owner {
				...User
			}
		}
	}`

	res, err := graphql.Graph.GraphQL(ctx, gql, vars, nil)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	out := struct {
		MarketplaceListing []models.MarketplaceListing
	}{}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	if len(out.MarketplaceListing) < 1 {
		ServeError(w, "No marketplace listing found via given id", 500)
		return
	}

	listing := out.MarketplaceListing[0]

	if *listing.OwnerId == mid {
		ServeError(w, "User already owns the listed assets", 400)
		return
	}

	buyer, err := store.GetUser(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	if buyer.Balance == nil {
		ServeError(w, "User has no balance", 400)
		return
	}

	if *buyer.Balance < *listing.Price {
		ServeError(w, "User has insufficient balance", 400)
		return
	}

	//REPLACE - create one postgres transaction so it can revert both if one balance update fails
	err = store.AddBalanceToUser(ctx, *listing.OwnerId, *listing.Price)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.AddBalanceToUser(ctx, mid, 0-*listing.Price)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	if listing.NftCardCategory != nil {
		listing.NftCardCategory.OwnerId = &mid
		err = store.UpdateNftCardCategory(ctx, *listing.NftCardCategory)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	if listing.NftCardDayMonth != nil {
		listing.NftCardDayMonth.OwnerId = &mid
		err = store.UpdateNftCardDayMonth(ctx, *listing.NftCardDayMonth)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	if listing.NftCardYear != nil {
		listing.NftCardYear.OwnerId = &mid
		err = store.UpdateNftCardYear(ctx, *listing.NftCardYear)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	if listing.NftCardTrigger != nil {
		listing.NftCardTrigger.OwnerId = &mid
		err = store.UpdateNftCardTrigger(ctx, *listing.NftCardTrigger)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	if listing.NftCardIdentity != nil {
		listing.NftCardIdentity.OwnerId = &mid
		err = store.UpdateNftCardIdentity(ctx, *listing.NftCardIdentity)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	if listing.NftCardPrediction != nil {
		listing.NftCardPrediction.OwnerId = &mid
		err = store.UpdateNftCardPrediction(ctx, *listing.NftCardPrediction)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	}

	err = store.DeleteMarketplaceListing(ctx, *listing.Id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)

}
