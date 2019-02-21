package config

type TranConfigs struct {
	ConfigList []string `toml:"configList"`
}

//配置检查并格式化
func (tc *TranConfigs) FormatConfig() {

}
