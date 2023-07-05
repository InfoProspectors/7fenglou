package utils
import (
	"github.com/InfoProspectors/7fenglou/types"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)
// fetchData从API中获取数据
func fetchData(url string) (*types.Response, error) {
	// 创建一个 http.Client 对象，用于发送请求
	client := &http.Client{
		Timeout: 30 * time.Second, // 设置超时时间为 30 秒
	}

	// 创建一个新的 http.Request 对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置随机的 User-Agent 请求头
	req.Header.Set("User-Agent", getRandomUserAgent())

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 响应
	response := &types.Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// saveJSONData将数据保存到JSON文件
func saveJSONData(data interface{}, filePath string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// createExcelSheet创建一个具有列名的Excel表格
func createExcelSheet(f *excelize.File, sheetName string, columns []string) error {
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	for i, column := range columns {
		cell := fmt.Sprintf("%s%d", string('A'+i), 1)
		f.SetCellValue(sheetName, cell, column)
	}

	return nil
}

// insertDataIntoExcel将数据插入到Excel文件中
//func insertDataIntoExcel(f *excelize.File, sheetName string, pageData []PageData) error {
//	row := 2
//	for _, data := range pageData {
//		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), data.ThreadId)
//		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), data.PostId)
//		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), data.UserId)
//		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), data.CategoryId)
//		// ... 继续设置其他单元格的值
//
//		row++
//	}
//
//	return nil
//}

// 获取随机 User-Agent
func getRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
		// 添加更多 User-Agent
	}
	return userAgents[rand.Intn(len(userAgents))]
}

// 获取随机 Accept-Language
func getRandomLanguage() string {
	languages := []string{
		"en-US,en;q=0.9",
		"zh-CN,zh;q=0.9",
		// 添加更多语言
	}
	return languages[rand.Intn(len(languages))]
}

// saveExcelFile保存Excel文件
func saveExcelFile(f *excelize.File, filePath string) error {
	err := f.SaveAs(filePath)
	if err != nil {
		return err
	}
	return nil
}

// writeErrorToFile将错误写入文件
func writeErrorToFile(err error, url string, fileName string) {
	file, err := os.OpenFile("log/err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将错误消息、URL和文件名写入文件
	errorMsg := fmt.Sprintf("[%s] %s\nURL: %s\nFile: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error(), url, fileName)
	_, err = file.WriteString(errorMsg)
	if err != nil {
		log.Fatal(err)
	}
}

// func readJSONAndSaveImages(filePath string) error {
// 	fmt.Println(filePath)
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return fmt.Errorf("无法打开文件：%s，错误：%s", filePath, err.Error())
// 	}
// 	defer file.Close()

// 	var data JSONData
// 	err = json.NewDecoder(file).Decode(&data)
// 	if err != nil {
// 		return fmt.Errorf("解析JSON数据时出错：%s", err.Error())
// 	}

// 	for _, item := range data.Data.PageData {
// 		for _, image := range item.Content.Indexes["101"].Body {
// 			imageURL := image.URL
// 			thumbURL := image.ThumbURL

// 			// 获取文件名
// 			fileName := filepath.Base(imageURL)

// 			// 创建目标文件夹
// 			targetDir := "data" + strings.TrimSuffix(filepath.Dir(imageURL), "/")
// 			err = os.MkdirAll(targetDir, os.ModePerm)
// 			if err != nil {
// 				return fmt.Errorf("无法创建目标文件夹：%s，错误：%s", targetDir, err.Error())
// 			}

// 			// 下载原图并保存
// 			err = downloadAndSaveImage(imageURL, filepath.Join(targetDir, fileName))
// 			if err != nil {
// 				return fmt.Errorf("保存原图时出错：%s", err.Error())
// 			}

// 			// 下载缩略图并保存
// 			thumbFileName := "thumb_" + fileName
// 			err = downloadAndSaveImage(thumbURL, filepath.Join(targetDir, thumbFileName))
// 			if err != nil {
// 				return fmt.Errorf("保存缩略图时出错：%s", err.Error())
// 			}
// 		}
// 	}

// 	return nil
// }

func downloadAndSaveImage(url string, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("获取图片数据时出错：%s", err.Error())
	}
	defer resp.Body.Close()

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取图片数据时出错：%s", err.Error())
	}

	err = ioutil.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return fmt.Errorf("保存图片时出错：%s", err.Error())
	}

	return nil
}
