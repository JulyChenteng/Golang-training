package main

import (
	"./fileop"
	"./search"
	"fmt"
	"time"
)

var (
	FileName = "E:\\CodeBase\\Solution\\dataprocess\\files\\score.txt"
)

func main() {
	start := time.Now()

	fileop.Split(FileName)

	count := search.FindScoresByGrade(666)
	fmt.Println("total : ", count)

	rank := search.GetRankingByStats("20180601526")
	fmt.Println("Ranking: ", rank)

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}
