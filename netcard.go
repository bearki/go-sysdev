package sysdev

import (
	"fmt"
	"strings"
)

// NetworkCardType 网卡类型
type NetworkCardType uint

// 网卡类型枚举
const (
	NetworkCardType_Unknown NetworkCardType = iota // 网卡类型：未知
	NetworkCardType_PCI                            // 网卡类型：PCI网卡
	NetworkCardType_USB                            // 网卡类型：USB网卡
)

func (val NetworkCardType) String() string {
	switch val {
	case NetworkCardType_PCI:
		return "PCI"
	case NetworkCardType_USB:
		return "USB"
	default:
		return "UNKNOW"
	}
}

// NetworkCardMacAddress 网卡MAC地址
type NetworkCardMacAddress [6]byte

func (val NetworkCardMacAddress) StringWithSep(sep string) string {
	var str [6]string
	for i, v := range val {
		str[i] = fmt.Sprintf("%02X", v)
	}
	return strings.Join(str[:], sep)
}

func (val NetworkCardMacAddress) String() string {
	return val.StringWithSep(":")
}

// NetworkCardInfo 网卡信息
type NetworkCardInfo struct {
	Name       string                // 网卡名称
	Type       NetworkCardType       // 网卡类型
	MacAddress NetworkCardMacAddress // 网卡MAC地址
}
