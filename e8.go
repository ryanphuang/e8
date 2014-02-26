package main

import (
	"bytes"
	// "fmt"
	"os"
	"strings"

	"github.com/h8liu/e8/e8asm"
	"github.com/h8liu/e8/img"
	"github.com/h8liu/e8/vm/mem"
)

func pu32(buf []byte, i uint32) {
	buf[0] = uint8(i)
	buf[1] = uint8(i >> 8)
	buf[2] = uint8(i >> 16)
	buf[3] = uint8(i >> 24)
}

func hello() []byte {
	buf := strings.NewReader(`
			add $1, $0, $0      ; init counter
			; addi $1, $1, 13
		loop:
			lbu $2, 0x2000($1)  ; load byte
			beq $2, $0, end     ; +5
		wait:
			lbu $3, 5           ; is output ready?
			bne $3, $0, wait    ; -2
			sb $2, 5            ; output byte
			; addi $1, $1, -1     ; update counter
			addi $1, $1, 1     ; update counter
			j loop              ; -7
		end:
			sb $0, 0x4($0)
	`)

	ret := new(bytes.Buffer)
	asm := &e8asm.Assembler{
		In:  buf,
		Out: ret,
	}
	e := asm.Assemble()
	if e != nil {
		panic(e)
	}

	return ret.Bytes()
}

func makeMap() []byte {
	str := "Hello, world.\n\000"

	ret := new(bytes.Buffer)
	w := img.NewWriter(ret)
	w.Write(mem.PageStart(1), hello())
	w.Write(mem.PageStart(2), []byte(str))

	return ret.Bytes()
}

func main() {
	c, e := img.Make(bytes.NewBuffer(makeMap()))
	if e != nil {
		panic(e)
	}

	c.Stdout = os.Stdout
	c.SetPC(mem.PageStart(1))

	c.Run(200)

	if !c.RIP() {
		panic("error occured")
	}
}
