package message

const (
	LoginMesType    = "loginMes"
	LoginResMesType = "loginResMes"
)

type PackageMessage struct {
	//32位,8字节
	DataSize int `json:"size"`
	//size 位数据
	DataArray []byte `json:"dataArray"`
}
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// LoginMes 登录消息
type LoginMes struct {
	UserId       int    `json:"userId"`
	UserPassword string `json:"userPassword"`
}

// LoginResMes 登录返回消息
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
