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
	"errors"
	"fmt"

	"github.com/algorand/go-algorand/protocol"
)

type (
	//ByteSignature using unspecified bound.
	//msgp:allocbound ByteSignature
	ByteSignature []byte

	// AlgorithmType enum type for signing algorithms
	AlgorithmType uint16
)

// all AlgorithmType enums
const (
	DilithiumType AlgorithmType = iota
	Ed25519Type

	MaxAlgorithmType
)

// IsValid verifies that the type of the algorithm is known
func (z AlgorithmType) IsValid() error {
	if z >= MaxAlgorithmType {
		return protocol.ErrInvalidObject
	}
	return nil
}

// Signer interface represents the possible operations that can be done with a signing key.
type Signer interface {
	Sign(message Hashable) ByteSignature
	SignBytes(message []byte) ByteSignature
	GetVerifyingKey() *GenericVerifyingKey
}

// ErrBadSignature represents a bad signature
var ErrBadSignature = fmt.Errorf("invalid signature")

// Verifier interface represents any algorithm that can verify signatures for a specific signing scheme.
type Verifier interface {
	// Verify and VerifyBytes returns error on bad signature, and any other problem.
	Verify(message Hashable, sig ByteSignature) error
	VerifyBytes(message []byte, sig ByteSignature) error
}

// GenericSigningKey holds a Signer, and the type of algorithm the Signer conforms with.
//msgp:postunmarshalcheck GenericSigningKey IsValid
type GenericSigningKey struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	Type AlgorithmType `codec:"sigType"`

	DilithiumSigner DilithiumSigner `codec:"ds"`
	Ed25519Singer   Ed25519Key      `codec:"edds"`
	invalidSinger   invalidSinger
}

// IsValid states whether the GenericSigningKey is valid, and is safe to use.
func (z *GenericSigningKey) IsValid() error {
	return z.Type.IsValid()
}

// GenericVerifyingKey is an abstraction of a key store of verifying keys.
// it can return the correct key according to the underlying algorithm.
// Implements Hashable too.
//
// NOTE: The GenericVerifyingKey key might not be a valid key if a malicious client sent it over the network
// make certain it is valid.
//msgp:postunmarshalcheck GenericVerifyingKey IsValid
type GenericVerifyingKey struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	Type AlgorithmType `codec:"type"`

	DilithiumPublicKey DilithiumVerifier `codec:"dpk"`
	Ed25519PublicKey   Ed25519PublicKey  `codec:"edpk"`
	invalidVerifier    invalidVerifier
}

// GetSigner fetches the Signer type that is stored inside this GenericSigningKey.
func (z *GenericSigningKey) GetSigner() Signer {
	switch z.Type {
	case DilithiumType:
		return &z.DilithiumSigner
	case Ed25519Type:
		return &z.Ed25519Singer
	default:
		return &z.invalidSinger
	}
}

// GetVerifier fetches the Verifier type that is stored inside this GenericVerifyingKey.
func (z *GenericVerifyingKey) GetVerifier() Verifier {
	switch z.Type {
	case DilithiumType:
		return &z.DilithiumPublicKey
	case Ed25519Type:
		return &z.Ed25519PublicKey
	default:
		return &z.invalidVerifier
	}
}

var errUnknownVerifier = errors.New("could not find stored Verifier")

var errUnknownSigner = errors.New("could not find stored signer")

var errNonExistingSignatureAlgorithmType = errors.New("signing algorithm type does not exist")

// NewSigner receives a type of signing algorithm and generates keys.
func NewSigner(t AlgorithmType) (*GenericSigningKey, error) {
	switch t {
	case DilithiumType:
		signer := NewDilithiumSigner().(*DilithiumSigner)
		return &GenericSigningKey{
			Type:            t,
			DilithiumSigner: *signer,
		}, nil
	case Ed25519Type:
		var seed Seed
		SystemRNG.RandBytes(seed[:])
		key := GenerateEd25519Key(seed)
		return &GenericSigningKey{
			Type:          t,
			Ed25519Singer: *key,
		}, nil
	default:
		return nil, errNonExistingSignatureAlgorithmType
	}
}
