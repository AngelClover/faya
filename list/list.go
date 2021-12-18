package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	//"github.com/gogf/gf/frame/g"
	//"github.com/gogf/gf/os/gtime"
)

var (
	listUrl = "http://57.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&fltt=2&invt=2&fid=f3&fs=m:0%2Bt:6,m:0%2Bt:13,m:0%2Bt:80,m:1%2Bt:2,m:1%2Bt:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152,f40,f38,f192"
)

type TimeObject struct {
	Code           string      `json:"f12"` // 代码
	Name           string      `json:"f14"` // 名称
	Price          interface{} `json:"f2"`  // 最新价
	DetP           interface{} `json:"f3"`  // 涨跌幅
	Det            interface{} `json:"f4"`  // 涨跌额
	Amount         interface{} `json:"f5"`  // 成交量（手）
	Money          interface{} `json:"f6"`  // 成交额
	Amplitude      interface{} `json:"f7"`  // 振幅
	Max            interface{} `json:"f15"` // 最高
	Min            interface{} `json:"f16"` // 最低
	TodayOpen      interface{} `json:"f17"` // 今开
	YesterdayClose interface{} `json:"f18"` // 昨收
	Ratio          interface{} `json:"f10"` // 量比
	Turnover       interface{} `json:"f8"`  // 换手率
	Pe             interface{} `json:"f9"`  // 市盈率
	Pb             interface{} `json:"f23"` // 市净率
}

type TimeList struct {
	Diff []*TimeObject `json:"diff"`
}

type TimeListWrapper struct {
	Data *TimeList `json:"data"`
}

func getDataTime(t *gtime.Time) (*gtime.Time, error) {
	res, err := gtime.ConvertZone(t.String(), "Asia/Shanghai")
	if err != nil {
		g.Log().Error(err)
	}
	hour := res.Hour()
	minute := res.Minute()
	second := res.Second()
	total := second + minute*60 + hour*3600
	bg1 := 9*3600 + 15*60
	ed1 := 11*3600 + 30*60
	bg2 := 13 * 3600
	ed2 := 15 * 3600
	re := total
	nt := t
	if total < bg1 {
		//return lastday not overwrite
		nt = t.AddDate(0, 0, -1)
		re = ed2 + 1
	}
	if total > ed1 && total < bg2 {
		re = ed1
	}
	if total > ed2 {
		re = ed2
	}
	y, m, d := nt.Date()
	h := re / 3600
	re -= h * 3600
	mm := re / 60
	s := re % 60
	ns := t.Nanosecond()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != err {
		g.Log().Error(err)
	}
	ret_time := time.Date(y, m, d, h, mm, s, ns, loc)

	return gtime.New(ret_time), err
}

func Get() []*TimeObject {
	fmt.Println("list.Get(")
	nosql := true
	
	content, had := db.Get("list")
	if had == false {

		resp, err := http.Get(listUrl)
		if err != nil {
			fmt.Println("http.get error", listUrl)
			return nil
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil.readAll error", resp.Body)
			return nil
		}

		if resp.StatusCode > 299 {
			fmt.Println("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
			return nil
			//return body, errors.New(strconv.Itoa(resp.StatusCode))
		}
		content = body
	}

	var resp2 TimeListWrapper
	err = json.Unmarshal(body, &resp2)
	if err != nil {
		fmt.Println("json unmarshal error : " + err.Error())
		return nil
	}

	writeTime := gtime.Now()
	contentTime, err := getDataTime(writeTime)
	fmt.Println(contentTime)

	for _, obj := range resp2.Data.Diff {
		fmt.Println(obj)
		if nosql {
		}
	}
	return resp2.Data.Diff
}
