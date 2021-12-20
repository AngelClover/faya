package data

import (
	"faya/list"
	"fmt"
)

var (
	dataMap = make(map[string]interface{}, 0)
)
func Get(code string, datatype string) interface{} {
	d, ok := dataMap[code]
	if !ok {
		dataMap[code] = make(map[string]interface{}, 0)
		d = dataMap[code]
	}
	dm := d.(map[string]interface{})
	res, okk := dm[datatype]
	if !okk {
		//fmt.Println("get data for ", code, datatype)
		if datatype == "rik" {
			dm[datatype] = list.RiKCode(code)
// 			fmt.Println("get data:", dm[datatype])
			res = dm[datatype]
		}else if datatype == "rik_reverse" {
			dm[datatype] = list.RiKCodeReverse(code)
// 			fmt.Println("get data:", dm[datatype])
			res = dm[datatype]
		}else {
			fmt.Println("get datatype error")
			res = nil
		}
	} else {
		fmt.Println("have stored data for ", code, datatype, "use it")
	}
	return res
}
