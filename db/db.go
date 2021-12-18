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
)

type cacheUnit struct {
	ctx string "json:ctx"
	date time.Date "json:date"
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

func Get(key string) (string, bool) {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("cannot found out", key)
		return "", false
	}
	
	var cc cacheUnit
	err = json.Unmarshal(val, &cc)
	if err != nil {
		fmt.Println("dbdata unmarshal error", err.Error)
		return "", false
	}
	if cc.Date != time.Today() {
		fmt.Println("db find content for key:", key, " but it is the date:", cc.Date)
		return "", false
	}
	return cc.context, true
}
