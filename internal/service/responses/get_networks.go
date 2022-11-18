package responses

import (
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

func NewGetNetworksDefaultResponse(data []data.Network) resources.NetworkListResponse {
	networks := make([]resources.Network, len(data))

	for i, value := range data {
		networks[i] = value.ResourceDefault()
	}

	return resources.NetworkListResponse{
		Data: networks,
	}
}

func NewGetNetworksDetailedResponse(data []data.Network) resources.NetworkDetailedListResponse {
	networks := make([]resources.NetworkDetailed, len(data))

	for i, value := range data {
		networks[i] = value.ResourceDetailed()
	}

	return resources.NetworkDetailedListResponse{
		Data: networks,
	}
}
