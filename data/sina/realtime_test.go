package sina

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestGetRealtime(t *testing.T) {
	str := `var hq_str_sh601003="柳钢股份,6.610,6.590,6.890,6.950,6.510,6.880,6.890,16908009,115327931.000,9000,6.880,15400,6.870,58600,6.860,58300,6.850,58800,6.840,59400,6.890,658900,6.900,653100,6.910,27100,6.920,62580,6.930,2021-08-12,11:30:00,00,";
var hq_str_sh601001="晋控煤业,9.300,9.500,9.340,9.440,9.130,9.340,9.350,47165453,436979711.000,300,9.340,9700,9.330,22400,9.320,23200,9.310,39400,9.300,1300,9.350,17000,9.360,12500,9.370,15400,9.380,42500,9.390,2021-08-12,11:30:00,00,";
var hq_str_hk00700="TENCENT,腾讯控股,482.600,484.000,491.200,478.600,486.800,2.800,0.579,486.60001,486.79999,3930914563,8102711,0.000,0.000,763.287,422.000,2021/08/12,11:59";
var hq_str_sh000001="上证指数,3522.7238,3532.6213,3528.2737,3538.3960,3513.4457,0,0,246361295,343695733032,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2021-08-12,11:35:03,00,";
var hq_str_hk01812="CHENMING PAPER,晨鸣纸业,4.630,4.620,4.660,4.600,4.650,0.030,0.649,4.63000,4.65000,12989960,2797250,0.000,0.000,9.878,3.033,2021/08/12,11:59";`

	hqList := strings.Split(str, ";\n")

	for _, item := range hqList{
		if len(item) < 10 {
			continue
		}
		r, _ := regexp.Compile(stockPatterString)
		arr := r.FindAllStringSubmatch(item, -1)
		if len(arr) < 1 {
			continue
		}
		// 1行1个实时行情数据
		hangqing := arr[0]
		data := strings.Split(hangqing[2], ",")
		code := hangqing[1]
		var rt interface{}
		var err error
		if strings.HasPrefix(code, "hk") {
			rt, err = GetHongkongRealtime(code, data)
		} else {
			rt, err = GetHushenRealtime(code, data)
		}
		_ = err

		fmt.Printf("%+v\n", rt)
	}

}
