package main

import (
	"context"
	"gameon-twotwentyk-api/connections"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/handlers"
	"gameon-twotwentyk-api/store"
	"gameon-twotwentyk-api/venly"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func main() {
	var err error

	// _, err = DecodeAggPack("./agg_pack.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// os.Exit(0)

	r := chi.NewRouter()

	corsParams := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsParams.Handler)
	r.Use(middleware.Logger)
	r.Use(jwtauth.Verifier(handlers.TokenAuth))
	r.Use(handlers.Authenticator)

	connections.InitPostgres()
	graphql.Init()

	store.RefreshCategoryMap(context.Background())
	store.RefreshCelebrityMap(context.Background())
	store.RefreshTriggerMap(context.Background())
	// store.RefreshNftCollectionMap(context.Background())

	c, err := venly.NewClient(venly.VenlyClientConfig{
		ClientId:     venly.VENLY_CLIENT_ID,
		ClientSecret: venly.VENLY_APP_SECRET,
	})
	if err != nil {
		log.Fatal(err)
	}

	venly.Global = c

	r.Get("/nft_card_day_month/{nft_card_day_month_id}", handlers.GetNftCardDayMonth)

	r.Route("/trigger", func(r chi.Router) {
		r.Get("/", handlers.ListTrigger)
	})

	r.Route("/category", func(r chi.Router) {
		r.Get("/", handlers.ListCategory)
	})

	r.Route("/celebrity", func(r chi.Router) {
		r.Get("/", handlers.ListCelebrity)
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewCelebrity)
			r.Route("/{celebrity_id}", func(r chi.Router) {
				r.Get("/", handlers.GetCelebrity)
				r.Put("/", handlers.UpdateCelebrity)
				r.Delete("/", handlers.DeleteCelebrity)
			})
		})
	})

	r.Route("/claim", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewClaim)
			r.Group(func(r chi.Router) {
				r.Use(handlers.RestrictAdmin)
				r.Get("/", handlers.ListClaim)

				r.Route("/{claim_id}", func(r chi.Router) {
					r.Get("/", handlers.GetClaim)
					r.Put("/", handlers.UpdateClaim)
					r.Delete("/", handlers.DeleteClaim)
					r.Post("/approve", handlers.ApproveClaim)
					r.Post("/reject", handlers.RejectClaim)
				})
			})

		})
	})

	r.Route("/article", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewArticle)
			r.Route("/{article_id}", func(r chi.Router) {
				r.Get("/", handlers.GetArticle)
				r.Put("/", handlers.UpdateArticle)
				r.Delete("/", handlers.DeleteArticle)
			})
		})
	})

	r.Route("/article_source", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewArticleSource)
			r.Route("/{article_source_id}", func(r chi.Router) {
				r.Get("/", handlers.GetArticleSource)
				r.Put("/", handlers.UpdateArticleSource)
				r.Delete("/", handlers.DeleteArticleSource)
			})
		})
	})

	r.Route("/me", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		r.Put("/", handlers.UpdateUser)
		r.Get("/", handlers.GetMyUserData)
		r.Get("/claim", handlers.ListClaimForUserById)
		r.Get("/nft", handlers.GetMyNfts)
		r.Get("/marketplace_listing", handlers.ListMarketplaceListingForUserById)
		r.Get("/recipe/identity", handlers.GetAvailableCelebrityRecipes)
		// r.Get("/recipe/prediction", handlers.GetMyPredictionRecipes)
		r.Get("/nft_card_identity", handlers.ListNftCardIdentityForUserById)
		r.Get("/nft_card_prediction", handlers.ListNftCardPredictionForUserById)
		r.Get("/nft_card_trigger", handlers.ListNftCardTriggerForUserById)
		r.Get("/nft_card_category", handlers.ListNftCardCategoryForUserById)
		r.Get("/nft_card_day_month", handlers.ListNftCardDayMonthForUserById)
		r.Get("/nft_card_year", handlers.ListNftCardYearForUserById)
		r.Get("/nft_card_crafting", handlers.ListNftCardCraftingForUserById)
		r.Get("/card_pack", handlers.ListCardPackForUserById)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
		r.Post("/register", handlers.Register)
		r.Post("/google", handlers.AuthGoogle)
		r.Post("/apple", handlers.AuthApple)
	})

	r.Route("/feed", func(r chi.Router) {
		r.Get("/", handlers.SearchArticles)
		r.Get("/personalised", handlers.GetArticlesPersonalised)
	})

	// blockchain
	//crafting

	r.Route("/nft", func(r chi.Router) {
		r.Post("/identity", handlers.CraftIdentity)
		r.Post("/prediction", handlers.CraftPrediction)
		r.Post("/trigger", handlers.NewNftCardTrigger)
		r.Post("/day_month", handlers.NewNftCardDayMonth)
		r.Post("/crafting", handlers.NewNftCardCrafting)
		r.Post("/category", handlers.NewNftCardCategory)
		r.Post("/year", handlers.NewNftCardYear)

		// r.Post("/recipe/prediction", handlers.GenerateNFTsForRecipePrediction)

	})

	r.Route("/marketplace_listing", func(r chi.Router) {
		r.Get("/", handlers.SearchMarketplaceListings)
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewMarketplaceListing)
			r.Route("/{marketplace_listing_id}", func(r chi.Router) {
				r.Get("/", handlers.GetMarketplaceListing)
				r.Put("/", handlers.UpdateMarketplaceListing)
				r.Delete("/", handlers.DeleteMarketplaceListing)
				r.Post("/buy", handlers.BuyMarketplaceListing)
			})
		})
	})

	r.Route("/card_series", func(r chi.Router) {
		r.Get("/", handlers.ListCardSeries)
		r.Route("/{card_series_id}", func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/buy", handlers.BuyCardSeriesPack)
		})
	})

	r.Route("/card_pack", func(r chi.Router) {
		r.Route("/{card_pack_id}", func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/open", handlers.OpenCardPack)
		})
	})

	r.Route("/nft_card_identity", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		r.Route("/{nft_card_identity_id}", func(r chi.Router) {
			r.Put("/", handlers.UpdateNftCardIdentity)
		})

	})

	r.Route("/card_collection", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		r.Use(handlers.RestrictAdmin)
		r.Get("/", handlers.ListCardCollection)
		r.Post("/", handlers.CreateNftCollection)
		r.Route("/{card_collection_id}", func(r chi.Router) {
			r.Get("/", handlers.GetCardCollection)
			r.Put("/", handlers.UpdateCardCollection)
		})
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		// r.Use(handlers.RestrictAdmin)
		r.Post("/generate/identity", handlers.GenerateNFTsForRecipeIdentity)
		r.Post("/generate/triggers", handlers.GenerateAllTriggers)
		// r.Post("/add_balance", handlers.AddBalanceToUser)
	})

	r.Route("/webhook", func(r chi.Router) {
		r.Route("/moonpay", func(r chi.Router) {
			r.Route("/nft", func(r chi.Router) {
				r.Get("/asset_info/{contract_address}/{token_id}", handlers.WebhookMoonpayGetNftInfo)

			})

			r.Route("/transaction", func(r chi.Router) {
				r.Post("/", handlers.WebhookMoonpayTransaction)
			})
		})
	})

	r.Route("/ws", func(r chi.Router) {
		// r.Connect("/", handlers.WebsocketUpgrade)
	})

	err = http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatalf("Error serving HTTP handlers: %v", err)
	}

}
