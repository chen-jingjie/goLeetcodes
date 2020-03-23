package strStr_28

import (
	"strconv"
	"testing"
)

type exmple struct {
	v1 string
	v2 string
	r  int
}

func TestStrStr(t *testing.T) {
	v := []exmple{
		{"hello", "ll", 2},
		{"testeest", "st", 2},
		{"chenasijo", "nas", 3},
		{"aaaaa", "", 0},
		{"hghslfnsdjk", "slf", 3},
		{"gdsiogwoeng", "iog", 3},
		{"aaa", "aaaa", -1},
		{"mississippi", "issip", 4},
		{"mississippi", "pi", 9},
	}

	for _, e := range v {
		if StrStr(e.v1, e.v2) != e.r {
			t.Error("错误" + e.v1 + "," + e.v2 + "," + strconv.Itoa(e.r))
		}
	}
}
