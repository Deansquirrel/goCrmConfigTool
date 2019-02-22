package worker

import (
	"encoding/json"
	"errors"
	"github.com/Deansquirrel/goCrmConfigTool/global"
	"github.com/Deansquirrel/goCrmConfigTool/object"
	"github.com/Deansquirrel/goToolCommon"
	"io/ioutil"
	"strings"
)
import log "github.com/Deansquirrel/goToolLog"

const (
	TemplateFolderName  = "Template"
	OutFolderName       = "Out"
	TemplateFolderSplit = "###"
	ReplactFlag         = "###"
)

type Worker struct {
}

//转换
func (w *Worker) Tran() {
	err := CheckConfig()
	if err != nil {
		log.Error("检查配置时遇到异常：" + err.Error())
		return
	}

	//==================================================================================================================
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
	//==================================================================================================================
	tcList, err := w.getTranConfigList()
	if err != nil {
		log.Error("获取转换配置时遇到错误：" + err.Error())
	} else {
		for _, tc := range tcList {
			log.Debug(tc.TemplateFolder + " | " + tc.TemplateFileName + " | " + tc.OutFolder + " | " + tc.OutFilename)
		}
	}
	//==================================================================================================================

	kMap := GetTemplateValue(&global.SysConfig.TemplateValue)
	log.Info("转换内容列表：")
	for key, val := range kMap {
		log.Info(key + " - " + val)
	}
	for _, config := range tranConfigList {
		err = w.TranFile(config, kMap)
		if err != nil {
			log.Error("转换文件时发生错误：" + err.Error())
			return
		}
	}
	log.Info("success")
}

//获取替换内容列表
func (w *Worker) getTranConfigList() ([]object.TranConfig, error) {
	cPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		return nil, err
	}
	b, err := goToolCommon.PathExists(cPath + "\\" + TemplateFolderName)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, errors.New("模板路径不存在，请检查")
	}
	templatePath := cPath + "\\" + TemplateFolderName
	tcList, err := w.getTranConfigListWorker(templatePath, "")
	if err != nil {
		return nil, err
	}
	for _, tc := range tcList {
		b, err := json.Marshal(tc)
		if err != nil {
			log.Debug(err.Error())
		} else {
			log.Debug(string(b))
		}
	}
	return nil, nil
}

//按路径生成配置文件列表并返回文件夹
//path 路径
//basePath TemplateFolder基础路径
func (w *Worker) getTranConfigListWorker(path string, basePath string) ([]object.TranConfig, error) {
	folderList, fileList, err := goToolCommon.GetFolderAndFileList(path)
	if err != nil {
		return nil, err
	}
	tranConfigList := make([]object.TranConfig, 0)
	for _, file := range fileList {
		tc := object.TranConfig{
			TemplateFolder:   basePath,
			TemplateFileName: file,
			OutFolder:        basePath,
			OutFilename:      file,
		}
		tranConfigList = append(tranConfigList, tc)
	}
	if len(folderList) > 0 {
		for _, folder := range folderList {
			var newBasePath string
			if basePath != "" {
				newBasePath = basePath + TemplateFolderSplit + folder
			} else {
				newBasePath = folder
			}
			tcList, err := w.getTranConfigListWorker(path+"\\"+folder, newBasePath)
			if err != nil {
				return nil, err
			} else {
				for _, tc := range tcList {
					tranConfigList = append(tranConfigList, tc)
				}
			}
		}
	}
	return tranConfigList, nil
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

//获取模板文件路径
func (w *Worker) getTemplateFullPath2(folder string, file string) (string, error) {
	cPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		return "", err
	}
	return cPath + "\\" + TemplateFolderName + "\\" + strings.Replace(folder, TemplateFolderSplit, "\\", -1) + file, nil
}

//获取输出文件路径
func (w *Worker) getOutFullPath2(folder string, file string) (string, error) {
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
