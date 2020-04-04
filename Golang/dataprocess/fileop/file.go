package fileop

import (
	"../defs"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//打开文件，如果不存在就创建，打开的模式是可读可写，权限是644
func Openfile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return nil
	}

	return file
}

//统计单个文件中记录的总数
func GetRecordSize(fileName string) uint64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error!")
		return 0
	}

	var count uint64
	reader := bufio.NewReader(file)
	for {
		if _, err := reader.ReadString('\n'); err != nil {
			break
		}

		count++
	}

	return count
}

//返回一个map，以成绩为key，[]defs.Score为值
func GetScoreMap(fileName string) *map[int][]defs.Score {
	inputfile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return nil
	}

	defer inputfile.Close()

	scoremap := make(map[int][]defs.Score, 4096)
	reader := bufio.NewReader(inputfile)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		str := strings.Trim(input, "\n")

		s := strings.Fields(str)
		grade, _ := strconv.Atoi(s[1])
		score := defs.Score{
			StuID: s[0],
			Grade: grade,
		}

		scoremap[grade] = append(scoremap[grade], score)
	}

	return &scoremap
}

func GetAllFile(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Get all file error!")
		return nil
	}

	return files
}

func removeAllFile(dir string) {
	files := GetAllFile(dir)

	for _, f := range files {
		os.Remove(dir + f.Name())
	}
}

//获取文件所有内存，保存至slice中返回
func GetScores(fileName string) defs.Scores {
	inputfile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return nil
	}

	defer inputfile.Close()

	scores := make(defs.Scores, 0, 4096)
	reader := bufio.NewReader(inputfile)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		str := strings.Trim(input, "\n")

		s := strings.Fields(str)
		grade, _ := strconv.Atoi(s[1])
		score := defs.Score{
			StuID: s[0],
			Grade: grade,
		}

		scores = append(scores, score)
	}

	return scores
}
