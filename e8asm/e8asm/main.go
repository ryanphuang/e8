package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/e8asm"
	"github.com/h8liu/e8/img"
	"github.com/h8liu/e8/vm/mem"
)

var (
	stepLog = flag.Bool("l", false, "log every entry")
	stdout  = flag.String("o", "", "stdout file")
	regdmp  = flag.String("d", "", "register result dump")
	quiet   = flag.Bool("q", false, "hide stats")
)

func err(cond bool, e interface{}) {
	if !cond {
		return
	}
	fmt.Fprintln(os.Stderr, "error:", e)
	os.Exit(1)
}

func makeImage(buf []byte, out io.Writer) error {
	w := img.NewWriter(out)

	e := w.Write(mem.PageStart(1), []byte("Hello, world.\n\000"))
	if e != nil {
		return e
	}

	if e = w.Write(mem.PageStart(2), buf); e != nil {
		return e
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	err(len(args) == 0, "no input file")
	err(len(args) > 1, "too many input files")

	fin, e := os.Open(args[0])
	err(e != nil, e)
	defer fin.Close()

	asmBuf := new(bytes.Buffer)
	asm := &e8asm.Assembler{
		In:  fin,
		Out: asmBuf,
	}
	e = asm.Assemble()
	err(e != nil, e)

	imgBuf := new(bytes.Buffer)
	e = makeImage(asmBuf.Bytes(), imgBuf)
	err(e != nil, e)

	core, e := img.Make(bytes.NewBuffer(imgBuf.Bytes()))
	err(e != nil, e)

	if *stdout == "" {
		core.Stdout = os.Stdout
	} else {
		fout, e := os.Create(*stdout)
		err(e != nil, e)
		defer fout.Close()
		core.Stdout = fout
	}

	if *stepLog {
		core.Log = os.Stderr
	}
	core.SetPC(mem.PageStart(2))

	cycles := uint32(0)
	for !core.Halted() {
		cycles += core.Run(1000)
	}

	if !*quiet {
		fmt.Fprintf(os.Stderr, "(%d cycles)\n", cycles)
	}

	if *regdmp == "-" {
		core.PrintTo(os.Stdout)
	} else if *regdmp != "" {
		fdmp, e := os.Create(*regdmp)
		err(e != nil, e)
		defer fdmp.Close()
		core.PrintTo(fdmp)
	}
}