package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewMarketplaceListing(ctx context.Context, data *models.MarketplaceListing) error {
	err := graphql.NewMarketplaceListing(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetMarketplaceListing(ctx context.Context, id int64) (models.MarketplaceListing, error) {
	data, err := graphql.GetMarketplaceListing(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteMarketplaceListing(ctx context.Context, id int64) error {
	err := graphql.DeleteMarketplaceListing(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMarketplaceListing(ctx context.Context, data models.MarketplaceListing) error {
	err := graphql.UpdateMarketplaceListing(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListMarketplaceListingByOwnerId(ctx context.Context, id int64) ([]models.MarketplaceListing, error) {
	data, err := graphql.ListMarketplaceListingByOwnerId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func SearchMarketplaceListings(ctx context.Context, q string, nft_collection_id int64, nft_type_ids []int64, limit int, id int64, with_owner bool) ([]models.MarketplaceListing, error) {
	data, err := graphql.SearchMarketplaceListings(ctx, q, nft_collection_id, nft_type_ids, limit, id, with_owner)
	if err != nil {
		return data, err
	}

	return data, nil
}

// func GetMarketplaceListingByNftIdAndCollectionIdWithOwner(ctx context.Context, nft_collection_id int64, nft_id int64) (models.MarketplaceListing, error) {
// 	data, err := graphql.GetMarketplaceListingByNftIdAndCollectionIdWithOwner(ctx, nft_collection_id, nft_id)
// 	if err != nil {
// 		return data, err
// 	}

// 	return data, nil
// }
