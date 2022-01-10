package main

import (
	// 	"faya/strategy"

	"faya/function"
	"fmt"
	//sadas
)

//Faya

func main() {
	fmt.Println("faya")
// 	function.Prefill()
	function.TuishiProducer()

	//view.PlotRik("300949", "2022-01-06")
// 	view.PlotRik("300949")

//   	l := list.Get()
// 	l = filter.HoldFilter(l)
// 	l = filter.Filter300(l)

// 	l = filter.RecentZtFilter(l)
// 	//view.Plot(l[0])
// 	function.Chi(l)


/*
	for _, o := range l{
		//list.GetBkCode(o.Code)
		p := list.RiKCodeReverse(o.Code)
		if len(p) <= 0 || p[0].Date != "2022-01-10" {
			fmt.Println(o)
			if len(p) > 1 {
				fmt.Println(p[0])
			}
		}
	}
 	*/

// 	list.GetBkCode("301111")
 	//list.MinCode("301111")
// 	strategy.ZtReview()
// 	strategy.Day5Viewer()
// 	strategy.LianXuXiaoYangXian()

	/*
	a := list.RiKCode("300364")
	features.GetDay5(a)
	features.GetDay5Det(a)
// 	fmt.Println(t)
	for i:= 0; i < 10 && i < len(a); i = i + 1{
		fmt.Println(a[i])
	}
	*/

}
