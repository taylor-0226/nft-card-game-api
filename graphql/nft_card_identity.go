package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_identity = ReflectToFragment(models.NftCardIdentityData{})
)

func NewNftCardIdentity(ctx context.Context, data *models.NftCardIdentity) error {
	q := `
		mutation CreateNftCardIdentity {
			nft_card_identity(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardIdentity
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
		NftCardIdentity []models.NftCardIdentity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardIdentity) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardIdentity[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardIdentity(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardIdentity {
			nft_card_identity(where: { id: { eq: $id } }) {
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
		NftCardIdentity []models.NftCardIdentity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardIdentity) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardIdentity(ctx context.Context, data models.NftCardIdentity) error {
	q := `
		mutation UpdateNftCardIdentity {
			nft_card_identity(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardIdentity
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
		NftCardIdentity []models.NftCardIdentity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardIdentity) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardIdentity(ctx context.Context, id int64) (models.NftCardIdentity, error) {
	var data models.NftCardIdentity

	q := fragment_nft_card_identity + `
			query GetNftCardIdentity {
			nft_card_identity(where: { id: { eq: $id } }) {
				...NftCardIdentity
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
		NftCardIdentity []models.NftCardIdentity
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardIdentity) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardIdentity[0]

	return data, nil
}

func ListNftCardIdentityByOwnerId(ctx context.Context, id int64, filters models.NftCardIdentityFilter) ([]models.NftCardIdentity, error) {
	var out []models.NftCardIdentity

	q := fragment_nft_card_identity + `query ListNftCardIdentityByOwnerId {
		nft_card_identity(%s) {
			...NftCardIdentity
		}
	}`

	input := struct {
		Id           int64    `json:"id"`
		Rarities     []int64  `json:"rarities"`
		Categories   []string `json:"categories"`
		Celebrities  []string `json:"celebrities"`
		CardSeriesId int64    `json:"card_series_id"`
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

	if filters.Celebrities != nil {
		input.Celebrities = *filters.Celebrities
		filter_params = append(filter_params, "celebrity_name: { in: $celebrities }")
	}

	if filters.CardSeriesId != nil {
		input.CardSeriesId = *filters.CardSeriesId
		filter_params = append(filter_params, "card_series_id: { eq: $card_series_id }")
	}

	if filters.Rarities != nil {
		input.Rarities = *filters.Rarities
		filter_params = append(filter_params, "rarity: { in: $rarities }")
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
		NftCardIdentity []models.NftCardIdentity
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardIdentity

	return out, nil
}
