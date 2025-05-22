package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"example.com/downloader/filemanager"
)

func main() {
	maxWorkers := flag.Int("workers", 4, "Number of concurrent download workers")
	inputFilePath := flag.String("input", "urls.txt", "Path to input file")
	outputDirPath := flag.String("output", "downloaded", "Path to outpu directory")
	flag.Parse()

	fm := filemanager.New(*inputFilePath, *outputDirPath)
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

	semaphore := make(chan struct{}, *maxWorkers)

	for _, url := range lines {

		wg.Add(1)
		semaphore <- struct{}{}
		go func(url string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			err = downloadFromUrl(url, *outputDirPath)
			if err != nil {
				fmt.Println(err)
			}
		}(url)

		downloadFromUrl(url, *outputDirPath)
	}
	wg.Wait()
}

func downloadFromUrl(url string, outputDirPath string) error {
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
