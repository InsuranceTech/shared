package depth

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DepthData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Symbol":
			z.Symbol, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Symbol")
				return
			}
		case "Bids":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Bids")
				return
			}
			if cap(z.Bids) >= int(zb0002) {
				z.Bids = (z.Bids)[:zb0002]
			} else {
				z.Bids = make([]*DepthPrice, zb0002)
			}
			for za0001 := range z.Bids {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Bids", za0001)
						return
					}
					z.Bids[za0001] = nil
				} else {
					if z.Bids[za0001] == nil {
						z.Bids[za0001] = new(DepthPrice)
					}
					err = z.Bids[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Bids", za0001)
						return
					}
				}
			}
		case "Asks":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Asks")
				return
			}
			if cap(z.Asks) >= int(zb0003) {
				z.Asks = (z.Asks)[:zb0003]
			} else {
				z.Asks = make([]*DepthPrice, zb0003)
			}
			for za0002 := range z.Asks {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Asks", za0002)
						return
					}
					z.Asks[za0002] = nil
				} else {
					if z.Asks[za0002] == nil {
						z.Asks[za0002] = new(DepthPrice)
					}
					err = z.Asks[za0002].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Asks", za0002)
						return
					}
				}
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
func (z *DepthData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Symbol"
	err = en.Append(0x83, 0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Symbol)
	if err != nil {
		err = msgp.WrapError(err, "Symbol")
		return
	}
	// write "Bids"
	err = en.Append(0xa4, 0x42, 0x69, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Bids)))
	if err != nil {
		err = msgp.WrapError(err, "Bids")
		return
	}
	for za0001 := range z.Bids {
		if z.Bids[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bids[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Bids", za0001)
				return
			}
		}
	}
	// write "Asks"
	err = en.Append(0xa4, 0x41, 0x73, 0x6b, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Asks)))
	if err != nil {
		err = msgp.WrapError(err, "Asks")
		return
	}
	for za0002 := range z.Asks {
		if z.Asks[za0002] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Asks[za0002].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Asks", za0002)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DepthData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Symbol"
	o = append(o, 0x83, 0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
	o = msgp.AppendString(o, z.Symbol)
	// string "Bids"
	o = append(o, 0xa4, 0x42, 0x69, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bids)))
	for za0001 := range z.Bids {
		if z.Bids[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bids[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Bids", za0001)
				return
			}
		}
	}
	// string "Asks"
	o = append(o, 0xa4, 0x41, 0x73, 0x6b, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Asks)))
	for za0002 := range z.Asks {
		if z.Asks[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Asks[za0002].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Asks", za0002)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DepthData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Symbol":
			z.Symbol, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Symbol")
				return
			}
		case "Bids":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bids")
				return
			}
			if cap(z.Bids) >= int(zb0002) {
				z.Bids = (z.Bids)[:zb0002]
			} else {
				z.Bids = make([]*DepthPrice, zb0002)
			}
			for za0001 := range z.Bids {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bids[za0001] = nil
				} else {
					if z.Bids[za0001] == nil {
						z.Bids[za0001] = new(DepthPrice)
					}
					bts, err = z.Bids[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Bids", za0001)
						return
					}
				}
			}
		case "Asks":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Asks")
				return
			}
			if cap(z.Asks) >= int(zb0003) {
				z.Asks = (z.Asks)[:zb0003]
			} else {
				z.Asks = make([]*DepthPrice, zb0003)
			}
			for za0002 := range z.Asks {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Asks[za0002] = nil
				} else {
					if z.Asks[za0002] == nil {
						z.Asks[za0002] = new(DepthPrice)
					}
					bts, err = z.Asks[za0002].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Asks", za0002)
						return
					}
				}
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
func (z *DepthData) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Symbol) + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Bids {
		if z.Bids[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bids[za0001].Msgsize()
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Asks {
		if z.Asks[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.Asks[za0002].Msgsize()
		}
	}
	return
}
