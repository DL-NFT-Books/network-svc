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

func CreateNetwork(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewCreateNetworkRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	networkData := request.Data.Attributes

	createdNetworkId, err := helpers.NetworksQ(r).Insert(data.Network{
		Name:              networkData.Name,
		ChainID:           networkData.ChainId,
		RpcUrl:            networkData.RpcUrl,
		WebSocketURL:      networkData.WsUrl,
		FactoryAddress:    networkData.FactoryAddress,
		FactoryName:       networkData.FactoryName,
		FactoryVersion:    networkData.FactoryVersion,
		NativeTokenName:   networkData.TokenName,
		NativeTokenSymbol: networkData.TokenSymbol,
	})
	if err != nil {
		logger.WithError(err).Debug("failed to insert new network")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	request.Data.Key = resources.NewKeyInt64(createdNetworkId, resources.NETWORKS)
	ape.Render(w, resources.NetworkDetailedResponse{
		Data: request.Data,
	})
}
