package ens

import (
	"strings"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// NameHash generates a hash from a name that can be used to
// look up the name in ENS
func NameHash(name string) (hash [32]byte) {
	if name != "" {
		parts := strings.Split(name, ".")
		for i := len(parts) - 1; i >= 0; i-- {
			hash = nameHashPart(hash, parts[i])
		}
	}
	return
}

func nameHashPart(currentHash [32]byte, name string) (hash [32]byte) {
	sha := sha3.NewKeccak256()
	sha.Write(currentHash[:])
	nameSha := sha3.NewKeccak256()
	nameSha.Write([]byte(name))
	nameHash := nameSha.Sum(nil)
	sha.Write(nameHash)
	sha.Sum(hash[:0])
	return
}
