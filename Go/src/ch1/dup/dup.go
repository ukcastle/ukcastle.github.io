package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {

	case "1":
		dup1()
	case "2":
		dup2()
	case "3":
		dup3()
	case "p1":
		dupPrac1()
	}

}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	printLines(counts)
}

func dup2() {
	counts := make(map[string]int)

	files := os.Args[2:]

	if len(files) <= 0 {
		countLines(os.Stdin, counts, 0)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, 0)
			f.Close()
		}
	}

	printLines(counts)

}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[2:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	printLines(counts)
}

func dupPrac1() {
	//중복된 줄이 있는 파일명을 출력하라

	counts := make(map[string]int)
	files := os.Args[2:]

	if len(files) <= 0 {
		countLines(os.Stdin, counts, 0)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts, 1) {
				fmt.Println("파일 이름: " + arg)
			}
			printLines(counts)
			counts = make(map[string]int)
			f.Close()
		}
	}

	printLines(counts)
}

func countLines(f *os.File, counts map[string]int, mod int) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	for _, n := range counts {
		if n > 1 {
			switch mod {
			case 0:
				fmt.Println(f.Name())
			case 1:
				return true
			}
		}
	}

	return false
}

func printLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
