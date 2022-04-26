package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	FengdanUrlPart1 = "http://push2.eastmoney.com/api/qt/stock/get?ut=fa5fd1943c7b386f172d6893dbfba10b&invt=2&fltt=2&fields=f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f127,f199,f128,f193,f196,f194,f195,f197,f80,f280,f281,f282,f284,f285,f286,f287,f292&secid="
)



type FengdanUnit struct {
	Sell2Price float64 `json:"f37"`
	Sell2 int `json:"f38"`
	Sell1Price float64 `json:"f39"`
	Sell1 int `json:"f40"`

	Buy1Price float64 `json:"f19"`
	Buy1 int `json:"f20"`
	Buy2Price float64 `json:"f17"`
	Buy2 int `json:"f18"`

	Weicha int `json:"f192"`
	Bk string `json:"f127"`
	Dy string `json:"f128"`
}
type FengdanUnitString struct {
	//f31 -> f40, sell5 -> sell1
	Sell2Price interface{} `json:"f37"`
	Sell2 interface{} `json:"f38"`
	Sell1Price interface{} `json:"f39"`
	Sell1 interface{} `json:"f40"`

	//f11 -> f20, buy5 to buy1
	Buy1 interface{} `json:"f20"`
	Buy1Price interface{} `json:"f19"`
	Buy2 interface{} `json:"f18"`
	Buy2Price interface{} `json:"f17"`

	Weicha interface{} `json:"f192"`
	Bk interface{} `json:"f127"`
	Dy interface{} `json:"f128"`

	Lowest interface{} `json:"f43"`
	Highest interface{} `json:"f44"`
	Close interface{} `json:"f45"`
	Open interface{} `json:"f46"`
	Amount interface{} `json:"f47"`
	Money interface{} `json:"f48"`

	Waipan interface{} `json:"f49"` //waipan
	Liangbi interface{} `json:"f50"` //liangbi
	Uplimit interface{} `json:"f51"`
	Downlimit interface{} `json:"f52"`

	GainPerStock interface{} `json:"f55"` //meigushouyi
	Code interface{} `json:"f57"` 
	Name interface{} `json:"f58"` 
	YesterdayClose interface{} `json:"f60"`
	Avg interface{} `json:"f71"`

	DataTime interface{} `json:"f80"`
	ZongGuBen interface{} `json:"f84"`
	LiutongGuBen interface{} `json:"f85"`
	//`json:"f86"`
	JingZiChanPerStock interface{} `json:"f92"`
	//`json:"f104"`
	//`json:"f105"`
	ZongShiZhi interface{} `json:"f116"`
	LiutongShiZhi interface{} `json:"f117"`


	//f135 - f149
	Zhuliliuru interface{} `json:"f135"`
	Zhuliliuchu interface{} `json:"f136"`
	Zhulijingliuru interface{} `json:"f137"`

	ZhuliliuruTeda interface{} `json:"f138"`
	ZhuliliuruDadan interface{} `json:"f141"`
	ZhuliliuruZhongdan interface{} `json:"f144"`
	ZhuliliuruXiaodan interface{} `json:"f147"`

	ZhuliliuchuTeda interface{} `json:"f139"`
	ZhuliliuchuDadan interface{} `json:"f142"`
	ZhuliliuchuZhongdan interface{} `json:"f145"`
	ZhuliliuchuXiaodan interface{} `json:"f148"`
	
	ZhulijingliuruTeda interface{} `json:"f140"`
	ZhulijingliuruDadan interface{} `json:"f143"`
	ZhulijingliuruZhongdan interface{} `json:"f146"`
	ZhulijingliuruXiaodan interface{} `json:"f149"`

	//f161 -> f177
	//f183 -> f199

}
func GetInt(t *int, s interface{}) {
// 	fmt.Println(s, "type:", reflect.TypeOf(s))
	switch s.(type) {
	case int:
// 		fmt.Println("int")
		tt, ok := s.(int)
		if !ok {
			tt = 0
		}
		*t = tt
	case float64:
// 		fmt.Println("float64")
		tt, ok := s.(float64)
		if !ok {
			tt = 0
		}
		*t = int(tt)
	case string:
// 		fmt.Println("string")
		*t = 0
	default:
// 		fmt.Println("default")
		*t = 0
	}
}
func GetFloat(t *float64, s interface{}) {
// 	fmt.Println(s, "type:", reflect.TypeOf(s))
	switch s.(type) {
	case int:
// 		fmt.Println("int")
		tt, ok := s.(int)
		if !ok {
			tt = 0
		}
		*t = float64(tt)
	case float64:
// 		fmt.Println("float64")
		tt, ok := s.(float64)
		if !ok {
			tt = 0
		}
		*t = tt
	case string:
// 		fmt.Println("string")
		*t = 0
	default:
// 		fmt.Println("default")
		*t = 0
	}
}
func GetString(t *string, s interface{}) {
// 	fmt.Println(s, "type:", reflect.TypeOf(s))
	switch s.(type) {
	case int:
// 		fmt.Println("int")
		tt, ok := s.(int)
		if !ok {
			tt = 0
		}
		*t = strconv.Itoa(tt)
	case float64:
// 		fmt.Println("float64")
		tt, ok := s.(float64)
		if !ok {
			tt = 0.0
		}
		*t = fmt.Sprintf("%v", tt)
	case string:
// 		fmt.Println("string")
		tt, ok := s.(string)
		if !ok {
			tt = ""
		}
		*t = tt
	default:
// 		fmt.Println("default")
		*t = ""
	}
}
func (o *FengdanUnit) UnmarshalJSON(data []byte) error {
// 	fmt.Println("UnmarshalJSON fengdanUnit (")
	type Fengdan FengdanUnit
	var str FengdanUnitString
	if err := json.Unmarshal(data, &str); err != nil{
		return err
	}
	//fmt.Printf("%+v\n", str)
	GetInt(&o.Buy1, str.Buy1)
	GetInt(&o.Buy2, str.Buy2)
	GetInt(&o.Sell1, str.Sell1)
	GetInt(&o.Sell2, str.Sell2)

	GetFloat(&o.Buy1Price, str.Buy1Price)
	GetFloat(&o.Buy2Price, str.Buy2Price)
	GetFloat(&o.Sell1Price, str.Sell1Price)
	GetFloat(&o.Sell2Price, str.Sell2Price)

	GetInt(&o.Weicha, str.Weicha)

	GetString(&o.Bk, str.Bk)
	GetString(&o.Dy, str.Dy)

	return nil
}


type FengdanResponse struct {
	Rc int `json:"rc"`
	Rt int `json:"rt"`
	Svr int `json:"svr"`
	Lt int `json:"lt"`
	Full int `json:"full"`
	Data *FengdanUnit `json:"data"`     // data can be null in return
}

func PrettyPrint(str []byte){
	var prettyJSON bytes.Buffer
	//if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
	if err := json.Indent(&prettyJSON, str, "", "    "); err != nil {
		fmt.Println(err)
    }
	fmt.Println(prettyJSON.String())
}


func FengdanCode(code string) *FengdanUnit{
	FengdanUrl := FengdanUrlPart1 + GetSecid(code)

	timeSpend := time.Since(lastWebVisitTime)
	//fmt.Println("timespend:", timeSpend," for web visit time interval:", webVisitInterval, "last:", lastWebVisitTime)
	if timeSpend < webVisitInterval{
		//fmt.Println("sleep for web visit time interval:", webVisitInterval)
		time.Sleep(webVisitInterval - timeSpend)
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(FengdanUrl)

	lastWebVisitTime = time.Now()
	if err != nil {
		fmt.Println("http.get error", FengdanUrl, err)
		return nil
	}
	defer resp.Body.Close()
 	//fmt.Println(resp)

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

	content := body
 	//fmt.Println(string(content))
	//PrettyPrint(content)


	var resp2 FengdanResponse
	err = json.Unmarshal(content, &resp2)

	if err != nil {
		fmt.Println("parsing fengdan response json unmarshal error : " + err.Error())
		return nil
	}

	fd := resp2.Data

	/*
	var fd FengdanUnit
	if resp2.Data == nil {
		return //ret
	}

	err = json.Unmarshal(content, &fd)
	if err != nil {
		fmt.Println("parsing fengdan unit json unmarshal error : " + err.Error())
		return //nil
	}
	*/
	//fmt.Println(fd)
	return fd

}

