package cliapp

import (
	"fmt"
	"os"
	"os/signal"
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

func shutdownProcess(pid int) error {
	// Find the process by PID
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("error finding process: %v", err)
	}

	// Send SIGTERM signal to the process
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		return fmt.Errorf("error sending SIGTERM signal: %v", err)
	}

	return nil
	// Wait for the process to exit
	// waitForProcessExit(process)
	// fmt.Println("Process with PID", pid, "has gracefully shut down.")
}

func waitForProcessExit(process *os.Process) {
	// Create a channel to receive signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGCHLD)

	// Wait for the process to exit
	<-sigCh
}
