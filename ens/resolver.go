package ens

import (
	"bytes"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/orinocopay/go-etherutils/ens/enscontract"
	"github.com/orinocopay/go-etherutils/ens/resolvercontract"
)

var zeroHash = make([]byte, 32)
var zeroAddress = common.HexToAddress("00")

// Resolve resolves an ENS name in to an Etheruem address
func Resolve(name string) (addr common.Address, err error) {
	nameHash := NameHash(name)
	if bytes.Compare(nameHash[:], zeroHash) == 0 {
		err = errors.New("Bad name")
	} else {
		addr, err = resolveHash(nameHash)
	}
	return
}

func resolveHash(nameHash [32]byte) (address common.Address, err error) {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://api.orinocopay.com:8545/")
	if err != nil {
		return zeroAddress, err
	}

	// Instantiate the ENS contract
	ens, err := enscontract.NewEnscontract(common.HexToAddress("314159265dd8dbb310642f98f50c066173c1259b"), conn)
	if err != nil {
		return zeroAddress, err
	}

	// Check that this name is owned
	ownerAddress, err := ens.Owner(nil, nameHash)
	if err != nil {
		return zeroAddress, err
	}
	if bytes.Compare(ownerAddress.Bytes(), zeroAddress.Bytes()) == 0 {
		return zeroAddress, errors.New("unregistered name")
	}

	// Obtain the resolver for this name
	resolverAddress, err := ens.Resolver(nil, nameHash)
	if err != nil {
		return zeroAddress, err
	}
	if bytes.Compare(resolverAddress.Bytes(), zeroAddress.Bytes()) == 0 {
		return zeroAddress, errors.New("no resolver")
	}

	// Instantiate the resolver contract
	resolver, err := resolvercontract.NewResolvercontract(resolverAddress, conn)
	if err != nil {
		return zeroAddress, err
	}

	// Resolve the name
	address, err = resolver.Addr(nil, nameHash)
	if err != nil {
		return zeroAddress, err
	}
	if bytes.Compare(address.Bytes(), zeroAddress.Bytes()) == 0 {
		return zeroAddress, errors.New("no address")
	}

	return
}
