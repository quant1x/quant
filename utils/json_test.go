package utils

import (
	"fmt"
	"log"
	"testing"

	"github.com/mymmsc/gox/fastjson"
)

func TestFastJson(t *testing.T) {
	var p fastjson.Parser
	v, err := p.Parse(`{foo:"bar", "baz": 123}`)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}

	fmt.Printf("foo=%s, baz=%d\n", v.GetStringBytes("foo"), v.GetInt("baz"))
}
