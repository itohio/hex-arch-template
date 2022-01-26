package main

import (
	"log"
	"syscall/js"
)

func init() {
	log.SetOutput(&jsWriter{
		Value: js.Global().Get("console"),
		fname: "log",
	})

}

type jsWriter struct {
	js.Value
	fname string
}

func (j jsWriter) Write(b []byte) (int, error) {
	j.Value.Call(j.fname, string(b))
	return len(b), nil
}
