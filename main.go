package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "fork":
		fork()
	default:
		panic("nope")
	}
}

func run() {
	log.Println("run() executing", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"fork"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	attempt(
		"run() Run",
		cmd.Run(),
	)
}

func fork() {
	log.Println("fork() executing", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	chrootDir := os.Getenv("HOME") + "minirootfs"
	attempt(
		"fork() Chroot",
		syscall.Chroot(chrootDir),
	)

	attempt(
		"fork() Chdir",
		os.Chdir("/"),
	)

	attempt(
		"fork() Mount",
		syscall.Mount("proc", "proc", "proc", 0, ""),
	)

	attempt(
		"fork() Run",
		cmd.Run(),
	)

	attempt(
		"fork() Unmount",
		syscall.Unmount("proc", 0),
	)
}

func attempt(cmd string, err error) {
	if err != nil {
		log.Println("CMD:", cmd)
		log.Println("ERR:", err)
	}
}
