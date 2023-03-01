package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
)

const (
	networksTableName            = "networks"
	networksIdColumn             = "id"
	networksNameColumn           = "name"
	networksRpcUrlColumn         = "rpc_url"
	networksWsUrlColumn          = "ws_url"
	networksFactoryAddressColumn = "factory_address"
	networksFactoryNameColumn    = "factory_name"
	networksFactoryVersionColumn = "factory_version"
	networksFirstBlockColumn     = "first_block"
	networksTokenNameColumn      = "token_name"
	networksTokenSymbolColumn    = "token_symbol"
	networksDecimalsColumn       = "decimals"
	networksChainIdColumn        = "chain_id"
)

func NewNetworksQ(db *pgdb.DB) data.NetworksQ {
	return &NetworksQ{
		db:            db.Clone(),
		selectBuilder: squirrel.Select("*").From(networksTableName),
	}
}

type NetworksQ struct {
	db            *pgdb.DB
	selectBuilder squirrel.SelectBuilder
}

func (n *NetworksQ) New() data.NetworksQ {
	return NewNetworksQ(n.db)
}

func (n *NetworksQ) Insert(data ...data.Network) ([]int64, error) {
	var id []int64
	stmt := squirrel.
		Insert(networksTableName).Columns(
		networksNameColumn,
		networksRpcUrlColumn,
		networksWsUrlColumn,
		networksFactoryAddressColumn,
		networksFactoryNameColumn,
		networksFactoryVersionColumn,
		networksFirstBlockColumn,
		networksTokenNameColumn,
		networksTokenSymbolColumn,
		networksDecimalsColumn,
		networksChainIdColumn)
	for _, network := range data {
		stmt = stmt.Values(getValuesFromNetwork(network)...)
	}
	err := n.db.Select(&id, stmt.Suffix("returning id"))
	return id, err
}

func getValuesFromNetwork(network data.Network) []interface{} {
	var res []interface{}
	res = append(res,
		network.Name,
		network.RpcUrl,
		network.WebSocketURL,
		network.FactoryAddress,
		network.FactoryName,
		network.FactoryVersion,
		network.FirstBlock,
		network.NativeTokenName,
		network.NativeTokenSymbol,
		network.Decimals,
		network.ChainID)
	return res
}

func (n *NetworksQ) Get() (*data.Network, error) {
	var result data.Network

	err := n.db.Get(&result, n.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (n *NetworksQ) Select() ([]data.Network, error) {
	var result []data.Network

	err := n.db.Select(&result, n.selectBuilder)
	return result, err
}

func (n *NetworksQ) FilterByChainID(chainId int64) data.NetworksQ {
	n.selectBuilder = n.selectBuilder.Where(squirrel.Eq{
		networksChainIdColumn: chainId,
	})

	return n
}
