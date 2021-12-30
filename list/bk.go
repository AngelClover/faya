package list

import (
	"encoding/json"
	"faya/db"

	//"faya/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	BkUrlPart = "http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=2&fields=f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f127,f199,f128,f193,f196,f194,f195,f197,f80,f280,f281,f282,f284,f285,f286,f287,f292&secid="
// 	lastWebVisitTime = time.Now()
// 	webVisitInterval = 100 * time.Millisecond
)
/*
	f20应该是买1  
f18 是买2
委差是f192
f40是卖1，f38是卖2
f19是买1多少钱

"f127":"文化传媒","f128":"广东板块"
*/

type BkDetails struct {
	Bk string `json:"f127"`
	Region string `json:"f128"`
}
type BkResponse struct {
	Rc int `json:"rc"`
	Rt int `json:"rt"`
	Svr int `json:"svr"`
	Lt int `json:"lt"`
	Full int `json:"full"`
	Data *BkDetails `json:"data"`
}


func GetBkCode(code string) string {
	BkUrl := BkUrlPart + GetSecid(code)
	var content []byte
	cacheKey := code + "bk"
	contentStr, had := db.Get(cacheKey)
	if had {
		content = []byte(contentStr)
	}else {
		timeSpend := time.Since(lastWebVisitTime)
		fmt.Println("timespend:", timeSpend," for web visit time interval:", webVisitInterval, "last:", lastWebVisitTime)
		if timeSpend < webVisitInterval{
			fmt.Println("sleep for web visit time interval:", webVisitInterval)
			time.Sleep(webVisitInterval)
		}
// 		fmt.Println(BkUrl)
		resp, err := http.Get(BkUrl)
		lastWebVisitTime = time.Now()
		if err != nil {
			fmt.Println("http.get error", BkUrl, err)
			return ""
		}
		defer resp.Body.Close()
// 	 	fmt.Println(resp)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil.readAll error", resp.Body)
			return ""
		}

		if resp.StatusCode > 299 {
			fmt.Println("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
			return ""
			//return body, errors.New(strconv.Itoa(resp.StatusCode))
		}
		content = body
		db.Insert(cacheKey, string(body))
	}
	var resp2 BkResponse
// 	fmt.Println(content)
	err := json.Unmarshal(content, &resp2)
	if err != nil {
		fmt.Println("json unmarshal error : " + err.Error())
		return ""
	}
// 	fmt.Println(resp2.Data)

// 	fmt.Println(code, resp2.Data.Bk, resp2.Data.Region)
	return resp2.Data.Bk
}
func GetBk(o *TimeObject) string{
	return GetBkCode(o.Code)
}
