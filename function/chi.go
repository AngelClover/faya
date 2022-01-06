package function

import (
	"faya/list"
	"fmt"
	"time"
)



func Chi(l []*list.TimeObject) {
	for ;list.IsOpeningTime() ; time.Sleep(5 * time.Second){
		fmt.Println("------------------")
		//lr := list.GetRealtimeList()
		ll := list.GetRealtimeInfo(l)
		
		for _, o := range ll{
			lm := list.RealtimeMinCode(o.Code)
			if len(lm) > 0 {
			//fmt.Println(o, "len: ", len(lm), lm[len(lm) - 1])
			m := lm[len(lm) - 1]
			fmt.Println(o.Code, o.Name, o.DetP, " | ", o.Turnover, o.Amount, o.Money, " | ", m.Amount)
		}

		}

	}
}
