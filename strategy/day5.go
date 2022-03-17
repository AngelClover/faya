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
	Amount []int
	Turnover []float64
	RecentTO float64
	Bk string
}
func Day5Viewer() []O {
 	fmt.Println("vim-go")
	//view.PlotRik("300949")
	l := list.Get()
	//l = filter.Filter300(l)
	l = filter.LajiFilter(l)
	l = filter.RecentZtFilter(l)
	//view.Plot(l[0])

	oo := make([]O, 0)
	for _, op := range l {
		//fmt.Println(*op)
		a := list.RiKCodeReverse(op.Code)
		features.GetDay5(a)
		features.GetRecentTurnover(a)
		depa := features.GetDay5Det(a)
		ret := features.GetZtDays(op, a)
		bk := features.GetBk(op)
		if len(depa) > 2{
			var o O
			o.Code = op.Code
			o.Name = op.Name
			o.Ztdays = ret
			o.Detp = depa
			o.Bk = bk
			o.Amount = []int{a[0].Amount, a[1].Amount, a[2].Amount}
			o.Turnover = []float64{a[0].Turnover, a[1].Turnover, a[2].Turnover}
			o.RecentTO = a[0].Features["RecentTurnover"].(float64)

			oo = append(oo, o)
			//fmt.Printf("%s %s %v => %.2f %.2f %.2f\n", op.Code, op.Name, ret, depa[2], depa[1], depa[0])
		} else {
			fmt.Println(op.Code, op.Name, ret, depa, bk)
		}
	}
	sort.Slice(oo, func(i, j int) bool {
		return oo[i].Detp[0] > oo[j].Detp[0]
	})
	Day5Print(oo)
	/*
	a := list.RiKCode("300364")
	features.GetDay5(a)
	features.GetDay5Det(a)
// 	fmt.Println(t)
	for i:= 0; i < 10 && i < len(a); i = i + 1{
		fmt.Println(a[i])
	}
	*/
	return oo

}

func Day5Print(oo [] O){
	for _, o := range oo{
		flag := ""
		if o.Turnover[0] < o.Turnover[2]*0.7 {
			flag = flag + "s "
		}
		fmt.Printf("%s %s %s %v => %.2f %.2f %.2f  ||%.2f|| %.2f %.2f %.2f %s\n", o.Code, o.Name, o.Bk, o.Ztdays, o.Detp[2], o.Detp[1], o.Detp[0], o.RecentTO,  o.Turnover[2], o.Turnover[1], o.Turnover[0], flag)
	}
}

func Day5DowngradeViewer() []O {
	l := Day5Viewer()
	fmt.Println("------------Day5 Downgrade-------")
	ret := make([]O, 0)
	for _, o:= range l{
		if len(o.Ztdays) > 0 && o.Ztdays[0] == 2 {
			if (o.Detp[2] > o.Detp[1] && o.Detp[1] > o.Detp[0]){
				ret = append(ret, o)
			}
		}

	}
	Day5Print(ret)

	return ret
}
