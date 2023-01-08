package sina

import (
	"encoding/json"
	"fmt"
	"github.com/mymmsc/gox/fastjson"
	"github.com/quant1x/quant/stock"
	"github.com/quant1x/quant/utils/http"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	fmt.Println(createUrl("sh000001", stock.ONE_HOUR, stock.DEFAULT_DATALEN))
	fmt.Println(createUrl("sh000001", stock.ONE_DAY, stock.DEFAULT_DATALEN))
	fmt.Println(createUrl("sh000001", stock.ONE_WEEK, stock.DEFAULT_DATALEN))
}

func TestHistory(t *testing.T) {
	kl, _ := GetHistory("sh000001", 1)
	fmt.Printf("%+v\n", kl)
}

type Transport struct {
	Time  string
	MAC   string
	Id    string
	Rssid string
}

func Test1(t *testing.T) {
	var st []Transport
	t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}
	st = append(st, t1)
	t2 := Transport{Time: "66", MAC: "77", Id: "88", Rssid: "99"}
	st = append(st, t2)
	buf, _ := json.Marshal(st)
	fmt.Println(string(buf))

	var str = string(buf)
	var st1 []Transport
	err := json.Unmarshal([]byte(str), &st1)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(st1)
	fmt.Println(st1[0].Time)

	var Msg []map[string]string
	json.Unmarshal([]byte(str), &Msg)
	fmt.Println(Msg)
}

func TestHistory4(t *testing.T) {
	fullCode := "sh000001"
	fullCode = "hk00700"
	datalen := 1
	url := createUrl(fullCode, stock.ONE_DAY, datalen)

	data, err := http.HttpGet(url)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	fmt.Printf("json=[%s]\n", data)
	var kl []stock.History
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		t.Errorf("data[%s], error=[%+v]\n", data, err)
	}
	if v.Type() == fastjson.TypeArray {
		va, err := v.Array()
		if err != nil {
			t.Errorf("%+v\n", err)
		}
		count := len(va)
		for i := 0; i < count; i++ {
			obj := va[i]
			oo, err := obj.Object()
			if err != nil {
				continue
			}
			fmt.Printf("%+v\n", oo)
			oo.Visit(func(key []byte, v *fastjson.Value) {
				fmt.Printf("%+v, key=%s,value=%v\n", v.Type(), key, v.String())
			})
			/*s := obj.GetString("ma_price5")
			fmt.Printf("%s\n", s)
			s = obj.GetString("ma_volume5")
			fmt.Printf("%s\n", s)*/
		}
	}
	fmt.Println(kl)
}
