package main

import (
	"TextSearch/app"
	"flag"
	"fmt"
	"os"
	"time"
)

var path = flag.String("path", "", "查找的根路径")
var key = flag.String("key", "", "要查找的key")

func main() {

	dir, _ := os.Getwd()
	filename := dir + "/mr-out"
	_, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	start := time.Now()

	flag.Parse()
	fmt.Printf("find path: %s\n", *path)
	fmt.Printf("find key: %s\n", *key)

	app.Start(*path, *key)

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("程序总耗时: %s\n", elapsed)
}
