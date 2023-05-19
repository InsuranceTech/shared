package scanner

type ConditionOperator int

const (
	CO_Less           ConditionOperator = 0 // <
	CO_LessOrEqual    ConditionOperator = 1 // <=
	CO_Greater        ConditionOperator = 2 // >
	CO_GreaterOrEqual ConditionOperator = 3 // >=
	CO_Range          ConditionOperator = 4 // Aralık İçi
	CO_OutRange       ConditionOperator = 5 // Aralık Dışı
	CO_Equal          ConditionOperator = 6 // ==
	CO_NotEqual       ConditionOperator = 7 // !=
	CO_CrossUp        ConditionOperator = 8 // Yukarı Kesiş
	CO_CrossDown      ConditionOperator = 9 // Aşağı Kesiş
)

// Condition Fixed Values
const (
	CO_FV_CANDLE_OPEN  = "{Candle.Open}"
	CO_FV_CANDLE_HIGH  = "{Candle.High}"
	CO_FV_CANDLE_LOW   = "{Candle.Low}"
	CO_FV_CANDLE_CLOSE = "{Candle.Close}"
	CO_FV_VALUE        = "{Value}"
)

type Condition struct {
	IndicatorID int
	FuncName    string
	Source      string
	Target      string
	Value1      float64
	Value2      float64
	Operator    ConditionOperator
}
