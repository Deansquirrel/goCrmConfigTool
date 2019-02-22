package global

import (
	"context"
	"github.com/Deansquirrel/goCrmConfigTool/config"
)

const (
	Version = "1.0.0 Build20190222"
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()
