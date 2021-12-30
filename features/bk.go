package features

import (
	"faya/data"
	"faya/list"
	"fmt"
)

func GetBk(o *list.TimeObject) string{
	content := data.Get(o.Code, "bk")
	cc, ok := content.(string)
	if !ok {
		fmt.Println("type assertion error:%s", content)
	}
	return cc
}
