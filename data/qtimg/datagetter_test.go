package qtimg

import (
	"fmt"
	"github.com/melman-go/aliopengo/util"
	"testing"
)

func TestTokenValidate(t *testing.T) {
	code := "sh600600"
	//	realTime := GetRealtime(code)
		pk:= GetPK(code)
	//	funFlow := GetFundFlow(code)
	//	info := GetInfo(code)
	//	fmt.Println(util.JsonEncodeS(realTime))
		fmt.Println(util.JsonEncodeS(pk))
	//	fmt.Println(util.JsonEncodeS(funFlow))
	//	fmt.Println(util.JsonEncodeS(info))
	list := GetDaily(code, 2001)
	//list := GetWeekly(code)
	fmt.Println(util.JsonEncodeS(list))
}

func TestGetRealtime(t *testing.T) {
	code := "sz002378"
	realTime := GetRealtime(code)
	fmt.Println(util.JsonEncodeS(realTime))
}
