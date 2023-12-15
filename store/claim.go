package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewClaim(ctx context.Context, data *models.Claim) error {
	err := graphql.NewClaim(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetClaim(ctx context.Context, id int64) (models.Claim, error) {
	data, err := graphql.GetClaim(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteClaim(ctx context.Context, id int64) error {
	err := graphql.DeleteClaim(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateClaim(ctx context.Context, data models.Claim) error {
	err := graphql.UpdateClaim(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListClaimByArticleId(ctx context.Context, id int64) ([]models.Claim, error) {
	data, err := graphql.ListClaimByArticleId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func ListClaimByClaimerId(ctx context.Context, id int64) ([]models.Claim, error) {
	data, err := graphql.ListClaimByClaimerId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func ListClaim(ctx context.Context) ([]models.Claim, error) {
	data, err := graphql.ListClaim(ctx)
	if err != nil {
		return data, err
	}

	return data, nil
}
