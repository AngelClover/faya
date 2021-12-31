package features

import (
	"faya/filter"
	"faya/list"
	"strings"
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

func getTime(str string) string {
	ll := strings.Split(str, " ")
	ret := str
	if len(ll) > 1 {
		ret = ll[1]
	}
	return ret
}
func GetFengbanTime(o *list.TimeObject, mlist []*list.MinUnit, uplimit float64) (string,string) {
	ll := len(mlist)
	first := "none"
	for i := 0; i < ll; i = i + 1 {
		if mlist[i].Close >= uplimit * 0.999 {
			first = mlist[i].DateTime
			break
		}
	}
	last := "none"
	for i := ll - 1; i >= 0; i = i - 1 {
		if mlist[i].Close < uplimit * 0.999 {
			j := i + 1
			if j >= ll {
				j = j - 1
			}
			last = mlist[j].DateTime
			break
		}
	}
	return getTime(first),getTime(last)
}
