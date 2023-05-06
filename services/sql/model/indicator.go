package model

type Indicator struct {
	tableName       struct{}        `pg:"indicators"`
	ID              int64           `pg:"id,pk"`
	Category        []int64         `pg:"category,array"`
	Type            int16           `pg:"type_of"`
	Name            string          `pg:"name"`
	Description     string          `pg:"description"`
	FuncName        string          `pg:"func_name"`
	CanBooster      bool            `pg:"can_booster"`
	CanScanner      bool            `pg:"can_adv_filter"`
	CanAlarm        bool            `pg:"can_alarm"`
	CanBoosterTable bool            `pg:"can_booster_tbl"`
	BundleFuncName  string          `pg:"bundle_func_name"`
	BundleOutName   string          `pg:"bundle_out_name"`
	Status          bool            `pg:"status"`
	Outs            []*IndicatorOut `pg:"rel:has-many,join_fk:indicator_id"`
}

func (i *Indicator) IsBundle() bool {
	return i.BundleFuncName != ""
}

func (i *Indicator) HasOut(name string) bool {
	return i.GetOut(name) != nil
}

func (i *Indicator) GetOut(name string) *IndicatorOut {
	if i.Outs == nil {
		return nil
	}
	for _, out := range i.Outs {
		if out.Name == name {
			return out
		}
	}
	return nil
}
