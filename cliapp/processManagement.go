package cliapp

import (
	"fmt"
	"os"
	"syscall"
)

func runProcess() (*int, error) {
	var pid int
	var system_process = &syscall.SysProcAttr{Noctty: true}
	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
		Sys: system_process,
	}
	process, err := os.StartProcess("./MythiccBotHyper", []string{"MythiccBotHyper", "start", "attached"}, &attr)
	if err == nil {
		pid = process.Pid
		// It is not clear from docs, but Release actually detaches the process
		err = process.Release()
		if err != nil {
			return nil, err
		}

	} else {
		return nil, err
	}

	return &pid, nil
}

func isProcessRunning(pid int) bool {
	// Try to find the process by its PID
	process, err := os.FindProcess(pid)
	if err != nil {
		// If an error occurs, assume the process is not running
		return false
	}

	// Send signal 0 to check if the process exists
	err = process.Signal(os.Signal(syscall.Signal(0)))
	return err == nil
}

func getCommandFromPID(pid int) (string, error) {
	// Construct the path to the command line file
	cmdLinePath := fmt.Sprintf("/proc/%d/cmdline", pid)

	// Read the contents of the command line file
	cmdLineBytes, err := os.ReadFile(cmdLinePath)
	if err != nil {
		return "", err
	}

	// Convert the null-separated byte slice to a string
	// The command line arguments are separated by null bytes in the file
	cmdLine := string(cmdLineBytes)

	return cmdLine, nil
}
