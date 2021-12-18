package filter

import (
	"faya/list"
	"strings"
)

func Filter300(input []*list.TimeObject) []*list.TimeObject {
	ret := make([]*list.TimeObject, 0)
	for _, obj := range input {
		if strings.Index(obj.Code, "300") == 0 {
			//fmt.Println(*obj)
			ret = append(ret, obj)
		}
	}
	return ret
}
