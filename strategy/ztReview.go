package strategy

import (
	"faya/features"
	"faya/filter"
	"faya/list"
	"fmt"
	"sort"
)


type ZtD struct {
	Code string
	Name string
	Days int
	Det float64
	Succ bool
}

func ZtViewer() {
	l := list.Get()
	zto := make([]ZtD, 0)
	for _, op := range l {
		a := list.RiKCodeReverse(op.Code)
		ztdayscount := features.GetZtDaysCount(op, a)
		if ztdayscount > 0 {
			var d ZtD
			d.Code = op.Code
			d.Name = op.Name
			d.Days = ztdayscount
			zto = append(zto ,d)
		}
	}
	fmt.Println("total zt: ", len(zto))
	sort.Slice(zto, func(i, j int) bool {
		return zto[i].Days > zto[j].Days || 
		zto[i].Days == zto[j].Days && zto[i].Code < zto[j].Code
	})
	for _, o := range zto {
		fmt.Printf("%s %s ^ %d zt\n", o.Code, o.Name, o.Days)
	}
}

func ZtReview() {
	l := list.Get()
	zto := make([]ZtD, 0)
	for _, op := range l {
		a := list.RiKCodeReverse(op.Code)
		ztdayscount := features.GetZtDaysCount(op, a)
		base := 0
		if len(a) > 1{
			base = a[1].Features["ZtDaysCount"].(int)
		}
		if base > 0 || ztdayscount > 0 {
			var d ZtD
			d.Code = op.Code
			d.Name = op.Name
			d.Days = base
			d.Det = a[0].Det
			if filter.ZtJudge(d.Code, d.Det) {
				d.Succ = true
			}
			zto = append(zto, d)
		}
	}
	fmt.Println("total zt: ", len(zto))
	sort.Slice(zto, func(i, j int) bool {
		return zto[i].Days > zto[j].Days || 
		zto[i].Days == zto[j].Days && zto[i].Det > zto[j].Det || 
		zto[i].Days == zto[j].Days && zto[i].Det == zto[j].Det && zto[i].Code < zto[j].Code

	})
	/*
	for _, o := range zto {
		fmt.Printf("%s %s ^ %d zt\n", o.Code, o.Name, o.Days)
	}
	*/
	for i := 0; i < len(zto); i = i + 1 {
		if i > 1 && zto[i].Days != zto[i - 1].Days {
			fmt.Print("\n")
		}
		o := zto[i]
		str := "失败"
		if o.Succ {
			str = "成功"
		}
		fmt.Printf("%s %s ^ %d进%d%s %f %s\n", o.Code, o.Name, o.Days, o.Days+1, str, o.Det, list.GetBkCode(o.Code))
	}
}
