package vm

import (
	"bytes"
	"testing"
	// "os"
)

func TestHelloWorld(t *testing.T) {
	c := NewVM()
	// c.Log = os.Stdout
	out := new(bytes.Buffer)
	str := "Hello, world.\n"
	c.Stdout = out

	dpage := NewDataPage()
	copy(dpage.bytes, []byte(str+"\000"))

	ipage := NewDataPage()
	c.Map(PageStart(1), ipage)
	c.Map(PageStart(2), dpage)

	a := &Align{ipage}

	offset := uint32(0)
	w := func(i uint32) uint32 {
		ret := offset
		a.WriteU32(offset, i)
		offset += 4
		return ret
	}

	/*
			add $1, $0, $0		; init counter
		loop:
		wait:
			lbu $2, $1[0x2000] 	; load byte
			beq $2, $0, end 	; +5
			lbu $3, $0[0x5]    	; is output ready?
			bne $3, $0, wait 	; -4
			sb $2, [0x5]  		; output byte
			addi $1, $1, 1		; increase counter
			j loop 				; -7
		end:
			sb $0, [0x4]
	*/

	w(Rinst(1, 0, 0, FnAdd))       // 000
	w(Iinst(OpLbu, 1, 2, 0x2000))  // 004
	w(Iinst(OpBeq, 2, 0, 0x0005))  // 008
	w(Iinst(OpLbu, 0, 3, 0x0005))  // 00c
	w(Iinst(OpBne, 3, 0, 0xfffc))  // 010
	w(Iinst(OpSb, 0, 2, 0x0005))   // 014
	w(Iinst(OpAddi, 1, 1, 0x0001)) // 018
	w(Jinst(-7))                   // 01c
	w(Iinst(OpSb, 0, 0, 0x0004))   // 020

	c.SetPC(PageStart(1))
	c.Run(1000)

	if !c.Halt || c.HaltValue != 0 || c.AddrError {
		t.Fail()
	}

	if out.String() != str {
		t.Fail()
	}
}
