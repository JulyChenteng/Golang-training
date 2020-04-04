package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIMake(t *testing.T) {
	sli1 := IMake([]int{}, 4, 8)
	if sli1.cap > 0 {
		sli1 = IAppend(sli1, 'a')
		printSlice(sli1)
		t.Log("Success")
	} else {
		t.FailNow()
	}

	sli2 := IMake([]string{}, 0, 80)
	if sli2.cap > 0 {
		sli2 := IAppend(sli2, "hello", "world")
		printSlice(sli2)
		t.Log("Success")
	} else {
		t.FailNow()
	}

	printMap()
}

func printSlice(slice MySlice) {
	ptr := *(*reflect.Value)(slice.array)
	array := ptr.Elem()

	//fmt.Println(array.Kind())    //array
	for i := 0; i < slice.len; i++ {
		fmt.Print(array.Index(i), "  ")
	}

	fmt.Println()
	fmt.Println("length: ", slice.len, "capacity: ", slice.cap)
	fmt.Println()
}

func printMap() {
	slice := IMake([]map[int]string{}, 8, 20)

	ptr := *(*reflect.Value)(slice.array)
	array := ptr.Elem()

	//fmt.Println(array.Kind())    //array

	for i := 0; i < slice.len; i++ {
		array.Index(i).Set(reflect.ValueOf(make(map[int]string)))
	}

	mp := map[int]string{
		44: "wangge",
		46: "cat",
	}

	slice = IAppend(slice, map[int]string{
		31: "alice",
		32: "charlie",
	}, map[int]string{
		33: "CT",
		35: "hello",
	}, map[int]string{
		66: "world",
		99: "dsds",
	}, mp)

	for i := 0; i < slice.len; i++ {
		for _, key := range array.Index(i).MapKeys() {
			fmt.Println(key, "->", array.Index(i).MapIndex(key), "  ")
		}
		fmt.Println()
	}

	fmt.Println("length: ", slice.len, "capacity: ", slice.cap)
	fmt.Println()
}

func TestILen(t *testing.T) {
	s1 := IMake([]string{}, 5, 10)

	if ILen(s1) == 5 {
		//printSlice(s1)
		t.Log("success！")
	} else {
		t.Error("failed！")
	}
}

func TestICap(t *testing.T) {
	s1 := IMake([]float32{}, 6, 12)

	if ICap(s1) == 12 {
		//printSlice(s1)
		t.Log("success!")
	} else {
		t.Error("failed!")
	}
}

func TestIAppend(t *testing.T) {
	sli := IMake([]int{}, 4, 8)
	sli1 := IAppend(sli, 1, 3, 4, 5, 8)

	printSlice(sli)
	printSlice(sli1)
}
