package filter

import "faya/list"

var (
	holdlist = []string {"603222", "601107", "002247", "001317", "301111", "300052"}
)


func HoldFilter(input []*list.TimeObject) []*list.TimeObject{
	ret := make([]*list.TimeObject, 0)
	for _, obj := range input{
		hold := false
		for _,j := range holdlist {
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
