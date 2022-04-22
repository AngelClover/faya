package function

import (
	"faya/features"
	"faya/filter"
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
		ll, _ := list.GetRealtimeInfo(l)
		
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
	longtoulist := make([]*list.TimeObject, 0)
	for _,o := range lz {
		//test filter
		if o.Days > 1 {
			var x list.TimeObject
			x.Code = o.Code
			longtoulist = append(longtoulist, &x)
		}
	}
	//for hold list
	for _,o := range filter.Holdlist{
		var x list.TimeObject
		x.Code = o
		longtoulist = append(longtoulist, &x)
	}
	for ;mockOpeningTime || list.IsOpeningTime() ; time.Sleep(20 * time.Second){
		fmt.Println("------------------")
		fmt.Println(time.Now().In(time.FixedZone("UTC+8", +8*60*60)))

		//lr := list.GetRealtimeList()
		ll, fullsetList := list.GetRealtimeInfo(longtoulist)
		
		lastztdays := -1
		for _, o := range ll{
			showMinFeature := false
			rk := list.RiKCodeReverse(o.Code)
			features.GetRecentTurnover(rk)
			rto := rk[0].Features["RecentTurnover"].(float64)
			//rtop :=  (o.Turnover).(float64) / rto * 100.0
			var rtop float64
			ro, ok :=(o.Turnover).(float64) 
			if ok {
				rtop = ro / rto * 100
			} else {
				ro = 0
				rtop = 0
			}

			am, ok :=(o.Amount).(float64) 
			if !ok {
				am = 0
			}

			mo, ok :=(o.Money).(float64) 
			if !ok {
				mo = 0
			}


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
					fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", ro, am/1e4, mo/1e8, " | ", rtop)
			}

		}

		fullsetList = filter.LajiFilter(fullsetList)
		fullsetList = filter.STFilter(fullsetList)
		strategy.Analysis(fullsetList)
		fmt.Println(time.Now())

		if mockOpeningTime {
			break
		}

	}
}
