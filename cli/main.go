package main

import (
	"os"

	"github.com/loomnetwork/zombie_battleground/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
