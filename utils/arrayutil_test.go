package utils

import (
	"fmt"
	"testing"
)

func TestStringCaseConvert(t *testing.T) {
	a1 := []string{"a", "b", "b"}
	a2 := NewArrayUtil().ArrayToStr(a1)
	fmt.Println(a2)
}
