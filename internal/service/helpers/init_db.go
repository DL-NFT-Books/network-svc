package helpers

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/network-svc/internal/config"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
	"gitlab.com/tokend/nft-books/network-svc/internal/data/postgres"
)

type InitBDer struct {
	logger    *logan.Entry
	networksQ data.NetworksQ
	initData  []data.Network
}

func NewInitDBer(cfg config.Config) *InitBDer {
	return &InitBDer{
		logger:    cfg.Log(),
		networksQ: postgres.NewNetworksQ(cfg.DB()),
		initData:  cfg.InitialNetworks(),
	}
}

func (i *InitBDer) Run() error {
	i.logger.Info("Start to initial database")
	if len(i.initData) > 0 {
		return i.networksQ.InitNetworksQ(i.initData)
		i.logger.Info("Successfully init")
	}
	i.logger.Info("No data to init")
	return nil
}
