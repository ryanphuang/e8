package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/asm"
	"github.com/h8liu/e8/img"
	"github.com/h8liu/e8/vm/mem"
)

var (
	stepLog  = flag.Bool("l", false, "log every entry")
	stdout   = flag.String("o", "", "stdout file")
	regdmp   = flag.String("d", "", "register result dump")
	quiet    = flag.Bool("q", false, "hide stats")
	maxCycle = flag.Int("n", 0, "max cycles")
)

func err(cond bool, e interface{}) {
	if !cond {
		return
	}
	s := fmt.Sprintf("%v", e)
	if s != "" {
		fmt.Fprintln(os.Stderr, "error: ", s)
	}
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

func assemble(fname string) ([]byte, error) {
	fin, e := os.Open(fname)
	if e != nil {
		return nil, e
	}
	defer fin.Close()

	asmBuf := new(bytes.Buffer)
	asm := &asm.Assembler{
		In:       fin,
		Out:      asmBuf,
		Filename: fname,
	}
	e = asm.Assemble()
	return asmBuf.Bytes(), e
}

func main() {
	flag.Parse()
	args := flag.Args()

	err(len(args) == 0, "no input file")
	err(len(args) > 1, "too many input files")

	asmBuf, e := assemble(args[0])
	err(asmBuf == nil && e != nil, e)
	err(e != nil, "") // assemble error already reported, so just exit

	imgBuf := new(bytes.Buffer)
	e = makeImage(asmBuf, imgBuf)
	err(e != nil, e)

	core, e := img.Make(imgBuf)
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

	cycles := 0
	max := *maxCycle
	for !core.Halted() {
		run := 1000
		if max > 0 {
			run = max - cycles
			if run > 1000 {
				run = 1000
			}
		}

		cycles += core.Run(run)
		if max > 0 && cycles >= max {
			break
		}
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
