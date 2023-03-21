package main

import (
	"os"

	"github.com/dl-nft-books/network-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
