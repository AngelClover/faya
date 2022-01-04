package function

import (
	"faya/list"
	"fmt"
)

func Prefill() {
	l := list.Get()
	ll := len(l)
	i := 0
	for _,o := range l {
		fmt.Println("prefilling ", i, ll)
		list.RiK(o)
		list.Min(o)
		//list.GetBk(o)
		i = i + 1
	}
}
