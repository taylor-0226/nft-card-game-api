package main

import (
	"encoding/json"
	"fmt"
	"gameon-twotwentyk-api/models"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

type AggPack struct {
	Tier [][]map[int64]interface{}

	// 3 x tier -> 12 x card type -> array of card amount values for each card pack in collection
}

type CollectionCards []map[int64]models.CardPackConfig

func DecodeAggPack(path string) (CollectionCards, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(b))

	var agg_pack AggPack
	err = json.Unmarshal(b, &agg_pack.Tier)
	if err != nil {
		return nil, err
	}

	var cards []map[int64]models.CardPackConfig
	for t, tier := range agg_pack.Tier {
		cards = append(cards, make(map[int64]models.CardPackConfig))
		t_name := ""
		switch t {
		case 0:
			t_name = "standard"
		case 1:
			t_name = "premium"
		case 2:
			t_name = "elite"
		}
		for cti, card_type := range tier {
			for i, v := range card_type {
				if fmt.Sprintf("%T", v) == "float64" {
					if v.(float64) == 0 {
						continue
					}
				}

				var card models.CardPackConfig
				c, ok := cards[t][i]
				if ok {
					card = c
					fmt.Printf("Card %d already exists in tier %s, adding %v\n", i, t_name, v)
					spew.Dump(c)
				}

				card.Changed++
				card.CardPackId = i
				card.Tier = t_name

				switch cti {
				case 0:
					card.Contains.Rare.Year = int64(v.(float64))
				case 1:
					card.Contains.Rare.DayMonth = int64(v.(float64))
				case 2:
					card.Contains.Rare.Category = int64(v.(float64))
				case 3:
					if fmt.Sprintf("%T", v) == "string" {
						card.Contains.Rare.Trigger = v.(string)
					}
				case 4:
					card.Contains.Uncommon.Year = int64(v.(float64))
				case 5:
					card.Contains.Uncommon.DayMonth = int64(v.(float64))
				case 6:
					card.Contains.Uncommon.Category = int64(v.(float64))
				case 7:
					if fmt.Sprintf("%T", v) == "string" {
						card.Contains.Uncommon.Trigger = v.(string)
					}
				case 8:
					card.Contains.Core.Year = int64(v.(float64))
				case 9:
					card.Contains.Core.DayMonth = int64(v.(float64))
				case 10:
					card.Contains.Core.Category = int64(v.(float64))
				case 11:
					if fmt.Sprintf("%T", v) == "string" {
						card.Contains.Core.Trigger = v.(string)
					}
				}

				cards[t][i] = card
			}
		}

	}

	js, err := json.Marshal(cards)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(js))

	err = ioutil.WriteFile("agg_out.json", js, 0644)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
