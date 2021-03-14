package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "1":
		fetch1()
	case "p1":
		fetchPrac1()
	case "p2":
		fetchPrac2()
	case "p3":
		fetchPrac3()
	}

}

func fetch1() {

	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)
	}
}

func fetchPrac1() {
	// io.Copy(dst, src)를 이용하고 os.Stdout을 경로로 사용
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		//b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		//fmt.Printf("%s\n", b)
	}
}
func fetchPrac2() {
	//http 누락되었을 경우 붙이기
	for _, url := range os.Args[2:] {

		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

func fetchPrac3() {
	//응답코드 출력
	for _, url := range os.Args[2:] {

		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		//io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
