package vm

import (
	"bytes"
	"testing"

	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"

	// "os"
)

func TestHelloWorld(t *testing.T) {
	c := NewVM()
	// c.Log = os.Stdout
	out := new(bytes.Buffer)
	str := "Hello, world.\n"
	c.Stdout = out

	dpage := mem.NewPage()
	copy(dpage.Bytes(), []byte(str+"\000"))

	ipage := mem.NewPage()
	c.Map(mem.PageStart(1), ipage)
	c.Map(mem.PageStart(2), dpage)

	a := &mem.Align{ipage}

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
			lbu $2, $1[0x2000] 	; load byte
			beq $2, $0, end 	; +5
		wait:
			lbu $3, $0[0x5]    	; is output ready?
			bne $3, $0, wait 	; -2
			sb $2, $0[0x5]		; output byte
			addi $1, $1, 1		; increase counter
			j loop 				; -7
		end:
			sb $0, [0x4]
	*/

	Rinst := inst.Rinst
	Iinst := inst.Iinst
	Jinst := inst.Jinst

	w(Rinst(1, 0, 0, inst.FnAdd))       // 000
	w(Iinst(inst.OpLbu, 1, 2, 0x2000))  // 004
	w(Iinst(inst.OpBeq, 2, 0, 0x0005))  // 008
	w(Iinst(inst.OpLbu, 0, 3, 0x0005))  // 00c
	w(Iinst(inst.OpBne, 3, 0, 0xfffe))  // 010
	w(Iinst(inst.OpSb, 0, 2, 0x0005))   // 014
	w(Iinst(inst.OpAddi, 1, 1, 0x0001)) // 018
	w(Jinst(-7))                        // 01c
	w(Iinst(inst.OpSb, 0, 0, 0x0004))   // 020

	c.SetPC(mem.PageStart(1))
	left := c.Run(1000)

	if left < 850 {
		t.Fail()
	}

	if !c.Halt || c.HaltValue != 0 || c.AddrError {
		t.Fail()
	}

	if out.String() != str {
		t.Fail()
	}
}
