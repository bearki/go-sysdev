package sysdevmanager

const (
	_ errno = iota // 占位

	ErrGetNetworkCardInfoFailed // 获取网卡信息失败
)

// 错误码变量名映射
var errVarName = map[errno]string{
	ErrGetNetworkCardInfoFailed: "ErrGetNetworkCardInfoFailed",
}
