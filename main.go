package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"example.com/downloader/filemanager"
)

const (
	inputFilePath = "./urls.txt"
	maxWorkers    = 4
)

func main() {
	fm := filemanager.New(inputFilePath, "downloaded/")
	lines, _ := fm.ReadFile()
	fmt.Println(lines)

	fm.CreateEmptyDir()

	for _, url := range lines {
		downloadFromUrl(url)
	}
}

func downloadFromUrl(url string) error {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()

	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]

	filepath := filepath.Join("downloaded/", filename)
	fmt.Println(filepath)
	out, err := os.Create(filepath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
