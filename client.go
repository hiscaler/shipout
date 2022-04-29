package shipout

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/shipout-go/config"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// 返回代码
const (
	OK = 200 // 无错误
)

var ErrNotFound = errors.New("shipout: not found")

type queryDefaultValues struct {
	PageNo   int // 当前页
	PageSize int // 每页数据量
}

type ShipOut struct {
	Debug              bool               // 是否调试模式
	Client             *resty.Client      // HTTP 客户端
	Logger             *log.Logger        // 日志
	EnableCache        bool               // 是否激活缓存
	QueryDefaultValues queryDefaultValues // 查询默认值
}

func NewShipOut(config config.Config) *ShipOut {
	logger := log.New(os.Stdout, "[ ShipOut ] ", log.LstdFlags|log.Llongfile)
	ttInstance := &ShipOut{
		Debug:  config.Debug,
		Logger: logger,
		QueryDefaultValues: queryDefaultValues{
			PageNo:   1,
			PageSize: 100,
		},
	}
	timestamp := time.Now().UnixMicro()
	sign := ""
	client := resty.New().
		SetDebug(config.Debug).
		SetBaseURL("https://open.shipout.com/api/").
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": config.Authorization,
			"timestamp":     strconv.Itoa(int(timestamp)),
			"appKey":        config.AppKey,
			"sign":          sign,
			"version":       "1.0.0",
		}).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			return nil
		})
	if config.Debug {
		client.SetBaseURL("https://opendev.shipout.com/api/")
		client.EnableTrace()
	}
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	ttInstance.Client = client
	return ttInstance
}

// NormalResponse Normal API response
type NormalResponse struct {
	Result    string        `json:"result"`
	ErrorCode string        `json:"errorCode"`
	Message   string        `json:"message"`
	ZhMessage string        `json:"zhMessage"`
	ErrorType string        `json:"errorType"`
	Data      []interface{} `json:"data"`
}

// ErrorWrap 错误包装
func ErrorWrap(code string, message string) error {
	if code == "" {
		return nil
	}

	message = strings.TrimSpace(message)
	if message == "" {
	}
	return fmt.Errorf("%s: %s", code, message)
}
