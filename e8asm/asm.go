package e8asm

import (
	"io"
)

func Assemble(in io.Reader, out io.Writer) error {
	lexer := newLexer(in)

	ast, err := parse(lexer)
	if err != nil {
		return err
	}

	err = generate(ast, out)
	return err
}
