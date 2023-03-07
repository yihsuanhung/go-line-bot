package conf

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	payload *viper.Viper
	mu      sync.Mutex
}

func NewConfig(o ...Option) *config {
	opts := mergeOptions(o...)
	// 之後會加 Default Value，以及錯誤處理
	var vip = viper.New()
	// runtime_viper.AddConfigPath(".")
	vip.SetConfigName("setting")
	vip.SetConfigType(opts.Type)
	vip.AddConfigPath(opts.Path)
	if err := vip.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config for data source: %v", err))
	}
	return &config{
		payload: vip,
		mu:      sync.Mutex{},
	}
}

func (c *config) Unmarshal(result interface{}) {
	if err := c.payload.Unmarshal(result); err != nil {
		panic(err)
	}
}
