package features

import (
	"faya/filter"
	"faya/list"
)

// last day is [0]
func GetZtDays(o *list.TimeObject, rklist []*list.RiKUnit) []int {
	ll := len(rklist)

	ret := make([]int, 0)
	for i := 0; i < ll; i = i + 1{
		rk := rklist[i]
		if filter.ZtJudge(o.Code, rk.Det) {
			ret = append(ret, i)
		}
	}

	return ret
}
