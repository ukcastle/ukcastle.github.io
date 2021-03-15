package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	switch os.Args[1] {
	case "1":
		fetchAll()

	}
}

func fetchAll() {
	start := time.Now()
	ch := make(chan string) // 문자열의 채널을 만듬

	for _, url := range os.Args[2:] {
		go fetch(url, ch) // go 루틴 시작
	}
	for range os.Args[2:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
