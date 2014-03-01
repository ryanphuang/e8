package asm

import (
	"fmt"
)

func ef(s string, args ...interface{}) error {
	return fmt.Errorf(s, args...)
}
