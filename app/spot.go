package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type KeyValue struct {
	KV map[string][]string
}

func Map(filename string) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filename)
	}

	res := KeyValue{
		KV: make(map[string][]string),
	}
	m := make(map[string]bool)
	words := strings.FieldsFunc(string(content), func(x rune) bool { return !unicode.IsLetter(x) })
	for _, w := range words {
		m[w] = true
	}
	for w := range m {
		res.KV[filename] = append(res.KV[filename], w)
	}
	Inter <- res
}

func ToFile() {
	dir, _ := os.Getwd()
	dir += "/mr-out"

	file, err := os.OpenFile(dir, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for {
		select {
		case kv := <-Inter:
			enc := json.NewEncoder(file)
			enc.Encode(kv)
		case <-Stop:
			return
		}
	}

}

func FilterFile(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".md" || ext == ".go" || ext == ".cpp" || ext == ".c" || ext == ".h" {
		//fmt.Printf("文件类型匹配：%s\n", ext)
		return true
	}
	//fmt.Printf("文件类型不匹配：%s\n", ext)
	return false
}

func FilterDirs(root string) (files []string) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Printf("The path %q does not exist\n", root)
		return
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil
			}
			return err
		}
		if info.IsDir() {
			// 获取最后一个路径元素作为文件夹名
			dirName := filepath.Base(path)
			// 过滤掉特定的系统文件夹
			if dirName == "$RECYCLE.BIN" || dirName == "System Volume Information" {
				return filepath.SkipDir // 跳过该目录
			}
		} else {
			// 打印非目录（即文件）的路径
			// fmt.Println(path)
			if FilterFile(path) {
				files = append(files, path)
			}

		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
	}
	return
}
