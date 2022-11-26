package data

import (
	"faya/list"
	"fmt"
	"sync"
)

var (
	//dataMap = make(map[string]interface{}, 0)
	dataMap = sync.Map{}
)
func Get(code string, datatype string) interface{} {
	d, ok := dataMap.Load(code)
	if !ok {
		//dataMap[code] = make(map[string]interface{}, 0)
		dataMap.Store(code, sync.Map{})
		d, _ = dataMap.Load(code)
	}
	//dm := d.(map[string]interface{})
	dm := d.(sync.Map)
	res, okk := dm.Load(datatype)
	if !okk {
		//fmt.Println("get data for ", code, datatype)
		if datatype == "rik" {
			dm.Store(datatype, list.RiKCode(code))
			res, _ = dm.Load(datatype)
// 			fmt.Println("get data:", dm[datatype])
		}else if datatype == "rik_reverse" {
			dm.Store(datatype, list.RiKCodeReverse(code))
			res, _ = dm.Load(datatype)
// 			fmt.Println("get data:", dm[datatype])
		}else if datatype == "bk" {
			dm.Store(datatype, list.GetBkCode(code))
			res, _ = dm.Load(datatype)
		}else {
			fmt.Println("get datatype error")
			res = nil
		}
	} else {
		fmt.Println("have stored data for ", code, datatype, "use it")
	}
	return res
}
