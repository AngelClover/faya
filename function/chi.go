package function

import (
	"faya/features"
	"faya/list"
	"faya/strategy"
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

func Show(l []strategy.ZtD){
	
	for i, o := range l {
		if i == 0 || l[i].Days != l[i-1].Days {
			fmt.Printf("\n")
		}
		fmt.Printf("%s %s %d %s\n", o.Code, o.Name, o.Days + 1, list.GetBkCode(o.Code))
	}
}

func ShowChi(lz []strategy.ZtD){
	l := make([]*list.TimeObject, 0)
	for _,o := range lz {
		//test filter
		if o.Days > 1 {
			var x list.TimeObject
			x.Code = o.Code
			l = append(l, &x)
		}
	}
	for ;mockOpeningTime || list.IsOpeningTime() ; time.Sleep(5 * time.Second){
		fmt.Println("------------------")
		//lr := list.GetRealtimeList()
		ll := list.GetRealtimeInfo(l)
		
		lastztdays := -1
		for _, o := range ll{
			showMinFeature := false
			rk := list.RiKCodeReverse(o.Code)
			features.GetRecentTurnover(rk)
			rto := rk[0].Features["RecentTurnover"].(float64)
			rtop :=  (o.Turnover).(float64) / rto * 100.0

			ztdays := func () int{
				for _, oo := range lz{
					if o.Code == oo.Code {
						return oo.Days + 1
					}
				}
				return 0
			}()
			if ztdays != lastztdays {
				fmt.Println("")
				if lastztdays == -1 {
					fmt.Println("code  name      zt   detp  | turnover  amount(10k) Money(e) | to_ratio")

				}
				lastztdays = ztdays
			}

			if showMinFeature {
				lm := list.RealtimeMinCode(o.Code)
				if len(lm) > 0 {
					//fmt.Println(o, "len: ", len(lm), lm[len(lm) - 1])
					m := lm[len(lm) - 1]
					fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
					//fmt.Println(o.Code, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
				}
			} else {
					fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", rtop)
			}

		}
		if mockOpeningTime {
			break
		}

	}
}
