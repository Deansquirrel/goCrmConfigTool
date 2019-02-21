package worker

import (
	"github.com/Deansquirrel/goCrmConfigTool/global"
	"github.com/Deansquirrel/goCrmConfigTool/object"
	"github.com/Deansquirrel/goToolCommon"
	"io/ioutil"
	"strings"
)
import log "github.com/Deansquirrel/goToolLog"

type Worker struct {
}

//转换
func (w *Worker) Tran() {
	err := CheckConfig()
	if err != nil {
		log.Error("检查配置时遇到异常：" + err.Error())
		return
	}
	tranConfigList := make([]*object.TranConfig, 0)
	for _, config := range global.SysConfig.TranConfigs.ConfigList {
		c, err := object.NewTranConfig(config)
		if err != nil {
			log.Error("获取转换配置时遇到错误：" + err.Error())
			return
		} else {
			tranConfigList = append(tranConfigList, c)
		}
	}
	kMap := GetTemplateValue(&global.SysConfig.TemplateValue)
	for _, config := range tranConfigList {
		err = w.TranFile(config, kMap)
		if err != nil {
			log.Error("转换文件时发生错误：" + err.Error())
			return
		}
	}
	log.Info("success")
}

//根据配置替换并生成文件
func (w *Worker) TranFile(config *object.TranConfig, kMap map[string]string) error {
	tPath, err := w.getTemplateFullPath(config.TemplateFolder, config.TemplateFileName)
	if err != nil {
		return err
	}
	oPath, err := w.getOutFullPath(config.OutFolder, config.OutFilename)
	if err != nil {
		return err
	}
	data, err := w.getFileDate(tPath)
	if err != nil {
		return err
	}
	sData := string(data)
	for k, v := range kMap {
		sData = strings.Replace(sData, "###"+k+"###", v, -1)
	}
	return w.writeOutFile(oPath, []byte(sData))
}

//获取模板文件路径
func (w *Worker) getTemplateFullPath(folder string, file string) (string, error) {
	cPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		return "", err
	}
	folder = goToolCommon.CheckAndDeleteFirstChar(folder, "\\")
	folder = goToolCommon.CheckAndDeleteLastChar(folder, "\\")
	return cPath + "\\" + "Template" + "\\" + folder + "\\" + file, nil
}

//获取输出文件路径
func (w *Worker) getOutFullPath(folder string, file string) (string, error) {
	cPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		return "", err
	}
	folder = goToolCommon.CheckAndDeleteFirstChar(folder, "\\")
	folder = goToolCommon.CheckAndDeleteLastChar(folder, "\\")

	err = goToolCommon.CheckAndCreateFolder(cPath + "\\" + "Out")
	if err != nil {
		return "", err
	}

	err = goToolCommon.CheckAndCreateFolder(cPath + "\\" + "Out" + "\\" + folder)
	if err != nil {
		return "", err
	}
	return cPath + "\\" + "Out" + "\\" + folder + "\\" + file, nil
}

//读取模板文件数据
func (w *Worker) getFileDate(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

//文件输出
func (w *Worker) writeOutFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
