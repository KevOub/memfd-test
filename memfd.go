// //https://go.dev/play/p/Lv9D8pg-5mt
// // https://github.com/golang/go/issues/51547
// package main

// import (
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
// const (
// 	data = "#!/bin/bash\necho Hello, world!"
// )

// // func fexecv(fd int, argv []byte, envp []byte) (int, error) {
// // 	r0, _, _ := unix.Syscall(281, 5, unix.AT_EMPTY_PATH)
// // }

// // referencing this https://code.woboq.org/userspace/glibc/sysdeps/unix/sysv/linux/fexecve.c.html
// func fexecv(fd int) {
// 	syscall.Syscall(unix.SYS_EXECVEAT, uintptr(fd), 0, 0)

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
// 	fmt.Println(string(buf))
// 	fexecv(fd)
// 	// cmd := "#!/bin/bash\necho Hello, world!"

// 	// for i, c := range data {
// 	// 	tmpByte := []byte{byte(c)}

// 	// 	_, err := unix.Pwrite(fd, tmpByte, int64(i))

// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}

// 	// }

// 	// exec.Command("echo", "hello world").Run()
// 	// syscall.Exec()

// 	// pid, err := syscall.ForkExec("", nil, &syscall.ProcAttr{Dir: "", Files: []uintptr{uintptr(fd)}})
// 	// if err != nil {
// 	// 	log.Printf("Exeerror\n%s\n", err)
// 	// }

// 	// fmt.Printf("Test %v", pid)

// 	// pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
// 	// // in child
// 	// if pid == 0 {
// 	// 	syscall.ForkExec()

// 	// } else if pid == 1 {
// 	// 	// in parent

// 	// }

// 	// fmt.Printf("%v\n", fd)

// }
