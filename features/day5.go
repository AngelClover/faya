package features

import (
	"faya/list"
)

//last day is [0]
func GetDay5(rklist []*list.RiKUnit) {
	ll := len(rklist)
	for i := 0; i < ll; i = i + 1{
		rk := rklist[i]
		sum := 0.0
		if 5 + i > ll {
			break
		}
		for j := 0; j < 5 && j + i < ll; j = j + 1{
			sum += rklist[i + j].Close
		}
		rk.Features["Day5"] = sum / 5
	}
}
func GetDay5Det(rklist []*list.RiKUnit) []float64 {
	ll := len(rklist)
	ret := make([]float64, 0)
	for i := 0; i < ll && i + 5 < ll; i = i + 1{
		rk := rklist[i]
		base := rk.Features["Day5"].(float64)
		det := rk.Close - base
		detp := det / base * 100
		rk.Features["Day5Det"] = detp
		if i < 5{
			ret = append(ret, detp)
		}
	}
	return ret
}
