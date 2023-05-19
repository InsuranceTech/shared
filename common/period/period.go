//msgp:tuple Period
//go:generate msgp -o=Period_Msgpack.go -tests=false
package period

import (
	"strconv"
	"time"
)

type Period int16

const (
	NonePeriod Period = 0
	M1         Period = 1
	M3         Period = 3
	M5         Period = 5
	M10        Period = 10
	M15        Period = 15
	M20        Period = 20
	M30        Period = 30
	M45        Period = 45
	H1         Period = 60
	H2         Period = 120
	H3         Period = 180
	H4         Period = 240
	H6         Period = 360
	H8         Period = 480
	H12        Period = 720
	D1         Period = 1440
)

var AllPeriods = []Period{M1, M3, M5, M10, M15, M20, M30, M45, H1, H2, H3, H4, H6, H8, H12, D1}

// ToString : Binancedeki kullanım şekli : 1m, 5m, 1h, 1d /*
func (p Period) ToString() string {
	if p >= 1 && p <= 59 {
		return strconv.Itoa(int(p)) + "m"
	}
	if p >= 60 && p <= 720 {
		return strconv.Itoa(int(p/60)) + "h"
	}
	if p == 1440 {
		return strconv.Itoa(int(p/1440)) + "d"
	}
	return ""
}

func (p Period) Minute() int {
	return int(p)
}

func (p Period) MinuteStr() string {
	return strconv.Itoa(int(p))
}

func (p Period) IsTriggerTime(time *time.Time) bool {
	if p == M1 {
		return time.Second() == 0
	}
	if p == M3 {
		return time.Minute()%3 == 0 && time.Second() == 0
	}
	if p == M5 {
		return time.Minute()%5 == 0 && time.Second() == 0
	}
	if p == M10 {
		return time.Minute()%10 == 0 && time.Second() == 0
	}
	if p == M15 {
		return time.Minute()%15 == 0 && time.Second() == 0
	}
	if p == M20 {
		return time.Minute()%20 == 0 && time.Second() == 0
	}
	if p == M30 {
		return time.Minute()%30 == 0 && time.Second() == 0
	}
	if p == M45 {
		dayMinute := (time.Hour() * 60) + time.Minute()
		return dayMinute%45 == 0 && time.Second() == 0
	}
	if p == H1 {
		return time.Minute() == 0 && time.Second() == 0
	}
	if p == H2 {
		return time.Hour()%2 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == H3 {
		return time.Hour()%3 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == H4 {
		return time.Hour()%4 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == H6 {
		return time.Hour()%6 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == H8 {
		return time.Hour()%8 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == H12 {
		return time.Hour()%12 == 0 && time.Minute() == 0 && time.Second() == 0
	}
	if p == D1 {
		return time.Hour() == 0 && time.Minute() == 0 && time.Second() == 0
	}
	return false
}

func (p Period) GetOpenDate(date *time.Time) *time.Time {
	downRound := func(minute int, par int) int {
		for i := minute; i >= 0; i-- {
			if i%par == 0 {
				return i
			}
		}
		return 0
	}

	var t time.Time
	location := date.Location()
	if p == M1 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, location)
	}
	if p == M3 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 3), 0, 0, location)
	}
	if p == M5 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 5), 0, 0, location)
	}
	if p == M10 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 10), 0, 0, location)
	}
	if p == M15 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 15), 0, 0, location)
	}
	if p == M20 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 20), 0, 0, location)
	}
	if p == M30 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), downRound(date.Minute(), 30), 0, 0, location)
	}
	if p == M45 {
		//panic("Error Combine Candles to 45M")
		dayMinutes := (date.Hour() * 60) + date.Minute()
		dayOfFixMinute := downRound(dayMinutes, 45)
		t = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, location)
		d := time.Duration(dayOfFixMinute) * time.Minute
		t = t.Add(d)
	}
	if p == H1 {
		t = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), 0, 0, 0, location)
	}
	if p == H2 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 2), 0, 0, 0, location)
	}
	if p == H3 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 3), 0, 0, 0, location)
	}
	if p == H4 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 4), 0, 0, 0, location)
	}
	if p == H6 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 6), 0, 0, 0, location)
	}
	if p == H8 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 8), 0, 0, 0, location)
	}
	if p == H12 {
		t = time.Date(date.Year(), date.Month(), date.Day(), downRound(date.Hour(), 12), 0, 0, 0, location)
	}
	if p == D1 {
		t = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, location)
	}
	return &t
}

func (p Period) GetCloseDate(date *time.Time) *time.Time {
	duration := time.Duration(p.Minute()) * time.Minute
	s1 := time.Duration(time.Second) * -1
	t2 := p.GetOpenDate(date).Add(duration).Add(s1)
	return &t2
}

func (p Period) GetNextOpenDate(date *time.Time) *time.Time {
	duration := time.Duration(p.Minute()) * time.Minute
	t2 := p.GetOpenDate(date).Add(duration)
	return &t2
}

func (p Period) GetCloseDateForPeriod(date *time.Time, forPeriod Period) *time.Time {
	duration := (time.Duration(forPeriod.Minute()) * time.Minute) * -1
	t := p.GetNextOpenDate(date).Add(duration)
	return &t
}
