package function

import (
	"faya/features"
	"faya/filter"
	"faya/list"
	"fmt"
)

// var targetFeatures = ["RecentTurnover"]
var test = true
var globalDate string = ""

func hide(x interface{}){
	return 
}

type TimePoint struct {
	Code string
	Datetime string
	O *list.TimeObject
	Rik *list.RiKUnit
	Min *list.MinUnit
	RikArray []*list.RiKUnit
	MinArray []*list.MinUnit
}

func Backtrace(date string){
	fmt.Println("for date:", date)
	globalDate = date
	l := list.Get()
	//yesterday's ?
	var Good,Bad []*TimePoint

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
					var one []*TimePoint
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
							y := &TimePoint{
								Code:o.Code,
								Datetime:x.DateTime,
								O:o,
								Rik:rkd,
								RikArray: rk,
								Min:x,
								MinArray: min,
							}
							one = append(one, y)
							
// 							fmt.Printf("%+v\n", x)
							fmt.Printf("%s %s %s %s a%.2fpr m%.2fpr m%.2fpy m%.2fe | %.2f\n", o.Code, o.Name, bk, x.DateTime,
							float64(amt_sum)*100.0/float64(amt),
							mny_sum*100.0/mny,
							mny_sum*100.0/rkd.Money,
							rkd.Money/100000000,
							status)
// 							break
						}
					}
					for i,x := range one {
						if status == 0 && i == len(one) - 1 {
							Good = append(Good, x)
// 							fmt.Printf("Good:%+v\n", x)
						}else {
							Bad = append(Bad, x)
// 							fmt.Printf("Bad:%+v\n", x)
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
	fmt.Printf("Good:%d Bad:%d\n", len(Good), len(Bad))

	fmt.Printf("Good:\n")
	for _,x := range Good {
// 		fmt.Printf("%d %s %s :\n %+v\n %+v\n %+v\n", i, x.Code, x.Datetime, x.O, x.Rik, x.Min)
		extractData(x)
	}

	fmt.Printf("Bad:\n")
	for _,x := range Bad {
// 		fmt.Printf("%d %s %s :\n %+v\n %+v\n %+v\n", i, x.Code, x.Datetime, x.O, x.Rik, x.Min)
		extractData(x)
	}
}
func extractData(x *TimePoint){

	rk := x.RikArray
	rkd := x.Rik


	bk := list.GetBkCode(x.Code)

	features.GetRecentTurnover(rk)
	//lastClose := x.O.YesterdayClose
	if len(rk) > 0{
	}else {
		fmt.Printf("%+v\n", x)
		panic("len(rk) == 0")
	}

	//highestdetp :=  (rkd.High - lastClose) * 100.0 / lastClose

	rto := rkd.Features["RecentTurnover"].(float64)
	amt := rkd.Features["RecentAmount"].(int)
	mny := rkd.Features["RecentMoney"].(float64)
// 	hide(rtop)
	status := (float64)(0.0)
	status = (rkd.Close / rkd.High - 1) * 100.0

	amt_sum := 0
	mny_sum := float64(0.0)
	to_sum := 0
	if len(x.MinArray) <= 0 {
		fmt.Printf("%+v\n", x)
		panic("len(min) == 0")
	}

	for _, o := range x.MinArray{
		if o == x.Min {
			break
		}
		amt_sum += o.Amount
		mny_sum += o.Money
// 		to_sum += o.Turnover
	}
	rtop :=  float64(to_sum) / rto * 100.0
	hide(rtop)

	fmt.Printf("%s %s %s %s a%.2fpr m%.2fpr | m%.2fpt m%.2fe->m%.2fe | t%.2fpr| %.2f\n", x.Code, x.O.Name, bk, x.Min.DateTime,
	float64(amt_sum)*100.0/float64(amt),
	mny_sum*100.0/mny,
	mny_sum*100.0/rkd.Money,
	mny_sum/100000000,
	rkd.Money/100000000,
	//rtop*100,
	100,
	status)
}
