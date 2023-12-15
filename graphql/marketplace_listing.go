package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	Fragment_marketplace_listing = ReflectToFragment(models.MarketplaceListingData{})
)

func NewMarketplaceListing(ctx context.Context, data *models.MarketplaceListing) error {
	q := `
		mutation CreateMarketplaceListing {
			marketplace_listing(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.MarketplaceListing
	}{
		Data: *data,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.MarketplaceListing[0].Id

	data.Id = &id

	return nil
}

func DeleteMarketplaceListing(ctx context.Context, id int64) error {
	q := `
		mutation DeleteMarketplaceListing {
			marketplace_listing(delete: true, where: { id: { eq: $id } }) {
				id
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateMarketplaceListing(ctx context.Context, data models.MarketplaceListing) error {
	q := `
		mutation UpdateMarketplaceListing {
			marketplace_listing(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.MarketplaceListing
	}{
		Id: *data.Id,
	}

	data.Id = nil
	input.Data = data

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetMarketplaceListing(ctx context.Context, id int64) (models.MarketplaceListing, error) {
	var data models.MarketplaceListing

	q := Fragment_marketplace_listing + `
			query GetMarketplaceListing {
			marketplace_listing(where: { id: { eq: $id } }) {
				...MarketplaceListing
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.MarketplaceListing) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.MarketplaceListing[0]

	return data, nil
}

func ListMarketplaceListingByOwnerId(ctx context.Context, id int64) ([]models.MarketplaceListing, error) {
	var out []models.MarketplaceListing

	q := Fragment_marketplace_listing + `query ListMarketplaceListingByOwnerId {
		marketplace_listing(where: { owner_id: { eq: $id }}) {
						...MarketplaceListing
		}
	}`

	input := struct {
		Id int64 `json:"id"`
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		MarketplaceListing []models.MarketplaceListing
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	// if len(ret.MarketplaceListing) < 1 {
	// 	return out, errors.New("Object not found")
	// }

	out = ret.MarketplaceListing

	return out, nil
}

func SearchMarketplaceListings(ctx context.Context, q string, card_collection_id int64, nft_type_ids []int64, limit int, listing_id int64, with_owner bool) ([]models.MarketplaceListing, error) {
	var out []models.MarketplaceListing
	var q_nft string
	card_collection_q := "or: {nft_card_crafting: {card_series: { card_collection_id: $cid }}, card_pack: {card_series: { card_collection_id: $cid }}, nft_card_day_month: {card_series: { card_collection_id: $cid }}, nft_card_year: {card_series: { card_collection_id: $cid }}, nft_card_category: {card_series: { card_collection_id: $cid }}, nft_card_identity: {card_series: { card_collection_id: $cid }}, nft_card_prediction: {card_series: { card_collection_id: $cid }}, nft_card_trigger: {card_series: { card_collection_id: $cid }}}"
	filters := fmt.Sprintf("where: { and: { nft_type_id: {in: $nids}, is_listed: true, %s }}", card_collection_q)

	fragments := Fragment_marketplace_listing

	for _, v := range nft_type_ids {
		switch v {
		case models.NFT_TYPE_ID_CATEGORY:
			fragments += `
		
		` + fragment_nft_card_category
			q_nft += `
			nft_card_category {
			...NftCardCategory
		}`
		case models.NFT_TYPE_ID_CRAFTING:
			fragments += `
		
		` + fragment_nft_card_crafting
			q_nft += `
			nft_card_crafting {
			...NftCardCrafting
		}`
		case models.NFT_TYPE_ID_PREDICTION:
			fragments += `
		` + fragment_nft_card_prediction
			q_nft += `
			nft_card_prediction
		{
			...NftCardPrediction
		}`
		case models.NFT_TYPE_ID_IDENTITY:
			fragments += `
		
		` + fragment_nft_card_identity
			q_nft += `
			nft_card_identity
		{
			...NftCardIdentity
		}`
		case models.NFT_TYPE_ID_DAY_MONTH:
			fragments += `
		
		` + fragment_nft_card_day_month
			q_nft += `
			nft_card_day_month
		{
			...NftCardDayMonth
		}`
		case models.NFT_TYPE_ID_TRIGGER:
			fragments += `
		
		` + fragment_nft_card_trigger
			q_nft += `
			nft_card_trigger
		{
			...NftCardTrigger
		}`
		case models.NFT_TYPE_ID_YEAR:
			fragments += `
		
		` + fragment_nft_card_year
			q_nft += `
			nft_card_year
		{
			...NftCardYear
		}`
		}
	}

	input := struct {
		Limit int
		Nids  []int64
		Id    int64
		Cid   int64
	}{
		Limit: limit,
		Nids:  nft_type_ids,
		Cid:   card_collection_id,
	}

	if listing_id != 0 {
		filters = fmt.Sprintf("where: { and: { nft_type_id: {in: $nids}, nft_card_category_id: $id, %s }}", card_collection_q)
		input.Id = 14
	}

	if with_owner {
		fragments += `
		
		` + Fragment_user
		q_nft = `owner
		{
			...User
		}`
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	gq := fragments + fmt.Sprintf(`query SearchMarketplaceListings {
		marketplace_listing(%s, limit: $limit) {
						...MarketplaceListing
						%s
					}
				}`, filters, q_nft)

	// fmt.Println(gq)

	res, err := Graph.GraphQL(ctx, gq, js, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		MarketplaceListing []models.MarketplaceListing
	}{}

	// js, err = res.Data.MarshalJSON()
	// if err != nil {
	// 	return out, err
	// }

	// spew.Dump(js)

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.MarketplaceListing

	return out, nil
}

// func GetMarketplaceListingByNftIdAndCollectionIdWithOwner(ctx context.Context, nft_collection_id int64, nft_id int64) (models.MarketplaceListing, error) {

// 	q := ``

// 	return out, nil
// }
