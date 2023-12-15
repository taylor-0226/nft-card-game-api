package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCelebritiesWithBirthdates() (map[string]models.Celebrity, error) {
	out := make(map[string]models.Celebrity)
	var err error
	file, err := os.Open("celebrities_with_birthdates.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		name := strings.ToLower(line[1])

		// c_raw := line[2]
		// categories := strings.Split(c_raw, "|")

		b_raw := line[3]
		b_split := strings.Split(b_raw, "-")
		if len(b_split) < 3 {
			continue
		}

		year, err := strconv.ParseInt(b_split[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		month, err := strconv.ParseInt(b_split[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		day, err := strconv.ParseInt(b_split[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		// ccategories := models.Strings(categories)

		out[name] = models.Celebrity{
			CelebrityData: models.CelebrityData{
				Id:         &id,
				Name:       &name,
				BirthMonth: &month,
				BirthYear:  &year,
				BirthDay:   &day,
			},
		}

		// store.NewCelebrity(context.TODO(), &models.Celebrity{
		// 	CelebrityData: models.CelebrityData{
		// 		Id:         &id,
		// 		Name:       &name,
		// 		BirthMonth: &month,
		// 		BirthYear:  &year,
		// 		BirthDay:   &day,
		// 		Categories: &ccategories,
		// 	},
		// })

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// for _, c := range categories {
		// 	new := models.Category{
		// 		CategoryData: models.CategoryData{
		// 			Name: &c,
		// 		},
		// 	}

		// 	store.NewCategory(context.TODO(), &new)
		// }

	}

	return out, nil
}

func ImportArticles() {

}

func ImportCelebrities() error {
	var err error
	var categories []string
	celebrities, err := GetCelebritiesWithBirthdates()
	if err != nil {
		return err
	}

	file, err := os.Open("celebrities.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(file))
	count := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		count++

		if count == 1 {
			categories = line

			for _, c := range categories {
				new := models.Category{
					CategoryData: models.CategoryData{
						Name: &c,
					},
				}
				err := store.NewCategory(context.Background(), &new)
				if err != nil {
					log.Println(err)
				}
			}

		} else {
			actor := line[0]
			if actor != "" {
				name := strings.ToLower(actor)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[0])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}

			}

			musician := line[1]
			if musician != "" {
				name := strings.ToLower(musician)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[1])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			social_media_personality := line[2]
			if social_media_personality != "" {
				name := strings.ToLower(social_media_personality)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[2])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			entertainer := line[3]
			if entertainer != "" {
				name := strings.ToLower(entertainer)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[3])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			athlete := line[4]
			if athlete != "" {
				name := strings.ToLower(athlete)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[4])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			politician := line[5]
			if politician != "" {
				name := strings.ToLower(politician)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[5])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			business_leader := line[6]
			if business_leader != "" {
				name := strings.ToLower(business_leader)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[6])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

			infamous_person := line[7]
			if infamous_person != "" {
				name := strings.ToLower(infamous_person)
				o, ok := celebrities[name]
				if ok {
					category := strings.ToLower(categories[7])
					new := models.Celebrity{
						CelebrityData: models.CelebrityData{
							Name:       &name,
							BirthDay:   o.BirthDay,
							BirthMonth: o.BirthMonth,
							BirthYear:  o.BirthYear,
							Category:   &category,
						},
					}

					err := store.NewCelebrity(context.Background(), &new)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}

		}
	}

	return nil
}

func ImportTriggers() error {
	var err error
	var triggers []string

	file, err := os.Open("triggers.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(file))
	count := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		count++

		if count == 1 {
			triggers = line
		} else {
			fmt.Println(line)
			major := line[0]
			if major != "" {
				name := strings.ToLower(major)
				trigger := strings.ToLower(triggers[0])
				new := models.Trigger{
					TriggerData: models.TriggerData{
						Name: &name,
						Tier: &trigger,
					},
				}

				err := store.NewTrigger(context.Background(), &new)
				if err != nil {
					log.Println(err)
				}
			}

			minor_1 := line[1]
			if minor_1 != "" {
				name := strings.ToLower(minor_1)
				trigger := strings.ToLower(triggers[1])
				new := models.Trigger{
					TriggerData: models.TriggerData{
						Name: &name,
						Tier: &trigger,
					},
				}

				err := store.NewTrigger(context.Background(), &new)
				if err != nil {
					log.Println(err)
				}
			}

			minor_2 := line[2]
			if minor_2 != "" {
				name := strings.ToLower(minor_2)
				trigger := strings.ToLower(triggers[2])
				new := models.Trigger{
					TriggerData: models.TriggerData{
						Name: &name,
						Tier: &trigger,
					},
				}

				err := store.NewTrigger(context.Background(), &new)
				if err != nil {
					log.Println(err)
				}
			}
		}

	}

	return nil
}
