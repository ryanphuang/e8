package main

import (
	"bytes"
	"os"

	"github.com/h8liu/e8/img"
	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"
)

func pu32(buf []byte, i uint32) {
	buf[0] = uint8(i)
	buf[1] = uint8(i >> 8)
	buf[2] = uint8(i >> 16)
	buf[3] = uint8(i >> 24)
}

func hello() []byte {
	ret := new(bytes.Buffer)
	buf := make([]byte, 4)

	w := func(i inst.Inst) {
		pu32(buf, i.U32())
		ret.Write(buf)
	}

	Rinst := inst.Rinst
	Iinst := inst.Iinst
	Jinst := inst.Jinst

	/*
			add $1, $0, $0		; init counter
		loop:
			lbu $2, 0x2000($1)	; load byte
			beq $2, $0, end 	; +5
		wait:
			lbu $3, 0x5($0)    	; is output ready?
			bne $3, $0, wait 	; -2
			sb $2, 0x5($0)		; output byte
			addi $1, $1, 1		; increase counter
			j loop 				; -7
		end:
			sb $0, 0x4($0)
	*/

	w(Rinst(0, 0, 1, inst.FnAdd))       // 000
	w(Iinst(inst.OpLbu, 1, 2, 0x2000))  // 004
	w(Iinst(inst.OpBeq, 2, 0, 0x0005))  // 008
	w(Iinst(inst.OpLbu, 0, 3, 0x0005))  // 00c
	w(Iinst(inst.OpBne, 3, 0, 0xfffe))  // 010
	w(Iinst(inst.OpSb, 0, 2, 0x0005))   // 014
	w(Iinst(inst.OpAddi, 1, 1, 0x0001)) // 018
	w(Jinst(-7))                        // 01c
	w(Iinst(inst.OpSb, 0, 0, 0x0004))   // 020

	return ret.Bytes()
}

func makeMap() []byte {
	str := "Hello, world.\n"

	ret := new(bytes.Buffer)
	w := img.NewWriter(ret)
	w.Write(mem.PageStart(1), hello())
	w.Write(mem.PageStart(2), []byte(str+"\000"))

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
