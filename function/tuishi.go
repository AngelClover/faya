package function

import (
	"faya/list"
	"fmt"
	"time"
)

func TuishiProducer() {
  	l := list.Get()
	daishangshi := make([]*list.TimeObject, 0)
	tuishi := make([]*list.TimeObject, 0)
	for _, o := range l{
		p := list.RiKCodeReverse(o.Code)
				
		if len(p) <= 0 {
			daishangshi = append(daishangshi, o)
		} else {
			boundry := time.Now().AddDate(0, 0, -20)
			tdate,_  := time.Parse("2006-01-02", p[0].Date)

			if tdate.Before(boundry) {
				fmt.Println(o, p[0])
				tuishi = append(tuishi, o)
			}
		}
	}

	ret := ""
	for _, o := range daishangshi {
		ret = ret + "\"" + o.Code + "\","
		//fmt.Println("\"",o.Code,"\",")
	}
	fmt.Println("daishangshi:", ret)
	ret = ""
	for _, o := range tuishi {
		ret = ret + "\"" + o.Code + "\","
// 		fmt.Println("\"",o.Code,"\",")
	}
	fmt.Println("tuishi:", ret)
}
