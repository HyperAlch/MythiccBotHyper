package botprocess

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type BotProcess struct {
	PID  int
	Name string
}

func (record *BotProcess) GetProcessRecords() error {
	// Get the pid string from the file
	pidStr, err := record.readFileFunc(".pid")
	if err != nil {
		return err
	}

	// Turn the pid into an int
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return err
	}

	// Get the executable name from the file
	exeName, err := record.readFileFunc(".exeName")
	if err != nil {
		return err
	}

	record.PID = pid
	record.Name = exeName

	return nil
}

func (_ *BotProcess) readFileFunc(fileName string) (string, error) {
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
