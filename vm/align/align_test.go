package align

import (
	"testing"
)

func TestAlign(t *testing.T) {
	o := func(b bool) {
		if !b {
			t.Fail()
		}
	}

	o(U32(3) == 0)
	o(U16(3) == 2)
	o(U32(1024) == 1024)
	o(U32(1025) == 1024)
	o(U32(1026) == 1024)
	o(U32(1027) == 1024)
}
