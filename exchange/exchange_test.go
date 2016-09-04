package main

import (
	// "bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

var in io.Reader
var out io.Writer

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
	os.Exit(0)
}

func setUp() {
	fmt.Println("testSetup")
	in = os.Stdin
	out = os.Stdout
}

func tearDown() {
	fmt.Println("tearDown")
}

func ExampleMain() {
	//https://gobyexample.com/command-line-flags
	//http://stackoverflow.com/questions/17412908/how-do-i-unit-test-command-line-flags-in-go
	exchange = "text"
	usage := `Format type. Must be "text", "json" or "hash". Defaults to "text".`
	flag.Var(&exchange, "format", usage)
	flag.Var(&exchange, "f", usage+" (shorthand)")

	main()
	// Output:
	//Exchange 프로그램을 시작합니다.
}

func ExampleInitMsg() {
	initMsg()
	// Output:
	//어떤 작업을 하시겠습니까?
	//1.환전, 2.더하기, 3.곱하기
}

type task string

func (t *task) String() string {
	return fmt.Sprint(*t)
}

func (t *task) Set(value string) error {
	var taskDefault = "default"
	if len(*t) > 0 && string(*t) != taskDefault {
		return errors.New("format flag already set")
	}
	if value != taskDefault && value != "plus" {
		return errors.New("Invalid Format Type")
	}
	*t = "exchange"
	return nil
}

var exchange task

func Example() {
}

// type readerHelper struct {
// 	reader io.Reader
// 	err    error
// }
//
// func TestNewWriter(t *testing.T) {
// 	var helper = newWriter(out)
// 	var msg = "테스트 합니다. "
// 	var length, err = io.WriteString(helper.writer, msg)
//
// 	if length != len(msg) || err != nil {
// 		t.Error("io.WriteString failed")
// 	}
//
// 	io.Copy(out, in)
// 	buf := make([]byte, len(msg))
// 	length, err = in.Read(buf)
//
// 	if length != len(msg) || err != nil {
// 		t.Error("io.Read failed")
// 	}
//
// 	if string(buf) != msg {
// 		t.Error("io Reading Error", string(buf), msg)
// 	}
//
// }
// func TestFiveDollarInput(t *testing.T) {
// 	entered := "12"
//
// 	r := bufio.NewReader(in)
// 	w := bufio.NewWriter(out)
// 	rw := bufio.NewReadWriter(r, w)
//
// 	var writeErr = question(rw, "어떤 작업을 하시겠습니까?")
//
// 	if writeErr != nil {
// 		t.Error("writeErr Error", writeErr)
// 	}
// 	// var delim byte = ' '
// 	// t.Error(rw.ReadString(delim))
// 	// var answer, readErr = answer(*rw, entered, t)
// 	var reader = answer(r, entered, t)
// 	if reader.err != nil {
// 		t.Error("Calling answer() is failed", reader.err)
// 	}
// 	// var length, err = r.Read(buf)
// 	buf := make([]byte, len(entered))
// 	var length, err = reader.reader.Read(buf)
// 	if length != len(entered) {
// 		t.Error("answer() returns wrong string : ", err)
// 	}
// }
//
// func answer(r io.Reader, entered string, t *testing.T) *readerHelper {
// 	buf := make([]byte, len(entered))
// 	var length, err = r.Read(buf)
// 	if length != len(entered) {
// 		t.Error("answer() returns wrong string : ", err)
// 	}
// 	return &readerHelper{reader: r}
// }
