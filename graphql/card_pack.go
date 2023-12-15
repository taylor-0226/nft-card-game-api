package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_card_pack = ReflectToFragment(models.CardPackData{})
)

func NewCardPack(ctx context.Context, data *models.CardPack) error {
	q := `
		mutation CreateCardPack {
			card_pack(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.CardPack
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
		CardPack []models.CardPack
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardPack) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.CardPack[0].Id

	data.Id = &id

	return nil
}

func DeleteCardPack(ctx context.Context, id int64) error {
	q := `
		mutation DeleteCardPack {
			card_pack(where: { id: { eq: $id } }) {
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
		CardPack []models.CardPack
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardPack) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateCardPack(ctx context.Context, data models.CardPack) error {
	q := `
		mutation UpdateCardPack {
			card_pack(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.CardPack
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
		CardPack []models.CardPack
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.CardPack) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetCardPack(ctx context.Context, id int64) (models.CardPack, error) {
	var data models.CardPack

	q := fragment_card_pack + `
			query GetCardPack {
			card_pack(where: { id: { eq: $id } }) {
				...CardPack
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
		CardPack []models.CardPack
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.CardPack) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.CardPack[0]

	return data, nil
}

func ListCardPackByCardSeriesId(ctx context.Context, id int64) ([]models.CardPack, error) {
	var out []models.CardPack

	q := fragment_card_pack + `query ListCardPackByCardSeriesId(where: { card_series_id: { eq: $id }}) {
						...CardPack
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
		CardPack []models.CardPack
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.CardPack) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.CardPack

	return out, nil
}

func ListCardPackByOwnerId(ctx context.Context, id int64) ([]models.CardPack, error) {
	var out []models.CardPack

	q := fragment_card_pack + `query ListCardPackByOwnerId {
	card_pack(where: { owner_id: { eq: $id }}) {
						...CardPack
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
		CardPack []models.CardPack
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.CardPack

	return out, nil
}
