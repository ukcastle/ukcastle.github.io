package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	switch os.Args[1] {
	case "1":
		echo1()
	case "2":
		echo2()
	case "3":
		echo3()
	case "p1":
		echoPrac1()
	case "p2":
		echoPrac2()
	case "p3":
		echoPrac3()
	default:
		fmt.Println("다시 선택")
	}
}

func echo1() {

	var s string

	for i := 2; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}

	fmt.Println(s)
}

func echo2() {

	var s string

	for _, arg := range os.Args[2:] {
		s += arg + " "
	}

	fmt.Println(s)
}

func echo3() {

	fmt.Println(strings.Join(os.Args[2:], " "))
}

func echoPrac1() {

	// 0번 인자도 같이 출력
	fmt.Println(strings.Join(os.Args[:], " "))

}

func echoPrac2() {
	// for range 인덱스와 값을 한줄에 하나씩 출력
	for index, arg := range os.Args[2:] {
		fmt.Println("Index : " + strconv.Itoa(index) + " Arg : " + arg)
	}
}

func echoPrac3() {
	//strings.Join과 비효율적인 버전의 속도 차이 측정

	var lists []string
	for i := 0; i < 100000; i++ {
		lists = append(lists, os.Args[i%(len(os.Args)-2)+2])
	}

	before := time.Now()
	var s string
	for _, arg := range lists {
		s += arg + " "
	}

	var badVer = time.Since(before)

	before = time.Now()

	s = strings.Join(lists[:], " ")

	fmt.Println(badVer)
	fmt.Println(time.Since(before))

	/*
		badver 11.65초
		join 4.2176 ms ㄷㄷㄷㄷ
	*/
}
