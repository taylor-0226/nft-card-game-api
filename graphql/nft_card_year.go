package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_year = ReflectToFragment(models.NftCardYearData{})
)

func NewNftCardYear(ctx context.Context, data *models.NftCardYear) error {
	q := `
		mutation CreateNftCardYear {
			nft_card_year(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardYear
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
		NftCardYear []models.NftCardYear
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardYear) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardYear[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardYear(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardYear {
			nft_card_year(where: { id: { eq: $id } }) {
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
		NftCardYear []models.NftCardYear
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardYear) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardYear(ctx context.Context, data models.NftCardYear) error {
	q := `
		mutation UpdateNftCardYear {
			nft_card_year(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardYear
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
		NftCardYear []models.NftCardYear
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardYear) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardYear(ctx context.Context, id int64) (models.NftCardYear, error) {
	var data models.NftCardYear

	q := fragment_nft_card_year + `
			query GetNftCardYear {
			nft_card_year(where: { id: { eq: $id } }) {
				...NftCardYear
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
		NftCardYear []models.NftCardYear
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardYear) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardYear[0]

	return data, nil
}

func ListNftCardYearByOwnerId(ctx context.Context, id int64, filters models.NftCardYearFilter) ([]models.NftCardYear, error) {
	var out []models.NftCardYear

	q := fragment_nft_card_year + `query ListNftCardYearByOwnerId {
		nft_card_year(%s) {
			...NftCardYear
		}
	}`

	input := struct {
		Id           int64   `json:"id"`
		Year         int64   `json:"year"`
		Rarities     []int64 `json:"rarities"`
		CardSeriesId int64   `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.Year != nil {
		input.Year = *filters.Year
		filter_params = append(filter_params, "year: { eq: $year }")
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
		NftCardYear []models.NftCardYear
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardYear

	return out, nil
}
