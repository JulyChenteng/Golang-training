package search

import (
	"fmt"
	"os"
	"sync"

	"../defs"
	"../fileop"
	"../workers"
	"bufio"
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

const (
	MAXGRADE = 1000
	POOLSIZE = 100

	RESULT_DIR  = "E:\\CodeBase\\Solution\\dataprocess\\files\\"
	RESULT_FILE = "E:\\CodeBase\\Solution\\dataprocess\\files\\result.txt"

	BUFFERSIZE = 4064
)

func FindScoresByGrade(grade int) int {
	files := fileop.GetAllFile(fileop.OUTPUTDIR)
	result := fileop.Openfile(RESULT_DIR + strconv.Itoa(grade) + ".txt")
	writer := bufio.NewWriter(result)

	defer result.Close()

	var stuCount, size int
	if len(files) > POOLSIZE {
		size = POOLSIZE
	} else {
		size = len(files)
	}

	//开启WorkerPool
	pool := workers.New(size)
	var mutex sync.Mutex

	for _, info := range files {
		pool.Run(workers.Job{
			Data: info,
			Proc: func(data interface{}) {
				if file, ok := data.(os.FileInfo); ok {
					scoreMap := fileop.GetScoreMap(fileop.OUTPUTDIR + file.Name())

					for _, v := range (*scoreMap)[grade] {
						mutex.Lock()
						writer.WriteString(v.ToString())
						if writer.Buffered() >= BUFFERSIZE {
							writer.Flush()
						}
						mutex.Unlock()
					}

					mutex.Lock()
					writer.Flush()
					stuCount += len((*scoreMap)[grade]) //统计所有文件中该成绩的人数
					mutex.Unlock()
				}
			},
		})
	}
	pool.Shutdown() //关闭通道，等待goroutine全部执行完

	return stuCount
}

//////////////////////////////////////////////////////////////////////////////////////////
func GetRankingByStats(stuId string) int {
	var rank = 0

	grade := FindGradeByStuId(stuId)
	if grade != -1 {
		counts := doStats()

		if counts[grade] > fileop.FILESIZE {
			return GetRankingByMerge(stuId)
		}

		for i := grade + 1; i <= MAXGRADE; i++ {
			rank += counts[i]
		}

		//成绩相同的排序再查找
		scores := GetScoresByGrade(grade)
		sort.Sort(scores)

		index := sort.Search(len(scores), func(i int) bool {
			return strings.Compare(scores[i].StuID, stuId) >= 0 && grade == scores[i].Grade
		})

		//fmt.Println(index, " ", scores[index])
		rank += index + 1

		return rank
	} else {
		fmt.Println("please input right StuID!")
		return 0
	}
}

//根据成绩做统计，统计每个成绩存在多少人
func doStats() []int {
	files := fileop.GetAllFile(fileop.OUTPUTDIR)
	counts := make([]int, MAXGRADE+1)

	var mutex sync.Mutex

	var size int
	if len(files) > POOLSIZE {
		size = POOLSIZE
	} else {
		size = len(files)
	}

	//开启WorkerPool
	pool := workers.New(size)

	for _, info := range files {
		pool.Run(workers.Job{
			Data: info,
			Proc: func(data interface{}) {
				if file, ok := data.(os.FileInfo); ok {
					scoreMap := fileop.GetScoreMap(fileop.OUTPUTDIR + file.Name())

					//遍历map，统计人数
					for k, v := range *scoreMap {
						mutex.Lock()
						counts[k] += len(v)
						mutex.Unlock()
					}

				}
			},
		})
	}
	pool.Shutdown() //关闭通道，等待goroutine全部执行完

	return counts
}

//根据成绩查找到所有与此成绩相等的记录
func GetScoresByGrade(grade int) defs.Scores {
	files := fileop.GetAllFile(fileop.OUTPUTDIR)
	score := make([]defs.Score, 0, 1024)

	var size int
	if len(files) > POOLSIZE {
		size = POOLSIZE
	} else {
		size = len(files)
	}

	var mutex sync.Mutex
	pool := workers.New(size)

	for _, info := range files {
		pool.Run(workers.Job{
			Data: info,
			Proc: func(data interface{}) {
				if info, ok := data.(os.FileInfo); ok {
					scoreMap := fileop.GetScoreMap(fileop.OUTPUTDIR + info.Name())

					mutex.Lock()
					score = append(score, (*scoreMap)[grade]...)
					mutex.Unlock()
				}
			},
		})
	}
	pool.Shutdown() //关闭通道，等待goroutine全部执行完

	return score
}

//根据学号查找该学生的成绩
func FindGradeByStuId(stuId string) int {
	scores := fileop.GetScores(fileop.GetAddrByStuId(stuId))

	sort.Sort(defs.SortByStuID{Scores: scores})

	index := sort.Search(len(scores), func(i int) bool {
		return strings.Compare(scores[i].StuID, stuId) >= 0
	})

	if index < len(scores) && strings.Compare(scores[index].StuID, stuId) == 0 {
		//fmt.Println(index, " ", addr, " ", scores[index])
		return scores[index].Grade
	} else {
		return -1
	}
}

////////////////////////////////////////////////////////////////////////////////////////////
func GetRankingByMerge(stuId string) int {
	Sort()
	Merge()

	inputFile, err := os.Open(RESULT_FILE)
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return 0
	}

	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)
	var count = 0

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		count++

		str := strings.Trim(input, "\n")
		s := strings.Fields(str)
		if strings.Compare(stuId, s[0]) == 0 {
			break
		}
	}

	return count
}

func Merge() {
	files := fileop.GetAllFile(fileop.OUTPUTDIR)
	result := fileop.Openfile(RESULT_FILE)
	writer := bufio.NewWriter(result)

	defer result.Close()

	fileList := make([]*os.File, 0, len(files))
	for _, info := range files {
		file, err := os.Open(fileop.OUTPUTDIR + info.Name())
		if err != nil {
			fmt.Println("open file error!")
			return
		}

		fileList = append(fileList, file)
	}

	readerList := make([]*bufio.Reader, 0, len(files))
	for _, f := range fileList {
		readerList = append(readerList, bufio.NewReader(f))
	}

	records := make(defs.Records, 0, len(fileList))
	//先从已排序好文件中取出第一条记录
	for i := 0; i < len(fileList); i++ {
		str, err := readerList[i].ReadString('\n')
		if err != nil {
			continue
		}

		s := strings.Fields(str)
		grade, _ := strconv.Atoi(s[1])
		score := defs.Score{
			StuID: s[0],
			Grade: grade,
		}

		records = append(records, defs.Record{Score: score, FileIndex: i})
	}

	//然后建堆，每次取出堆定元素，再调整
	heap.Init(&records)
	for len(records) != 0 {
		record := heap.Pop(&records) //取出堆顶元素

		if rec, ok := record.(defs.Record); ok {
			writer.WriteString(rec.Score.ToString())
			if writer.Buffered() >= BUFFERSIZE {
				writer.Flush()
			}

			//从该记录所在的文件中再取下一条记录
			input, err := readerList[rec.FileIndex].ReadString('\n')
			if err != nil {
				continue
			}
			str := strings.Trim(input, "\n")

			s := strings.Fields(str)
			grade, _ := strconv.Atoi(s[1])
			score := defs.Score{
				StuID: s[0],
				Grade: grade,
			}

			heap.Push(&records, defs.Record{Score: score, FileIndex: rec.FileIndex})
		}
	}

	writer.Flush()

	for _, f := range fileList {
		f.Close()
	}
}

//func Merge() {
//	files := fileop.GetAllFile(fileop.OUTPUTDIR)
//	result := fileop.Openfile(RESULT_FILE)
//	writer := bufio.NewWriter(result)
//
//	fileList := make([]*os.File, 0, len(files))
//	for _, info := range files {
//		file, err := os.Open(fileop.OUTPUTDIR + info.Name())
//		if err != nil {
//			fmt.Println("open file error!")
//			return
//		}
//
//		fileList = append(fileList, file)
//	}
//
//	readerList := make([]*bufio.Reader, 0, len(files))
//	for  _, f := range fileList {
//		readerList = append(readerList, bufio.NewReader(f))
//	}
//
//	records := make(defs.Records, 0, len(fileList))
//
//	for i := 0; i < len(fileList); i++ {
//		str, err :=  readerList[i].ReadString('\n')
//		if err != nil {
//			continue
//		}
//
//		s := strings.Fields(str)
//		grade, _ := strconv.Atoi(s[1])
//		score := defs.Score{
//			StuID: s[0],
//			Grade: grade,
//		}
//
//		records = append(records, defs.Record{Score: score, FileIndex: i})
//	}
//
//	for len(records) != 0 {
//		sort.Sort(records)
//
//		rec := records[0]
//		records = records[1:]
//
//		writer.WriteString(rec.Score.ToString())
//		writer.Flush()
//
//		str, err :=  readerList[rec.FileIndex].ReadString('\n')
//		if err != nil {
//			continue
//		}
//
//		s := strings.Fields(str)
//		grade, _ := strconv.Atoi(s[1])
//		score := defs.Score{
//			StuID: s[0],
//			Grade: grade,
//		}
//
//		records =  append(records, defs.Record{Score: score, FileIndex: rec.FileIndex})
//	}
//
//	for _, f := range fileList {
//		f.Close()
//	}
//}

//对每个小文件排序
func Sort() {
	files := fileop.GetAllFile(fileop.OUTPUTDIR)

	var size int
	if len(files) > POOLSIZE {
		size = POOLSIZE
	} else {
		size = len(files)
	}

	//开启WorkerPool
	pool := workers.New(size)
	for _, info := range files {
		pool.Run(workers.Job{
			Data: info,
			Proc: func(data interface{}) {
				if info, ok := data.(os.FileInfo); ok {
					scores := fileop.GetScores(fileop.OUTPUTDIR + info.Name())
					sort.Sort(scores)

					inputFile, err := os.OpenFile(fileop.OUTPUTDIR+info.Name(), os.O_WRONLY, 0600)
					if err != nil {
						fmt.Println("open file error !")
						return
					}

					defer inputFile.Close()

					writer := bufio.NewWriter(inputFile)

					for _, v := range scores {
						writer.WriteString(v.ToString())
						if writer.Buffered() > BUFFERSIZE {
							writer.Flush()
						}
					}
					writer.Flush()
				}
			},
		})
	}
	pool.Shutdown() //关闭通道，等待goroutine全部执行完
}
