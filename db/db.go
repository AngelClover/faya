package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	read = true
	write = true
	clear = false
	doNotCheckTime = true
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

/*
api level cache
*/
func Get(key string) (string, bool) {
	if read == false {
		return "", false
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	val, err := rdb.Get(ctx, key).Result()
	//exsistance
	if err == redis.Nil {
		fmt.Println("key:", key, "does not exist")
		return "", false
	}else if err != nil {
		fmt.Println("cannot found out", key, "should panic")
		return "", false
	}
	fmt.Println("cache key:", key, "content:", val)
	
	//valid judge
	var cc cacheUnit
	valb := []byte(val)
	err = json.Unmarshal(valb, &cc)
	if err != nil {
		fmt.Println("dbdata unmarshal error", err.Error)
		return "", false
	}
	// expire judge
	fmt.Println("get content time:", cc.Tm)
	if doNotCheckTime == false {
		yn, mn, dn := time.Now().Date()
		yr, mr, dr := cc.Tm.Date()
		if yn != yr || mn != mr || dn != dr {
			fmt.Println("db find content for key:", key, " but it is the date:", cc.Tm)
			return "", false
		}
	}
	// read correct from db
	fmt.Println("get ", key, "from db")
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
