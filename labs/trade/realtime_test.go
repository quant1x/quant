package internal

import (
	"fmt"
	"testing"
)

func TestBatchSnapShot(t *testing.T) {
	data := BatchSnapShot([]string{"sh600600"})
	fmt.Printf("%+v\n", data)
}
