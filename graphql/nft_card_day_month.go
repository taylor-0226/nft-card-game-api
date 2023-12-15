package graphql

import (
	"context"
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_day_month = ReflectToFragment(models.NftCardDayMonthData{})
)

func NewNftCardDayMonth(ctx context.Context, data *models.NftCardDayMonth) error {
	q := `
		mutation CreateNftCardDayMonth {
			nft_card_day_month(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardDayMonth
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
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardDayMonth[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardDayMonth(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }) {
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
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardDayMonth(ctx context.Context, data models.NftCardDayMonth) error {
	q := `
		mutation UpdateNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardDayMonth
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
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardDayMonth(ctx context.Context, id int64) (models.NftCardDayMonth, error) {
	var data models.NftCardDayMonth

	q := fragment_nft_card_day_month + `
			query GetNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }) {
				...NftCardDayMonth
			}
		}
		`

	input := struct {
		Id int64 `json:"id"`
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
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardDayMonth) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardDayMonth[0]

	return data, nil
}

func ListNftCardDayMonthByOwnerId(ctx context.Context, id int64, filters models.NftCardDayMonthFilter) ([]models.NftCardDayMonth, error) {
	var out []models.NftCardDayMonth

	q := fragment_nft_card_day_month + `query ListNftCardDayMonthByOwnerId {
		nft_card_day_month(%s) {
			...NftCardDayMonth
		}
	}`

	input := struct {
		Id           int64   `json:"id"`
		Rarities     []int64 `json:"rarities"`
		Day          int64   `json:"day"`
		Month        int64   `json:"month"`
		CardSeriesId int64   `json:"card_series_id"`
	}{
		Id: id,
	}

	filter_params := []string{}
	q_filters := "where: {%s}"

	filter_params = append(filter_params, "owner_id: { eq: $id }")

	if filters.Day != nil {
		input.Day = *filters.Day
		filter_params = append(filter_params, "day: { eq: $day }")
	}

	if filters.Month != nil {
		input.Month = *filters.Month
		filter_params = append(filter_params, "month: { eq: $month }")
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
		NftCardDayMonth []models.NftCardDayMonth
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.NftCardDayMonth

	return out, nil
}
