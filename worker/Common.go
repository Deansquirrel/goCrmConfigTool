package worker

import (
	"errors"
	"github.com/Deansquirrel/goCrmConfigTool/config"
	"github.com/Deansquirrel/goCrmConfigTool/global"
)

//将TemplateValue转换为map[string]string
func GetTemplateValue(tv *config.TemplateValue) map[string]string {
	result := make(map[string]string)
	for index := range tv.KeyList {
		result[tv.KeyList[index]] = tv.ValueList[index]
	}
	return result
}

//检查配置
func CheckConfig() error {
	if len(global.SysConfig.TemplateValue.KeyList) != len(global.SysConfig.TemplateValue.ValueList) {
		return errors.New("TemplateValue中key和value的数量不一致")
	}
	if len(global.SysConfig.TemplateValue.KeyList) < 1 || len(global.SysConfig.TemplateValue.ValueList) < 1 {
		return errors.New("TemplateValue中key和value不能为空")
	}
	return nil
}
