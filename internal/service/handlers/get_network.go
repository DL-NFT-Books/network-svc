package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
	"gitlab.com/tokend/nft-books/network-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/network-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

func GetNetworkDefaultByChainID(w http.ResponseWriter, r *http.Request) {
	network := GetNetworkByChainId(w, r)
	if network == nil {
		return
	}
	ape.Render(w, resources.NetworkResponse{
		Data: network.ResourceDefault(),
	})
}

func GetNetworkByChainId(w http.ResponseWriter, r *http.Request) *data.Network {
	logger := helpers.Log(r)

	request, err := requests.NewNetworkByChainIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return nil
	}

	network, err := helpers.NetworksQ(r).FilterByChainID(request.Id).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get network")
		ape.RenderErr(w, problems.InternalError())
		return nil
	}
	if network == nil {
		ape.RenderErr(w, problems.NotFound())
		return nil
	}

	return network
}
