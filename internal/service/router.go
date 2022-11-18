package service

import (
	"gitlab.com/tokend/nft-books/network-svc/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/networks", func(r chi.Router) {
		r.Post("/", nil)
		r.Get("/", nil)

		r.Route("/detailed", func(r chi.Router) {
			r.Get("/", nil)
		})
	})

	return r
}
