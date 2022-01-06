package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)




func RealtimeMinCode(code string) []*MinUnit {
// 	fmt.Println("list.RealtimeMinCode(", code)

	if !IsOpeningTime() {
		fmt.Println("called the wrong api, it is not opening hour now")
		return nil
	}

	timeSpend := time.Since(lastRealtimeApiVisitTime)
// 	fmt.Println("timespend:", timeSpend," for realtime api visit time interval:", realtimeApiInterval, "last:", lastRealtimeApiVisitTime)
	if timeSpend < realtimeApiInterval{
// 		fmt.Println("sleep for realtime api visit time interval:", realtimeApiInterval)
		time.Sleep(realtimeApiInterval - timeSpend)
	}

	MinUrl := MinUrlPart1 + GetSecid(code) + MinUrlPart2

	resp, err := http.Get(MinUrl)
	// 	 	fmt.Println(resp)

	lastRealtimeApiVisitTime = time.Now()
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
	
// 	fmt.Println(content)
	var resp2 MinResponse
	err = json.Unmarshal(body, &resp2)

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
