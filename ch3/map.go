package main

import (
	"fmt"
)

func zeroMap() {
	var nilMap map[string]int
	fmt.Println(nilMap)
	fmt.Println(nilMap == nil)
	totalWins := map[string]int{}
	fmt.Println(totalWins)
	totalWins["hello"] = 22
	totalWins["world"] = 11
	totalWins["hello"]++
	fmt.Println(totalWins)
	delete(totalWins, "hello")

	teams := map[string][]string{
		"joy":  []string{"hello", "world"},
		"nice": []string{"ni", "ce"},
	}
	fmt.Println(teams)
}

func okMapPrac() {
	var strIntMap = map[string]int{
		"hello": 5,
		"world": 2,
	}
	v, ok := strIntMap["hello"]
	fmt.Println(v, ok)
	v, ok = strIntMap["world"]
	fmt.Println(v, ok)
	v, ok = strIntMap["noValue"]
	fmt.Println(v, ok)
}

// set은 없지만 type을 키로 bool 으로 초기화 한 맵을 만듬 (boolean의 제로값 = false)
func makeMapLikeSet() {
	var intBooleanMap = map[int]bool{}
	var intArr = []int{5, 3, 1, 1, 1, 1, 1, 1, 56, 6, 78, 99}
	for _, v := range intArr {
		intBooleanMap[v] = true
	}
	fmt.Println(len(intBooleanMap), len(intArr))
	fmt.Println(intArr[5])
	fmt.Println(intBooleanMap[99])
	fmt.Println(intBooleanMap[98])
}

func useStructInsteadBoolean() {
	intSet := map[int]struct{}{}
	vals := []int{5, 1, 2, 3, 45, 6, 5, 11, 1, 1, 2, 3, 2}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
