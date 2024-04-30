package main

import (
	"fmt"

	goi18n "github.com/bearki/go-i18n/v2"
	"github.com/bearki/go-sysdev"
)

func main() {
	goi18n.SetDefault(goi18n.ZH_CN)
	h := sysdev.New()
	res, err := h.GetNetworkCardInfo()
	fmt.Println(res, err)
}
