package main

import (
	"fmt"

	goi18n "github.com/bearki/go-i18n/v2"
	"github.com/bearki/go-sysdev"
)

func main() {
	goi18n.SetDefault(goi18n.ZH_CN)
	res, err := sysdev.GetNetworkCardInfo()
	if err != nil {
		panic(err)
	}
	for _, item := range res {
		fmt.Println(item)
	}
}
