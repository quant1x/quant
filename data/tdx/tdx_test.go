package tdx

import (
	"fmt"
	"testing"
)

func TestGetKLine(t *testing.T) {
	data := GetKLine("000002", 0, 1)
	fmt.Println(data)
}
