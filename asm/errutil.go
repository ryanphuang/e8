package asm

import (
	"fmt"
)

func ef(s string, args ...interface{}) error {
	return fmt.Errorf(s, args...)
}

func lef(s string, args ...interface{}) (*Line, error) {
	return nil, fmt.Errorf(s, args...)
}
