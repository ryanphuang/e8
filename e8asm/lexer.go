package e8asm

import (
	"io"
	"bufio"
)

type lexer struct {
	scanner *bufio.Scanner
}

func newLexer(in io.Reader) *lexer {
	ret := new(lexer)
	ret.scanner = bufio.NewScanner(in)
	return ret
}

