package main

import (
	"context"
	"encoding/json"
	"github.com/Deansquirrel/goCrmConfigTool/common"
	"github.com/Deansquirrel/goCrmConfigTool/global"
	"github.com/Deansquirrel/goCrmConfigTool/worker"
	log "github.com/Deansquirrel/goToolLog"
)

func main() {
	//==================================================================================================================
	log.StdOut = true
	log.Level = log.LevelInfo
	log.Info("程序启动")
	defer log.Info("程序退出")
	//==================================================================================================================
	//==================================================================================================================
	config, err := common.GetSysConfig("config.toml")
	if err != nil {
		log.Error("加载配置文件时遇到错误：" + err.Error())
		return
	}
	config.FormatConfig()
	global.SysConfig = config
	//==================================================================================================================
	err = common.RefreshSysConfig(*global.SysConfig)
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
		return
	}
	global.Ctx, global.Cancel = context.WithCancel(context.Background())
	//==================================================================================================================
	log.Info("Version " + global.Version)
	//==================================================================================================================
	c, err := json.Marshal(global.SysConfig)
	if err != nil {
		log.Info("显示设置时遇到错误：" + err.Error())
	} else {
		log.Info("设置：" + string(c))
	}
	//==================================================================================================================
	w := worker.Worker{}
	w.Tran()
	//==================================================================================================================
	//time.AfterFunc(time.Second*30, func() {
	//	global.Cancel()
	//})
	//==================================================================================================================
	//select {
	//case <-global.Ctx.Done():
	//}
}
