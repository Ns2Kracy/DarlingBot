package DarlingBot

// Server 运行环境
type Server interface {
}

// APICaller api调用器
type APICaller interface {
	CallAPI(api APIRequest) (APIResponse, error)
}
