package app

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Inter chan KeyValue
var Stop chan struct{}

func Start(root string, key string) {
	start := time.Now()
	files := FilterDirs(root)
	fmt.Println(len(files))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("搜索文件耗时: %s\n", elapsed)

	Inter = make(chan KeyValue)
	Stop = make(chan struct{})

	go ToFile()

	for _, filename := range files {
		wg.Add(1)
		go Map(filename)
	}
	wg.Wait()
	close(Stop)
}
