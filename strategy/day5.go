package strategy

import (
	"faya/features"
	"faya/filter"
	"faya/list"
	"fmt"
	"sort"
)

type O struct {
	Code string
	Name string
	Ztdays []int
	Detp []float64
}
func Day5Viewer() {
	fmt.Println("vim-go")
	//view.PlotRik("300949")
	l := list.Get()
	//l = filter.Filter300(l)
	l = filter.RecentZtFilter(l)
	//view.Plot(l[0])

	oo := make([]O, 0)
	for _, op := range l {
		//fmt.Println(*op)
		a := list.RiKCodeReverse(op.Code)
		features.GetDay5(a)
		depa := features.GetDay5Det(a)
		ret := features.GetZtDays(op, a)
		if len(depa) > 2{
			var o O
			o.Code = op.Code
			o.Name = op.Name
			o.Ztdays = ret
			o.Detp = depa
			oo = append(oo, o)
			//fmt.Printf("%s %s %v => %.2f %.2f %.2f\n", op.Code, op.Name, ret, depa[2], depa[1], depa[0])
		} else {
			fmt.Println(op.Code, op.Name, ret, depa)
		}
	}
	sort.Slice(oo, func(i, j int) bool {
		return oo[i].Detp[0] > oo[j].Detp[0]
	})
	for _, o := range oo{
		fmt.Printf("%s %s %v => %.2f %.2f %.2f\n", o.Code, o.Name, o.Ztdays, o.Detp[2], o.Detp[1], o.Detp[0])
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
