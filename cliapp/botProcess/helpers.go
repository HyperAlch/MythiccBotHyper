package botprocess

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func writeFileFunc(pid int) error {
	pidFileContent := []byte(fmt.Sprintf("%v", pid))
	err2 := os.WriteFile("./.pid", pidFileContent, 0644)
	if err2 != nil {
		return err2
	}

	filename := []byte(filepath.Base(os.Args[0]))
	err2 = os.WriteFile("./.exeName", filename, 0644)
	if err2 != nil {
		return err2
	}

	return nil
}

func readFileFunc(fileName string) (string, error) {
	readFile, err := os.Open(fileName)

	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	if len(fileLines) < 1 {
		return "", errors.New("Record file empty " + fileName)
	}
	return fileLines[0], nil
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
