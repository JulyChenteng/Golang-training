package main

import (
	"fmt"
	"sort"
)

/*
import "fmt"

var arr = [...]int{1, 4, 16, 64}

const MAXVALUE = 1024

func getCount(res int) int {
	count := 0

	if res < arr[0] || res == MAXVALUE {
		return count
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if res >= arr[i] {
			count += res / arr[i]
			res = res % arr[i]
		}
	}

	return count
}

func main() {
	total := 1024

	var price int
	fmt.Scanf("%d", &price)
	if price > MAXVALUE || price < 0 {
		panic("Input is illeage!")
	}

	cnt := getCount(total - price)
	fmt.Println(cnt)
}
*/

func main() {
	var totalCnt, avg int
	fmt.Scanf("%d %d", &totalCnt, &avg)

	var arr sort.IntSlice
	var tmp int

	for i := 0; i < totalCnt; i++ {
		fmt.Scanf("%d", &tmp)
		arr = append(arr, tmp)
	}
	sort.Sort(arr)
	fmt.Println(arr)

	totalLen := sum(arr)
	if totalLen/avg == 1 {
	}
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
