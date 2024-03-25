package sysdevmanager

import (
	_ "embed"
	"fmt"

	goi18n "github.com/bearki/go-i18n/v2"
	"gopkg.in/yaml.v3"
)

// 包名，以便于错误码能正常描述
const packageDescription = "system device manager"

//go:embed errno_i18n.yml
var errYml []byte

// 错误码描述映射
var errMap = make(map[string]map[goi18n.Code]string)

func init() {
	_ = yaml.Unmarshal(errYml, errMap)
	errYml = nil
}

// 错误码变量名映射
var errVarName = map[errno]string{
	ErrInputParam:   "ErrInputParam",
	ErrGetClassDevs: "ErrGetClassDevs",
}

func (e errno) Error() string {
	// 是否存在该错误码
	errName, ok := errVarName[e]
	if !ok {
		// 不存在的错误码
		return fmt.Sprintf("unknown %s errno: %d", packageDescription, e)
	}
	// 错误码是否有对应的国际化描述
	if errMsgMap, ok := errMap[errName]; ok {
		// 是否有对应语言的描述
		if len(errMsgMap) > 0 {
			errMsg, ok := errMsgMap[goi18n.GetEnv()]
			if ok {
				return errMsg
			} else {
				for _, msg := range errMsgMap {
					return msg
				}
			}
		}
	}
	// 错误码没有对应的国际化描述
	return fmt.Sprintf("raw %s errno: %d", packageDescription, e)
}
