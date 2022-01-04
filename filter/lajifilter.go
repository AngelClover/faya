package filter

import (
	"faya/list"
	"fmt"
	"strings"
)
func LajiFilter(input []* list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, o:= range input{
		if strings.Index(o.Code, "000562") == 0 {
			fmt.Println("Kick out", o)
			continue
		}
		ret = append(ret, o)
	}
	return ret
}
