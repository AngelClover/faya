package serve

import (
	"encoding/json"
	"faya/features"
	"faya/list"
	"fmt"
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
	am_now float64 `json:"now"`
	am_last float64 `json:"last"`
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
		if len(a) > 2{
			var o O
			o.Code = op.Code
			o.Name = op.Name
			o.Bk = ""
			o.Amount = []int{a[0].Amount, a[1].Amount, a[2].Amount}
			am_now, ok := op.Amount.(float64)
			o.am_now = am_now
			if !ok {
				//fmt.Println(op.Amount, "->", am_now, "not ok")
				o.am_now = 0
			}
			o.am_last = float64(a[0].Amount)
			o.CheckedDate = a[0].Date
			if !isOpening {
				o.am_last = float64(a[1].Amount)
				o.CheckedDate = a[1].Date
			}
			o.RecentAmountRatio = o.am_now / o.am_last
			if o.RecentAmountRatio >= 2{
				ans.OO = append(ans.OO, o)
			}
			//fmt.Println(o, am_last, am_now)
		}



		counter += 1
		if counter > countlimit {
			//break
		}
	}
	fmt.Println("total length:", counter)

	sort.Slice(ans.OO, func(i, j int) bool {
		return ans.OO[i].RecentAmountRatio > ans.OO[j].RecentAmountRatio
	})


	AmountRatioHeadPrint(ans.OO)
	//fmt.Println(ans)
	b, err := json.Marshal(ans.OO)
	if err != nil {
		fmt.Println(err)
		return []byte("")
	}else {
		return []byte("{\"tm\":\"" + ans.Tm + "\",\"data\":" + string(b) + "}")
	}
	

	//lc := list.GetStamp(contrastDate)

}



func AmountRatioHeadPrint(oo [] O){
	counter := 0
	counterlimit := 10000
	for _, o := range oo{
		//fmt.Printf("%s %s %s %v => %.2f %.2f %.2f  ||%.2f|| %.2f %.2f %.2f %s\n", o.Code, o.Name, o.Bk, o.Ztdays, o.Detp[2], o.Detp[1], o.Detp[0], o.RecentTO,  o.Turnover[2], o.Turnover[1], o.Turnover[0], flag)
		fmt.Printf("%s %s %s %v %s | %f -> %f\n", o.Code, o.Name, o.Bk, o.RecentAmountRatio, o.CheckedDate, o.am_last, o.am_now )

		counter += 1
		if counter > counterlimit {
			break
		}
		if o.RecentAmountRatio < 2{
			break
		}
	}
}
