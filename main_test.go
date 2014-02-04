package gocamp

import (
	"fmt"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	list := []string{"go", "ca", "mp", "now"}

	msg := Encode(list)
  actual := Decode(msg)

	fmt.Println(actual)
}
