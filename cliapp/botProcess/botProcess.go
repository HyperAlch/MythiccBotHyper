package botprocess

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

type BotProcess struct {
	pid  int
	name string
}

func (botProcess *BotProcess) PID() int {
	return botProcess.pid
}

func (botProcess *BotProcess) Name() string {
	return botProcess.name
}

func SpawnProcess() (*BotProcess, error) {
	botProcess, err := FetchProcess()
	if botProcess != nil {
		if botProcess.IsRunning() {
			return nil, fmt.Errorf("bot with PID %v is still running", botProcess.pid)
		}
	} else {
		return nil, err
	}

	runErr := botProcess.run()
	if runErr != nil {
		return nil, runErr
	}

	writeErr := writeFileFunc(botProcess.pid)
	if writeErr != nil {
		return nil, writeErr
	}

	return botProcess, nil
}

func FetchProcess() (*BotProcess, error) {
	botProcess := BotProcess{
		pid:  0,
		name: "",
	}

	// Get the pid string from the file
	pidStr, err := readFileFunc(".pid")
	if err != nil {
		return nil, err
	}

	// Turn the pid into an int
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return nil, err
	}

	// Get the executable name from the file
	exeName, err := readFileFunc(".exeName")
	if err != nil {
		return nil, err
	}

	botProcess.pid = pid
	botProcess.name = exeName

	return &botProcess, nil
}

func (botProcess *BotProcess) run() error {
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
			return err
		}

	} else {
		return err
	}

	botProcess.pid = pid
	return nil
}

func (botProcess *BotProcess) IsRunning() bool {
	// Try to find the process by its PID
	process, err := os.FindProcess(botProcess.pid)
	if err != nil {
		// If an error occurs, assume the process is not running
		return false
	}

	// Send signal 0 to check if the process exists
	err = process.Signal(os.Signal(syscall.Signal(0)))
	pidValid := err == nil

	execNameLength := len(botProcess.name)
	execName, err := getCommandFromPID(botProcess.pid)
	if len(execName) < execNameLength || err != nil {
		return false
	}
	execNameValid := botProcess.name == execName[0:execNameLength]

	return pidValid && execNameValid
}

func (botProcess *BotProcess) Stop() error {
	// Find the process by PID
	process, err := os.FindProcess(botProcess.pid)
	if err != nil {
		return fmt.Errorf("error finding process: %v", err)
	}

	// Send SIGTERM signal to the process
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		return fmt.Errorf("error sending SIGTERM signal: %v", err)
	}

	return nil
}
