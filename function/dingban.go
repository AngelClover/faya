package function

import (
	"faya/list"
	"fmt"
	"time"
)

var (
	//mockOpeningTime = false
)


func Dingban() {
	for ;mockOpeningTime || list.IsOpeningTime() ; time.Sleep(5 * time.Second){
		fmt.Println("------------------")
		ll := list.GetRealtimeList()
		for _, o := range ll{
			if (o.DetP.(float64) < 9.7){
				break
			}
			fmt.Println(o.Code, o.Name, o.DetP)
		}
	}
}
