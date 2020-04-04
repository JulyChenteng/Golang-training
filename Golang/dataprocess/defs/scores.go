package defs

import "strings"

type Scores []Score

func (scores Scores) Len() int {
	return len(scores)
}

func (scores Scores) Swap(i, j int) {
	scores[i], scores[j] = scores[j], scores[i]
}

//按成绩从高到低排序，成绩相等则按照学号从小到大排序
func (scores Scores) Less(i, j int) bool {
	if scores[i].Grade != scores[j].Grade {
		return scores[i].Grade > scores[j].Grade
	} else {
		return strings.Compare(scores[i].StuID, scores[j].StuID) <= 0
	}
}

////根据成绩排序
//type SortByGrade struct {
//	Scores
//}
//
////按照成绩降序排列
//func (sort SortByGrade) Less(i, j int) bool {
//	return sort.Scores[i].Grade > sort.Scores[j].Grade
//}
//
//根据StuID排序
type SortByStuID struct {
	Scores
}

//按学号升序排列
func (sort SortByStuID) Less(i, j int) bool {
	return strings.Compare(sort.Scores[i].StuID, sort.Scores[j].StuID) <= 0
}
