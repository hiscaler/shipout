package config

type Config struct {
	Debug         bool   // 是否为调试模式
	Authorization string // 商家鉴权信息
	AppKey        string // 获取 API Key 中生成的开发者鉴权信息中的 appKey
	EnableCache   bool   // 是否激活缓存
}
