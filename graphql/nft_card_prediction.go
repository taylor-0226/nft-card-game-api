package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_prediction = ReflectToFragment(models.NftCardPredictionData{})
)

func NewNftCardPrediction(ctx context.Context, data *models.NftCardPrediction) error {
	q := `
		mutation CreateNftCardPrediction {
			nft_card_prediction(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardPrediction[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardPrediction(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }) {
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardPrediction(ctx context.Context, data models.NftCardPrediction) error {
	q := `
		mutation UpdateNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardPrediction(ctx context.Context, id int64) (models.NftCardPrediction, error) {
	var data models.NftCardPrediction

	q := fragment_nft_card_prediction + `
			query GetNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }) {
				...NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardPrediction) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardPrediction[0]

	return data, nil
}

func ListNftCardPredictionByOwnerId(ctx context.Context, id int64, filters models.NftCardPredictionFilter) ([]models.NftCardPrediction, error) {
	var out []models.NftCardPrediction

	q := fragment_nft_card_prediction + `query ListNftCardPredictionByOwnerId {
		nft_card_prediction(%s) {
			...NftCardPrediction
		}
	}`

	input := struct {
		Id           int64    `json:"id"`
		Rarities     []int64  `json:"rarities"`
		Celebrities  []string `json:"celebrities"`
		Triggers     []string `json:"triggers"`
		CardSeriesId int64    `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.Triggers != nil {
		input.Triggers = *filters.Triggers
		// filter_params = append(filter_params, "triggers: { in: $trigs }")
		filter_params = append(filter_params, `triggers: { in: $triggers }`)
	}

	if filters.Celebrities != nil {
		input.Celebrities = *filters.Celebrities
		filter_params = append(filter_params, "celebrity_name: { in: $celebrities }")
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

	fmt.Println(q)

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		NftCardPrediction []models.NftCardPrediction
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardPrediction

	return out, nil
}
