package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Commit 结构体表示一个提交
type Commit struct {
	ID        string   `json:"id"`        // 提交 ID
	Message   string   `json:"message"`   // 提交消息
	Author    string   `json:"author"`    // 提交作者
	Timestamp int64    `json:"timestamp"` // 提交时间戳
	Files     []string `json:"files"`     // 修改的文件列表
}

// Data 结构体表示数据对象
type Data struct {
	PageData     []PageData `json:"pageData"`     // 页面数据列表
	CurrentPage  int        `json:"currentPage"`  // 当前页码
	PerPage      int        `json:"perPage"`      // 每页数量
	TotalCount   int        `json:"totalCount"`   // 总数据数量
	FirstPageUrl string     `json:"firstPageUrl"` // 第一页的URL
	NextPageUrl  string     `json:"nextPageUrl"`  // 下一页的URL
	PrePageUrl   string     `json:"prePageUrl"`   // 上一页的URL
	PageLength   int        `json:"pageLength"`   // 当前页数据长度
	TotalPage    int        `json:"totalPage"`    // 总页数
}

type Response struct {
	Code        int    `json:"Code"`
	Message     string `json:"Message"`
	Data        Data   `json:"Data"`
	RequestId   string `json:"RequestId"`
	RequestTime string `json:"RequestTime"`
}

// PageData 结构体表示页面数据
type PageData struct {
	//ThreadId           int    `json:"threadId"`           // 线程ID
	//PostId             int    `json:"postId"`             // 帖子ID
	//UserId             int    `json:"userId"`             // 用户ID
	//CategoryId         int    `json:"categoryId"`         // 分类ID
	//ParentCategoryId   int    `json:"parentCategoryId"`   // 父级分类ID
	//TopicId            int    `json:"topicId"`            // 主题ID
	//CategoryName       string `json:"categoryName"`       // 分类名称
	//ParentCategoryName string `json:"parentCategoryName"` // 父级分类名称
	//Title              string `json:"title"`              // 标题
	//ViewCount          int    `json:"viewCount"`          // 浏览次数
	//IsApproved         int    `json:"isApproved"`         // 是否已审核
	//IsStick            bool   `json:"isStick"`            // 是否置顶
	//IsDraft            bool   `json:"isDraft"`            // 是否为草稿
	//IsSite             bool   `json:"isSite"`             // 是否为站点
	//IsAnonymous        bool   `json:"isAnonymous"`        // 是否匿名
	//IsFavorite         bool   `json:"isFavorite"`         // 是否为收藏
	//Price              int    `json:"price"`              // 价格
	//AttachmentPrice    int    `json:"attachmentPrice"`    // 附件价格
	//PayType            int    `json:"payType"`            // 支付类型
	//Paid               bool   `json:"paid"`               // 是否已支付
	//IsLike             bool   `json:"isLike"`             // 是否点赞
	//IsReward           bool   `json:"isReward"`           // 是否打赏
	//CreatedAt          string `json:"createdAt"`          // 创建时间
	//IssueAt            string `json:"issueAt"`            // 发布时间
	//UpdatedAt          string `json:"updatedAt"`          // 更新时间
	//DiffTime           string `json:"diffTime"`           // 时间差
	User struct {
		Nickname string `json:"nickname"` // 用户昵称
	} `json:"user"` // 用户信息
	Group struct {
		// 定义 "group" 对象中的字段，如果在 JSON 中可用。
	} `json:"group"` // 群组信息
	//DisplayTag struct {
	//	IsPrice   bool        `json:"isPrice"`   // 是否含有价格标签
	//	IsEssence bool        `json:"isEssence"` // 是否为精华帖子
	//	IsRedPack interface{} `json:"isRedPack"` // 是否为红包帖子（接口类型）
	//	IsReward  interface{} `json:"isReward"`  // 是否为悬赏帖子（接口类型）
	//	IsVote    bool        `json:"isVote"`    // 是否为投票帖子
	//} `json:"displayTag"` // 显示标签信息

	DisplayTag struct{} `json:"displayTag"`

	//Position struct {
	//	Longitude string `json:"longitude"` // 经度
	//	Latitude  string `json:"latitude"`  // 纬度
	//	Address   string `json:"address"`   // 地址
	//	Location  string `json:"location"`  // 位置
	//} `json:"position"` // 位置信息
	Position struct{}    `json:"position"`
	Ability  interface{} `json:"ability"`
	//Ability  struct {
	//	CanEdit               bool `json:"canEdit"`               // 是否可以编辑
	//	CanDelete             bool `json:"canDelete"`             // 是否可以删除
	//	CanEssence            bool `json:"canEssence"`            // 是否可以设置为精华
	//	CanStick              bool `json:"canStick"`              // 是否可以置顶
	//	CanReply              bool `json:"canReply"`              // 是否可以回复
	//	CanViewPost           bool `json:"canViewPost"`           // 是否可以查看帖子
	//	CanFreeViewPost       bool `json:"canFreeViewPost"`       // 是否可以免费查看帖子
	//	CanViewVideo          bool `json:"canViewVideo"`          // 是否可以查看视频
	//	CanViewAttachment     bool `json:"canViewAttachment"`     // 是否可以查看附件
	//	CanDownloadAttachment bool `json:"canDownloadAttachment"` // 是否可以下载附件
	//} `json:"ability"` // 用户权限
	Content interface{} `json:"content"`

	// Content struct {
	// 	Text    string `json:"text"` // 文本内容
	// 	Indexes map[string]struct {
	// 		TomId     string `json:"tomId"`     // Tom ID
	// 		Operation string `json:"operation"` // 操作
	// 		Body      []struct {
	// 			Id         int    `json:"id"`         // ID
	// 			Order      int    `json:"order"`      // 排序
	// 			Type       int    `json:"type"`       // 类型
	// 			TypeId     int    `json:"typeId"`     // 类型ID
	// 			IsRemote   bool   `json:"isRemote"`   // 是否远程
	// 			IsApproved int    `json:"isApproved"` // 是否已审核
	// 			URL        string `json:"url"`        // URL
	// 			ThumbURL   string `json:"thumbUrl"`   // 缩略图URL
	// 			Attachment string `json:"attachment"` // 附件
	// 			Extension  string `json:"extension"`  // 扩展名
	// 			FileName   string `json:"fileName"`   // 文件名
	// 			FilePath   string `json:"filePath"`   // 文件路径
	// 			FileSize   int    `json:"fileSize"`   // 文件大小
	// 			FileType   string `json:"fileType"`   // 文件类型
	// 			FileWidth  int    `json:"fileWidth"`  // 文件宽度
	// 			FileHeight int    `json:"fileHeight"` // 文件高度
	// 			NeedPay    int    `json:"needPay"`    // 需要付费
	// 		} `json:"body"` // 正文主体
	// 		PriceList []string    `json:"priceList"` // 价格列表
	// 		Plugin    interface{} `json:"_plugin"`   // 插件
	// 		ThreadId  int         `json:"threadId"`  // 帖子ID
	// 	} `json:"indexes"` // 索引信息
	// } `json:"content"` // 帖子内容

	Freewords       int  `json:"freewords"`       // 免费字数
	UserStickStatus bool `json:"userStickStatus"` // 用户置顶状态
}

type JSONData struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    struct {
		PageData []struct {
			Content struct {
				Indexes map[string]struct {
					Body []struct {
						ThumbURL string `json:"thumbUrl"`
						URL      string `json:"url"`
					} `json:"body"`
				} `json:"indexes"`
			} `json:"content"`
		} `json:"pageData"`
	} `json:"Data"`
}

// fetchData从API中获取数据
func fetchData(url string) (*Response, error) {
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
	response := &Response{}
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
	fileList, err := ioutil.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	fileCount := len(fileList)

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
	response, err := fetchData(url)
	if err != nil {
		errLogger.Println("获取数据时出错:", err)
	}

	// 将数据保存为JSON文件
	filePath := fmt.Sprintf("data/7fenglou%d.json", index)

	// 在保存数据前添加运行日志记录
	runLogger.Printf("开始保存数据，URL：%s，文件路径：%s\n", url, filePath)

	err = saveJSONData(response, filePath)
	if err != nil {
		writeErrorToFile(err, url, filePath)
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
			response, err = fetchData(url)
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
			err = saveJSONData(response, filePath)
			if err != nil {
				writeErrorToFile(err, url, filePath)
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
		if (elapsedTime > timeout||weitrue) {
			break // 退出循环
		}
	}

	//// 创建Excel文件
	//f := excelize.NewFile()
	//defer func() {
	//	if err := f.Close(); err != nil {
	//		log.Println(err)
	//	}
	//}()
	//
	//// 添加一个工作表并设置列名
	//err = createExcelSheet(f, "Threads", []string{"Thread ID", "Post ID", "User ID", "Category ID"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 将数据插入Excel文件
	//err = insertDataIntoExcel(f, "Threads", pageData)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 保存Excel文件
	//err = saveExcelFile(f, "7fenglou.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("Data saved successfully.")
}
