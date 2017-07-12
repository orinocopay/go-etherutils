package etherutils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

// NetworkID returns the network ID (also known as the chain ID) for this chain.
func NetworkID(ctx context.Context, c *rpc.Client) (*big.Int, error) {
	version := big.NewInt(0)

	var ver string
	err := c.CallContext(ctx, &ver, "net_version")
	if err == nil {
		version.SetString(ver, 10)
	}

	return version, err
}
