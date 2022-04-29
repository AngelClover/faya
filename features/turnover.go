package features

import (
	"faya/list"
)

func GetRecentTurnover(rklist []*list.RiKUnit){
	ll := len(rklist)
	for i := 0; i < ll; i = i + 1{
		rk := rklist[i]
		mto := 0.0
		amt := 0
		mny := 0.0
		for j := 1; j < 10 && j + i < ll; j = j + 1{
			if mto < rklist[i + j].Turnover {
				mto = rklist[i + j].Turnover
			}
			if amt < rklist[i + j].Amount {
				amt = rklist[i + j].Amount
			}
			if mny < rklist[i + j].Money {
				mny = rklist[i + j].Money
			}
		}
		rk.Features["RecentTurnover"] = mto
		rk.Features["RecentAmount"] = amt
		rk.Features["RecentMoney"] = mny
	}
}
