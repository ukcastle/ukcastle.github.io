package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"알고리즘": {"자료 구조"},
	"전산학":  {"선형대수"},
	"컴파일러": {
		"자료 구조",
		"형식 언어",
		"컴퓨터 구조",
	},
	"자료 구조":    {"이산수학"},
	"데이터베이스":   {"자료 구조"},
	"이산 수학":    {"프로그래밍 기초"},
	"형식 언어":    {"이산수학"},
	"네트워크":     {"운영체제"},
	"운영체제":     {"자료 구조", "컴퓨터 구조"},
	"프로그래밍 언어": {"자료 구조", "컴퓨터 구조"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) (order []string) {
	seen := make(map[string]bool)

	var visitAll func(items []string) // 먼저 변수를 선언해야 재귀 호출 가능

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return
}
