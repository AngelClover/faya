package filter

import "faya/list"

var (
	Holdlist = []string {"600683" }
)


func HoldFilter(input []*list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, obj := range input{
		hold := false
		for _,j := range Holdlist {
			if j == obj.Code || j == obj.Name {
				hold = true
				break
			}
		}
		if hold {
			ret = append(ret, obj)
		}
	}
	return ret
}
