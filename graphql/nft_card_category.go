package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_category = ReflectToFragment(models.NftCardCategoryData{})
)

func NewNftCardCategory(ctx context.Context, data *models.NftCardCategory) error {
	q := `
		mutation CreateNftCardCategory {
			nft_card_category(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardCategory[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardCategory(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }) {
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardCategory(ctx context.Context, data models.NftCardCategory) error {
	q := `
		mutation UpdateNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardCategory(ctx context.Context, id int64) (models.NftCardCategory, error) {
	var data models.NftCardCategory

	q := fragment_nft_card_category + `
			query GetNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }) {
				...NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardCategory) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardCategory[0]

	return data, nil
}

func ListNftCardCategoryByOwnerId(ctx context.Context, id int64, filters models.NftCardCategoryFilter) ([]models.NftCardCategory, error) {
	var out []models.NftCardCategory

	q := fragment_nft_card_category + `query ListNftCardCategoryByOwnerId {
		nft_card_category(%s) {
			...NftCardCategory
		}
	}`

	input := struct {
		Id           int64   `json:"id"`
		Rarities     []int64 `json:"rarities"`
		Categories   []string `json:"categories"`
		CardSeriesId int64   `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.Categories != nil {
		input.Categories = *filters.Categories
		filter_params = append(filter_params, "category: { in: $categories }")
	}

	if filters.Rarities != nil {
		input.Rarities = *filters.Rarities
		filter_params = append(filter_params, "rarity: { in: $rarities }")
	}

	if filters.CardSeriesId != nil {
		input.CardSeriesId = *filters.CardSeriesId
		filter_params = append(filter_params, "card_series_id: { eq: $card_series_id }")

	}

	if len(filter_params) > 1 {
		q_filters = "where: {and: {%s}}"
	}

	var q_filter_inner string
	for i, v := range filter_params {
		if i == 0 {
			q_filter_inner = v
			continue
		}

		q_filter_inner += ", " + v
	}

	q_filters = fmt.Sprintf(q_filters, q_filter_inner)

	q = fmt.Sprintf(q, q_filters)

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		NftCardCategory []models.NftCardCategory
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardCategory

	return out, nil
}
