package sysdev

import (
	"github.com/bearki/go-sysdev/sysdevmanager"
	"github.com/bearki/go-sysdev/windows"
)

// New 创建系统设备管理器
func New() sysdevmanager.Manager {
	return windows.New()
}
