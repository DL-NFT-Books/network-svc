package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
)

const (
	networksTableName = "networks"
	chainIdColumn     = "chain_id"
)

func NewNetworksQ(db *pgdb.DB, init ...data.Network) data.NetworksQ {
	if len(init) > 0 {
		stmt := squirrel.
			Insert(networksTableName).Columns(
			"name",
			"chain_id",
			"rpc_url",
			"ws_url",
			"factory_address",
			"factory_name",
			"factory_version",
			"first_block",
			"token_name",
			"token_symbol")
		for _, network := range init {
			stmt = stmt.Values(
				network.Name,
				network.ChainID,
				network.RpcUrl,
				network.WebSocketURL,
				network.FactoryAddress,
				network.FactoryName,
				network.FactoryVersion,
				network.FirstBlock,
				network.NativeTokenName,
				network.NativeTokenSymbol)
		}
		err := db.Exec(stmt)
		if err != nil {
			panic(errors.Wrap(err, "failed to initialize network db"))
			return nil
		}
	}
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

func (n *NetworksQ) Insert(data data.Network) (int64, error) {
	clauses := structs.Map(data)
	var id int64

	stmt := squirrel.
		Insert(networksTableName).
		SetMap(clauses).
		Suffix("returning id")
	err := n.db.Get(&id, stmt)

	return id, err
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
		chainIdColumn: chainId,
	})

	return n
}
