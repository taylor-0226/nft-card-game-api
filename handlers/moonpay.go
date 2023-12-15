package handlers

import (
	"encoding/json"
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
)

const (
	MOONPAY_SECRET_KEY      = "sk_test_cHMS4IN3BLApimnz8C0RYtyObpbOT4H"
	MOONPAY_WEBHOOK_KEY     = "wk_test_BWUpCi6aQXMPoGo5yKA8ZkjRjlFvD60J"
	MOONPAY_PUBLISHABLE_KEY = "pk_test_PaUTi3HVAHvclaZTMJS0TNTfMIrpPj"
)

func WebhookMoonpayGetNftInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contract_address := chi.URLParam(r, "contract_address")

	token_id_q := chi.URLParam(r, "token_id")

	//0 doesn't pass through json encoding - fix in custom encoderS
	token_id := int64(14)

	if !strings.Contains(token_id_q, "_") {
		var err error
		token_id, err = strconv.ParseInt(token_id_q, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

	// listing_id_q := r.URL.Query().Get("listingId")
	// buyer_wallet_address := r.URL.Query().Get("walletAddress")

	// nft_collection, ok := store.GetNftCollectionByContractAddress(ctx, contract_address)
	// if !ok {
	// 	ServeError(w, fmt.Sprintf("Unable to find nft collection for contract address (%s)", contract_address), http.StatusBadRequest)
	// 	return
	// }

	nft_collection := struct {
		Id   int64
		Name string
	}{
		Id:   models.NFT_TYPE_ID_CATEGORY,
		Name: "category",
	}

	out := struct {
		TokenId           string `json:"tokenId"`
		ContractAddress   string `json:"contractAddress"`
		Name              string `json:"name"`
		Collection        string `json:"collection"`
		ImageUrl          string `json:"imageUrl"`
		ExplorerUrl       string `json:"explorerUrl"`
		Price             string `json:"price"`
		PriceCurrencyCode string `json:"priceCurrencyCode"`
		Quantity          int64  `json:"quantity"`
		SellerAddress     string `json:"sellerAddress"`
		SellType          string `json:"sellType"`
		Flow              string `json:"flow"`
		Network           string `json:"network"`
	}{
		TokenId:           token_id_q,
		ContractAddress:   contract_address,
		PriceCurrencyCode: "USD",
		SellType:          "Secondary",
		Flow:              "Lite",
		Network:           "polygon",
		Quantity:          1,
	}

	switch nft_collection.Id {
	case models.NFT_TYPE_ID_CATEGORY:
		listings, err := store.SearchMarketplaceListings(ctx, "", 1, []int64{nft_collection.Id}, 1, token_id, true)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		listing := listings[0]

		nft, err := store.GetNftCardCategory(ctx, token_id)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		rarity, ok := store.NftRarityMap[*nft.Rarity]
		if !ok {
			ServeError(w, fmt.Sprintf("Error getting rarity from map with id (%d)", rarity), http.StatusBadRequest)
			return
		}

		out.Name = fmt.Sprintf("%s %s Category Card", *nft.Category, rarity)
		out.Collection = nft_collection.Name
		out.SellerAddress = *listing.Owner.WalletAddress
		out.Price = strconv.FormatInt(*listing.Price, 10)
		break
	case models.NFT_TYPE_ID_CRAFTING:
	case models.NFT_TYPE_ID_DAY_MONTH:
	case models.NFT_TYPE_ID_IDENTITY:
	case models.NFT_TYPE_ID_PREDICTION:
	case models.NFT_TYPE_ID_TRIGGER:
	case models.NFT_TYPE_ID_YEAR:
	default:
		ServeError(w, "NFT collection for given id is not currently supported.", http.StatusServiceUnavailable)
		return
	}

	// ServeJSON(w, out)

	js, err := json.Marshal(out)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	// w.Write(js)

	fmt.Println(string(js))

	w.Write([]byte(`{
		"tokenId":"109",
		"contractAddress":"0x2180c6ecf2f770bd51dbb0d779cc81048899",
		"name":"MoonRocket",
		"collection":"MoonPay Special Collection",
		"imageUrl":"https://b1ic5m3wqxz9zd.cloudfront.net/f793.jpg",
		"explorerUrl":"",
		"price":0.1,
		"priceCurrencyCode":"ETH",
		"quantity":1,
		"sellerAddress":"0xt246e19c76a23068fa235e1673c10opecfbeb7hf",
		"sellType":"Primary",
		"flow":"Lite",
		"network":"Ethereum"
	 }`))

}

func WebhookMoonpayTransaction(w http.ResponseWriter, r *http.Request) {
	input := json.RawMessage{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	spew.Dump(input)

	w.WriteHeader(200)
}

// func WebhookMoonpayNotifyDeliverNft(w http.ResponseWriter, r *http.Request) {
// 	contract_address := chi.URLParam(r, "contract_address")
// 	token_id := chi.URLParam(r, "token_id")

// 	input := struct {
// 		Mode                string
// 		BuyerWalletAddress  string
// 		PriceCurrencyCode   string
// 		Price               int64
// 		Quantity            string
// 		SellerWalletAddress string
// 		ListingId           string
// 	}{}

// 	w.WriteHeader(200)

// }

// func WebhookMoonpayGetNftTransactionStatus(w http.ResponseWriter, r *http.Request) {
// 	ids_q := r.URL.Query().Get("id")
// 	ids := strings.Split(ids_q, ",")

// 	out := struct {
// 		Id              string
// 		Status          string
// 		TransactionHash string
// 		StatusChangedAt time.Time
// 		failureReason   string
// 		tokenId         string
// 	}{}
// }
