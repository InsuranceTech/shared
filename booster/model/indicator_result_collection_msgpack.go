package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *IndicatorResultCollection) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Results":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Results")
				return
			}
			if cap(z.Results) >= int(zb0002) {
				z.Results = (z.Results)[:zb0002]
			} else {
				z.Results = make([]*IndicatorResult, zb0002)
			}
			for za0001 := range z.Results {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Results", za0001)
						return
					}
					z.Results[za0001] = nil
				} else {
					if z.Results[za0001] == nil {
						z.Results[za0001] = new(IndicatorResult)
					}
					err = z.Results[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Results", za0001)
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
func (z *IndicatorResultCollection) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Results"
	err = en.Append(0x81, 0xa7, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Results)))
	if err != nil {
		err = msgp.WrapError(err, "Results")
		return
	}
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Results[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Results", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IndicatorResultCollection) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Results"
	o = append(o, 0x81, 0xa7, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Results)))
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Results[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Results", za0001)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IndicatorResultCollection) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Results":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Results")
				return
			}
			if cap(z.Results) >= int(zb0002) {
				z.Results = (z.Results)[:zb0002]
			} else {
				z.Results = make([]*IndicatorResult, zb0002)
			}
			for za0001 := range z.Results {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Results[za0001] = nil
				} else {
					if z.Results[za0001] == nil {
						z.Results[za0001] = new(IndicatorResult)
					}
					bts, err = z.Results[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Results", za0001)
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
func (z *IndicatorResultCollection) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Results[za0001].Msgsize()
		}
	}
	return
}
