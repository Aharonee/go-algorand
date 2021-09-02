// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package crypto

import (
	"crypto/sha512"
	"errors"
	"hash"

	"github.com/algonathan/sumhash"
	"golang.org/x/crypto/sha3"
)

var compressor sumhash.LookupTable

func init() {
	C := 4
	N := 14
	shk := sha3.NewShake256()
	seed := []byte("I have nothing up my sleeve...")
	_, err := shk.Write(seed)
	if err != nil {
		panic(err)
	}
	compressor = sumhash.RandomMatrix(shk, N, C).LookupTable()
}

// HashType enum type for signing algorithms
type HashType uint64

// types of hashes
const (
	Sha512_256 HashType = iota
	Sumhash
)

//size of each hash
const (
	Sha512_256Size    = 32
	SumhashDigestSize = 112
)

// HashFactory is responsible for generating new hashes accordingly to the type it stores.
type HashFactory struct {
	_struct  struct{} `codec:",omitempty,omitemptyarray"`
	HashType HashType `codec:"t"`
}

var errUnknownHash = errors.New("unknown hash type")

// NewHash generates a new hash.Hash to use.
func (h HashFactory) NewHash() (hash.Hash, error) {
	switch h.HashType {
	case Sha512_256:
		return sha512.New512_256(), nil
	case Sumhash:
		return sumhash.New(compressor), nil
	default:
		return nil, errUnknownHash
	}
}

// HashSum Makes it easier to sum using hash interface and Hashable interface
func HashSum(hsh hash.Hash, h Hashable) []byte {
	rep := HashRep(h)
	return HashBytes(hsh, rep)
}

// HashBytes Makes it easier to sum using hash interface.
func HashBytes(hash hash.Hash, m []byte) []byte {
	hash.Reset()
	hash.Write(m)
	outhash := hash.Sum(nil)
	return outhash
}
