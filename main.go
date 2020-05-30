package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			run()
		case "fork":
			fork()
		default:
			log.Fatalln("run or fork are the only commands for dkr")
		}
	} else {
		log.Fatalln("no arguments found, use run or fork to boot dkr")
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

	var home string
	var chrootDir string

	minirootfsHome := os.Getenv("MINIROOTFS_HOME")
	if minirootfsHome != "" {
		home = minirootfsHome
		chrootDir = home + "/minirootfs"
	} else {
		home = os.Getenv("HOME")

		if home == "/" {
			chrootDir = home + "minirootfs"
		} else {
			chrootDir = home + "/minirootfs"
		}
	}

	log.Println("chrootDir:", chrootDir)

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
