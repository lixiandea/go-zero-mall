/**
* @Author:Dijiang
* @Description:
* @Date: Created in 18:11 2022/8/4
* @Modified By: Dijiang
 */

package main

type rateLimiter struct {
	limiters map[string]chan int
}

var globalLimiter *rateLimiter

func init() {
	globalLimiter = &rateLimiter{limiters: make(map[string]chan int, 16)}
	go func() {
		for _, ch := range globalLimiter.limiters {
			ch <- 1
		}
	}()
}

func NewLimiter(key string, rateNum int) {
	globalLimiter.limiters[key] = make(chan int, rateNum)
}
