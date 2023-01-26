package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

func GetNetworkDetailedByChainID(w http.ResponseWriter, r *http.Request) {
	network := GetNetworkByChainId(w, r)
	if network == nil {
		return
	}
	ape.Render(w, resources.NetworkDetailedResponse{
		Data: network.ResourceDetailed(),
	})
}
