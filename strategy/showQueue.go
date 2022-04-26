package strategy

import (
	"faya/list"
	"fmt"
	"sort"
)




type BkStatUnit struct {
	BkName string
	Count int
	List []*list.TimeObject
}


func showQueue(l []*list.TimeObject, showDet bool) {
	fmt.Println("showQ(")
	var a []*BkStatUnit
	for _,o := range l{
		bk := list.GetBk(o)
		//fmt.Printf("%s %s |", o.Name, bk)
		var b *BkStatUnit = nil
		for _,x := range a {
			if x.BkName == bk {
				b = x
				break
			}
		}
		if b == nil {
			b = &BkStatUnit{
				BkName: bk,
				Count: 0,
			}
			a = append(a, b)
		}
		b.Count += 1
		b.List = append(b.List, o)
	}
	sort.Slice(a, func(i, j int) bool {
		//not stable for ==
		return a[i].Count > a[j].Count
	})
	for _,b := range a {
		//fmt.Println(b)
		fmt.Printf("%s %d :", b.BkName, b.Count)
		for _,x := range b.List {
			fmt.Printf(" %s", x.Name)
			if showDet {
				fmt.Printf(" %f", x.DetP)
			}
		}
		if b.Count > 1 {
			fmt.Printf("\n")
		}else {
			fmt.Printf(" | ")
		}

	}
	//fmt.Printf("\n")
}
