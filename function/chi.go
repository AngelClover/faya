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
	holdlist := make([]*list.TimeObject, 0)
	for _,o := range filter.Holdlist{
		var x list.TimeObject
		x.Code = o
		holdlist = append(holdlist, &x)
	}
	for ;mockOpeningTime || list.IsOpeningTime() ; time.Sleep(20 * time.Second){
		fmt.Println("------------------")
		fmt.Println(time.Now().In(time.FixedZone("UTC+8", +8*60*60)))

		//lr := list.GetRealtimeList()

		showDraft := func (li []*list.TimeObject, showMinFeature bool) {

			lastztdays := -1
			for _, o := range li{
				//showMinFeature := false
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
// 				if showMinFeature {
// 					fmt.Printf("recentTO:%v RO:%v RTOP:%v Amout:%v money:%v\n", rto, ro, rtop, am, mo)
// 				}


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
// 						fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
						fmt.Printf("%v %v zt%v %v | t%v a%.2fm m%.2fe | am:%v rt%.2fp\n", o.Code, o.Name, ztdays, o.DetP, o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, m.Amount, rtop)
						//fmt.Println(o.Code, o.DetP, " | ", o.Turnover, o.Amount.(float64)/1e4, o.Money.(float64)/1e8, " | ", m.Amount, rtop)
					}
					fd := list.FengdanCode(o.Code)
					fmt.Printf("%v(%v) a%v(%v)[m%.2fm] > %v(%v) %v(%v)\n", fd.Buy2, fd.Buy2Price,  fd.Buy1, fd.Buy1Price, fd.Buy1Price * float64(fd.Buy1) *100/1e4,
					fd.Sell1, fd.Sell1Price, fd.Sell2, fd.Sell2Price)
				} else {
					fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", ro, am/1e4, mo/1e8, " | ", rtop)
				}

			}
		}
		ll, fullsetList := list.GetRealtimeInfo(longtoulist)
		showDraft(ll, false)

		fullsetList = filter.LajiFilter(fullsetList)
		fullsetList = filter.STFilter(fullsetList)
		strategy.Analysis(fullsetList)

		if len(holdlist) > 0{
			hl, _:= list.GetRealtimeInfo(holdlist)
			showDraft(hl, true)
		}

		fmt.Println(time.Now())

		if mockOpeningTime {
			break
		}

	}
}
