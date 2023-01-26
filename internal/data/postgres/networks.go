package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
	"log"
)

const (
	networksTableName = "networks"
	chainIdColumn     = "chain_id"
)

func NewNetworksQ(db *pgdb.DB) data.NetworksQ {
	return &NetworksQ{
		db:            db.Clone(),
		selectBuilder: squirrel.Select("*").From(networksTableName),
	}
}

func (n *NetworksQ) InitNetworksQ(init []data.Network, log *logan.Entry) error {
	res, err := n.Get()
	if err != nil {
		return errors.Wrap(err, "failed to check if networks exists")
	}
	if res != nil {
		log.Warn("network db is already init")
		return nil
	}
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
		"token_symbol",
		"decimals")
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
			network.NativeTokenSymbol,
			network.Decimals)
	}
	return n.db.Exec(stmt)
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
	log.Println("n.selectBuilder.ToSql()")
	log.Println(n.selectBuilder.ToSql())
	return result, err
}

func (n *NetworksQ) FilterByChainID(chainId int64) data.NetworksQ {
	n.selectBuilder = n.selectBuilder.Where(squirrel.Eq{
		chainIdColumn: chainId,
	})

	return n
}
