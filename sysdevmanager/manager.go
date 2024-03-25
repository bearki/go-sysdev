package sysdevmanager

// Manager 系统设备管理器接口声明
type Manager interface {
	// GetNetworkCardInfo 获取网卡信息
	//
	//	@param	网卡信息列表
	//	@return	异常信息
	GetNetworkCardInfo() ([]NetworkCardInfo, error)
}
