package filemanager

import (
	"bufio"
	"errors"
	"os"
)

type Filemanager struct {
	inputFilePath string
	outputDirPath string
}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		inputFilePath: inputPath,
		outputDirPath: outputPath,
	}
}

func (fm Filemanager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.inputFilePath)

	if err != nil {
		return nil, errors.New("failed to open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read input file")
	}

	return lines, nil
}

func (fm Filemanager) CreateEmptyDir() error {
	err := os.MkdirAll(fm.outputDirPath, 0777)
	return err
}
