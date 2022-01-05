package strategy

import (
	"faya/features"
	"faya/filter"
	"faya/list"
	"fmt"
	"sort"
)


type LY struct {
	Code string
	Name string
	Days int
}
func LianXuXiaoYangXian() {
	l := list.Get()
	l = filter.LajiFilter(l)
	ret := make([]LY, 0)
	for _, o := range l {
		a := list.RiKCodeReverse(o.Code)
		yxdayscount := features.GetYXDaysCount(o, a)
		var d LY
		d.Code = o.Code
		d.Name = o.Name
		d.Days = yxdayscount
		if yxdayscount > 3 {
			ret = append(ret ,d)
		}
	}
	fmt.Println("total lxxyx: ", len(ret))
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Days > ret[j].Days || 
		ret[i].Days == ret[j].Days && ret[i].Code < ret[j].Code
	})
	for _, o := range ret {
		fmt.Printf("%s %s ^ %d lxxyx", o.Code, o.Name, o.Days)
		a := list.RiKCodeReverse(o.Code)
		for i := o.Days - 1; i >= 0; i = i - 1 {
			fmt.Printf(" %.2f(%.2f->%.2f)", a[i].Det, a[i].Open, a[i].Close)
		}
		fmt.Printf("\n")
	}
}
