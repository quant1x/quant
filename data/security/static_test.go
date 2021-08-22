package security

import (
	"testing"
)

func TestGetStaticBasic(t *testing.T) {
	market := "sh"
	list, err := getStaticBasic(market)
	if err != nil {
		t.Fatalf("获取%s失败", market)
	}
	t.Logf("%+v", list)
}

func TestGetBasicInfo(t *testing.T) {
	code := "sh600600"
	info, err := GetBasicInfo(code)
	if err != nil {
		t.Fatalf("获取%s失败, %+v", code, err)
	}
	t.Logf("%+v", info)
}
