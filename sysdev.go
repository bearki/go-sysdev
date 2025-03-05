package sysdev

/*
#cgo pkg-config: sysdev
#cgo CFLAGS: -I/home/WORK/Project/sysdev/install/gnu/libsysdev_linux_x86_64/include
#include "sysdev_helper.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

// GetNetworkCardInfo 获取网卡信息
//
//	@param	网卡信息列表
//	@return	异常信息
func GetNetworkCardInfo() ([]NetworkCardInfo, error) {
	// 声明响应
	var replyList *C.NetworkCardInfo = nil
	var replyListSize C.size_t = 0
	// 获取网卡信息
	code := C.SysDevGetNetworkCardInfo(&replyList, &replyListSize)
	if code != C.StatusCode_Success {
		return nil, errors.Join(ErrGetNetworkCardInfoFailed, fmt.Errorf("return code: %d", code))
	}

	// 延迟释放
	defer C.SysDevFreeNetworkCardInfo(replyList, replyListSize)

	// 检查释放有获取到
	if replyListSize <= 0 {
		return nil, nil
	}

	// 执行拷贝
	resList := make([]NetworkCardInfo, replyListSize)
	for i := 0; i < int(replyListSize); i++ {
		// 提取单个元素
		item := C.getNetworkCardInfoListItem(replyList, C.size_t(i))
		// 构建信息
		var info NetworkCardInfo
		// 提取网卡名称
		if item.netCardName != nil {
			info.Name = C.GoString(item.netCardName)
		}
		// 转换网卡类型
		info.Type = NetworkCardType_Unknown
		switch item.netCardType {
		case C.NetworkCardType_PCI:
			info.Type = NetworkCardType_PCI
		case C.NetworkCardType_USB:
			info.Type = NetworkCardType_USB
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
