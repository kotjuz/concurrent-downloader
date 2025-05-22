# Concurrent File Downloader

A simple Go program for **concurrently downloading files from provided URLs**, with a limit on the number of simultaneous downloads (*workers*). This project demonstrates **concurrency in Go** using goroutines, channels, and sync package.

---

## Features

- Reads list of URLs from a file
- Concurrent downloading with a configurable number of workers
- Saves downloaded files into a specified directory
- Configurable via CLI flags

---

# CLI flags
-input string   
    Path to the file containing URLs (default: "urls.txt")

-output string   
    Directory to save downloaded files (default: "downloaded")

-workers int   
    Number of concurrent download workers (default: 4)


-help   
    Display flags


---

## Example usage

```bash
go run main.go -input urls.txt -output downloaded -workers 4
```