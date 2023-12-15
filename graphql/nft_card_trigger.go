package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_trigger = ReflectToFragment(models.NftCardTriggerData{})
)

func NewNftCardTrigger(ctx context.Context, data *models.NftCardTrigger) error {
	q := `
		mutation CreateNftCardTrigger {
			nft_card_trigger(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardTrigger
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
		NftCardTrigger []models.NftCardTrigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardTrigger) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardTrigger[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardTrigger(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardTrigger {
			nft_card_trigger(where: { id: { eq: $id } }) {
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
		NftCardTrigger []models.NftCardTrigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardTrigger) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardTrigger(ctx context.Context, data models.NftCardTrigger) error {
	q := `
		mutation UpdateNftCardTrigger {
			nft_card_trigger(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardTrigger
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
		NftCardTrigger []models.NftCardTrigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardTrigger) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardTrigger(ctx context.Context, id int64) (models.NftCardTrigger, error) {
	var data models.NftCardTrigger

	q := fragment_nft_card_trigger + `
			query GetNftCardTrigger {
			nft_card_trigger(where: { id: { eq: $id } }) {
				...NftCardTrigger
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
		NftCardTrigger []models.NftCardTrigger
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardTrigger) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardTrigger[0]

	return data, nil
}

func ListNftCardTriggerByOwnerId(ctx context.Context, id int64, filters models.NftCardTriggerFilter) ([]models.NftCardTrigger, error) {
	var out []models.NftCardTrigger

	q := fragment_nft_card_trigger + `query ListNftCardTriggerByOwnerId {
		nft_card_trigger(%s) {
			...NftCardTrigger
		}
	}`

	input := struct {
		Id           int64    `json:"id"`
		Triggers     []string `json:"triggers"`
		Categories   []int64  `json:"categories"`
		Tiers        []string `json:"tiers"`
		Rarities     []int64  `json:"rarities"`
		CardSeriesId int64    `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.Triggers != nil {
		input.Triggers = *filters.Triggers
		filter_params = append(filter_params, "trigger: { in: $triggers }")
	}

	if filters.Tiers != nil {
		input.Tiers = *filters.Tiers
		filter_params = append(filter_params, "tier: { in: $tiers }")
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
		NftCardTrigger []models.NftCardTrigger
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardTrigger

	return out, nil
}
