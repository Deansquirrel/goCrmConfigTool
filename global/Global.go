package global

import (
	"context"
	"github.com/Deansquirrel/goCrmConfigTool/config"
)

const (
	Version = "0.0.0 Build20190218"
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()
