package handlers

import (
	"gameon-twotwentyk-api/models"
)

type IdentityRecipe struct {
	Day         int
	Month       int
	Year        int
	Category    string
	CelebrityId int64
	Celebrity   models.Celebrity
}

// func GetIdentityRecipes(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	identities, err := store.ListIdentity(ctx)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var out []IdentityRecipe
// 	for _, c := range identities {
// 		new := IdentityRecipe{
// 			Day:        c.Birthdate.Day(),
// 			Month:      int(c.Birthdate.Month()),
// 			Year:       c.Birthdate.Year(),
// 			Category:   c.Category,
// 			IdentityId: *c.Id,
// 			Identity:   c,
// 		}
// 		out = append(out, new)
// 	}

// 	ServeJSON(w, out)
// }

func ListIdentities() {

}

// func GetMyIdentityRecipes(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	mid := ctx.Value(models.CTX_user_id).(int64)

// 	user, err := store.GetUser(ctx, mid)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	data, err := nft.GetWalletNfts(user.WalletAddress)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	//get existing cards - day, month, year, category

// 	nft_card_day := data.NftCardDay
// 	nft_card_month := data.NftCardMonth
// 	nft_card_year := data.NftCardYear
// 	nft_card_category := data.NftCardCategory

// 	var day_arr []int
// 	var month_arr []int
// 	var year_arr []int
// 	var category_arr []string

// 	for _, n := range nft_card_day {
// 		day_arr = append(day_arr, n.Day)
// 	}

// 	for _, n := range nft_card_month {
// 		month_arr = append(month_arr, n.Month)
// 	}

// 	for _, n := range nft_card_year {
// 		year_arr = append(year_arr, n.Year)
// 	}

// 	for _, n := range nft_card_category {
// 		category_arr = append(category_arr, n.Category)
// 	}

// 	identities, err := store.ListIdentityByArrays(ctx, day_arr, month_arr, year_arr, category_arr)
// 	if err != nil {
// 		ServeError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	ServeJSON(w, identities)

// }
