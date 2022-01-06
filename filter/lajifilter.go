package filter

import (
	"faya/list"
	"fmt"
	"strings"
)
var (
	lajiList = []string {"000562", "002070"}
)
func LajiFilter(input []* list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, o:= range input{
		ok := true
		for _, j := range lajiList {
			if strings.Index(o.Code, j) == 0 {
				fmt.Println("Kick out", o)
				ok = false
				break
			}
		}
		if ok {
			ret = append(ret, o)
		}
	}
	return ret
}
