package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 147쪽 연습문제
// 1번 findlinks 루프대신 재귀 호출로 연결리스트 n.FirstChild를 탐색하게 변경하라
// 2번 원소 이름(p, div, span 등)과 HTML 문서 트리 내에서 원소 개수의 맵을 생성하는 함수 작성
// 3번 HTML 문서 트리 안에 있는 모든 텍스트 노드의 내용을 출력하는 함수 작성
// 4번 문서에서 이미지, 스크립트, 스타일시트 같은 다른 종류의 링크도 추출하게 visit 함수 확장
func main() {

	url := os.Args[2]

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	switch os.Args[1] {
	case "0":
		findLink(resp.Body)

	case "1":
		p1_findLink(resp.Body)
	}

	resp.Body.Close()

}

func findLink(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func p1_findLink(r io.Reader) {
	// 1번 findlinks 루프대신 재귀 호출로 연결리스트 n.FirstChild를 탐색하게 변경하라
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func p1_vist(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links

}
