package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	lastRealtimeApiVisitTime = time.Now()
	realtimeApiInterval = 1000 * time.Millisecond
	mockOpeningTime = true
)

//realtime api need no cache
func GetRealtimeList() []*TimeObject {
// 	fmt.Println("list.GetRealtimeList(")


	if !mockOpeningTime && !IsOpeningTime() {
		fmt.Println("called the wrong api, it is not opening hour now")
		return nil
	}



	timeSpend := time.Since(lastRealtimeApiVisitTime)
// 	fmt.Println("timespend:", timeSpend," for realtime api visit time interval:", realtimeApiInterval, "last:", lastRealtimeApiVisitTime)
	if timeSpend < realtimeApiInterval{
// 		fmt.Println("sleep for realtime api visit time interval:", realtimeApiInterval)
		time.Sleep(realtimeApiInterval - timeSpend)
	}


	resp, err := http.Get(listUrl)
	lastRealtimeApiVisitTime = time.Now()

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

	var resp2 TimeListWrapper
	err = json.Unmarshal(body, &resp2)
	if err != nil {
		fmt.Println("json unmarshal error : " + err.Error())
		return nil
	}

// 	writeTime := gtime.Now()
// 	contentTime, err := getDataTime(writeTime)
// 	fmt.Println("get content time:", contentTime)

	for _, obj := range resp2.Data.Diff {
		//fmt.Println(obj)
		test(obj)
	}
// 	fmt.Println("get list length:", len(resp2.Data.Diff))
	return resp2.Data.Diff

}

//TODO: optimizing
func GetRealtimeInfo(li []*TimeObject) []*TimeObject{
	l := GetRealtimeList()
	ret := make([]*TimeObject, 0)
	for _,o := range li {
		var j* TimeObject
		found := false
		for _,oo := range l {
			if o.Code == oo.Code {
				j = oo
				found = true
				break
			}
		}
		if found {
			ret = append(ret, j)
		}
	}
	return ret
}
