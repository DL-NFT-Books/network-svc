package cli

import (
	"gitlab.com/distributed_lab/logan/v3"
	"github.com/dl-nft-books/network-svc/internal/config"
	"github.com/dl-nft-books/network-svc/internal/service"
	"github.com/dl-nft-books/network-svc/internal/service/helpers"

	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
)

var (
	app = kingpin.New("network-svc", "")

	// Run commands
	runCmd     = app.Command("run", "run command")
	serviceCmd = runCmd.Command("service", "run service")
	dbCmd      = runCmd.Command("init-db", "init db")

	// Migration commands
	migrateCmd     = app.Command("migrate", "migrate command")
	migrateUpCmd   = migrateCmd.Command("up", "migrate db up")
	migrateDownCmd = migrateCmd.Command("down", "migrate db down")
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		service.Run(cfg)
	case dbCmd.FullCommand():
		err = helpers.NewInitDBer(cfg).Run()
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}
	return true
}
