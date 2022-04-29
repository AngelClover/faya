package function

import (
	"faya/features"
	"faya/filter"
	"faya/list"
	"fmt"
)

// var targetFeatures = ["RecentTurnover"]
var test = true

func hide(x interface{}){
	return 
}

func Backtrace(date string){
	fmt.Println("for date:", date)
	l := list.Get()
	//yesterday's ?
	for _, o := range l {
		rk := list.RiKCodeReverse(o.Code)

		if len(rk) > 0{

			lastClose := -1.0
			rkd := rk[0]
			for i,_ := range rk {
				if rk[i].Date == date {
					rkd = rk[i]
					if i + 1 < len(rk){
						lastClose = rk[i + 1].Close
					}
					break
				}
			}

			if lastClose >= 0 {
				detp :=  (rkd.High - lastClose) * 100.0 / lastClose

				//fmt.Println("(", o.Code, detp)
				if filter.ZtJudge(o.Code, detp) {
					// 		for _, f := range targetFeatures {
					// 		}
					features.GetRecentTurnover(rk)
					rto := rkd.Features["RecentTurnover"].(float64)
					amt := rkd.Features["RecentAmount"].(int)
					mny := rkd.Features["RecentMoney"].(float64)
					rtop :=  rkd.Turnover / rto * 100.0
					hide(rtop)
// 					fmt.Println(rtop)
// 					fmt.Printf("%+v\n", o)
// 					fmt.Printf("%+v rtop:%f\n", rkd, rtop)
					
					bk := list.GetBkCode(o.Code)
					status := (float64)(0.0)
					status = (rkd.Close / rkd.High - 1) * 100.0
					//fmt.Println(o.Code, o.Name, bk, status, rkd.Money, rtop, amt)
					min := list.MinCodeDate(o.Code, date)
					amt_sum := 0
					mny_sum := float64(0.0)
					if len(min) <= 0 {
						continue
					}
					for j, x := range min {
						amt_sum += x.Amount
						mny_sum += x.Money
						/*
						if o.Code == "603421" {
							fmt.Printf("%+v\n", x)
							fmt.Println(amt_sum, mny_sum)
							fmt.Println(rto, amt, mny, rtop)
						}
						*/
						if x.High == rkd.High &&
						(j == 0 || min[j - 1].High != x.High) {
// 							fmt.Printf("%+v\n", x)
							fmt.Printf("%s %s %s %s %.2f %.2f %.2f %.2f | %.2f\n", o.Code, o.Name, bk, x.DateTime, float64(amt_sum)*100.0/float64(amt), mny_sum*100.0/mny, mny_sum*100.0/rkd.Money, rkd.Money/100000000, status)
// 							break
						}
					}
				}
				/*
				if test {
					break
				}
				*/
			}
		}

	}
	
}
