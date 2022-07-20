/**
* @Author:Dijiang
* @Description:
* @Date: Created in 18:15 2022/7/7
* @Modified By: Dijiang
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func wxRedPacketSplit(num, m int) []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, m)
	i := 0
	for ; i < m-1; i++ {
		left := num - m + i
		if left <= 0 {
			break
		}
		res[i] = rand.Intn(left) + 1
		num -= res[i]
	}
	for i < m-1 {
		res[i] = 1
		num -= res[i]
		i++
	}
	res[m-1] = num
	return res
}

func main() {
	res := wxRedPacketSplit(4, 4)
	for _, v := range res {
		fmt.Println(v)
	}
}
