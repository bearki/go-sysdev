package test

import (
	"fmt"
	"testing"

	goi18n "github.com/bearki/go-i18n/v2"
	"github.com/bearki/go-sysdev"
	"github.com/bearki/go-sysdev/sysdevmanager"
)

func TestNetworkCard(t *testing.T) {
	goi18n.SetDefault(goi18n.ZH_CN)
	fmt.Println(sysdevmanager.ErrGetNetworkCardInfoFailed)
	h := sysdev.New()
	res, err := h.GetNetworkCardInfo()
	fmt.Println(res, err)
}
