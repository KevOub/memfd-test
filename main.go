package main

import (
	_ "embed"

	"github.com/KevOub/memfd-test/pkg/dropper"
)

var (
	//go:embed samplebins/hello
	data []byte
)

func main() {

	dropper.LoadAndExec(data)

}
