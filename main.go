package main

import (
	"faya/filter"
	"faya/list"

	//asdas
	"fmt"
	//sadas
)

//Faya

func main() {
	fmt.Println("vim-go")
	l := list.Get()
	l = filter.Filter300(l)
	a := filter.RecentZtFilter(l)
	for _, op := range a {
		fmt.Println(*op)
	}
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
