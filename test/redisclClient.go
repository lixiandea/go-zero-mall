/**
* @Author:Dijiang
* @Description:
* @Date: Created in 12:59 2022/6/29
* @Modified By: Dijiang
 */

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	if rdb == nil {
		panic("client start error")
	}
}
func main() {
	defer rdb.Close()
	ctx := context.Background()
	rdb.Set(ctx, "lixiande", "whatfuck", 10000000000)
	res, err := rdb.Get(ctx, "lixiande").Result()
	if err != nil {
		println("err %v", err.Error())
	}
	println(res)
	duration, err := rdb.TTL(ctx, "lixiande").Result()
	println(duration)
	for i := 0; i < 100; i++ {
		rdb.ZAdd(ctx, "ranktest", &redis.Z{Score: float64(i), Member: fmt.Sprintf("test%d", i)})
	}
	result, err := rdb.ZRange(ctx, "ranktest", 0, 100).Result()
	for _, v := range result {
		println(v)
	}
	zs, err := rdb.ZRangeWithScores(ctx, "ranktest", 0, 100).Result()
	if err != nil {
		println("err %v", err.Error())

	}
	for _, v := range zs {
		println("%s, %s", v.Score, v.Member)
	}

}
