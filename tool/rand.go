package tool

import (
	"math/rand"
	"sort"
	"time"
)

// 返回从[min, max)1个随机整数
func RandInt(min, max int) int {
	if max < min {
		return 0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn((max - min)) + min

	return num
}

// 返回从[min, max)不重复的n个随机从小到大有序整数切片
func RandInts(min, max, n int) []int {
	if max < min || (max-min) < n {
		return nil
	}

	res := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(res) < n {
		num := r.Intn((max - min)) + min

		exist := false
		for _, v := range res {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			res = append(res, num)
		}
	}
	sort.Ints(res)
	return res
}

func RandFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
