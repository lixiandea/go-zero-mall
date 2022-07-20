/**
* @Author:Dijiang
* @Description:
* @Date: Created in 22:49 2022/7/18
* @Modified By: Dijiang
 */

package main

import "sort"

func numberOfPairs(nums []int) []int {
	record := make(map[int]int, 0)
	for _, v := range nums {
		if r, ok := record[v]; ok {
			record[v] = r + 1
		} else {
			record[v] = 1
		}
	}
	res, left := 0, 0
	//result := make([]int, 0)
	for _, v := range record {
		res += v / 2
		left += v % 2
	}
	return []int{res, left}

}

func maximumSum(nums []int) int {
	record := make(map[int][]int)
	for _, v := range nums {
		s := calSum(v)
		if _, ok := record[s]; ok {
			record[s] = append(record[s], v)
		} else {
			record[s] = []int{v}
		}
	}
	res := -1
	for _, v := range record {
		sort.Ints(v)
		if len(v) >= 2 {
			res = max(res, v[len(v)-1]+v[len(v)-2])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func calSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for _, v := range intervals {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			if v[0] > res[len(res)-1][1] {
				res = append(res, v)
			} else {
				res[len(res)-1][1] = max(res[len(res)-1][1], v[1])
			}
		}
	}
	return res
}

func isIsomorphic(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	t2s, s2t := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		if (len(s2t) > 0 && s2t[s[i]] != t[i]) || (len(t2s) > 0 && t2s[t[i]] != s[i]) {
			return false
		}
		s2t[s[i]] = t[i]
		t2s[t[i]] = s[i]
	}
	return true
}

func isSubsequence(s string, t string) bool {
	s, t = t, s
	index1, index2 := 0, 0
	for index2 < len(t) && index1 < len(s) {
		if s[index1] == t[index2] {
			index1++
			index2++
		} else {
			index1++
		}
	}
	if index2 == len(t) {
		return true
	}
	return false
}

func getRow(rowIndex int) []int {
	temp1 := make([]int, rowIndex+1)
	temp2 := make([]int, rowIndex+1)
	curIndex := 0
	for ; curIndex < rowIndex; curIndex++ {
		temp2[0] = 1
		for i := 1; i < curIndex-1; i++ {
			temp2[i] = temp1[i-1] + temp1[i]
		}
		temp2[curIndex] = 1
		temp1, temp2 = temp2, temp1
	}
	return temp1
}
func main() {
	isSubsequence("abc", "ahbgc")
}
