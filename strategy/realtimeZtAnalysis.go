package strategy

import (
	"faya/features"
	"faya/list"
	"fmt"
	"time"
)

type Event struct{
	Time time.Time
	Code string
	Name string
	Message string
	Det float64
}

var (
	Q []Event = make([]Event, 0, 1000)
	head int = 0
	tail int = 0
	lastList []*list.TimeObject
)

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


func isZt(o *list.TimeObject) bool {
	if o == nil {
		return false
	}
	/*
	baseprice, ok := o.YesterdayClose.(float64)
	if !ok {
		return false
	}
	*/
	//uplimit := baseprice * (1 + UpPercent(o.Code) * 0.01)
	uplimit := UpPercent(o.Code)
	detp, ok := o.DetP.(float64)
	if !ok {
		return false
	}
	return detp >= (uplimit - 0.2)
}
func nearZt(o *list.TimeObject) bool {
	baseprice, ok := o.YesterdayClose.(float64)
	if !ok {
		return false
	}
	splitline := baseprice * (1 + UpPercent(o.Code) * 0.6 * 0.01)
	detp, ok := o.DetP.(float64)
	if !ok {
		return false
	}
	/*
	if o.Code == "000812" {
		fmt.Println("Alan debug", o.Code, " ", o.Name," ",baseprice, " ", detp, "    ", splitline)
		fmt.Println(o)
	}
	*/
	//no change
	if detp >= UpPercent(o.Code) - 0.2 {
		return false
	}
	return baseprice * (1 + detp*0.01) >= splitline
}
// return true will be filtered
func filtered(o *list.TimeObject) bool {
	uplimit := UpPercent(o.Code)
	detp, ok := o.DetP.(float64)
	if !ok {
		return true
	}
	if  detp < uplimit * 0.5 {
		return true
	} else {
		return false
	}
}


func OutputAnalysis(o *list.TimeObject){
	ztdays := "xx"

	rk := list.RiKCodeReverse(o.Code)
	features.GetRecentTurnover(rk)

	ztdayscount := features.GetZtDaysCount(o, rk)

	if len(rk) <= 0{
		return
	}

	rto := rk[0].Features["RecentTurnover"].(float64)
	//baseztcount := rk[0].Features["ZtDaysCount"].(int)
	
	var rtop float64
	ro, ok :=(o.Turnover).(float64) 
	if ok {
		rtop = ro / rto * 100
	} else {
		ro = 0
		rtop = 0
	}
	am, ok :=(o.Amount).(float64) 
	if !ok {
		am = 0
	}

	mo, ok :=(o.Money).(float64) 
	if !ok {
		mo = 0
	}
	//rtop :=  (o.Turnover).(float64) / rto * 100.0

	fmt.Println(o.Code, o.Name, ztdays, o.DetP, " | ", ro,  am/1e4, mo/1e8, " | ", rtop, " | ", ztdayscount)

}

func Analysis(l []*list.TimeObject){

	ztc := 0
	nearztc := 0

	//fl = [isZt]

	for _,o := range l{
		if isZt(o) {
			ztc++
		}
		if nearZt(o) {
			nearztc++
		}
	}
	fmt.Println("ztc:", ztc, "  nearztc:", nearztc, "   total:", len(l))
	var ztQ []*list.TimeObject
	for _,o := range l{
		if isZt(o) {
			//OutputAnalysis(o)
			//fmt.Printf("%s ", o.Name)
			ztQ = append(ztQ, o)
		}
	}
	showQueue(ztQ, false)

	fmt.Println("------------up6")
	var nearQ []*list.TimeObject
	for _,o := range l{
		if nearZt(o) {
			//OutputAnalysis(o)
			//fmt.Printf("%s %f ", o.Name, o.DetP)
			nearQ = append(nearQ, o)
		}
	}
	showQueue(nearQ, true)

	
	for _, o := range l{
		if !filtered(o){
			var oo *list.TimeObject = nil
			for _, oi := range lastList {
				if oi.Code == o.Code {
					oo = oi
					break
				}
			}
			//if oo != nil {
				if isZt(oo) != isZt(o) {
					m := ""
					if isZt(o) {
						m = "feng"
					}else {
						m = "kai"
					}

					det, ok := o.DetP.(float64)
					if !ok {
						det = 0
					}

					x := &Event{
						Code: o.Code,
						Name: o.Name,
						Time: time.Now(),
						Message: m,
						Det: det,
					}
					Q = append(Q, *x)
					tail++
				}
			//}
		}
	}
	fmt.Println("queue[", head, ",", tail, "]")
	for i := Max(tail - 10, 0); i < tail; i++ {
		if Q[i].Message == "feng" {
			fd := list.FengdanCode(Q[i].Code)
			fmt.Printf("%v %d(%f) %s ||", Q[i].Time, fd.Buy1, float64(fd.Buy1)*fd.Buy1Price/10000, fd.Bk)
			for _,o := range l{
				if o.Code == Q[i].Code {
					OutputAnalysis(o)
					break
				}
			}
			//fmt.Printf("%+v\n",  fd)
			//fmt.Println(fd.Buy1, " | ", fd.Buy2, " | ", fd.Weicha, " | ", fd.Bk)
		} else {
			fmt.Println(Q[i])
		}
	}

	lastList = l
}
