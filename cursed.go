// // https://go.dev/play/p/Lv9D8pg-5mt
// // https://github.com/golang/go/issues/51547
// package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"log"
// 	"syscall"

// 	"golang.org/x/sys/unix"
// )

// // func fexecv(fd int, argv []byte, envp []byte) (int, error) {
// // 	r0, _, err := unix.Syscall(unix.SYS_, uintptr(flags), 0, 0)
// // 	if r0 == 0 {
// // 		return 0, err
// // 	}
// // 	return int(r0), nil
// // }
// // const (

// // data = "#!/bin/bash\necho Hello, test!\necho $$\nsleep 60"
// // )

// var (

// 	//go:embed samplebins/hello
// 	data []byte
// )

// // func fexecv(fd int, argv []byte, envp []byte) (int, error) {
// // 	r0, _, _ := unix.Syscall(281, 5, unix.AT_EMPTY_PATH)
// // }

// // referencing this https://code.woboq.org/userspace/glibc/sysdeps/unix/sysv/linux/fexecve.c.html
// // we abuse the fact /proc/self/fd/NUM will reference a custom file descriptor
// func fexecv(fd int) {

// 	// load the code in memory
// 	cmd := fmt.Sprintf("/proc/self/fd/%d", fd)
// 	// name of the program
// 	args := []string{"[kworker/u!0]", ""}
// 	envv := []string{}

// 	// output, err := exec.Command(test).CombinedOutput()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Print(string(output))
// 	syscall.Exec(cmd, args, envv)
// }

// func main() {

// 	fd, err := unix.MemfdCreate("", 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = unix.Ftruncate(fd, int64(len(data)))
// 	if err != nil {
// 		panic(err)
// 	}
// 	buf, err := unix.Mmap(fd, 0, len(data), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
// 	if err != nil {

// 		panic(err)
// 	}
// 	defer unix.Munmap(buf)

// 	copy(buf, data)
// 	// fmt.Println(buf)
// 	fexecv(fd)

// }
