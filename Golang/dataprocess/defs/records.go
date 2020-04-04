package defs

import "strings"

/*
 * 为归并排序而定义
 */
type Record struct {
	Score
	FileIndex int //用于标识该记录存在于那个文件中
}

type Records []Record

func (records Records) Len() int {
	return len(records)
}

func (records Records) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}

func (records Records) Less(i, j int) bool {
	if records[i].Grade != records[j].Grade {
		return records[i].Grade > records[j].Grade
	} else {
		return strings.Compare(records[i].StuID, records[j].StuID) <= 0
	}
}

func (records *Records) Push(data interface{}) {
	*records = append(*records, data.(Record))
}

func (records *Records) Pop() (data interface{}) {
	n := len(*records)
	data = (*records)[n-1] //返回删除的元素
	*records = (*records)[:n-1]

	return data
}
