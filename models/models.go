package models

import (
	"time"
)

type Strings []string
type Ints []int64
type ctxkey int64

const (
	CTX_is_auth       = ctxkey(0)
	CTX_user_id       = ctxkey(1)
	CTX_user_role_ids = ctxkey(2)
	CTX_user_timezone = ctxkey(3)
)

type NftCardYearData struct {
	Rarity       *int64
	CardSeriesId *int64
	Id           *int64
	Year         *int64
	OwnerId      *int64
	IsCrafted    *bool
}

type NftCardYear struct {
	NftCardYearData
	Owner              *User
	CardSeries         *CardSeries
	MarketplaceListing []*MarketplaceListing
}

type IdentityData struct {
	Name       *string
	BirthDay   *int64
	BirthMonth *int64
	BirthYear  *int64
	Category   *string
	Id         *int64
}

type Identity struct {
	IdentityData
}

type TriggerData struct {
	Name *string
	Tier *string
	Id   *int64
}

type Trigger struct {
	TriggerData
}

type MarketplaceOfferData struct {
	Status               *int64
	MarketplaceListingId *int64
	Amount               *int64
	BuyerId              *int64
	Id                   *int64
}

type MarketplaceOffer struct {
	MarketplaceOfferData
	MarketplaceListing *MarketplaceListing
	Buyer              *User
}

type CardSeriesData struct {
	Rarity           *int64
	PctMonth         *int64
	CardAmount       *int64
	PctYear          *int64
	Id               *int64
	Quantity         *int64
	Name             *string
	PctEvent         *int64
	CostUsd          *int64
	CardCollectionId *int64
	PctIdentity      *int64
	PctDay           *int64
}

type CardSeries struct {
	CardSeriesData
	NftCardCategory   []*NftCardCategory
	NftCardCrafting   []*NftCardCrafting
	NftCardDayMonth   []*NftCardDayMonth
	NftCardYear       []*NftCardYear
	CardCollection    *CardCollection
	NftCardIdentity   []*NftCardIdentity
	NftCardPrediction []*NftCardPrediction
	NftCardTrigger    []*NftCardTrigger
	CardPack          []*CardPack
}

type ClaimData struct {
	Id              *int64
	Status          *int64
	CreatedAt       *time.Time
	ClaimerId       *int64
	NftPredictionId *int64
	ArticleId       *int64
}

type Claim struct {
	ClaimData
	Claimer *User
	Article *Article
}

type CategoryData struct {
	Id   *int64
	Name *string
}

type Category struct {
	CategoryData
}

type NftCardTriggerData struct {
	Trigger      *string
	Id           *int64
	IsCrafted    *bool
	Tier         *string
	CardSeriesId *int64
	OwnerId      *int64
	Rarity       *int64
}

type NftCardTrigger struct {
	NftCardTriggerData
	CardSeries         *CardSeries
	Owner              *User
	MarketplaceListing []*MarketplaceListing
}

type CardPackData struct {
	CardSeriesId *int64
	Cards        *CardPackCards
	OwnerId      *int64
	IsOpened     *bool
	Tier         *int64
	Id           *int64
}

type CardPack struct {
	CardPackData
	CardSeries *CardSeries
	Owner      *User
}

type NftCardDayMonthData struct {
	Id           *int64
	IsCrafted    *bool
	CardSeriesId *int64
	Rarity       *int64
	Day          *int64
	Month        *int64
	OwnerId      *int64
}

type NftCardDayMonth struct {
	NftCardDayMonthData
	CardSeries         *CardSeries
	Owner              *User
	MarketplaceListing []*MarketplaceListing
}

type NftCardCraftingData struct {
	Id           *int64
	IsCrafted    *bool
	OwnerId      *int64
	Rarity       *int64
	CardSeriesId *int64
}

type NftCardCrafting struct {
	NftCardCraftingData
	Owner              *User
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
}

type NftCardPredictionData struct {
	Rarity        *int64
	CardSeriesId  *int64
	Id            *int64
	IsClaimed     *bool
	OwnerId       *int64
	Triggers      *Strings
	CelebrityName *string
}

type NftCardPrediction struct {
	NftCardPredictionData
	Owner              *User
	CardSeries         *CardSeries
	MarketplaceListing []*MarketplaceListing
}

type ArticleData struct {
	ArticleSourceId *int64
	Url             *string
	ThumbnailSrc    *string
	Title           *string
	Tags            *Strings
	Id              *int64
	Excerpt         *string
	CreatedAt       *time.Time
}

type Article struct {
	ArticleData
	Claim         []*Claim
	ArticleSource *ArticleSource
}

type CelebrityData struct {
	Category         *string
	EligibleTriggers *Strings
	Id               *int64
	Name             *string
	BirthDay         *int64
	BirthMonth       *int64
	BirthYear        *int64
}

type Celebrity struct {
	CelebrityData
}

type ArticleSourceData struct {
	Id   *int64
	Name *string
}

type ArticleSource struct {
	ArticleSourceData
	Article []*Article
}

type MarketplaceListingData struct {
	NftCardPredictionId *int64
	Price               *int64
	NftCardCraftingId   *int64
	NftCardIdentityId   *int64
	NftTypeId           *int64
	CreatedAt           *time.Time
	IsListed            *bool
	NftCardTriggerId    *int64
	NftCardDayMonthId   *int64
	CardPackId          *int64
	NftCardYearId       *int64
	Id                  *int64
	OwnerId             *int64
	NftCardCategoryId   *int64
}

type MarketplaceListing struct {
	MarketplaceListingData
	IsOwned           bool
	CardPack          *CardPack
	NftCardTrigger    *NftCardTrigger
	NftCardDayMonth   *NftCardDayMonth
	NftCardPrediction *NftCardPrediction
	NftCardYear       *NftCardYear
	NftCardIdentity   *NftCardIdentity
	Owner             *User
	NftCardCrafting   *NftCardCrafting
	NftCardCategory   *NftCardCategory
	MarketplaceOffer  []*MarketplaceOffer
}

type CardCollectionData struct {
	Id   *int64
	Name *string
}

type CardCollection struct {
	CardCollectionData
	CardSeries []*CardSeries
}

type UserData struct {
	VenlyId        *string
	CreatedAt      *time.Time
	Password       *string
	Id             *int64
	Email          *string
	ExternalAuthId *string
	WalletAddress  *string
	PhoneNumber    *string
	Name           *string
	Username       *string
	RoleIds        *Ints
	Balance        *int64
}

type User struct {
	UserData
	NftCardYear        []*NftCardYear
	MarketplaceOffer   []*MarketplaceOffer
	Claim              []*Claim
	MarketplaceListing []*MarketplaceListing
	NftCardTrigger     []*NftCardTrigger
	NftCardPrediction  []*NftCardPrediction
	NftCardCategory    []*NftCardCategory
	NftCardIdentity    []*NftCardIdentity
	CardPack           []*CardPack
	NftCardCrafting    []*NftCardCrafting
	NftCardDayMonth    []*NftCardDayMonth
}

type NftCardCategoryData struct {
	IsCrafted    *bool
	Id           *int64
	Category     *string
	OwnerId      *int64
	Rarity       *int64
	CardSeriesId *int64
}

type NftCardCategory struct {
	NftCardCategoryData
	CardSeries         *CardSeries
	MarketplaceListing []*MarketplaceListing
	Owner              *User
}

type NftCardIdentityData struct {
	Year          *int64
	IsCrafted     *bool
	OwnerId       *int64
	Category      *string
	CardSeriesId  *int64
	Id            *int64
	Day           *int64
	Rarity        *int64
	Month         *int64
	CelebrityName *string
}

type NftCardIdentity struct {
	NftCardIdentityData
	MarketplaceListing []*MarketplaceListing
	Owner              *User
	CardSeries         *CardSeries
}
