package common

import "math"

type CandleArr []*float64

const _defaultReturn = 0.0 // Değer bulunamaması veya olmaması halinde 0 veya sonradan nil döndürebiliriz

/*
Sum Tüm elemanların toplamını verir
*/
func (arr *CandleArr) Sum() float64 {
	total := 0.0
	for i := 0; i < len(*arr); i++ {
		total += *(*arr)[i]
	}
	return total
}

/*
Last Son elemanı verir
*/
func (arr *CandleArr) Last() float64 {
	len := len(*arr)
	if len == 0 {
		return _defaultReturn
	}
	return *(*arr)[len-1]
}

/*
Len Eleman sayısını verir
*/
func (arr *CandleArr) Len() int {
	return len(*arr)
}

/*
Get TradingView PineScript'teki kullanım şekliyle çalışır
0 : Son elemanı getirir,
1 : bir önceki
..
*/
func (arr *CandleArr) Get(revIndex int) float64 {
	_index := len(*arr) - 1 - revIndex
	if _index < 0 {
		return _defaultReturn
	}
	return *(*arr)[_index]
}

func (arr *CandleArr) Max() (max float64) {
	max = float64(math.MinInt64)
	for i := 0; i < len(*arr); i++ {
		if *(*arr)[i] > max {
			max = *(*arr)[i]
		}
	}
	return
}

func (arr *CandleArr) Min() (min float64) {
	min = math.MaxFloat64
	for i := 0; i < len(*arr); i++ {
		if *(*arr)[i] < min {
			min = *(*arr)[i]
		}
	}
	return
}

func (arr *CandleArr) Cross(line2 float64, prevLine2 float64) float64 {
	if (*arr).Last() > line2 && (*arr).Get(1) < prevLine2 {
		return 1
	} else if (*arr).Last() < line2 && (*arr).Get(1) > prevLine2 {
		return -1
	}
	return 0
}

func (arr *CandleArr) CrossUp(line2 float64, prevLine2 float64) float64 {
	if (*arr).Last() > line2 && (*arr).Get(1) < prevLine2 {
		return 1
	}
	return 0
}

func (arr *CandleArr) CrossDown(line2 float64, prevLine2 float64) float64 {
	if (*arr).Last() < line2 && (*arr).Get(1) > prevLine2 {
		return 1
	}
	return 0
}
