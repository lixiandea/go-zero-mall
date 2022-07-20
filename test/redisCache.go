/**
* @Author:Dijiang
* @Description:
* @Date: Created in 15:06 2022/7/5
* @Modified By: Dijiang
 */

package main

import (
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	sqlConn     *sqlx.SqlConn
	redisClient *redis.Client
)

func main() {
	res, err := hex.DecodeString("00000000")
	if err != nil {
		return
	}
	print(len(res))
	//print(hex.DecodeString("00000000"))
}
