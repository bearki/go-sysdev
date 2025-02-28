package test

import (
	"fmt"
	"testing"

	goi18n "github.com/bearki/go-i18n/v2"
	"github.com/bearki/go-sysdev"
)

func TestNetworkCard(t *testing.T) {
	goi18n.SetDefault(goi18n.ZH_CN)
	fmt.Println(sysdev.ErrGetNetworkCardInfoFailed)
	res, err := sysdev.GetNetworkCardInfo()
	if err != nil {
		panic(err)
	}
	for _, item := range res {
		fmt.Println(item)
	}
}
