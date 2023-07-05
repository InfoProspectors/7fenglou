// Commit 结构体表示一个提交
package types

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
