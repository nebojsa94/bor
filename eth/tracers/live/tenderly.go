package live

import (
	gethtracers "github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/tenderly/live-tracer-ethereum/live"
)

// init registers the tenderly tracer hooks to the live directory.
func init() {
	gethtracers.LiveDirectory.Register("tenderly", live.NewTenderlyTracerHooks)
}
