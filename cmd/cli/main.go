package main

import (
	"os"

	"github.com/ethereum/go-ethereum/internal/cli"
	"github.com/ethereum/go-ethereum/params"

	// Force-load the tracer engines to trigger registration
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"
	_ "github.com/ethereum/go-ethereum/eth/tracers/live"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"
	_ "github.com/linxGnu/grocksdb"
)

func main() {
	params.UpdateBorInfo()
	os.Exit(cli.Run(os.Args[1:]))
}
