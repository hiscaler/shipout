package shipout

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/cryptox"
	"github.com/hiscaler/shipout-go/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 返回代码
const (
	OK = 200 // 无错误
)

func init() {
	extra.RegisterFuzzyDecoders()
}

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
	soInstance := &ShipOut{
		Debug:  config.Debug,
		Logger: logger,
		QueryDefaultValues: queryDefaultValues{
			PageNo:   1,
			PageSize: 100,
		},
	}
	client := resty.New().
		SetDebug(config.Debug).
		SetBaseURL("https://open.shipout.com/api/").
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"appKey":       config.AppKey,
		}).
		SetAuthToken(config.Authorization).
		SetAllowGetMethodPayload(true).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			headers := map[string]string{
				"timestamp": strconv.Itoa(int(time.Now().UnixMicro())),
				"version":   "1.0.0",
				"path":      request.URL,
			}
			keys := make([]string, len(headers))
			i := 0
			for k := range headers {
				keys[i] = k
				i++
			}
			sort.Strings(keys)
			sb := strings.Builder{}
			for _, key := range keys {
				sb.WriteString(key)
				sb.WriteString(headers[key])
			}
			sb.WriteString(config.SecretKey)
			headers["sign"] = strings.ToUpper(cryptox.Md5(sb.String()))
			request.SetHeaders(headers)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsSuccess() {
				r := struct {
					Result    string `json:"result"`
					ErrorCode string `json:"ErrorCode"`
					Message   string `json:"message"`
				}{}
				if err = jsoniter.Unmarshal(response.Body(), &r); err == nil {
					code := r.ErrorCode
					if code == "" {
						code = r.Result
					}
					err = ErrorWrap(code, r.Message)
				}
			}
			if err != nil {
				logger.Printf("OnAfterResponse error: %s", err.Error())
			}
			return
		})
	if config.Debug {
		client.SetBaseURL("https://opendev.shipout.com/api/")
		client.EnableTrace()
	}
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	soInstance.Client = client
	return soInstance
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
	if code == "" || code == "OK" {
		return nil
	}

	message = strings.TrimSpace(message)
	if message == "" {
	}
	return fmt.Errorf("%s: %s", code, message)
}
