package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	doormanCfg "github.com/dl-nft-books/doorman/connector/config"
	"github.com/dl-nft-books/network-svc/internal/data/mem"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	doormanCfg.DoormanConfiger
	mem.InitialNetworker
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	doormanCfg.DoormanConfiger
	mem.InitialNetworker
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:           getter,
		Databaser:        pgdb.NewDatabaser(getter),
		Copuser:          copus.NewCopuser(getter),
		Listenerer:       comfig.NewListenerer(getter),
		Logger:           comfig.NewLogger(getter, comfig.LoggerOpts{}),
		DoormanConfiger:  doormanCfg.NewDoormanConfiger(getter),
		InitialNetworker: mem.NewInitialNetworker(getter),
	}
}
