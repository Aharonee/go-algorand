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

package merklearray

import (
	"fmt"
	"github.com/algorand/go-algorand/protocol"
	"math/bits"
)

type vectorCommitmentArray struct {
	array     Array
	pathLen   uint8
	paddedLen uint64
}

func getBottomElement() []byte {
	return []byte(protocol.MerkleBottomLeaf)
}

func generateVectorCommitmentArray(innerArray Array) *vectorCommitmentArray {
	arrayLen := innerArray.Length()
	if arrayLen == 0 || arrayLen == 1 {
		return &vectorCommitmentArray{array: innerArray, pathLen: 1, paddedLen: 1}
	}

	path := uint8(bits.Len64(arrayLen - 1))
	var fullSize uint64
	// if only one bit is set then this is a power of 2 number
	// if not, we round up the number to the closest power of 2
	if bits.OnesCount64(arrayLen) == 1 {
		fullSize = arrayLen
	} else {
		fullSize = 1 << path
	}

	return &vectorCommitmentArray{array: innerArray, pathLen: path, paddedLen: fullSize}
}

func (vc *vectorCommitmentArray) Length() uint64 {
	return vc.paddedLen
}

func (vc *vectorCommitmentArray) Marshal(pos uint64) ([]byte, error) {
	lsbIndex := msbToLsbIndex(pos, vc.pathLen)
	if lsbIndex >= vc.paddedLen {
		return nil, fmt.Errorf("vectorCommitmentArray.Get(%d): out of bounds, full size %d", pos, vc.paddedLen)
	}
	// try to enforce the DS
	if lsbIndex < vc.array.Length() {
		// assert that this doesn't start with MB
		return vc.array.Marshal(lsbIndex)
	}

	return getBottomElement(), nil
}

func msbToLsbIndex(msbIndex uint64, pathLen uint8) uint64 {
	return bits.Reverse64(msbIndex) >> (64 - pathLen)
}
