package helpers

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"github.com/dl-nft-books/network-svc/internal/config"
	"github.com/dl-nft-books/network-svc/internal/data"
	"github.com/dl-nft-books/network-svc/internal/data/postgres"
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
		networks, err := i.networksQ.Select()
		if err != nil {
			return errors.Wrap(err, "failed to get networks")
		}
		// If the database is already initialized skipping initializing
		if len(networks) > 0 {
			i.logger.Info("DB is already initialized")
			return nil
		}
		if _, err := i.networksQ.Insert(i.initData...); err != nil {
			return err
		}

		i.logger.Info("Successfully init")
		return nil
	}
	i.logger.Info("No data to init")
	return nil
}
