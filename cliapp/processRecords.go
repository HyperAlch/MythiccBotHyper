package cliapp

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type ProcessRecord struct {
	pid  int
	name string
}

func getProcessRecords() (*ProcessRecord, error) {
	readFileFunc := func(fileName string) (string, error) {
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

	// Create a ProcessRecord using the information from the files
	processRecord := ProcessRecord{
		pid:  pid,
		name: exeName,
	}

	return &processRecord, nil
}
