package features

import (
	"faya/list"
)

func GetRecentTurnover(rklist []*list.RiKUnit){
	ll := len(rklist)
	for i := 0; i < ll; i = i + 1{
		rk := rklist[i]
		mto := 0.0
		for j := 0; j < 10 && j + i < ll; j = j + 1{
			if mto < rklist[i + j].Turnover {
				mto = rklist[i + j].Turnover
			}
		}
		rk.Features["RecentTurnover"] = mto
	}
}
