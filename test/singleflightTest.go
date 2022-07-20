/**
* @Author:Dijiang
* @Description:
* @Date: Created in 11:15 2022/7/25
* @Modified By: Dijiang
 */

package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
	"time"
)

var (
	g            singleflight.Group
	ErrCacheMiss = errors.New("cache miss")
	count        = 0
)

func load(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		v, err, _ := g.Do(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
				time.Sleep(100 * time.Millisecond)
				g.Forget(key)
			}
			setCache(key, data)
			return data, nil
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		data = v.(string)
	}
	return data, nil
}

func loadFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

func setCache(key string, data interface{}) {
	fmt.Println("set key ", key, " cache to ", data)
}
func loadFromDB(key string) (string, error) {
	log.Println("query db")
	count++
	//unix := strconv.Itoa(int(time.Now().UnixNano()))
	//if rand.Int31n(10) > 5 {
	//	return "", ErrCacheMiss
	//}
	return "", ErrCacheMiss
	//return unix, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	// 模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := load("key")
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
	print(count)
}
