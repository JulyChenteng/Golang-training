package fileop

import (
	"fmt"
	"testing"
)

//func BenchmarkSplit(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Split("E:\\CodeBase\\Solution\\score.txt")
//	}
//}

func TestSplit(t *testing.T) {
	Split("E:\\CodeBase\\Solution\\dataprocess\\files\\score.txt")
}

func TestGetRecordSize(t *testing.T) {
	if size := GetRecordSize("E:\\CodeBase\\Solution\\dataprocess\\files\\score.txt"); size != 0 {
		fmt.Println(size)
		t.Log("success")
	} else {
		t.Fail()
	}
}

func TestGetTotalSize(t *testing.T) {
	if size := GetTotalSize(); size != 0 {
		fmt.Println("total size: ", size)
		t.Log("Success")
	} else {
		t.Fail()
	}
}
