package stateproof

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"sort"

	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/crypto/merklearray"
	"github.com/algorand/go-algorand/data/basics"
)

// The following msgp objects are implemented in this file:
// Builder
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//
// MessageHash
//      |-----> (*) MarshalMsg
//      |-----> (*) CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> (*) Msgsize
//      |-----> (*) MsgIsZero
//
// Reveal
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//
// StateProof
//      |-----> (*) MarshalMsg
//      |-----> (*) CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> (*) Msgsize
//      |-----> (*) MsgIsZero
//
// sigslotCommit
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *Builder) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0004Len := uint32(7)
	var zb0004Mask uint16 /* 11 bits */
	if (*z).Data == (MessageHash{}) {
		zb0004Len--
		zb0004Mask |= 0x4
	}
	if (*z).LnProvenWeight == 0 {
		zb0004Len--
		zb0004Mask |= 0x8
	}
	if len((*z).Participants) == 0 {
		zb0004Len--
		zb0004Mask |= 0x10
	}
	if (*z).Parttree == nil {
		zb0004Len--
		zb0004Mask |= 0x20
	}
	if (*z).ProvenWeight == 0 {
		zb0004Len--
		zb0004Mask |= 0x40
	}
	if (*z).Round == 0 {
		zb0004Len--
		zb0004Mask |= 0x80
	}
	if (*z).StrengthTarget == 0 {
		zb0004Len--
		zb0004Mask |= 0x400
	}
	// variable map header, size zb0004Len
	o = append(o, 0x80|uint8(zb0004Len))
	if zb0004Len != 0 {
		if (zb0004Mask & 0x4) == 0 { // if not empty
			// string "data"
			o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
			o = msgp.AppendBytes(o, ((*z).Data)[:])
		}
		if (zb0004Mask & 0x8) == 0 { // if not empty
			// string "lnprv"
			o = append(o, 0xa5, 0x6c, 0x6e, 0x70, 0x72, 0x76)
			o = msgp.AppendUint64(o, (*z).LnProvenWeight)
		}
		if (zb0004Mask & 0x10) == 0 { // if not empty
			// string "parts"
			o = append(o, 0xa5, 0x70, 0x61, 0x72, 0x74, 0x73)
			if (*z).Participants == nil {
				o = msgp.AppendNil(o)
			} else {
				o = msgp.AppendArrayHeader(o, uint32(len((*z).Participants)))
			}
			for zb0003 := range (*z).Participants {
				o = (*z).Participants[zb0003].MarshalMsg(o)
			}
		}
		if (zb0004Mask & 0x20) == 0 { // if not empty
			// string "parttree"
			o = append(o, 0xa8, 0x70, 0x61, 0x72, 0x74, 0x74, 0x72, 0x65, 0x65)
			if (*z).Parttree == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).Parttree.MarshalMsg(o)
			}
		}
		if (zb0004Mask & 0x40) == 0 { // if not empty
			// string "prv"
			o = append(o, 0xa3, 0x70, 0x72, 0x76)
			o = msgp.AppendUint64(o, (*z).ProvenWeight)
		}
		if (zb0004Mask & 0x80) == 0 { // if not empty
			// string "rnd"
			o = append(o, 0xa3, 0x72, 0x6e, 0x64)
			o = msgp.AppendUint64(o, (*z).Round)
		}
		if (zb0004Mask & 0x400) == 0 { // if not empty
			// string "str"
			o = append(o, 0xa3, 0x73, 0x74, 0x72)
			o = msgp.AppendUint64(o, (*z).StrengthTarget)
		}
	}
	return
}

func (_ *Builder) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Builder)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Builder) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0004 int
	var zb0005 bool
	zb0004, zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0004, zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0004 > 0 {
			zb0004--
			bts, err = msgp.ReadExactBytes(bts, ((*z).Data)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Data")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Round")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			var zb0006 int
			var zb0007 bool
			zb0006, zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Participants")
				return
			}
			if zb0006 > StateProofTopVoters {
				err = msgp.ErrOverflow(uint64(zb0006), uint64(StateProofTopVoters))
				err = msgp.WrapError(err, "struct-from-array", "Participants")
				return
			}
			if zb0007 {
				(*z).Participants = nil
			} else if (*z).Participants != nil && cap((*z).Participants) >= zb0006 {
				(*z).Participants = ((*z).Participants)[:zb0006]
			} else {
				(*z).Participants = make([]basics.Participant, zb0006)
			}
			for zb0003 := range (*z).Participants {
				bts, err = (*z).Participants[zb0003].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Participants", zb0003)
					return
				}
			}
		}
		if zb0004 > 0 {
			zb0004--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).Parttree = nil
			} else {
				if (*z).Parttree == nil {
					(*z).Parttree = new(merklearray.Tree)
				}
				bts, err = (*z).Parttree.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Parttree")
					return
				}
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).LnProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "LnProvenWeight")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).ProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "ProvenWeight")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).StrengthTarget, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "StrengthTarget")
				return
			}
		}
		if zb0004 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0004)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0005 {
			(*z) = Builder{}
		}
		for zb0004 > 0 {
			zb0004--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "data":
				bts, err = msgp.ReadExactBytes(bts, ((*z).Data)[:])
				if err != nil {
					err = msgp.WrapError(err, "Data")
					return
				}
			case "rnd":
				(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Round")
					return
				}
			case "parts":
				var zb0008 int
				var zb0009 bool
				zb0008, zb0009, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Participants")
					return
				}
				if zb0008 > StateProofTopVoters {
					err = msgp.ErrOverflow(uint64(zb0008), uint64(StateProofTopVoters))
					err = msgp.WrapError(err, "Participants")
					return
				}
				if zb0009 {
					(*z).Participants = nil
				} else if (*z).Participants != nil && cap((*z).Participants) >= zb0008 {
					(*z).Participants = ((*z).Participants)[:zb0008]
				} else {
					(*z).Participants = make([]basics.Participant, zb0008)
				}
				for zb0003 := range (*z).Participants {
					bts, err = (*z).Participants[zb0003].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Participants", zb0003)
						return
					}
				}
			case "parttree":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).Parttree = nil
				} else {
					if (*z).Parttree == nil {
						(*z).Parttree = new(merklearray.Tree)
					}
					bts, err = (*z).Parttree.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Parttree")
						return
					}
				}
			case "lnprv":
				(*z).LnProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LnProvenWeight")
					return
				}
			case "prv":
				(*z).ProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "ProvenWeight")
					return
				}
			case "str":
				(*z).StrengthTarget, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "StrengthTarget")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *Builder) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Builder)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Builder) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 4 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for zb0003 := range (*z).Participants {
		s += (*z).Participants[zb0003].Msgsize()
	}
	s += 9
	if (*z).Parttree == nil {
		s += msgp.NilSize
	} else {
		s += (*z).Parttree.Msgsize()
	}
	s += 6 + msgp.Uint64Size + 4 + msgp.Uint64Size + 4 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Builder) MsgIsZero() bool {
	return ((*z).Data == (MessageHash{})) && ((*z).Round == 0) && (len((*z).Participants) == 0) && ((*z).Parttree == nil) && ((*z).LnProvenWeight == 0) && ((*z).ProvenWeight == 0) && ((*z).StrengthTarget == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageHash) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *MessageHash) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*MessageHash)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageHash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (_ *MessageHash) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*MessageHash)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageHash) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (32 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *MessageHash) MsgIsZero() bool {
	return (*z) == (MessageHash{})
}

// MarshalMsg implements msgp.Marshaler
func (z *Reveal) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Part.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if ((*z).SigSlot.Sig.MsgIsZero()) && ((*z).SigSlot.L == 0) {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "p"
			o = append(o, 0xa1, 0x70)
			o = (*z).Part.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "s"
			o = append(o, 0xa1, 0x73)
			// omitempty: check for empty values
			zb0002Len := uint32(2)
			var zb0002Mask uint8 /* 3 bits */
			if (*z).SigSlot.L == 0 {
				zb0002Len--
				zb0002Mask |= 0x2
			}
			if (*z).SigSlot.Sig.MsgIsZero() {
				zb0002Len--
				zb0002Mask |= 0x4
			}
			// variable map header, size zb0002Len
			o = append(o, 0x80|uint8(zb0002Len))
			if (zb0002Mask & 0x2) == 0 { // if not empty
				// string "l"
				o = append(o, 0xa1, 0x6c)
				o = msgp.AppendUint64(o, (*z).SigSlot.L)
			}
			if (zb0002Mask & 0x4) == 0 { // if not empty
				// string "s"
				o = append(o, 0xa1, 0x73)
				o = (*z).SigSlot.Sig.MarshalMsg(o)
			}
		}
	}
	return
}

func (_ *Reveal) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Reveal)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Reveal) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			var zb0003 int
			var zb0004 bool
			zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "SigSlot")
					return
				}
				if zb0003 > 0 {
					zb0003--
					bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array", "Sig")
						return
					}
				}
				if zb0003 > 0 {
					zb0003--
					(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array", "L")
						return
					}
				}
				if zb0003 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0003)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "SigSlot")
					return
				}
				if zb0004 {
					(*z).SigSlot = sigslotCommit{}
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot")
						return
					}
					switch string(field) {
					case "s":
						bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot", "Sig")
							return
						}
					case "l":
						(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot", "L")
							return
						}
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot")
							return
						}
					}
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Part.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Part")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = Reveal{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "s":
				var zb0005 int
				var zb0006 bool
				zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0005, zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "SigSlot")
						return
					}
					if zb0005 > 0 {
						zb0005--
						bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array", "Sig")
							return
						}
					}
					if zb0005 > 0 {
						zb0005--
						(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array", "L")
							return
						}
					}
					if zb0005 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0005)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "SigSlot")
						return
					}
					if zb0006 {
						(*z).SigSlot = sigslotCommit{}
					}
					for zb0005 > 0 {
						zb0005--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot")
							return
						}
						switch string(field) {
						case "s":
							bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "SigSlot", "Sig")
								return
							}
						case "l":
							(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "SigSlot", "L")
								return
							}
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "SigSlot")
								return
							}
						}
					}
				}
			case "p":
				bts, err = (*z).Part.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Part")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *Reveal) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Reveal)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Reveal) Msgsize() (s int) {
	s = 1 + 2 + 1 + 2 + (*z).SigSlot.Sig.Msgsize() + 2 + msgp.Uint64Size + 2 + (*z).Part.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Reveal) MsgIsZero() bool {
	return (((*z).SigSlot.Sig.MsgIsZero()) && ((*z).SigSlot.L == 0)) && ((*z).Part.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *StateProof) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0004Len := uint32(7)
	var zb0004Mask uint8 /* 8 bits */
	if (*z).PartProofs.MsgIsZero() {
		zb0004Len--
		zb0004Mask |= 0x1
	}
	if (*z).SigProofs.MsgIsZero() {
		zb0004Len--
		zb0004Mask |= 0x2
	}
	if (*z).SigCommit.MsgIsZero() {
		zb0004Len--
		zb0004Mask |= 0x8
	}
	if len((*z).PositionsToReveal) == 0 {
		zb0004Len--
		zb0004Mask |= 0x10
	}
	if len((*z).Reveals) == 0 {
		zb0004Len--
		zb0004Mask |= 0x20
	}
	if (*z).MerkleSignatureSaltVersion == 0 {
		zb0004Len--
		zb0004Mask |= 0x40
	}
	if (*z).SignedWeight == 0 {
		zb0004Len--
		zb0004Mask |= 0x80
	}
	// variable map header, size zb0004Len
	o = append(o, 0x80|uint8(zb0004Len))
	if zb0004Len != 0 {
		if (zb0004Mask & 0x1) == 0 { // if not empty
			// string "P"
			o = append(o, 0xa1, 0x50)
			o = (*z).PartProofs.MarshalMsg(o)
		}
		if (zb0004Mask & 0x2) == 0 { // if not empty
			// string "S"
			o = append(o, 0xa1, 0x53)
			o = (*z).SigProofs.MarshalMsg(o)
		}
		if (zb0004Mask & 0x8) == 0 { // if not empty
			// string "c"
			o = append(o, 0xa1, 0x63)
			o = (*z).SigCommit.MarshalMsg(o)
		}
		if (zb0004Mask & 0x10) == 0 { // if not empty
			// string "pr"
			o = append(o, 0xa2, 0x70, 0x72)
			if (*z).PositionsToReveal == nil {
				o = msgp.AppendNil(o)
			} else {
				o = msgp.AppendArrayHeader(o, uint32(len((*z).PositionsToReveal)))
			}
			for zb0003 := range (*z).PositionsToReveal {
				o = msgp.AppendUint64(o, (*z).PositionsToReveal[zb0003])
			}
		}
		if (zb0004Mask & 0x20) == 0 { // if not empty
			// string "r"
			o = append(o, 0xa1, 0x72)
			if (*z).Reveals == nil {
				o = msgp.AppendNil(o)
			} else {
				o = msgp.AppendMapHeader(o, uint32(len((*z).Reveals)))
			}
			zb0001_keys := make([]uint64, 0, len((*z).Reveals))
			for zb0001 := range (*z).Reveals {
				zb0001_keys = append(zb0001_keys, zb0001)
			}
			sort.Sort(SortUint64(zb0001_keys))
			for _, zb0001 := range zb0001_keys {
				zb0002 := (*z).Reveals[zb0001]
				_ = zb0002
				o = msgp.AppendUint64(o, zb0001)
				o = zb0002.MarshalMsg(o)
			}
		}
		if (zb0004Mask & 0x40) == 0 { // if not empty
			// string "v"
			o = append(o, 0xa1, 0x76)
			o = msgp.AppendByte(o, (*z).MerkleSignatureSaltVersion)
		}
		if (zb0004Mask & 0x80) == 0 { // if not empty
			// string "w"
			o = append(o, 0xa1, 0x77)
			o = msgp.AppendUint64(o, (*z).SignedWeight)
		}
	}
	return
}

func (_ *StateProof) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProof)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateProof) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0004 int
	var zb0005 bool
	zb0004, zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0004, zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0004 > 0 {
			zb0004--
			bts, err = (*z).SigCommit.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SigCommit")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SignedWeight")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			bts, err = (*z).SigProofs.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SigProofs")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			bts, err = (*z).PartProofs.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PartProofs")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			(*z).MerkleSignatureSaltVersion, bts, err = msgp.ReadByteBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "MerkleSignatureSaltVersion")
				return
			}
		}
		if zb0004 > 0 {
			zb0004--
			var zb0006 int
			var zb0007 bool
			zb0006, zb0007, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Reveals")
				return
			}
			if zb0006 > MaxReveals {
				err = msgp.ErrOverflow(uint64(zb0006), uint64(MaxReveals))
				err = msgp.WrapError(err, "struct-from-array", "Reveals")
				return
			}
			if zb0007 {
				(*z).Reveals = nil
			} else if (*z).Reveals == nil {
				(*z).Reveals = make(map[uint64]Reveal, zb0006)
			}
			for zb0006 > 0 {
				var zb0001 uint64
				var zb0002 Reveal
				zb0006--
				zb0001, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Reveals")
					return
				}
				bts, err = zb0002.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Reveals", zb0001)
					return
				}
				(*z).Reveals[zb0001] = zb0002
			}
		}
		if zb0004 > 0 {
			zb0004--
			var zb0008 int
			var zb0009 bool
			zb0008, zb0009, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PositionsToReveal")
				return
			}
			if zb0008 > MaxReveals {
				err = msgp.ErrOverflow(uint64(zb0008), uint64(MaxReveals))
				err = msgp.WrapError(err, "struct-from-array", "PositionsToReveal")
				return
			}
			if zb0009 {
				(*z).PositionsToReveal = nil
			} else if (*z).PositionsToReveal != nil && cap((*z).PositionsToReveal) >= zb0008 {
				(*z).PositionsToReveal = ((*z).PositionsToReveal)[:zb0008]
			} else {
				(*z).PositionsToReveal = make([]uint64, zb0008)
			}
			for zb0003 := range (*z).PositionsToReveal {
				(*z).PositionsToReveal[zb0003], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "PositionsToReveal", zb0003)
					return
				}
			}
		}
		if zb0004 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0004)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0005 {
			(*z) = StateProof{}
		}
		for zb0004 > 0 {
			zb0004--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "c":
				bts, err = (*z).SigCommit.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "SigCommit")
					return
				}
			case "w":
				(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "SignedWeight")
					return
				}
			case "S":
				bts, err = (*z).SigProofs.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "SigProofs")
					return
				}
			case "P":
				bts, err = (*z).PartProofs.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "PartProofs")
					return
				}
			case "v":
				(*z).MerkleSignatureSaltVersion, bts, err = msgp.ReadByteBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "MerkleSignatureSaltVersion")
					return
				}
			case "r":
				var zb0010 int
				var zb0011 bool
				zb0010, zb0011, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Reveals")
					return
				}
				if zb0010 > MaxReveals {
					err = msgp.ErrOverflow(uint64(zb0010), uint64(MaxReveals))
					err = msgp.WrapError(err, "Reveals")
					return
				}
				if zb0011 {
					(*z).Reveals = nil
				} else if (*z).Reveals == nil {
					(*z).Reveals = make(map[uint64]Reveal, zb0010)
				}
				for zb0010 > 0 {
					var zb0001 uint64
					var zb0002 Reveal
					zb0010--
					zb0001, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Reveals")
						return
					}
					bts, err = zb0002.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Reveals", zb0001)
						return
					}
					(*z).Reveals[zb0001] = zb0002
				}
			case "pr":
				var zb0012 int
				var zb0013 bool
				zb0012, zb0013, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PositionsToReveal")
					return
				}
				if zb0012 > MaxReveals {
					err = msgp.ErrOverflow(uint64(zb0012), uint64(MaxReveals))
					err = msgp.WrapError(err, "PositionsToReveal")
					return
				}
				if zb0013 {
					(*z).PositionsToReveal = nil
				} else if (*z).PositionsToReveal != nil && cap((*z).PositionsToReveal) >= zb0012 {
					(*z).PositionsToReveal = ((*z).PositionsToReveal)[:zb0012]
				} else {
					(*z).PositionsToReveal = make([]uint64, zb0012)
				}
				for zb0003 := range (*z).PositionsToReveal {
					(*z).PositionsToReveal[zb0003], bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "PositionsToReveal", zb0003)
						return
					}
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *StateProof) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProof)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StateProof) Msgsize() (s int) {
	s = 1 + 2 + (*z).SigCommit.Msgsize() + 2 + msgp.Uint64Size + 2 + (*z).SigProofs.Msgsize() + 2 + (*z).PartProofs.Msgsize() + 2 + msgp.ByteSize + 2 + msgp.MapHeaderSize
	if (*z).Reveals != nil {
		for zb0001, zb0002 := range (*z).Reveals {
			_ = zb0001
			_ = zb0002
			s += 0 + msgp.Uint64Size + zb0002.Msgsize()
		}
	}
	s += 3 + msgp.ArrayHeaderSize + (len((*z).PositionsToReveal) * (msgp.Uint64Size))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *StateProof) MsgIsZero() bool {
	return ((*z).SigCommit.MsgIsZero()) && ((*z).SignedWeight == 0) && ((*z).SigProofs.MsgIsZero()) && ((*z).PartProofs.MsgIsZero()) && ((*z).MerkleSignatureSaltVersion == 0) && (len((*z).Reveals) == 0) && (len((*z).PositionsToReveal) == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *sigslotCommit) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).L == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Sig.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "l"
			o = append(o, 0xa1, 0x6c)
			o = msgp.AppendUint64(o, (*z).L)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "s"
			o = append(o, 0xa1, 0x73)
			o = (*z).Sig.MarshalMsg(o)
		}
	}
	return
}

func (_ *sigslotCommit) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigslotCommit)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *sigslotCommit) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Sig.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Sig")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).L, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "L")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = sigslotCommit{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "s":
				bts, err = (*z).Sig.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sig")
					return
				}
			case "l":
				(*z).L, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "L")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *sigslotCommit) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigslotCommit)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *sigslotCommit) Msgsize() (s int) {
	s = 1 + 2 + (*z).Sig.Msgsize() + 2 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *sigslotCommit) MsgIsZero() bool {
	return ((*z).Sig.MsgIsZero()) && ((*z).L == 0)
}
