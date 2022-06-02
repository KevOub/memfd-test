package main

import (
	_ "embed"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/KevOub/memfd-test/pkg/dropper"
)

var (
	//go:embed samplebins/hello
	data []byte
	// mainfile string
	mainfile = "cGFja2FnZSBtYWluCgppbXBvcnQgKAoJImZtdCIKCSJ0aW1lIgopCgpmdW5jIG1haW4oKSB7Cgljb3VudGVyIDo9IDAKCWZtdC5QcmludCgiSGVsbG8gRGF2ZSEiKQoJZm9yIHsKCQl0aW1lLlNsZWVwKHRpbWUuU2Vjb25kICogNSkKCQlmbXQuUHJpbnRmKCJDT1VOVEVSOiAlZFxuIiwgY291bnRlcikKCQljb3VudGVyICs9IDEKCX0KfQo="
)

func main() {

	// go flags for which file to b64 encode and put into mainfile
	go2build := flag.String("gofile", "", "go file to b64 encode and put into mainfile")
	file2build := flag.String("bin", "", "file to run from memory")
	mode := flag.String("mode", "", "mode to run dropper [build/load/loadexec]")

	flag.Parse()

	if *mode == "build" {
		if *go2build != "" {
			godata, err := ioutil.ReadFile(*go2build)
			if err != nil {
				fmt.Print(err)
			}
			tmp := base64.StdEncoding.EncodeToString([]byte(godata))
			mainfile = tmp
		}

		err := dropper.Build(mainfile)
		if err != nil {
			fmt.Print(err)

		}

	}

	if *mode == "loadexec" {

		if *file2build != "" {
			// open binary file and read it into data
			tmp, err := ioutil.ReadFile(*file2build)
			if err != nil {
				fmt.Print(err)
			}
			data = tmp
		}

		err := dropper.LoadAndExec(data)
		if err != nil {
			fmt.Print(err)
		}

	}

	if *mode == "load" {

		if *file2build != "" {
			// open binary file and read it into data
			tmp, err := ioutil.ReadFile(*file2build)
			if err != nil {
				fmt.Print(err)
			}
			data = tmp
		}

		// this returns a file descriptor to the loaded file
		_, err := dropper.Load(data)
		if err != nil {
			fmt.Print(err)
		}

	}

}
