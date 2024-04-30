package windows

/*
#cgo pkg-config: sysdev
#include "sysdev_helper.h"
*/
import "C"
import (
	"errors"
	"fmt"

	"github.com/bearki/go-sysdev/sysdevmanager"
)

// 转换状态码
func convertStatusCode(code C.StatusCode) error {
	switch code {
	// 成功
	case C.StatusCode_Success:
		return nil
	// 传入参数错误
	case C.StatusCode_ErrInputParam:
		return ErrInputParam
	// 获取设备信息集句柄失败
	case C.StatusCode_ErrGetClassDevs:
		return ErrGetClassDevs
	// 默认
	default:
		return fmt.Errorf("unknow system device manager status code: %d", code)
	}
}

// SystemDevice 系统设备管理器
type SystemDevice struct{}

// New 创建系统设备管理器实例
func New() sysdevmanager.Manager {
	return &SystemDevice{}
}

// GetNetworkCardInfo 获取网卡信息
//
//	@param	网卡信息列表
//	@return	异常信息
func (p *SystemDevice) GetNetworkCardInfo() ([]sysdevmanager.NetworkCardInfo, error) {
	// 声明响应
	var replyList *C.NetworkCardInfo = nil
	var replyListSize C.size_t = 0
	// 获取网卡信息
	code := C.SysDevGetNetworkCardInfo(&replyList, &replyListSize)
	if err := convertStatusCode(code); err != nil {
		return nil, errors.Join(sysdevmanager.ErrGetNetworkCardInfoFailed, err)
	}

	// 延迟释放
	defer C.SysDevFreeNetworkCardInfo(replyList, replyListSize)

	// 检查释放有获取到
	if replyListSize <= 0 {
		return nil, nil
	}

	// 执行拷贝
	resList := make([]sysdevmanager.NetworkCardInfo, replyListSize)
	for i := 0; i < int(replyListSize); i++ {
		// 提取单个元素
		item := C.getNetworkCardInfoListItem(replyList, C.size_t(i))
		// 构建信息
		var info sysdevmanager.NetworkCardInfo
		// 转换网卡类型
		info.Type = sysdevmanager.NetworkCardType_Unknown
		switch item.netCardType {
		case C.NetworkCardType_PCI:
			info.Type = sysdevmanager.NetworkCardType_PCI
		case C.NetworkCardType_USB:
			info.Type = sysdevmanager.NetworkCardType_USB
		}
		// 转换MAC地址
		for j, v := range item.macAddress {
			info.MacAddress[j] = byte(v)
		}
		// 拷贝
		resList[i] = info
	}

	// OK
	return resList, nil
}
