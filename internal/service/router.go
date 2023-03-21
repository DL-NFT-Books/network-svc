package service

import (
	"github.com/dl-nft-books/network-svc/internal/config"
	"github.com/dl-nft-books/network-svc/internal/data/postgres"
	"github.com/dl-nft-books/network-svc/internal/service/handlers"
	"github.com/dl-nft-books/network-svc/internal/service/helpers"
	"github.com/dl-nft-books/network-svc/internal/service/middlewares"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxNetworksQ(postgres.NewNetworksQ(cfg.DB())),
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
		),
	)

	r.Route("/integrations/networks", func(r chi.Router) {
		// basic info
		r.Get("/", handlers.GetNetworksDefault)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetNetworkDefaultByChainID)
		})

		// detailed info
		r.With(middlewares.CheckAccessToken).
			Post("/", handlers.CreateNetwork)

		r.With(middlewares.CheckAccessToken).
			Route("/detailed", func(r chi.Router) {
				r.Get("/", handlers.GetNetworksDetailed)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handlers.GetNetworkDetailedByChainID)
				})
			})
	})

	return r
}
