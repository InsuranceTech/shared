package depth

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DepthPrice) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Price":
			z.Price, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Price")
				return
			}
		case "Sale":
			z.Sale, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Sale")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DepthPrice) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Price"
	err = en.Append(0x82, 0xa5, 0x50, 0x72, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Price)
	if err != nil {
		err = msgp.WrapError(err, "Price")
		return
	}
	// write "Sale"
	err = en.Append(0xa4, 0x53, 0x61, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Sale)
	if err != nil {
		err = msgp.WrapError(err, "Sale")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DepthPrice) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Price"
	o = append(o, 0x82, 0xa5, 0x50, 0x72, 0x69, 0x63, 0x65)
	o = msgp.AppendFloat64(o, z.Price)
	// string "Sale"
	o = append(o, 0xa4, 0x53, 0x61, 0x6c, 0x65)
	o = msgp.AppendFloat64(o, z.Sale)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DepthPrice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Price":
			z.Price, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Price")
				return
			}
		case "Sale":
			z.Sale, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sale")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DepthPrice) Msgsize() (s int) {
	s = 1 + 6 + msgp.Float64Size + 5 + msgp.Float64Size
	return
}