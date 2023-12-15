package models

//used to generate the collection - saved in database and factory contract
type CollectionCreation struct {
	ID            *int64
	Name          string
	CardPackTiers []CardPackTier
	Status        int
	AggPackPath   *string
	CardPacks     []CardPackConfig
}

type CardPackTier struct {
	Name           string
	CardPackAmount int
	Guaranteed     struct {
		Rare     CardTypeValue
		Core     CardTypeValue
		Uncommon CardTypeValue
	}
	Probabilities struct {
		Rare     CardTypeValue
		Core     CardTypeValue
		Uncommon CardTypeValue
	}
}

type CardTypeValue struct {
	Category int64
	DayMonth int64
	Trigger  string
	Year     int64
}

//config for each card pack
type CardPackConfig struct {
	Changed int
	CardPackId int64
	Tier       string
	Contains   struct {
		Rare     CardTypeValue
		Core     CardTypeValue
		Uncommon CardTypeValue
	}
}
