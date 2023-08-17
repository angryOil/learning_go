package main

import "fmt"

func loopPrac() {
	forRange()
	forRangeUseOnlyKey()
	useOnlyValue()
	mapLoop()
	stringLoop()
	modifyInLoop()
	loopLabeling()
	forRange2()
}

func forRange() {
	evenVals := []int{2, 4, 6, 8, 12, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func forRangeUseOnlyKey() {
	dbNamesMap := map[string]bool{
		"maria":    true,
		"mysql":    false,
		"postgres": true,
	}
	for k := range dbNamesMap {
		fmt.Println(k)
	}
}

func useOnlyValue() {
	nums := []int{1, 4, 55, 6, 8, 3, 4, 6, 1}
	for _, v := range nums {
		fmt.Println(v)
	}
}

func mapLoop() {
	someMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	for i := 0; i < 4; i++ {
		fmt.Println("loop", i)
		for k, v := range someMap {
			fmt.Println(k, v)
		}
	}
}

func stringLoop() {
	samples := []string{"hello", "world", "joy", "온!!", "angryOil"}

	for _, sample := range samples {
		fmt.Println(sample)
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}

func modifyInLoop() {
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, v := range nums {
		v *= 3
	}
	fmt.Println(nums)
}

func loopLabeling() {
	samples := []string{"one", "two", "three"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'o' {
				continue outer
			}
		}
	}
}

func forRange2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range nums {
		if i == 2 {
			continue
		}
		if i == len(nums)-2 {
			break
		}
		fmt.Println(i, v)
	}
}
