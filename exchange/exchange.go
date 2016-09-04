package main

import (
	// "AgileGO/money"
	// "bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	s := "Exchange 프로그램을 시작합니다."
	src := strings.NewReader(s)
	io.Copy(os.Stdout, src)
	initMsg()
}

func initMsg() {
	fmt.Println("어떤 작업을 하시겠습니까?")
	fmt.Println("1.환전, 2.더하기, 3.곱하기")
}
