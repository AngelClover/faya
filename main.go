package main

import (
	"faya/list"
	"fmt"
	//sadas
)

//Faya

func main() {
	fmt.Println("vim-go")
	//view.PlotRik("300949")

  	//l := list.Get()
// 	//l = filter.Filter300(l)
// 	l = filter.RecentZtFilter(l)
// 	//view.Plot(l[0])
/*
	for _, o := range l{
		fmt.Println(o)
		list.GetBkCode(o.Code)
	}
	*/

// 	list.GetBkCode("301111")
 	list.MinCode("301111")


	//strategy.Day5Viewer()

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
