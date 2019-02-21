package config

import (
	"encoding/json"
)

type SysConfig struct {
	Total         Total         `toml:"total"`
	TemplateValue TemplateValue `toml:"templateValue"`
	TranConfigs   TranConfigs   `toml:"tranConfig"`
}

//返回配置字符串
func (sc *SysConfig) GetConfigStr() (string, error) {
	b, err := json.Marshal(sc)
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

//配置检查并格式化
func (sc *SysConfig) FormatConfig() {
	sc.Total.FormatConfig()
	sc.TemplateValue.FormatConfig()
}
