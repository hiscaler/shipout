package shipout

import (
	"fmt"
	"github.com/hiscaler/shipout-go/config"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
)

var shipOutClient *ShipOut

func TestMain(m *testing.M) {
	b, err := os.ReadFile("./config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	shipOutClient = NewShipOut(c)
	m.Run()
}
