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

type BkCombinedStatUnit struct {
	BkName string
	ZtCount int
	ZtList []*list.TimeObject
	NearCount int
	NearList []*list.TimeObject
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
func showCombinedQueue(zt []*list.TimeObject, near []*list.TimeObject) {
	var a []*BkCombinedStatUnit
	for _,o := range zt{
		bk := list.GetBk(o)
		var b *BkCombinedStatUnit = nil
		for _,x := range a {
			if x.BkName == bk {
				b = x
				break
			}
		}
		if b == nil {
			b = &BkCombinedStatUnit{
				BkName: bk,
				ZtCount: 0,
				NearCount: 0,
			}
			a = append(a, b)
		}
		b.ZtCount += 1
		b.ZtList = append(b.ZtList, o)
	}
	for _,o := range near{
		bk := list.GetBk(o)
		var b *BkCombinedStatUnit = nil
		for _,x := range a {
			if x.BkName == bk {
				b = x
				break
			}
		}
		if b == nil {
			b = &BkCombinedStatUnit{
				BkName: bk,
				ZtCount: 0,
				NearCount: 0,
			}
			a = append(a, b)
		}
		b.NearCount += 1
		b.NearList = append(b.NearList, o)
	}

	sort.Slice(a, func(i, j int) bool {
		//not stable for ==
		return a[i].ZtCount > a[j].ZtCount ||
		a[i].ZtCount == a[j].ZtCount && a[i].NearCount == a[j].NearCount
	})

	for _,b := range a {
		//fmt.Println(b)
		fmt.Printf("%s %d+%d :", b.BkName, b.ZtCount, b.NearCount)
		for _,x := range b.ZtList {
			fmt.Printf(" %s", x.Name)
		}
		fmt.Printf(" | ")
		for _,x := range b.NearList {
			fmt.Printf(" %s %f", x.Name, x.DetP)
		}

		if b.ZtCount > 1 {
			fmt.Printf("\n")
		}else {
			fmt.Printf(" | ")
		}

	}

}
