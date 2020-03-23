package flipEquiv_951

import "testing"

func TestFlipEquiv(t *testing.T) {
	if !FlipEquiv(nil, nil) {
		t.Error("错误")
	}
}
