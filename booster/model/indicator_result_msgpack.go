package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"time"

	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *IndicatorResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Symbol")
			return
		}
		z.Symbol = nil
	} else {
		if z.Symbol == nil {
			z.Symbol = new(symbol.Symbol)
		}
		err = z.Symbol.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Symbol")
			return
		}
	}
	z.IndicatorID, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "IndicatorID")
		return
	}
	z.FuncName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "FuncName")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Values")
		return
	}
	if z.Values == nil {
		z.Values = make(map[string][]float64, zb0002)
	} else if len(z.Values) > 0 {
		for key := range z.Values {
			delete(z.Values, key)
		}
	}
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 []float64
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Values")
			return
		}
		var zb0003 uint32
		zb0003, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, "Values", za0001)
			return
		}
		if cap(za0002) >= int(zb0003) {
			za0002 = (za0002)[:zb0003]
		} else {
			za0002 = make([]float64, zb0003)
		}
		for za0003 := range za0002 {
			za0002[za0003], err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Values", za0001, za0003)
				return
			}
		}
		z.Values[za0001] = za0002
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "LastCandle")
			return
		}
		z.LastCandle = nil
	} else {
		if z.LastCandle == nil {
			z.LastCandle = new(common.Candle)
		}
		err = z.LastCandle.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "LastCandle")
			return
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "PrevCandle")
			return
		}
		z.PrevCandle = nil
	} else {
		if z.PrevCandle == nil {
			z.PrevCandle = new(common.Candle)
		}
		err = z.PrevCandle.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "PrevCandle")
			return
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "UpdateTime")
			return
		}
		z.UpdateTime = nil
	} else {
		if z.UpdateTime == nil {
			z.UpdateTime = new(time.Time)
		}
		*z.UpdateTime, err = dc.ReadTime()
		if err != nil {
			err = msgp.WrapError(err, "UpdateTime")
			return
		}
	}
	z.Signal, err = dc.ReadInt16()
	if err != nil {
		err = msgp.WrapError(err, "Signal")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *IndicatorResult) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 8
	err = en.Append(0x98)
	if err != nil {
		return
	}
	if z.Symbol == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Symbol.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Symbol")
			return
		}
	}
	err = en.WriteInt64(z.IndicatorID)
	if err != nil {
		err = msgp.WrapError(err, "IndicatorID")
		return
	}
	err = en.WriteString(z.FuncName)
	if err != nil {
		err = msgp.WrapError(err, "FuncName")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Values)))
	if err != nil {
		err = msgp.WrapError(err, "Values")
		return
	}
	for za0001, za0002 := range z.Values {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Values")
			return
		}
		err = en.WriteArrayHeader(uint32(len(za0002)))
		if err != nil {
			err = msgp.WrapError(err, "Values", za0001)
			return
		}
		for za0003 := range za0002 {
			err = en.WriteFloat64(za0002[za0003])
			if err != nil {
				err = msgp.WrapError(err, "Values", za0001, za0003)
				return
			}
		}
	}
	if z.LastCandle == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.LastCandle.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "LastCandle")
			return
		}
	}
	if z.PrevCandle == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.PrevCandle.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "PrevCandle")
			return
		}
	}
	if z.UpdateTime == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteTime(*z.UpdateTime)
		if err != nil {
			err = msgp.WrapError(err, "UpdateTime")
			return
		}
	}
	err = en.WriteInt16(z.Signal)
	if err != nil {
		err = msgp.WrapError(err, "Signal")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IndicatorResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 8
	o = append(o, 0x98)
	if z.Symbol == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Symbol.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Symbol")
			return
		}
	}
	o = msgp.AppendInt64(o, z.IndicatorID)
	o = msgp.AppendString(o, z.FuncName)
	o = msgp.AppendMapHeader(o, uint32(len(z.Values)))
	for za0001, za0002 := range z.Values {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendArrayHeader(o, uint32(len(za0002)))
		for za0003 := range za0002 {
			o = msgp.AppendFloat64(o, za0002[za0003])
		}
	}
	if z.LastCandle == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.LastCandle.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "LastCandle")
			return
		}
	}
	if z.PrevCandle == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.PrevCandle.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "PrevCandle")
			return
		}
	}
	if z.UpdateTime == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendTime(o, *z.UpdateTime)
	}
	o = msgp.AppendInt16(o, z.Signal)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IndicatorResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Symbol = nil
	} else {
		if z.Symbol == nil {
			z.Symbol = new(symbol.Symbol)
		}
		bts, err = z.Symbol.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Symbol")
			return
		}
	}
	z.IndicatorID, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IndicatorID")
		return
	}
	z.FuncName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FuncName")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Values")
		return
	}
	if z.Values == nil {
		z.Values = make(map[string][]float64, zb0002)
	} else if len(z.Values) > 0 {
		for key := range z.Values {
			delete(z.Values, key)
		}
	}
	for zb0002 > 0 {
		var za0001 string
		var za0002 []float64
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Values")
			return
		}
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Values", za0001)
			return
		}
		if cap(za0002) >= int(zb0003) {
			za0002 = (za0002)[:zb0003]
		} else {
			za0002 = make([]float64, zb0003)
		}
		for za0003 := range za0002 {
			za0002[za0003], bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Values", za0001, za0003)
				return
			}
		}
		z.Values[za0001] = za0002
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.LastCandle = nil
	} else {
		if z.LastCandle == nil {
			z.LastCandle = new(common.Candle)
		}
		bts, err = z.LastCandle.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "LastCandle")
			return
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.PrevCandle = nil
	} else {
		if z.PrevCandle == nil {
			z.PrevCandle = new(common.Candle)
		}
		bts, err = z.PrevCandle.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "PrevCandle")
			return
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.UpdateTime = nil
	} else {
		if z.UpdateTime == nil {
			z.UpdateTime = new(time.Time)
		}
		*z.UpdateTime, bts, err = msgp.ReadTimeBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "UpdateTime")
			return
		}
	}
	z.Signal, bts, err = msgp.ReadInt16Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signal")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *IndicatorResult) Msgsize() (s int) {
	s = 1
	if z.Symbol == nil {
		s += msgp.NilSize
	} else {
		s += z.Symbol.Msgsize()
	}
	s += msgp.Int64Size + msgp.StringPrefixSize + len(z.FuncName) + msgp.MapHeaderSize
	if z.Values != nil {
		for za0001, za0002 := range z.Values {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.ArrayHeaderSize + (len(za0002) * (msgp.Float64Size))
		}
	}
	if z.LastCandle == nil {
		s += msgp.NilSize
	} else {
		s += z.LastCandle.Msgsize()
	}
	if z.PrevCandle == nil {
		s += msgp.NilSize
	} else {
		s += z.PrevCandle.Msgsize()
	}
	if z.UpdateTime == nil {
		s += msgp.NilSize
	} else {
		s += msgp.TimeSize
	}
	s += msgp.Int16Size
	return
}
