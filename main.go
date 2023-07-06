package main

import (
	"fmt"
	"github.com/InfoProspectors/7fenglou/utils"
	"github.com/joho/godotenv"
	// "io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	BaseURL   = "https://7fenglou.com/api/v3/thread.list"
	PageLimit = 50 // 每页返回的最大数据量
)

func main() {
	weitrue := false
	// 从环境变量获取运行时间，默认为30分钟
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	timeoutStr := os.Getenv("timeoutStr")
	if timeoutStr == "" {
		timeoutStr = "30"
	}

	// 检查运行时间是否包含后缀m
	if !strings.HasSuffix(timeoutStr, "m") {
		timeoutStr += "m"
	}

	// 解析运行时间
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatal("无效的运行时间:", err)
	}

	fmt.Printf("爬虫所需运行时间: %s\n", timeout.String())
	// 定义开始时间
	startTime := time.Now()

	// 如果目录不存在，则创建存储JSON文件的目录
	err = os.MkdirAll("data", 0755)
	if err != nil {
		log.Fatal(err)
	}
	// 创建运行日志文件目录
	err = os.MkdirAll("log", 0755)
	if err != nil {
		log.Fatal(err)
	}
	// 获取data/文件夹下的文件数量
	// fileList, err := ioutil.ReadDir("data/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fileCount := len(fileList)

	fileCount := 0
	// 创建运行日志文件
	runLogFile, err := os.OpenFile("log/run.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer runLogFile.Close()
	runLogger := log.New(runLogFile, "", log.Ldate|log.Ltime)

	// 创建成功日志文件
	successLogFile, err := os.OpenFile("log/success.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer successLogFile.Close()
	successLogger := log.New(successLogFile, "", log.Ldate|log.Ltime)

	// 创建错误日志文件
	errLogFile, err := os.OpenFile("log/err.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer errLogFile.Close()
	errLogger := log.New(errLogFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	// 以文件数量作为起始索引
	index := fileCount + 1
	url := fmt.Sprintf("%s?perPage=%d&page=%d", BaseURL, PageLimit, index)

	// 从API获取数据
	response, err := utils.FetchData(url)
	if err != nil {
		errLogger.Println("获取数据时出错:", err)
	}

	// 将数据保存为JSON文件
	filePath := fmt.Sprintf("data/7fenglou%d.json", index)

	// 在保存数据前添加运行日志记录
	runLogger.Printf("开始保存数据，URL：%s，文件路径：%s\n", url, filePath)

	err = utils.SaveJSONData(response, filePath)
	if err != nil {
		utils.WriteErrorToFile(err, url, filePath)
	} else {
		fmt.Println("当前请求的URL为：", url)
	}

	pageData := response.Data.PageData
	for {
		select {
		case <-time.After(timeout):
			fmt.Println("达到超时时间，退出循环")
			break
		default:
			// 等待一段时间以避免超过API的请求限制
			time.Sleep(1 * time.Second)

			// 增加索引
			index++

			// 使用新索引更新URL
			url = fmt.Sprintf("%s?perPage=%d&page=%d", BaseURL, PageLimit, index)
			// 从下一页获取数据
			response, err = utils.FetchData(url)
			if err != nil {
				errLogger.Println("获取数据时出错:", err)
				continue // 跳到下一次迭代
			}

			weitrue = response.Data.CurrentPage > response.Data.TotalPage
			if weitrue {
				fmt.Printf("当前页 %d/%d，是否为最后一页：%v\n", response.Data.CurrentPage, response.Data.TotalPage, response.Data.CurrentPage >= response.Data.TotalPage)
				break // 退出循环
			}
			// 检查响应和数据是否满足条件
			if response == nil || response.Data.PageData == nil || response.Data.NextPageUrl == "" ||
				(len(response.Data.PageData) == 0 && response.Data.PageLength == 0 || response.Data.CurrentPage >= response.Data.TotalPage) {
				break // 退出循环
			}
			// 将新页面的数据追加到现有数据中
			pageData = append(pageData, response.Data.PageData...)

			// 将数据保存为JSON文件
			filePath := fmt.Sprintf("data/7fenglou%d.json", index)

			// 检查文件是否存在
			if _, err := os.Stat(filePath); err == nil {
					fmt.Println("文件已存在，跳过保存：", filePath)
					continue // 跳到下一次迭代
			}
			err = utils.SaveJSONData(response, filePath)
			if err != nil {
				utils.WriteErrorToFile(err, url, filePath)
				continue // 跳到下一次迭代
			}
			// 在保存数据后添加成功日志记录
			successLogger.Printf("数据保存成功，URL：%s，文件路径：%s\n", url, filePath)
			// runLogger.Printf("进入图片保存逻辑，URL：%s，文件路径：%s\n", url, filePath)
			// err := readJSONAndSaveImages(filePath)
			if err != nil {
				fmt.Println(err)
			}
			// 计算进度百分比
			progress := float64(response.Data.CurrentPage-1) * float64(response.Data.PerPage) / float64(response.Data.TotalCount) * 100
			currentTime := time.Now()
			elapsedTime := currentTime.Sub(startTime)
			// 打印当前页和进度
			fmt.Printf("从第%d页获取数据（%.2f%%），当前请求的URL为：%s\n", response.Data.CurrentPage, progress, url)
			fmt.Printf("当前时间: %s, 运行时间: %s\n", currentTime.Format("2006-01-02 15:04:05"), elapsedTime.String())

			// 估算剩余时间
			remainingTime := timeout - elapsedTime
			fmt.Printf("预计剩余时间: %s\n", remainingTime.String())
		}
		elapsedTime := time.Since(startTime)
		if elapsedTime > timeout || weitrue {
			break // 退出循环
		}
	}
	fmt.Println("Data saved successfully.")
}
