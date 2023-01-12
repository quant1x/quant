package tdx

import (
	"fmt"
	"testing"
)

func TestGetKLine(t *testing.T) {
	data := GetKLine("600600", 0, 1)
	fmt.Println(data)
}
