package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_claim = ReflectToFragment(models.ClaimData{})
)

func NewClaim(ctx context.Context, data *models.Claim) error {
	q := `
		mutation CreateClaim {
			claim(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Claim
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
		Claim []models.Claim
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Claim) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Claim[0].Id

	data.Id = &id

	return nil
}

func DeleteClaim(ctx context.Context, id int64) error {
	q := `
		mutation DeleteClaim {
			claim(where: { id: { eq: $id } }) {
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
		Claim []models.Claim
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Claim) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateClaim(ctx context.Context, data models.Claim) error {
	q := `
		mutation UpdateClaim {
			claim(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Claim
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
		Claim []models.Claim
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Claim) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetClaim(ctx context.Context, id int64) (models.Claim, error) {
	var data models.Claim

	q := fragment_claim + `
			query GetClaim {
			claim(where: { id: { eq: $id } }) {
				...Claim
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
		Claim []models.Claim
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Claim) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Claim[0]

	return data, nil
}

func ListClaimByClaimerId(ctx context.Context, id int64) ([]models.Claim, error) {
	var out []models.Claim

	q := fragment_claim + `query ListClaimByClaimerId {
		claim(where: { claimer_id: { eq: $id }}) {
						...Claim
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
		Claim []models.Claim
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.Claim) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.Claim

	return out, nil
}

func ListClaimByArticleId(ctx context.Context, id int64) ([]models.Claim, error) {
	var out []models.Claim

	q := fragment_claim + `query ListClaimByArticleId {
		claim(where: { article_id: { eq: $id }}) {
						...Claim
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
		Claim []models.Claim
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.Claim

	return out, nil
}

func ListClaim(ctx context.Context) ([]models.Claim, error) {
	var out []models.Claim

	q := fragment_claim + `query ListClaim {
						claim {
						...Claim
						}
					}`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		Claim []models.Claim
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.Claim

	return out, nil
}
