// Copyright (C) 2019-2022 Algorand, Inc.
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

package merklesignature

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/crypto/merklearray"
)

type (
	// Signature represents a signature in the merkle signature scheme using falcon signatures as an underlying crypto scheme.
	// It consists of an ephemeral public key, a signature, a merkle verification path and an index.
	// The merkle signature considered valid only if the Signature is verified under the ephemeral public key and
	// the Merkle verification path verifies that the ephemeral public key is located at the given index of the tree
	// (for the root given in the long-term public key).
	// More details can be found on Algorand's spec
	Signature struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		Signature             crypto.FalconSignature      `codec:"sig"`
		VectorCommitmentIndex uint64                      `codec:"idx"`
		Proof                 merklearray.SingleLeafProof `codec:"prf"`
		VerifyingKey          crypto.FalconVerifier       `codec:"vkey"`
	}

	// Secrets contains the private data needed by the merkle signature scheme.
	Secrets struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		// these keys should be temporarily stored in memory until Persist is called,
		// in which they will be dumped into database and disposed of.
		// non-exported fields to prevent msgpack marshalling
		ephemeralKeys []crypto.FalconSigner

		SignerContext
	}

	// Signer represents the StateProof signer for a specified round.
	//msgp:ignore Signer
	Signer struct {
		SigningKey *crypto.FalconSigner

		// The requested round to be signed by this Signer. Use FirstRoundInKeyLifetime in order to calculate the actual round of the key.
		Round uint64

		SignerContext
	}

	// SignerContext contains all the immutable data and metadata related to merklesignature.Secrets (without the secret keys)
	SignerContext struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		FirstValid  uint64           `codec:"fv"`
		KeyLifetime uint64           `codec:"iv"`
		Tree        merklearray.Tree `codec:"tree"`
	}

	// Commitment represents the root of the vector commitment tree built upon the MSS keys.
	Commitment [MerkleSignatureSchemeRootSize]byte

	// Verifier is used to verify a merklesignature.Signature produced by merklesignature.Secrets.
	Verifier struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		Commitment  Commitment `codec:"cmt"`
		KeyLifetime uint64     `codec:"lf"`
	}

	//KeyRoundPair represents an ephemeral signing key with it's corresponding round
	KeyRoundPair struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		Round uint64               `codec:"rnd"`
		Key   *crypto.FalconSigner `codec:"key"`
	}
)

// KeyLifetimeDefault defines the default lifetime of a key in the merkle signature scheme (in rounds).
const KeyLifetimeDefault = 256

// SchemeSaltVersion is the current salt version of merkleSignature
const SchemeSaltVersion = byte(0)

// CryptoPrimitivesID is an identification that the Merkle Signature Scheme uses a subset sum hash function
// and a falcon signature scheme.
var CryptoPrimitivesID = uint16(0)

// Errors for the merkle signature scheme
var (
	ErrStartBiggerThanEndRound           = errors.New("cannot create Merkle Signature Scheme because end round is smaller then start round")
	ErrKeyLifetimeIsZero                 = errors.New("received zero KeyLifetime")
	ErrNoStateProofKeyForRound           = errors.New("no stateproof key exists for this round")
	ErrSignatureSchemeVerificationFailed = errors.New("merkle signature verification failed")
	ErrSignatureSaltVersionMismatch      = fmt.Errorf("the signature's salt version does not match")
)

// New creates secrets needed for the merkle signature scheme.
// This function generates one key for each round within the participation period [firstValid, lastValid] (inclusive bounds)
// which holds round % interval == 0.
// In case firstValid equals zero then signer will generate all keys from (0,Z], i.e will not generate key for round zero.
func New(firstValid, lastValid, keyLifetime uint64) (*Secrets, error) {
	if firstValid > lastValid {
		return nil, ErrStartBiggerThanEndRound
	}
	if keyLifetime == 0 {
		return nil, ErrKeyLifetimeIsZero
	}

	// calculates the number of indices from first valid round and up to lastValid.
	// writing this explicit calculation to avoid overflow.
	numberOfKeys := lastValid/keyLifetime - ((firstValid - 1) / keyLifetime)
	if firstValid == 0 {
		numberOfKeys = lastValid/keyLifetime + 1 // add 1 for round zero
	}

	keys, err := KeysBuilder(numberOfKeys)
	if err != nil {
		return nil, err
	}
	tree, err := merklearray.BuildVectorCommitmentTree(&committablePublicKeyArray{keys, firstValid, keyLifetime}, crypto.HashFactory{HashType: MerkleSignatureSchemeHashFunction})
	if err != nil {
		return nil, err
	}

	return &Secrets{
		ephemeralKeys: keys,
		SignerContext: SignerContext{
			FirstValid:  firstValid,
			KeyLifetime: keyLifetime,
			Tree:        *tree,
		},
	}, nil
}

// GetVerifier can be used to store the commitment and verifier for this signer.
func (s *Secrets) GetVerifier() *Verifier {
	return s.SignerContext.GetVerifier()
}

// GetVerifier can be used to store the commitment and verifier for this signer.
func (s *SignerContext) GetVerifier() *Verifier {
	var ver Verifier
	copy(ver.Commitment[:], s.Tree.Root())
	ver.KeyLifetime = s.KeyLifetime
	return &ver
}

// FirstRoundInKeyLifetime calculates the round of the valid key for a given round by lowering to the closest KeyLiftime divisor.
func (s *Signer) FirstRoundInKeyLifetime() (uint64, error) {
	if s.KeyLifetime == 0 {
		return 0, ErrKeyLifetimeIsZero
	}

	return firstRoundInKeyLifetime(s.Round, s.KeyLifetime), nil
}

// SignBytes signs a given message. The signature is valid on a specific round
func (s *Signer) SignBytes(msg []byte) (Signature, error) {
	key := s.SigningKey
	// Possible since there may not be a StateProof key for this specific round
	if key == nil {
		return Signature{}, ErrNoStateProofKeyForRound
	}

	if err := checkMerkleSignatureSchemeParams(s.FirstValid, s.Round, s.KeyLifetime); err != nil {
		return Signature{}, err
	}

	validKeyRound, err := s.FirstRoundInKeyLifetime()
	if err != nil {
		return Signature{}, err
	}
	merkleTreeIndex := roundToIndex(s.FirstValid, validKeyRound, s.KeyLifetime)

	proof, err := s.Tree.ProveSingleLeaf(merkleTreeIndex)
	if err != nil {
		return Signature{}, err
	}

	sig, err := key.SignBytes(msg)
	if err != nil {
		return Signature{}, err
	}

	return Signature{
		Signature:             sig,
		Proof:                 *proof,
		VerifyingKey:          *s.SigningKey.GetVerifyingKey(),
		VectorCommitmentIndex: merkleTreeIndex,
	}, nil
}

// GetAllKeys returns all stateproof secrets.
// An empty array will be return if no stateproof secrets are found
func (s *Secrets) GetAllKeys() []KeyRoundPair {
	NumOfKeys := uint64(len(s.ephemeralKeys))
	keys := make([]KeyRoundPair, NumOfKeys)
	for i := uint64(0); i < NumOfKeys; i++ {
		keyRound := KeyRoundPair{
			Round: indexToRound(s.FirstValid, s.KeyLifetime, i),
			Key:   &s.ephemeralKeys[i],
		}
		keys[i] = keyRound
	}
	return keys
}

// GetKey retrieves key from memory
// the function return nil if the key does not exists
func (s *Secrets) GetKey(round uint64) *crypto.FalconSigner {
	idx := roundToIndex(s.FirstValid, round, s.KeyLifetime)
	if idx >= uint64(len(s.ephemeralKeys)) || (round%s.KeyLifetime) != 0 {
		return nil
	}

	return &s.ephemeralKeys[idx]
}

// GetSigner returns the secret keys required for the specified round as well as the rest of the required state proof immutable data
func (s *Secrets) GetSigner(round uint64) *Signer {
	return &Signer{
		SigningKey:    s.GetKey(round),
		Round:         round,
		SignerContext: s.SignerContext,
	}
}

// IsEmpty returns true if the verifier contains an empty key
func (v *Commitment) IsEmpty() bool {
	return *v == [MerkleSignatureSchemeRootSize]byte{}
}

// IsSaltVersionEqual validates that the version of the signature is matching the expected version
func (s *Signature) IsSaltVersionEqual(version byte) error {
	if !s.Signature.IsSaltVersionEqual(version) {
		return ErrSignatureSaltVersionMismatch
	}
	return nil
}

// FirstRoundInKeyLifetime calculates the round of the valid key for a given round by lowering to the closest KeyLiftime divisor.
func (v *Verifier) FirstRoundInKeyLifetime(round uint64) (uint64, error) {
	if v.KeyLifetime == 0 {
		return 0, ErrKeyLifetimeIsZero
	}

	return firstRoundInKeyLifetime(round, v.KeyLifetime), nil
}

// VerifyBytes verifies that a merklesignature sig is valid, on a specific round, under a given public key
func (v *Verifier) VerifyBytes(round uint64, msg []byte, sig Signature) error {
	validKeyRound, err := v.FirstRoundInKeyLifetime(round)
	if err != nil {
		return err
	}

	ephkey := CommittablePublicKey{
		VerifyingKey: sig.VerifyingKey,
		Round:        validKeyRound,
	}

	// verify the merkle tree verification path using the ephemeral public key, the
	// verification path and the index.
	err = merklearray.VerifyVectorCommitment(
		v.Commitment[:],
		map[uint64]crypto.Hashable{sig.VectorCommitmentIndex: &ephkey},
		sig.Proof.ToProof(),
	)
	if err != nil {
		return fmt.Errorf("%w - %v", ErrSignatureSchemeVerificationFailed, err)
	}

	// verify that the signature is valid under the ephemeral public key
	err = sig.VerifyingKey.VerifyBytes(msg, sig.Signature)
	if err != nil {
		return fmt.Errorf("%w - %v", ErrSignatureSchemeVerificationFailed, err)
	}
	return nil
}

// GetFixedLengthHashableRepresentation returns the signature as a hashable byte sequence.
// the format details can be found in the Algorand's spec.
func (s *Signature) GetFixedLengthHashableRepresentation() ([]byte, error) {
	schemeType := make([]byte, 2)
	binary.LittleEndian.PutUint16(schemeType, CryptoPrimitivesID)
	sigBytes, err := s.Signature.GetFixedLengthHashableRepresentation()
	if err != nil {
		return nil, err
	}

	verifierBytes := s.VerifyingKey.GetFixedLengthHashableRepresentation()

	binaryVectorCommitmentIndex := make([]byte, 8)
	binary.LittleEndian.PutUint64(binaryVectorCommitmentIndex, s.VectorCommitmentIndex)

	proofBytes := s.Proof.GetFixedLengthHashableRepresentation()

	merkleSignatureBytes := make([]byte, 0, len(schemeType)+len(sigBytes)+len(verifierBytes)+len(binaryVectorCommitmentIndex)+len(proofBytes))
	merkleSignatureBytes = append(merkleSignatureBytes, schemeType...)
	merkleSignatureBytes = append(merkleSignatureBytes, sigBytes...)
	merkleSignatureBytes = append(merkleSignatureBytes, verifierBytes...)
	merkleSignatureBytes = append(merkleSignatureBytes, binaryVectorCommitmentIndex...)
	merkleSignatureBytes = append(merkleSignatureBytes, proofBytes...)
	return merkleSignatureBytes, nil
}
