package search

import (
	"fmt"
	"testing"

	"../fileop"
)

func TestFindScoresByGrade(t *testing.T) {
	count := FindScoresByGrade(666)

	if count != 0 {
		fmt.Println("total count : ", count)
		t.Log("success!")
	} else {
		t.Error("failed")
	}
}

func TestGetRankingByStats(t *testing.T) {
	fileop.Split("E:\\CodeBase\\Solution\\dataprocess\\files\\score.txt")

	rank := GetRankingByStats("20180601526")
	if rank != 0 {
		fmt.Println("rankings: ", rank)
		t.Log("Success!")
	} else {
		t.Error("Failed!")
	}
}

func TestGetRankingByMerge(t *testing.T) {
	rank := GetRankingByMerge("20180601526")
	if rank != 0 {
		fmt.Println("rankings: ", rank)
		t.Log("Success!")
	} else {
		t.Error("Failed!")
	}
}

func TestSort(t *testing.T) {
	Sort()
}

func TestMerge(t *testing.T) {
	Merge()
}

func TestGetSize(t *testing.T) {
	fmt.Println("total size: ", fileop.GetTotalSize())
	fmt.Println("result size: ", fileop.GetRecordSize("E:\\CodeBase\\Solution\\dataprocess\\files\\result.txt"))
}
