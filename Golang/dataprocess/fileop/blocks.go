package fileop

import (
	"../workers"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	FILESIZE  = 15000
	POOLSIZE  = 100
	OUTPUTDIR = "E:\\CodeBase\\Solution\\dataprocess\\files\\data\\"

	BUFFERSIZE = 4064
)

var primeList = [...]int{
	53, 97, 193, 389, 769,
	1543, 3079, 6151, 12289, 24593,
	49157, 98317, 196613, 393241, 786433,
	1572869, 3145739, 6291469, 12582917, 25165843,
	50331653, 100663319, 201326611, 402653189, 805306457,
	1610612741, 3221225473, 4294967291,
}
var (
	hashSize int
)

func Split(fileName string) bool {
	hashSize = getHashSize(fileName)
	removeAllFile(OUTPUTDIR)

	files := make([]*os.File, 0, hashSize)
	writers := make([]*bufio.Writer, 0, hashSize)

	for i := 0; i < hashSize; i++ {
		outputFile := Openfile(OUTPUTDIR + strconv.Itoa(i) + ".txt")
		files = append(files, outputFile)
	}

	for i := 0; i < len(files); i++ {
		writers = append(writers, bufio.NewWriter(files[i]))
	}

	pool := workers.New(POOLSIZE)

	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return false
	}

	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	//去除第一行
	_, err = reader.ReadString('\n')
	if err != nil {
		return false
	}

	var mutex sync.Mutex

	//循环按行读取内容，然后交给workerpool去处理
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		pool.Run(workers.Job{
			Data: str,
			Proc: func(data interface{}) {
				if str, ok := data.(string); ok {
					s := strings.Fields(str)

					addr := hash(s[0]) % hashSize
					writer := writers[addr]
					mutex.Lock()
					writer.WriteString(str)
					if writer.Buffered() >= BUFFERSIZE {
						writer.Flush()
					}
					mutex.Unlock()
				}
			},
		})
	}

	pool.Shutdown()

	//将缓冲区剩余的内容写入文件
	for i := 0; i < len(writers); i++ {
		writers[i].Flush()
	}

	for i := 0; i < len(files); i++ {
		files[i].Close()
	}

	return true
}

func hash(str string) int {
	h := 0

	for _, ch := range []byte(str) {
		h = 5*h + int(ch)
	}

	return h
}

func nextPrime(num int) int {
	index := sort.Search(len(primeList), func(i int) bool {
		return primeList[i] >= num
	})
	if index == 0 {
		return primeList[index]
	}

	return primeList[index-1]
}

//求文件分割时采用HASH值
func getHashSize(fileName string) int {
	recordSize := GetRecordSize(fileName)

	return nextPrime(int(recordSize / FILESIZE))
}

//根据学号得到该记录存放的文件名
func GetAddrByStuId(stuId string) string {
	addr := hash(stuId) % hashSize
	return OUTPUTDIR + strconv.Itoa(addr) + ".txt"
}

//统计所有文件的记录个数总数
func GetTotalSize() uint64 {
	files := GetAllFile(OUTPUTDIR)

	var size int
	if len(files) > POOLSIZE {
		size = POOLSIZE
	} else {
		size = len(files)
	}

	var mutex sync.Mutex
	pool := workers.New(size)

	var total uint64
	for _, info := range files {
		pool.Run(workers.Job{
			Data: info,
			Proc: func(data interface{}) {
				if file, ok := data.(os.FileInfo); ok {
					count := GetRecordSize(OUTPUTDIR + file.Name())

					//统计所有文件中该成绩的人数
					mutex.Lock()
					total += count
					mutex.Unlock()
				}
			},
		})
	}
	pool.Shutdown() //关闭通道，等待goroutine全部执行完

	return total
}
