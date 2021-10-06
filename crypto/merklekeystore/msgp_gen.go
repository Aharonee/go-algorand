package merklekeystore

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/go-algorand/crypto/merklearray"
	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// CommittablePublicKey
//           |-----> (*) MarshalMsg
//           |-----> (*) CanMarshalMsg
//           |-----> (*) UnmarshalMsg
//           |-----> (*) CanUnmarshalMsg
//           |-----> (*) Msgsize
//           |-----> (*) MsgIsZero
//
// Proof
//   |-----> (*) MarshalMsg
//   |-----> (*) CanMarshalMsg
//   |-----> (*) UnmarshalMsg
//   |-----> (*) CanUnmarshalMsg
//   |-----> (*) Msgsize
//   |-----> (*) MsgIsZero
//
// Signature
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//
// Signer
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//
// Verifier
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *CommittablePublicKey) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).VerifyingKey.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Round == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "pk"
			o = append(o, 0xa2, 0x70, 0x6b)
			o = (*z).VerifyingKey.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "rnd"
			o = append(o, 0xa3, 0x72, 0x6e, 0x64)
			o = msgp.AppendUint64(o, (*z).Round)
		}
	}
	return
}

func (_ *CommittablePublicKey) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*CommittablePublicKey)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CommittablePublicKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VerifyingKey")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Round")
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
			(*z) = CommittablePublicKey{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "pk":
				bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VerifyingKey")
					return
				}
			case "rnd":
				(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Round")
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

func (_ *CommittablePublicKey) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*CommittablePublicKey)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CommittablePublicKey) Msgsize() (s int) {
	s = 1 + 3 + (*z).VerifyingKey.Msgsize() + 4 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *CommittablePublicKey) MsgIsZero() bool {
	return ((*z).VerifyingKey.MsgIsZero()) && ((*z).Round == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *Proof) MarshalMsg(b []byte) []byte {
	return ((*(merklearray.Proof))(z)).MarshalMsg(b)
}
func (_ *Proof) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Proof)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Proof) UnmarshalMsg(bts []byte) ([]byte, error) {
	return ((*(merklearray.Proof))(z)).UnmarshalMsg(bts)
}
func (_ *Proof) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Proof)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Proof) Msgsize() int {
	return ((*(merklearray.Proof))(z)).Msgsize()
}

// MsgIsZero returns whether this is a zero value
func (z *Proof) MsgIsZero() bool {
	return ((*(merklearray.Proof))(z)).MsgIsZero()
}

// MarshalMsg implements msgp.Marshaler
func (z *Signature) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 4 bits */
	if (*z).ByteSignature.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Proof.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).VerifyingKey.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "bsig"
			o = append(o, 0xa4, 0x62, 0x73, 0x69, 0x67)
			o = (*z).ByteSignature.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "prf"
			o = append(o, 0xa3, 0x70, 0x72, 0x66)
			o = (*z).Proof.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "vkey"
			o = append(o, 0xa4, 0x76, 0x6b, 0x65, 0x79)
			o = (*z).VerifyingKey.MarshalMsg(o)
		}
	}
	return
}

func (_ *Signature) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signature)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Signature) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			bts, err = (*z).ByteSignature.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "ByteSignature")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Proof.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Proof")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VerifyingKey")
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
			(*z) = Signature{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "bsig":
				bts, err = (*z).ByteSignature.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "ByteSignature")
					return
				}
			case "prf":
				bts, err = (*z).Proof.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Proof")
					return
				}
			case "vkey":
				bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VerifyingKey")
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

func (_ *Signature) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signature)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Signature) Msgsize() (s int) {
	s = 1 + 5 + (*z).ByteSignature.Msgsize() + 4 + (*z).Proof.Msgsize() + 5 + (*z).VerifyingKey.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Signature) MsgIsZero() bool {
	return ((*z).ByteSignature.MsgIsZero()) && ((*z).Proof.MsgIsZero()) && ((*z).VerifyingKey.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *Signer) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(3)
	var zb0002Mask uint8 /* 7 bits */
	if (*z).Interval == 0 {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	if (*z).FirstValid == 0 {
		zb0002Len--
		zb0002Mask |= 0x10
	}
	if (*z).Tree.MsgIsZero() {
		zb0002Len--
		zb0002Mask |= 0x40
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "iv"
			o = append(o, 0xa2, 0x69, 0x76)
			o = msgp.AppendUint64(o, (*z).Interval)
		}
		if (zb0002Mask & 0x10) == 0 { // if not empty
			// string "rnd"
			o = append(o, 0xa3, 0x72, 0x6e, 0x64)
			o = msgp.AppendUint64(o, (*z).FirstValid)
		}
		if (zb0002Mask & 0x40) == 0 { // if not empty
			// string "tree"
			o = append(o, 0xa4, 0x74, 0x72, 0x65, 0x65)
			o = (*z).Tree.MarshalMsg(o)
		}
	}
	return
}

func (_ *Signer) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signer)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Signer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0002 int
	var zb0003 bool
	zb0002, zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0002, zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 > 0 {
			zb0002--
			(*z).FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "FirstValid")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			(*z).Interval, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Interval")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = (*z).Tree.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Tree")
				return
			}
		}
		if zb0002 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0002)
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
		if zb0003 {
			(*z) = Signer{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "rnd":
				(*z).FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "FirstValid")
					return
				}
			case "iv":
				(*z).Interval, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Interval")
					return
				}
			case "tree":
				bts, err = (*z).Tree.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tree")
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

func (_ *Signer) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signer)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Signer) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 3 + msgp.Uint64Size + 5 + (*z).Tree.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Signer) MsgIsZero() bool {
	return ((*z).FirstValid == 0) && ((*z).Interval == 0) && ((*z).Tree.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *Verifier) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(2)
	var zb0002Mask uint8 /* 3 bits */
	if (*z).Root == ([KeyStoreRootSize]byte{}) {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	if (*z).HasValidRoot == false {
		zb0002Len--
		zb0002Mask |= 0x4
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "r"
			o = append(o, 0xa1, 0x72)
			o = msgp.AppendBytes(o, ((*z).Root)[:])
		}
		if (zb0002Mask & 0x4) == 0 { // if not empty
			// string "vr"
			o = append(o, 0xa2, 0x76, 0x72)
			o = msgp.AppendBool(o, (*z).HasValidRoot)
		}
	}
	return
}

func (_ *Verifier) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Verifier)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Verifier) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0002 int
	var zb0003 bool
	zb0002, zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0002, zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = msgp.ReadExactBytes(bts, ((*z).Root)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Root")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			(*z).HasValidRoot, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "HasValidRoot")
				return
			}
		}
		if zb0002 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0002)
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
		if zb0003 {
			(*z) = Verifier{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "r":
				bts, err = msgp.ReadExactBytes(bts, ((*z).Root)[:])
				if err != nil {
					err = msgp.WrapError(err, "Root")
					return
				}
			case "vr":
				(*z).HasValidRoot, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "HasValidRoot")
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

func (_ *Verifier) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Verifier)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Verifier) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize + (KeyStoreRootSize * (msgp.ByteSize)) + 3 + msgp.BoolSize
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Verifier) MsgIsZero() bool {
	return ((*z).Root == ([KeyStoreRootSize]byte{})) && ((*z).HasValidRoot == false)
}
