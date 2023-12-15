package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_crafting = ReflectToFragment(models.NftCardCraftingData{})
)

func NewNftCardCrafting(ctx context.Context, data *models.NftCardCrafting) error {
	q := `
		mutation CreateNftCardCrafting {
			nft_card_crafting(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardCrafting
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
		NftCardCrafting []models.NftCardCrafting
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCrafting) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardCrafting[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardCrafting(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardCrafting {
			nft_card_crafting(where: { id: { eq: $id } }) {
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
		NftCardCrafting []models.NftCardCrafting
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCrafting) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardCrafting(ctx context.Context, data models.NftCardCrafting) error {
	q := `
		mutation UpdateNftCardCrafting {
			nft_card_crafting(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardCrafting
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
		NftCardCrafting []models.NftCardCrafting
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCrafting) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardCrafting(ctx context.Context, id int64) (models.NftCardCrafting, error) {
	var data models.NftCardCrafting

	q := fragment_nft_card_crafting + `
			query GetNftCardCrafting {
			nft_card_crafting(where: { id: { eq: $id } }) {
				...NftCardCrafting
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
		NftCardCrafting []models.NftCardCrafting
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardCrafting) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardCrafting[0]

	return data, nil
}

func ListNftCardCraftingByOwnerId(ctx context.Context, id int64, filters models.NftCardCraftingFilter) ([]models.NftCardCrafting, error) {
	var out []models.NftCardCrafting

	q := fragment_nft_card_crafting + `query ListNftCardCraftingByOwnerId {
					nft_card_crafting(%s) {
						...NftCardCrafting
					}
				}`

	input := struct {
		Id           int64   `json:"id"`
		Rarities     []int64 `json:"rarities"`
		CardSeriesId int64   `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.CardSeriesId != nil {
		filter_params = append(filter_params, "card_series_id: { eq: $card_series_id}")
		input.CardSeriesId = *filters.CardSeriesId
	}

	if filters.Rarities != nil {
		filter_params = append(filter_params, "rarity: { in: $rarities}")
		input.Rarities = *filters.Rarities
	}

	if filters.Status != nil {

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
		NftCardCrafting []models.NftCardCrafting
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardCrafting

	return out, nil
}
