package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/charlievieth/fastwalk"
)

func FiltFile(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".md" || ext == ".go" || ext == ".cpp" || ext == ".c" || ext == ".h" {
		//fmt.Printf("文件类型匹配：%s\n", ext)
		return true
	}
	//fmt.Printf("文件类型不匹配：%s\n", ext)
	return false
}

func main() {
	// 记录开始时间
	startTime := time.Now()

	// 初始化文件路径切片
	var filePaths []string

	// 定义遍历的根目录
	root := "D:\\workspace"

	// 创建 fastwalk 配置
	conf := fastwalk.Config{
		Follow:     false, // 是否跟随符号链接
		NumWorkers: 8,     // 设置并发数为 8
	}

	// 定义回调函数
	err := fastwalk.Walk(&conf, root, func(path string, de os.DirEntry, err error) error {
		if err != nil {
			//log.Printf("error accessing path %q: %v", path, err)
			return nil // 出错时继续遍历
		}
		if FiltFile(path) {
			// 收集文件路径
			filePaths = append(filePaths, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("fastwalk error: %v", err)
	}

	// 记录结束时间
	endTime := time.Now()

	// 打印开始和结束时间
	fmt.Printf("Start Time: %v\n", startTime)
	fmt.Printf("End Time: %v\n", endTime)

	// 打印运行时间
	fmt.Printf("Traversal completed in %v\n", endTime.Sub(startTime))

	// 打印收集到的文件路径数量
	fmt.Printf("Total files collected: %d\n", len(filePaths))

	// 打印内存使用情况
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory allocated: %v MB\n", memStats.Alloc/1024/1024)
	fmt.Printf("Memory allocated but not yet freed: %v MB\n", memStats.HeapAlloc/1024/1024)
	fmt.Printf("Total memory obtained from system: %v MB\n", memStats.Sys/1024/1024)

	// 如果需要，打印所有收集到的文件路径
	// for _, path := range filePaths {
	// 	fmt.Println(path)
	// }
}
