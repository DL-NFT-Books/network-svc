package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/network-svc/internal/service/responses"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

func GetNetworksDetailed(w http.ResponseWriter, r *http.Request) {
	networks := GetNetworks(w, r)
	if networks == nil {
		ape.Render(w, resources.NetworkDetailedListResponse{
			Data: make([]resources.NetworkDetailed, 0),
		})

		return
	}

	ape.Render(w, responses.NewGetNetworksDetailedResponse(networks))
}
