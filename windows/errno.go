package windows

const (
	_ errno = iota // 占位

	ErrInputParam   // 传入参数错误
	ErrGetClassDevs // 获取设备信息集句柄失败
)

// 错误码变量名映射
var errVarName = map[errno]string{
	ErrInputParam:   "ErrInputParam",
	ErrGetClassDevs: "ErrGetClassDevs",
}
