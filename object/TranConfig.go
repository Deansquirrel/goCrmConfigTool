package object

import (
	"github.com/kataras/iris/core/errors"
	"strconv"
	"strings"
)

type TranConfig struct {
	TemplateFolder   string
	TemplateFileName string
	OutFolder        string
	OutFilename      string
	TranKeyList      string
	TranValList      string
}

func NewTranConfig(config string) (*TranConfig, error) {
	cList := strings.Split(config, "|")
	if len(cList) != 4 {
		return nil, errors.New("配置字符串[" + config + "]转换异常，expect 6,got " + strconv.Itoa(len(cList)))
	}
	return &TranConfig{
		TemplateFolder:   cList[0],
		TemplateFileName: cList[1],
		OutFolder:        cList[2],
		OutFilename:      cList[3],
	}, nil
}
