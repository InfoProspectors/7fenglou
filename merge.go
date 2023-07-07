package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	dirPath := "data"               // 目录路径
	pattern := `^7fenglou\d+\.json` // 文件名匹配的正则表达式
	logFilePath := "log/file.log"   // 日志文件路径

	// 获取目录下的所有文件名
	fileNames, err := listFiles(dirPath)
	if err != nil {
		fmt.Printf("Failed to list files: %v\n", err)
		return
	}

	// 正则表达式匹配文件名
	matchingFiles := make([]string, 0)
	for _, fileName := range fileNames {
		match, _ := regexp.MatchString(pattern, fileName)
		if match {
			matchingFiles = append(matchingFiles, fileName)
		}
	}

	// 按文件名排序
	//sort.Strings(matchingFiles)

	// 打开日志文件
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// 将文件名写入日志文件
	for i, fileName := range matchingFiles {
		_, err := logFile.WriteString(fmt.Sprintf("File %d: %s\n", i+1, fileName))
		if err != nil {
			fmt.Printf("Failed to write to log file: %v\n", err)
			return
		}
	}

	fmt.Println("File names have been written to file.log.")
}

// 列出目录下的所有文件名
func listFiles(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}
