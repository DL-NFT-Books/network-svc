package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/dl-nft-books/network-svc/internal/data"
	"github.com/dl-nft-books/network-svc/internal/service/helpers"
	"github.com/dl-nft-books/network-svc/internal/service/responses"
	"github.com/dl-nft-books/network-svc/resources"
)

func GetNetworksDefault(w http.ResponseWriter, r *http.Request) {
	networks := GetNetworks(w, r)
	if networks == nil {
		ape.Render(w, resources.NetworkListResponse{
			Data: make([]resources.Network, 0),
		})

		return
	}

	ape.Render(w, responses.NewGetNetworksDefaultResponse(networks))
}

func GetNetworks(w http.ResponseWriter, r *http.Request) []data.Network {
	logger := helpers.Log(r)

	networks, err := helpers.NetworksQ(r).Select()
	if err != nil {
		logger.WithError(err).Error("failed to get networks")
		ape.RenderErr(w, problems.InternalError())
	}

	return networks
}
