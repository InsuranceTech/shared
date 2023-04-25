package common

type SourceType int

const (
	ST_OPEN   = 0
	ST_HIGH   = 1
	ST_LOW    = 2
	ST_CLOSE  = 3
	ST_VOLUME = 4
	ST_HL2    = 5
	ST_HLC3   = 6
	ST_HLCC4  = 7
	ST_OHLCV5 = 8
	ST_CUSTOM = 9
	ST_OHLC4  = 10
	//test for v1.1.3
	ST__V13
)
