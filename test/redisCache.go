/**
* @Author:Dijiang
* @Description:
* @Date: Created in 15:06 2022/7/5
* @Modified By: Dijiang
 */

package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	sqlConn     *sqlx.SqlConn
	redisClient *redis.Client
)

func init() {
	sqlConn := sqlx.NewMysql("root:Cz05180921.@tcp(mysql:3306)/mall@charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
}
