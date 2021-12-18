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

type RiKUnit struct {
	Date string
	Open float64
	Close float64
	High float64
	Low float64
	Amount int
	Money float64
	Wide float64
	Det float64
	DetPrice float64
	Turnover float64
	Features map[string]interface{}
}
type RiKDailyString struct {
	text string
}

type RiKDetails struct {
	Code string `json:"code"`
	Market int `json:"market"`
	Name string `json:"name"`
	Decimal int `json:"decimal"`
	Dktotal int `json:"dktotal"`
	PreKPrice float64 `json:"preKPrice"`
	Klines []string `json:"klines"`
}

var (
	RikUrlPart1 = "http://74.push2his.eastmoney.com/api/qt/stock/kline/get?secid="
	RikUrlPart2 = "&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&klt=101&fqt=0&end=20500101&lmt=120"
	lastWebVisitTime = time.Now()
	webVisitInterval = 100 * time.Millisecond
)

type RiKResponse struct {
	Rc int `json:"rc"`
	Rt int `json:"rt"`
	Svr int `json:"svr"`
	Lt int `json:"lt"`
	Full int `json:"full"`
	Data *RiKDetails `json:"data"`
}

func RiK(obj TimeObject) []*RiKUnit{
	return RiKCode(obj.Code)
}
func RiKCode(code string) []*RiKUnit {

	market := "0"
	if strings.Index(code, "6") == 0 {
		market = "1"
	}
	cacheKey := code + "rik"
	var content []byte
	contentStr, had := db.Get(cacheKey)
	if had == true {
		content = []byte(contentStr)
	}else {
		RikUrl := RikUrlPart1 + market + "." + code + RikUrlPart2
		//time control
		timeSpend := time.Since(lastWebVisitTime)
		fmt.Println("timespend:", timeSpend," for web visit time interval:", webVisitInterval, "last:", lastWebVisitTime)
		if timeSpend < webVisitInterval{
			fmt.Println("sleep for web visit time interval:", webVisitInterval)
			time.Sleep(webVisitInterval)
		}

		resp, err := http.Get(RikUrl)
		lastWebVisitTime = time.Now()
		if err != nil {
			fmt.Println("http.get error", listUrl)
			return nil
		}
		defer resp.Body.Close()
	// 	fmt.Println(resp)
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
		db.Insert(cacheKey, string(body))
	}

	//fmt.Println(content)
	var resp2 RiKResponse
	err := json.Unmarshal(content, &resp2)

	if err != nil {
		fmt.Println("json unmarshal error : " + err.Error())
		return nil
	}
	//fmt.Println(resp2.Data)
// 	fmt.Println(resp2.Data.Klines)
// 	fmt.Println(strings.Split(resp2.Data.Klines[0], ","))
	ret := make([]*RiKUnit, 0)
	for _, day := range resp2.Data.Klines {
		ll := strings.Split(day, ",")
		var rk RiKUnit
		rk.Date = ll[0]
		
		if  f, err := strconv.ParseFloat(ll[1], 64); err == nil {
			rk.Open = f
		} else {
			fmt.Println("1 parse error", ll[1] ,f,  ll )
			continue
		}
		if  f, err := strconv.ParseFloat(ll[2], 64); err == nil {
			rk.Close = f
		} else {
			fmt.Println("2 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[3], 64); err == nil {
			rk.High = f
		} else {
			fmt.Println("3 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[4], 64); err == nil {
			rk.Low = f
		} else {
			fmt.Println("4 parse error", ll)
			continue
		}
		if  i, err := strconv.Atoi(ll[5]); err == nil {
			rk.Amount = i
		} else {
			fmt.Println("5 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[6], 64); err == nil {
			rk.Money = f
		} else {
			fmt.Println("6 parse error", ll)
			continue
		}
		if  f, err := strconv.ParseFloat(ll[7], 64); err == nil {
			rk.Wide = f
		} else {
			fmt.Println("7 parse error", ll)
			continue
		}
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
		rk.Features = make(map[string]interface{})
		ret = append(ret, &rk)
	}
	return ret
// 	fmt.Println(ret)
// 	print(ret)
// 	return resp2.Data.Klines
// reverse

}
func RiKCodeReverse(code string) []*RiKUnit {
	ret := RiKCode(code)
	rev := make([]*RiKUnit, 0)
	for i := len(ret) - 1; i >= 0; i = i - 1 {
		rev = append(rev, ret[i])
	}
	return rev
}
