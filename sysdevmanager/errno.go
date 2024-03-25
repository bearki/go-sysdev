package sysdevmanager

// 错误码
type errno uint

// 错误码枚举
const (
	_               errno = iota // 占位
	ErrInputParam                // 传入参数错误
	ErrGetClassDevs              // 获取设备信息集句柄失败
)
