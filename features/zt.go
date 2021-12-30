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
func GetZtDaysCount(o *list.TimeObject, rklist []*list.RiKUnit) int {
	//a := GetZtDays(o, rklist)
	ll := len(rklist)

	ret := 0
	for i := 0; i < ll; i = i + 1 {
		j := 0
		//rk := rklist[i]
		for ;i + j < ll; j = j + 1 {
			if filter.ZtJudge(o.Code, rklist[i + j].Det) == false {
				break
			}
		}
		rklist[i].Features["ZtDaysCount"] = j
		if i == 0 {
			ret = j
		}
	}
	return ret 
}

