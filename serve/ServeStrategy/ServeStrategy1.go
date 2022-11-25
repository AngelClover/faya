package servestrategy

import (
	"encoding/json"
	"faya/db"
	"faya/features"
	"faya/list"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)


type ServeStrategy1 struct {
}



func getLastDate() string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	now := time.Now()
	y := now.Year()
	m := now.Month()
	d := now.Day()
	startTime := time.Date(y, m, d, 9, 15, 0, 0, loc)
	
// 	fmt.Println(time.Now().Date())
// 	fmt.Println(startTime)
// 	fmt.Println(startTime.Local())
	if !now.After(startTime) {
		now = now.AddDate(0, 0, -1)
		for ; now.Weekday() == time.Sunday || now.Weekday() == time.Saturday; {
			now = now.AddDate(0, 0, -1)
		}
	}
	return strings.Split(now.String(), " ")[0]
}
func getLastLastDate() string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	now := time.Now()
	y := now.Year()
	m := now.Month()
	d := now.Day()
	startTime := time.Date(y, m, d, 9, 15, 0, 0, loc)
	
// 	fmt.Println(time.Now().Date())
// 	fmt.Println(startTime)
// 	fmt.Println(startTime.Local())
	for i := 0; i < 2;i++ {
		if !now.After(startTime) {
			now = now.AddDate(0, 0, -1)
			for ; now.Weekday() == time.Sunday || now.Weekday() == time.Saturday; {
				now = now.AddDate(0, 0, -1)
			}
		}
	}
	return strings.Split(now.String(), " ")[0]
}
//sameday compare
func isBeforeOpening() bool{
	loc := time.FixedZone("UTC+8", +8*60*60)
	now := time.Now()
	y := now.Year()
	m := now.Month()
	d := now.Day()
	startTime := time.Date(y, m, d, 9, 15, 0, 0, loc)
	
// 	fmt.Println(time.Now().Date())
// 	fmt.Println(startTime)
// 	fmt.Println(startTime.Local())
	return now.Before(startTime)
}
//sameday compare
func isAfterClosing() bool{
	loc := time.FixedZone("UTC+8", +8*60*60)
	now := time.Now()
	y := now.Year()
	m := now.Month()
	d := now.Day()
	closeTime := time.Date(y, m, d, 15, 05, 0, 0, loc)
	
// 	fmt.Println(time.Now().Date())
// 	fmt.Println(startTime)
// 	fmt.Println(startTime.Local())
	return now.After(closeTime)
}

/*
if time.now
before opening, will contrast list.Get -> lastlast trading date,
between opening and closing, will contrast list.Get -> last trading date
after closing, will contrast list.Get -> last trading date
*/
type O struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Amount []int `json:"amount"`
	RecentAmountRatio float64 `json:"ratio"`
	Bk string `json:"bk"`
	CheckedDate string `json:"ckdate"`
	AmNow float64 `json:"now"`
	AmLast float64 `json:"last"`
	LastLastRatio float64 `json:"llr"` //amount yesterday / before
	LastLastRatioFlag bool `json:"llrf"`

	WeekAmount int `json:"wa"`
	LastWeekAmount int `json:"lwa"`
	WeekAmountRatio float64 `json:"war"`
	WeekAmountFlag bool `json:"waf"`

	DetP float64 `json:"detp"`
	DetpFlag bool `json:"detpf"`

	TenOnSeventy int `json:"tos"`
	TenOnSeventyFlag bool `json:"tosf"`
}

func (s *ServeStrategy1) Run() []byte{
	tm := time.Now()
	fmt.Println("now time:", tm)
	isOpening := list.IsOpeningTime()
	if (isOpening){
		fmt.Println("is Opening time")
	}else {
		fmt.Println("is NOT Opening time")
	}

	t := getLastDate()
	contrastDate := t
	if isBeforeOpening(){
		fmt.Println("is before opening")
		contrastDate = getLastLastDate()
	}else {
		fmt.Println("is NOT before opening")

	}
	fmt.Println("contrast date: ", contrastDate)

	var l []*list.TimeObject
	if isOpening {
		l = list.GetRealtimeList()
	}else {
		l = list.Get()
	}

	counter := 0
	countlimit := 10
	type ANS struct {
		Tm string `json:"tm"`
		OO []O `json:"data"`
	}
	var ans ANS
	ans.Tm = list.GetContentTime().String()
	ans.OO = make([]O, 0)
	for _, op := range l{
		//fmt.Println(op)
		a := list.RiKCodeReverse(op.Code)
		features.GetRecentTurnover(a)
		//bk := features.GetBk(op)
		ll := len(a)
		if ll > 2{
			var o O
			//basic feature
			o.Code = op.Code
			o.Name = op.Name
			o.Bk = ""
			o.Amount = []int{a[0].Amount, a[1].Amount, a[2].Amount}
			//daily amount
			am_now, ok := op.Amount.(float64)
			o.AmNow = am_now
			if !ok {
				//fmt.Println(op.Amount, "->", am_now, "not ok")
				o.AmNow = 0
			}
			o.AmLast = float64(a[0].Amount)
			o.CheckedDate = a[0].Date
			if !isOpening {
				o.AmLast = float64(a[1].Amount)
				o.CheckedDate = a[1].Date
			}
			o.RecentAmountRatio = o.AmNow / o.AmLast
			InQCondition := (o.RecentAmountRatio >= 2)
			if !InQCondition {
				continue
			}

			detpt, ok := op.DetP.(float64)
			if ok{
				o.DetP = detpt
			}else {
				o.DetP = 0
			}
			o.LastLastRatio = float64(a[1].Amount) / float64(a[0].Amount)
			if o.LastLastRatio < 2{
				o.LastLastRatioFlag = true
			}else {
				o.LastLastRatioFlag = false
			}
			//depP flag judgement
			if (o.DetP > -0.01) {
				o.DetpFlag = true
			}else {
				o.DetpFlag = false
			}

			//weekly amount
			amt, ok := op.Amount.(float64)
			if ok {
				o.WeekAmount = int(amt)
			}else{
				o.WeekAmount = 0
			}
			for i:=0 ;i < 5 && i < len(a); i++{
				if i == 0 && isOpening {
				}else {
					o.WeekAmount += a[i].Amount
				}
			}
			o.LastWeekAmount = 0
			for i:=0 ;i < 5 && i+5 < len(a); i++{
				o.LastWeekAmount += a[i+5].Amount
			}
			if o.LastWeekAmount == 0 {
				o.WeekAmountRatio = 0
			}else {
				o.WeekAmountRatio = float64(o.WeekAmount) / float64(o.LastWeekAmount)
			}
			if (o.WeekAmountRatio > 1.8){
				o.WeekAmountFlag = true
			}else {
				o.WeekAmountFlag = false
			}
			

			//golden crossing
			features.GetDayAvg(a, 10)
			features.GetDayAvg(a, 73)
			
			tenOnSeventy := func(rk *list.RiKUnit) bool{
				if rk.Features["Day10"] == 0 {
					return false
				}
				if rk.Features["Day73"] == 0 {
					return false
				}

				det := rk.Features["Day10"].(float64) - rk.Features["Day73"].(float64)
				if math.Abs(det) < 0.01 {
					return false
				}
				return det > 0
			}
			today := tenOnSeventy(a[0])
			length := 0
			for j := 0; j < ll; j++{
				rk := a[j]
				if tenOnSeventy(rk) == today{
					length += 1
				}else {
					break
				}
			}
			if today {
				o.TenOnSeventy = length
			}else {
				o.TenOnSeventy = -length
			}
			if o.TenOnSeventy >= -1 && o.TenOnSeventy < 5 {
				o.TenOnSeventyFlag = true
			}else {
				o.TenOnSeventyFlag = false
			}

			//put in queue condition
			if InQCondition {
				ans.OO = append(ans.OO, o)
			}
			//fmt.Println(o, am_last, am_now)
		}



		counter += 1
		if counter > countlimit {
			//break
		}
	}
	//fmt.Println("total length:", counter)

	sort.Slice(ans.OO, func(i, j int) bool {
		return ans.OO[i].RecentAmountRatio > ans.OO[j].RecentAmountRatio
	})

	fmt.Println("OO length:", len(ans.OO))

	AmountRatioHeadPrint(ans.OO)
	//fmt.Println(ans)
	ret := []byte("")
	b, err := json.Marshal(ans.OO)
	if err != nil {
		fmt.Println(err)
		ret = []byte("")
	}else {
		ret = []byte("{\"tm\":\"" + ans.Tm + "\",\"data\":" + string(b) + "}")

		cacheKey := "ss1"
		db.Insert(cacheKey, string(ret))
	}
	return ret
	

	//lc := list.GetStamp(contrastDate)

}

func (s *ServeStrategy1) GetCached() []byte {
	cacheKey := "ss1"
	contentStr, had := db.Get(cacheKey)
	var content []byte 
	if had == true {
		content = []byte(contentStr)
	} else {
		content = s.Run()
	}
	return content
}


func AmountRatioHeadPrint(oo [] O){
	counter := 0
	counterlimit := 10
	for _, o := range oo{
		//fmt.Printf("%s %s %s %v => %.2f %.2f %.2f  ||%.2f|| %.2f %.2f %.2f %s\n", o.Code, o.Name, o.Bk, o.Ztdays, o.Detp[2], o.Detp[1], o.Detp[0], o.RecentTO,  o.Turnover[2], o.Turnover[1], o.Turnover[0], flag)
		fmt.Printf("%s %s %s %v %s | %f -> %f\n", o.Code, o.Name, o.Bk, o.RecentAmountRatio, o.CheckedDate, o.AmLast, o.AmNow )

		counter += 1
		if counter > counterlimit {
			break
		}
	}
	fmt.Printf("total oo length:%d\n", len(oo))
}
