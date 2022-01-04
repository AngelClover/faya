package list

import (
	"encoding/json"
	"faya/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//chang 'get' to 'sse' will get event stream
var (
	MinUrlPart1 = "http://push2.eastmoney.com/api/qt/stock/trends2/get?fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&ut=fa5fd1943c7b386f172d6893dbfba10b&secid="
	MinUrlPart2 = "&ndays=1&iscr=0"
)

type MinUnit struct {
	DateTime string
	Open float64
	Close float64
	High float64
	Low float64
	Amount int
	Money float64
	Avg float64 // all day avg
	Features map[string]interface{}
}

type MinDetails struct {
	Code string `json:"code"`
	Market int `json:"market"`
	Name string `json:"name"`
	Type int `json:"type"`
	Status int `json:"status"`
	Decimal int `json:"decimal"`
	PreSettlement float64 `json:"preSettlement"`
	PreClose float64 `json:"preClose"`
	PrePrice float64 `json:"prePrice"`
	Beticks string `json:"beticks"`
	TrendsTotal int `json:"trendsTotal"`
	Time int `json:"time"`
	Kind int `json:"kind"`
	Trends []string `json:"trends"`
}

type MinResponse struct {
	Rc int `json:"rc"`
	Rt int `json:"rt"`
	Svr int `json:"svr"`
	Lt int `json:"lt"`
	Full int `json:"full"`
	Data *MinDetails `json:"data"`     // data can be null in return
}

func isOpeningTime() bool {
	loc := time.FixedZone("UTC+8", +8*60*60)
	now := time.Now()

	if now.Weekday() == time.Sunday || now.Weekday() == time.Saturday {
		return false
	}

	y := now.Year()
	m := now.Month()
	d := now.Day()
	startTime := time.Date(y, m, d, 9, 15, 0, 0, loc)
	endTime := time.Date(y, m, d, 15, 00, 0, 0, loc)

	if now.After(startTime) && endTime.After(now) {
		return true
	}
	return false
}

func getCacheKeyMin(code string) string {
	ret := code + "min"

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
// 	str := now.Date()
// 	fmt.Println(now.String())
// 	fmt.Println(strings.Split(now.String(), " ")[0])
	ret = ret + strings.Split(now.String(), " ")[0]
	return ret

}

func MinCode(code string) []*MinUnit {

	cacheKey := getCacheKeyMin(code)
	var content []byte
	contentStr, had := db.Get(cacheKey)
	if had == true {
		content = []byte(contentStr)
	}else {
		MinUrl := MinUrlPart1 + GetSecid(code) + MinUrlPart2
// 		fmt.Println(MinUrl)
		//time control
		timeSpend := time.Since(lastWebVisitTime)
		fmt.Println("timespend:", timeSpend," for web visit time interval:", webVisitInterval, "last:", lastWebVisitTime)
		if timeSpend < webVisitInterval{
			fmt.Println("sleep for web visit time interval:", webVisitInterval)
			time.Sleep(webVisitInterval)
		}

		resp, err := http.Get(MinUrl)
// 	 	fmt.Println(resp)

		lastWebVisitTime = time.Now()
		if err != nil {
			fmt.Println("http.get error", MinUrl, err)
			return nil
		}
		defer resp.Body.Close()
// 	 	fmt.Println(resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil.readAll error", resp.Body)
			return  nil
		}

		if resp.StatusCode > 299 {
			fmt.Println("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
			return nil
			//return body, errors.New(strconv.Itoa(resp.StatusCode))
		}
		
		content = body
		if !isOpeningTime(){
			db.Insert(cacheKey, string(body))
		}
	}

// 	fmt.Println(content)
	var resp2 MinResponse
	err := json.Unmarshal(content, &resp2)

	if err != nil {
		fmt.Println("json unmarshal error : " + err.Error())
		return nil
	}
// 	fmt.Println(resp2.Data)

	ret := make([]*MinUnit, 0)
	if resp2.Data == nil {
		return ret
	}
	for _, day := range resp2.Data.Trends {
		ll := strings.Split(day, ",")
		var min MinUnit
		min.DateTime = ll[0]
		
		if  f, err := strconv.ParseFloat(ll[1], 64); err == nil {
			min.Open = f
		} else {
			fmt.Println("1 parse error", ll[1] ,f,  ll )
			continue
		}
		if  f, err := strconv.ParseFloat(ll[2], 64); err == nil {
			min.Close = f
		} else {
			fmt.Println("2 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[3], 64); err == nil {
			min.High = f
		} else {
			fmt.Println("3 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[4], 64); err == nil {
			min.Low = f
		} else {
			fmt.Println("4 parse error", ll)
			continue
		}
		if  i, err := strconv.Atoi(ll[5]); err == nil {
			min.Amount = i
		} else {
			fmt.Println("5 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[6], 64); err == nil {
			min.Money = f
		} else {
			fmt.Println("6 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[7], 64); err == nil {
			min.Avg = f
		} else {
			fmt.Println("7 parse error", ll)
			continue
		}
		/*
		if  f, err := strconv.ParseFloat(ll[8], 64); err == nil {
			rk.Det = f
		} else {
			fmt.Println("8 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[9], 64); err == nil {
			rk.DetPrice = f
		} else {
			fmt.Println("9 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[10], 64); err == nil {
			rk.Turnover = f
		} else {
			fmt.Println("10 parse error", ll)
			continue
		}
		*/
		min.Features = make(map[string]interface{})
		ret = append(ret, &min)
// 		fmt.Println(min)
	}
	return ret

}
func MinCodeReverse(code string) []*MinUnit {
	ret := MinCode(code)
	rev := make([]*MinUnit, 0)
	for i := len(ret) - 1; i >= 0; i = i - 1 {
		rev = append(rev, ret[i])
	}
	return rev
}
func Min(obj *TimeObject) []*MinUnit{
	return MinCode(obj.Code)
}
