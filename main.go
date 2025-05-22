package main

import (
	"fmt"

	"example.com/downloader/filemanager"
)

const (
	inputFilePath = "./urls.txt"
	maxWorkers    = 4
)

func main() {
	fm := filemanager.New(inputFilePath, "./res.txt")
	lines, _ := fm.ReadFile()
	fmt.Println(lines)
}
