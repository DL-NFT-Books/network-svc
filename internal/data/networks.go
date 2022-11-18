package data

import "gitlab.com/tokend/nft-books/network-svc/resources"

type Network struct {
	ID             int64  `db:"id" structs:"-"`
	Name           string `db:"name" structs:"name"`
	ChainID        int64  `db:"chain_id" structs:"chain_id"`
	RpcUrl         string `db:"rpc_url" structs:"rpc_url"`
	WebSocketURL   string `db:"ws_url" structs:"ws_url"`
	FactoryAddress string `db:"factory_address" structs:"factory_address"`
}

type NetworksQ interface {
	New() NetworksQ

	Insert(data Network) (id int64, err error)
	Get() (*Network, error)
	Select() ([]Network, error)

	FilterByChainID(chainId int64) NetworksQ
}

func (n *Network) ResourceDefault() *resources.Network {
	return &resources.Network{
		Key: resources.NewKeyInt64(n.ID, resources.NETWORKS),
		Attributes: resources.NetworkAttributes{
			Name:           n.Name,
			ChainId:        int32(n.ChainID),
			FactoryAddress: n.FactoryAddress,
		},
	}
}

func (n *Network) ResourceDetailed() *resources.NetworkDetailed {
	return &resources.NetworkDetailed{
		Key: resources.NewKeyInt64(n.ID, resources.NETWORKS),
		Attributes: resources.NetworkDetailedAttributes{
			Name:           n.Name,
			ChainId:        int32(n.ChainID),
			RpcUrl:         n.RpcUrl,
			WsUrl:          n.WebSocketURL,
			FactoryAddress: n.FactoryAddress,
		},
	}
}
