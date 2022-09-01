package stateproof

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"sort"

	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/crypto/stateproof"
	"github.com/algorand/go-algorand/data/basics"
)

// The following msgp objects are implemented in this file:
// builder
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//
// sigFromAddr
//      |-----> (*) MarshalMsg
//      |-----> (*) CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> (*) Msgsize
//      |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *builder) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0003Len := uint32(5)
	var zb0003Mask uint8 /* 6 bits */
	if (*z).Builder == nil {
		zb0003Len--
		zb0003Mask |= 0x1
	}
	if len((*z).AddrToPos) == 0 {
		zb0003Len--
		zb0003Mask |= 0x4
	}
	if (*z).VotersHdr.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x8
	}
	if (*z).Message.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x10
	}
	if (*z).Proto.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x20
	}
	// variable map header, size zb0003Len
	o = append(o, 0x80|uint8(zb0003Len))
	if zb0003Len != 0 {
		if (zb0003Mask & 0x1) == 0 { // if not empty
			// string "Builder"
			o = append(o, 0xa7, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72)
			if (*z).Builder == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).Builder.MarshalMsg(o)
			}
		}
		if (zb0003Mask & 0x4) == 0 { // if not empty
			// string "addr"
			o = append(o, 0xa4, 0x61, 0x64, 0x64, 0x72)
			if (*z).AddrToPos == nil {
				o = msgp.AppendNil(o)
			} else {
				o = msgp.AppendMapHeader(o, uint32(len((*z).AddrToPos)))
			}
			zb0001_keys := make([]basics.Address, 0, len((*z).AddrToPos))
			for zb0001 := range (*z).AddrToPos {
				zb0001_keys = append(zb0001_keys, zb0001)
			}
			sort.Sort((zb0001_keys))
			for _, zb0001 := range zb0001_keys {
				zb0002 := (*z).AddrToPos[zb0001]
				_ = zb0002
				o = zb0001.MarshalMsg(o)
				o = msgp.AppendUint64(o, zb0002)
			}
		}
		if (zb0003Mask & 0x8) == 0 { // if not empty
			// string "hdr"
			o = append(o, 0xa3, 0x68, 0x64, 0x72)
			o = (*z).VotersHdr.MarshalMsg(o)
		}
		if (zb0003Mask & 0x10) == 0 { // if not empty
			// string "msg"
			o = append(o, 0xa3, 0x6d, 0x73, 0x67)
			o = (*z).Message.MarshalMsg(o)
		}
		if (zb0003Mask & 0x20) == 0 { // if not empty
			// string "proto"
			o = append(o, 0xa5, 0x70, 0x72, 0x6f, 0x74, 0x6f)
			o = (*z).Proto.MarshalMsg(o)
		}
	}
	return
}

func (_ *builder) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*builder)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *builder) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0003 int
	var zb0004 bool
	zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0003 > 0 {
			zb0003--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).Builder = nil
			} else {
				if (*z).Builder == nil {
					(*z).Builder = new(stateproof.Builder)
				}
				bts, err = (*z).Builder.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Builder")
					return
				}
			}
		}
		if zb0003 > 0 {
			zb0003--
			var zb0005 int
			var zb0006 bool
			zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "AddrToPos")
				return
			}
			if zb0005 > stateproof.StateProofTopVoters {
				err = msgp.ErrOverflow(uint64(zb0005), uint64(stateproof.StateProofTopVoters))
				err = msgp.WrapError(err, "struct-from-array", "AddrToPos")
				return
			}
			if zb0006 {
				(*z).AddrToPos = nil
			} else if (*z).AddrToPos == nil {
				(*z).AddrToPos = make(map[basics.Address]uint64, zb0005)
			}
			for zb0005 > 0 {
				var zb0001 basics.Address
				var zb0002 uint64
				zb0005--
				bts, err = zb0001.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AddrToPos")
					return
				}
				zb0002, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AddrToPos", zb0001)
					return
				}
				(*z).AddrToPos[zb0001] = zb0002
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).Proto.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Proto")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).VotersHdr.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VotersHdr")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).Message.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Message")
				return
			}
		}
		if zb0003 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0003)
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
		if zb0004 {
			(*z) = builder{}
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "Builder":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).Builder = nil
				} else {
					if (*z).Builder == nil {
						(*z).Builder = new(stateproof.Builder)
					}
					bts, err = (*z).Builder.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Builder")
						return
					}
				}
			case "addr":
				var zb0007 int
				var zb0008 bool
				zb0007, zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "AddrToPos")
					return
				}
				if zb0007 > stateproof.StateProofTopVoters {
					err = msgp.ErrOverflow(uint64(zb0007), uint64(stateproof.StateProofTopVoters))
					err = msgp.WrapError(err, "AddrToPos")
					return
				}
				if zb0008 {
					(*z).AddrToPos = nil
				} else if (*z).AddrToPos == nil {
					(*z).AddrToPos = make(map[basics.Address]uint64, zb0007)
				}
				for zb0007 > 0 {
					var zb0001 basics.Address
					var zb0002 uint64
					zb0007--
					bts, err = zb0001.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "AddrToPos")
						return
					}
					zb0002, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "AddrToPos", zb0001)
						return
					}
					(*z).AddrToPos[zb0001] = zb0002
				}
			case "proto":
				bts, err = (*z).Proto.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Proto")
					return
				}
			case "hdr":
				bts, err = (*z).VotersHdr.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VotersHdr")
					return
				}
			case "msg":
				bts, err = (*z).Message.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Message")
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

func (_ *builder) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*builder)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *builder) Msgsize() (s int) {
	s = 1 + 8
	if (*z).Builder == nil {
		s += msgp.NilSize
	} else {
		s += (*z).Builder.Msgsize()
	}
	s += 5 + msgp.MapHeaderSize
	if (*z).AddrToPos != nil {
		for zb0001, zb0002 := range (*z).AddrToPos {
			_ = zb0001
			_ = zb0002
			s += 0 + zb0001.Msgsize() + msgp.Uint64Size
		}
	}
	s += 6 + (*z).Proto.Msgsize() + 4 + (*z).VotersHdr.Msgsize() + 4 + (*z).Message.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *builder) MsgIsZero() bool {
	return ((*z).Builder == nil) && (len((*z).AddrToPos) == 0) && ((*z).Proto.MsgIsZero()) && ((*z).VotersHdr.MsgIsZero()) && ((*z).Message.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *sigFromAddr) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 4 bits */
	if (*z).SignerAddress.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Round.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Sig.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "a"
			o = append(o, 0xa1, 0x61)
			o = (*z).SignerAddress.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "r"
			o = append(o, 0xa1, 0x72)
			o = (*z).Round.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "s"
			o = append(o, 0xa1, 0x73)
			o = (*z).Sig.MarshalMsg(o)
		}
	}
	return
}

func (_ *sigFromAddr) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigFromAddr)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *sigFromAddr) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			bts, err = (*z).SignerAddress.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SignerAddress")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Round.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Round")
				return
			}
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
			(*z) = sigFromAddr{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "a":
				bts, err = (*z).SignerAddress.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "SignerAddress")
					return
				}
			case "r":
				bts, err = (*z).Round.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Round")
					return
				}
			case "s":
				bts, err = (*z).Sig.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sig")
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

func (_ *sigFromAddr) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigFromAddr)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *sigFromAddr) Msgsize() (s int) {
	s = 1 + 2 + (*z).SignerAddress.Msgsize() + 2 + (*z).Round.Msgsize() + 2 + (*z).Sig.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *sigFromAddr) MsgIsZero() bool {
	return ((*z).SignerAddress.MsgIsZero()) && ((*z).Round.MsgIsZero()) && ((*z).Sig.MsgIsZero())
}
