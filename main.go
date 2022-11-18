package main

import (
	"os"

	"gitlab.com/tokend/nft-books/network-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
