package types

// Commit 结构体表示提交信息
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

// Response 结构体表示响应信息
type Response struct {
	Code        int    `json:"Code"`        // 响应状态码
	Message     string `json:"Message"`     // 响应消息
	Data        Data   `json:"Data"`        // 数据对象
	RequestId   string `json:"RequestId"`   // 请求ID
	RequestTime string `json:"RequestTime"` // 请求时间
}

// PageData 结构体表示页面数据
type PageData struct {
	User            User        `json:"user"`            // 用户信息
	Group           Group       `json:"group"`           // 群组信息
	DisplayTag      struct{}    `json:"displayTag"`      // 显示标签信息
	Position        struct{}    `json:"position"`        // 位置信息
	Ability         interface{} `json:"ability"`         // 用户权限
	Content         interface{} `json:"content"`         // 帖子内容
	Freewords       int         `json:"freewords"`       // 免费字数
	UserStickStatus bool        `json:"userStickStatus"` // 用户置顶状态
}

// User 结构体表示用户信息
type User struct {
	Nickname string `json:"nickname"` // 用户昵称
}

// Group 结构体表示群组信息
type Group struct {
	// 在这里定义 "group" 对象中的字段，如果在 JSON 中可用。
}

// JSONData 结构体表示 JSON 数据对象
type JSONData struct {
	Code    int    `json:"Code"`    // 响应状态码
	Message string `json:"Message"` // 响应消息
	Data    struct {
		PageData []struct {
			Content struct {
				Indexes map[string]struct {
					Body []struct {
						ThumbURL string `json:"thumbUrl"` // 缩略图 URL
						URL      string `json:"url"`      // URL
					} `json:"body"` // 正文主体
				} `json:"indexes"` // 索引信息
			} `json:"content"` // 帖子内容
		} `json:"pageData"` // 页面数据列表
	} `json:"Data"` // 数据对象
}
