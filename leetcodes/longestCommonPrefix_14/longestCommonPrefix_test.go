package longestCommonPrefix_14

import (
	"fmt"
	"testing"
)

type example struct {
	strings []string
	prefix  string
}

var examples []example

func init() {
	examples = append(examples, example{[]string{"flower", "flow", "flight"}, "fl"})
	examples = append(examples, example{[]string{"dog", "raceway", "car"}, ""})
	examples = append(examples, example{[]string{}, ""})
	examples = append(examples, example{[]string{""}, ""})
	examples = append(examples, example{[]string{"flower", "flow", "", "car"}, ""})
	examples = append(examples, example{[]string{"dog", "dog", "dog"}, "dog"})
	examples = append(examples, example{[]string{"dog"}, "dog"})
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, e := range examples {
		if e.prefix != LongestCommonPrefix(e.strings) {
			t.Error("错误")
			fmt.Println(e.strings, LongestCommonPrefix(e.strings), e.prefix, e.prefix == LongestCommonPrefix(e.strings))
		}
	}
}
