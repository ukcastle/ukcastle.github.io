package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", true, "hint") // ,명령행인자, 호출값, 디폴트값, 힌트
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) // 명령행인자에 접근할 땐 포인터 호출

	if !(*n) {
		fmt.Println("dd")
	} else {
		fmt.Println("ss")
	}

	a := new(int)
	*a = 10

}
