package data

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

type Network struct {
	ID                int64  `db:"id" structs:"-"`
	Name              string `db:"name" structs:"name"`
	ChainID           int64  `db:"chain_id" structs:"chain_id"`
	RpcUrl            string `db:"rpc_url" structs:"rpc_url"`
	WebSocketURL      string `db:"ws_url" structs:"ws_url"`
	FactoryAddress    string `db:"factory_address" structs:"factory_address"`
	FactoryName       string `db:"factory_name" structs:"factory_name"`
	FactoryVersion    string `db:"factory_version" structs:"factory_version"`
	FirstBlock        int64  `db:"first_block" structs:"first_block"`
	NativeTokenName   string `db:"token_name" structs:"token_name"`
	NativeTokenSymbol string `db:"token_symbol" structs:"token_symbol"`
	Decimals          int64  `db:"decimals" structs:"decimals"`
}

type NetworksQ interface {
	New() NetworksQ

	Insert(data Network) (id int64, err error)
	InitNetworksQ(init []Network, log *logan.Entry) error
	Get() (*Network, error)
	Select() ([]Network, error)

	FilterByChainID(chainId int64) NetworksQ
}

func (n *Network) ResourceDefault() resources.Network {
	return resources.Network{
		Key: resources.NewKeyInt64(n.ID, resources.NETWORKS),
		Attributes: resources.NetworkAttributes{
			Name:           n.Name,
			ChainId:        n.ChainID,
			FactoryAddress: n.FactoryAddress,
			TokenName:      n.NativeTokenName,
			TokenSymbol:    n.NativeTokenSymbol,
			Decimals:       n.Decimals,
		},
	}
}

func (n *Network) ResourceDetailed() resources.NetworkDetailed {
	return resources.NetworkDetailed{
		Key: resources.NewKeyInt64(n.ID, resources.NETWORKS),
		Attributes: resources.NetworkDetailedAttributes{
			Name:           n.Name,
			ChainId:        n.ChainID,
			RpcUrl:         n.RpcUrl,
			WsUrl:          n.WebSocketURL,
			FactoryAddress: n.FactoryAddress,
			FactoryName:    n.FactoryName,
			FactoryVersion: n.FactoryVersion,
			FirstBlock:     n.FirstBlock,
			TokenName:      n.NativeTokenName,
			TokenSymbol:    n.NativeTokenSymbol,
			Decimals:       n.Decimals,
		},
	}
}
