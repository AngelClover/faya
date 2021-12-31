package filter

import (
	"faya/data"
	"faya/list"
	"fmt"
	"strings"
)

func ZtJudge(code string, detp float64) bool {
	if strings.Index(code, "30") == 0 && detp > 19 {
		return true
	}
	if strings.Index(code, "688") == 0 && detp > 19 {
		return true
	}
	if strings.Index(code, "8") == 0 && detp > 29 {
		return true
	}
	if strings.Index(code, "30") != 0 && strings.Index(code, "688") != 0 && detp > 9.6{
		return true
	}
	return false
}

func ZtFilter(input []*list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, obj := range input{
		per, ok := obj.DetP.(float64)
		if !ok {
			fmt.Println("not ok", obj)
			per = 0
		}
		if ZtJudge(obj.Code, per) {
			ret = append(ret, obj)
		}
	}
	return ret
}

func RecentZtFilter(input []* list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, obj := range input{
		a := data.Get(obj.Code, "rik_reverse")
		b, ok:= a.([]*list.RiKUnit)
		if !ok {
			fmt.Println("cast error", a)
		}
		if len(b) > 0 {
// 			fmt.Println(b[0])
			for j := 0; j < 3 && j < len(b); j = j + 1{
				if ZtJudge(obj.Code, b[j].Det) {
					//fmt.Println(obj.Code, obj.Name, "has zt for last", j, "days with det:", b[j].Det)
					ret = append(ret, obj)
					break
				}
			}
		}
	}
	fmt.Println("RecentZtFilter done :", len(ret))
	return ret
}
