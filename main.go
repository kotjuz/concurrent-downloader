package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"example.com/downloader/filemanager"
)

const (
	inputFilePath = "./urls.txt"
	maxWorkers    = 4
	outputDirPath = "downloaded/"
)

func main() {

	fm := filemanager.New(inputFilePath, outputDirPath)
	lines, err := fm.ReadFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = fm.CreateEmptyDir()

	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, maxWorkers)

	for _, url := range lines {

		wg.Add(1)
		semaphore <- struct{}{}
		go func(url string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			err = downloadFromUrl(url)
			if err != nil {
				fmt.Println(err)
			}
		}(url)

		downloadFromUrl(url)
	}
	wg.Wait()
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

	filepath := filepath.Join(outputDirPath, filename)

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
