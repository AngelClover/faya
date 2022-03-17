package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	read = true
	write = true
	clear = false
	doNotCheckTime = false
)

type cacheUnit struct {
	Content string "json:ctt"
	Tm time.Time "json:time"
}

func client() {

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
}

func getInstance() *redis.Client {
	var inc *redis.Client
	if inc == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		inc = rdb
	}
	return inc
}

/*
api level cache
*/
func Get(key string) (string, bool) {
	if read == false {
		return "", false
	}
	// 	time.Sleep(1 * time.Second)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	val, err := rdb.Get(ctx, key).Result()
	//exsistance
	if err == redis.Nil {
		fmt.Println("key:", key, "does not exist")
		return "", false
	} else if err != nil {
		fmt.Println("cannot found out", key, "should panic", err)
// 		panic(key + "not found")
		//return "", false
	}
//  	fmt.Println("cache key:", key, "content:", val)

	if len(val) < 3{
		fmt.Println("cache key:", key, "content:", val, "len:", len(val))
		return  "", false
	}
	
	//valid judge
	var cc cacheUnit
	valb := []byte(val)
	err = json.Unmarshal(valb, &cc)
	if err != nil {
		fmt.Println("dbdata unmarshal error", err.Error)
		fmt.Println(val)
		return "", false
	}
	doExpireJudge := true
	if strings.Index(key, "bk") > 4 {
		doExpireJudge = false
	}
	if doNotCheckTime {
		doExpireJudge = false
	}
	// expire judge
// 	fmt.Println("get content time:", cc.Tm)
	if doExpireJudge {
		now := time.Now()
		yn, mn, dn := now.Date()
		yr, mr, dr := cc.Tm.Date()
		loc := time.FixedZone("UTC+8", +8*60*60)

		startTime := time.Date(yn, mn, dn, 9, 15, 0, 0, loc)
		endTime := time.Date(yn, mn, dn, 15, 0, 0, 0, loc)
		
		//cc.Tm < now < starTime
		if startTime.After(now) {
			noww := now.AddDate(0, 0, -1)
			//yn, mn, dn = noww.Date()
			endd := endTime.AddDate(0, 0, -1)

			if (cc.Tm.After(noww) && cc.Tm.After(endd)) {
// 				fmt.Println(key, "expire 1", noww, "<", cc.Tm)
// 				fmt.Println(key, "expire 1", endd, "<", cc.Tm)
				return cc.Content, true
			}
		}

		//(data time, now) (yesterday noon, morning)


		if yn != yr || mn != mr || dn != dr {
			fmt.Println("db find content for key:", key, " but it is the date:", cc.Tm)
			return "", false
		}
		//add same day expire judge
		//cc.Tm < endTime < now
		if endTime.After(cc.Tm) && now.After(endTime) {
				fmt.Println(key, "expire 3")
			return "", false
		}
	}

	// read correct from db
//  	fmt.Println("get ", key, "from db")
	return cc.Content, true
}

func Insert(key string, val string) {
	if write == false {
		return 
	}
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	defer rdb.Close()
	if clear == true {
		err := rdb.Set(ctx, key, "", 0).Err()
		if err != nil {
			fmt.Println("err when clear key:", key)
			panic(err)
		}
	}

	var cc cacheUnit
	cc.Content = val
	cc.Tm = time.Now()
	cacheValueb, err := json.Marshal(cc)
	if err != nil {
		fmt.Println("cacheValue marcha err", err)
		panic(err)
	}

	cacheValue := string(cacheValueb)
	err = rdb.Set(ctx, key, cacheValue, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("set ", key, "succ at time:", cc.Tm)
}
