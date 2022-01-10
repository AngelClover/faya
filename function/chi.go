package function

import (
	"faya/features"
	"faya/list"
	"fmt"
	"time"
)

var (
	mockOpeningTime = false
)


func Chi(l []*list.TimeObject) {
	for ;mockOpeningTime || list.IsOpeningTime() ; time.Sleep(5 * time.Second){
		fmt.Println("------------------")
		//lr := list.GetRealtimeList()
		ll := list.GetRealtimeInfo(l)
		
		for _, o := range ll{
			lm := list.RealtimeMinCode(o.Code)
			rk := list.RiKCodeReverse(o.Code)
			features.GetRecentTurnover(rk)
			rto := rk[0].Features["RecentTurnover"].(float64)
			rtop :=  (o.Turnover).(float64) / rto * 100.0
			if len(lm) > 0 {
			//fmt.Println(o, "len: ", len(lm), lm[len(lm) - 1])
			m := lm[len(lm) - 1]
// 			fmt.Println(o.Code, o.Name, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
			fmt.Println(o.Code, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
		}

		}
		if mockOpeningTime {
			break
		}

	}
}
