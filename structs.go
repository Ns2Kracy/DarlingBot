package DarlingBot

import "github.com/tidwall/gjson"

// Config 配置文件
type Config struct {
	Nickname   []string `json:"nickname"`    // 机器人的昵称, 可以设置多个
	Command    string   `json:"command"`     // 触发命令
	SuperUsers []int64  `json:"super_users"` // 超级用户
	Server     []Server `json:"-"`           // 通信驱动
}

// APIRequest 请求
type APIRequest struct {
	Action string `json:"action"` // API终结点
	Params Params `json:"params"` // 参数
	Echo   string `json:"echo"`   // "回声"
}

// APIResponse 响应
type APIResponse struct {
	Status  string       `json:"status"`  // 状态, 表示 API 是否调用成功, 如果成功, 则是 OK
	Retcode int          `json:"retcode"` // 错误消息, 仅在 API 调用失败式有该字段
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"` // 对错误的详细解释(中文), 仅在 API 调用失败时有该字段
	Data    gjson.Result `json:"data"`    // 数据
	Echo    string       `json:"echo"`    // "回声", 如果请求时指定了 echo, 那么响应也会包含 echo
}
