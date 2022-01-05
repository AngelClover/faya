package features

import "faya/list"


func GetYXDaysCount(o *list.TimeObject, rklist []*list.RiKUnit) int {
	//a := GetZtDays(o, rklist)
	ll := len(rklist)

	ret := 0
	for i := 0; i < ll; i = i + 1 {
		j := 0
		//rk := rklist[i]
		for ;i + j < ll; j = j + 1 {
			if rklist[i + j].Close < rklist[i + j].Open {
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
