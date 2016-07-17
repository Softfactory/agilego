# Step 09. 실행가능한 프로그램

## 입출력 사용하기

다음 할일은 무엇인가요?
미리 정의된 것이 없다면 다소 혼란스러울 수 있습니다.
1부의 말미에 몇가지 수정할 사항을 추가했습니다.

> **TODO**

> CLI로 동작하는 exchange 프로그램 만들기

> 웹으로 동작하는 exchange-web 프로그램 만들기

> 실시간 환율 정보 가져오기

할일을 적어 보이지만 그렇게 쉬워보이지 않습니다.
첫번째 할일을 좀더 명확하게 정의해야 합니다.
첫번째 할일에 대한 상세한 요구사항을 나열하면 다음과 같이 될 것입니다.  

> **TODO**

> 1. CLI에서 통화기호를 물어보면 제공된 통화기호 중 원하는 통화기호를 선택해서 입력한다.
> 1. 다음으로 프로그램은 금액에 대해서 질문한다. 사용자는 원하는 금액을 입력한다.
> 1. 프로그램은 어떤 통화로 변환할 것인지 질의한다. 원하는 통화기호를 입력한다.
> 1. 사용자는 최종 입력한 통화기호로 연산된 금액을 볼 수 있다.

너무도 길고 긴 상호작용을 우리는 위험천만하게도 한줄의 한일로 끝내려고 했습니다.
또한가지 문제가 더 있는데 바로 앞의 할일은 너무도 장황해서 테스트로 정의해서 사용하기에는 적합하지 않습니다.
더하기 또는 빼기에 사용될 좀더 일반화된 할일을 정의해야 합니다.

> **TODO**

> 프로그램은 어떤 작업을 할지 사용자에게 물어본다. 사용자는 제시된 작업중 하나를 선택한다.

> 프로그램은 선택된 작업에 필요한 피연산자를 입력받는다.
'1번 피연산자를 입력하세요.' 와 같은 메시지가 출력되고 사용자가 입력한다.

> 프로그램에 피연산자 입력이 완료되면 프로그램은 작업에 대한  결과값을 출력한다.

### Main

이제 어떤 작업을 할지 사용자에게 물어보는 실행 가능한 프로그램을 만들 차례입니다.
`exchange` 폴더를 만들고 `exchange.go`와 `exchange_test.go` 파일을 만들어 넣습니다.
패키지는 `main` 입니다.
지금 단계면 테스트 코드를 아래와 같이 쉽게 작성할 수 있어야 합니다.
```go
import (
	"testing"
)

func TestMain(t *testing.T) {
	main()
}
```

아무것도 하지 않는 `main` 함수를 만들어서 출력을 확인합니다.
```go
func main() {
	fmt.Println("어떤 작업을 하시겠습니까?")
}
```
테스트가 성공합니다.
` go run exchange.go` 를 실행하면 우리가 원하는 메시지가 출력됩니다.
이번에 수행한 테스트가 안전하다고 생각하신다면 아직 TDD에 완벽히 적응했다고 보기 어렵습니다.
'어떤 작업을 하시겠습니까?' 도 테스트가 되어야 합니다.

### 콘솔 출력값 테스트

세상에나 콘솔 출력값을 중간에 가로체기라도 하려고 하는가하고 의심하실 필요가 없습니다.
아주 간단한 방식으로 테스트가 가능합니다.

```go
func ExampleMain() {
	main()
	// Output: 어떤 작업을 하시겠습니까?
}
```
`Output`의 결과와 다르면 테스트는 실패합니다.

```
=== RUN   ExampleMain
--- FAIL: ExampleMain (0.00s)
got:
Show me  the Money!
want:
Show me the Money!
FAIL
exit status 1
FAIL	AgileGO/exchange	0.001s
```
GO 설계자들이 원래 의도했던 것은 아니지만 테스트를 위한 방법으로는 아주 좋습니다.

### 테스트 셋업

이번 장에서는 GO의 테스트에 대해서 좀더 깊이있는 부분에 대해서 살펴봐야 합니다.
CLI로 실행되는 프로그램은 일반적인 소스코드를 테스트하는 방법과는 달라야 합니다.
`main` 은 실행프로그램에서 첫번째 호출되는 함수입니다.
`main`가 실행될 때 OS의 리소스를 가져옵니다.
일반적인 소스코드를 테스트하는 코드로 Runtime에 발생되는 리소스의 환경을 반영할 수 없습니다.   
GO는 테스트 전후에 필요한 작업을 수행할 수 있도록 `testing.M` 구조체를 가지고 있습니다.
메인 스레드에서 동작하는 코드를 테스트할 때도 이 구조체를 이용해야 합니다.

```go
func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
	os.Exit(0)
}

func setUp() {
	fmt.Println("testSetup")
}

func tearDown() {
	fmt.Println("tearDown")
}
```
테스트를 실행하면 다음과 같이 출력됩니다.
```go
testSetup
=== RUN   ExampleMain
--- PASS: ExampleMain (0.00s)
PASS
tearDown
ok  	AgileGO/exchange	0.001s
```

### 기본 입출력 인터페이스

GO는 `io.Reader`와 `io.Writer` 인터페이스를 이용해서 입출력을 처리합니다.
이들 인터페이스를 활용한 패키지 중의 하나가 `fmt` 입니다.
아래에서 `fmt.Println()` 소스를 보면 `Writer`로 `os.Stdout` 을 사용합니다.

```go
func Println(a ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a, true, true)
	n, err = w.Write(p.buf)
	p.free()
	return
}   
```
따라서 예제의 소스 코드를 아래와 같이 변경해도 동일한 결과를 얻을 수 있습니다.
```go
s := "어떤 작업을 하시겠습니까?"
src := strings.NewReader(s)
io.Copy(os.Stdout, src)
```
예제로 보여주기 위한 것이지 실제 코드에서 이렇게 코딩할 일은 많지 않을 것입니다.
 `fmt`를 사용하면 입출력을 구현하는데 문제가 없습니다.
소켓의 입출력을 직접 구현하거나 할때도 앞서 언급한 `io.Reader` 와 `io.Writer` 를 사용합니다.
`googollee/go-socket.io` 패키지중에서 `ioutil.go` 소스를 보면 이해가 더 쉽습니다.

```go
type writerHelper struct {
	writer io.Writer
	err    error
}

func newWriterHelper(w io.Writer) *writerHelper {
	return &writerHelper{
		writer: w,
	}
}
```
## CLI 프로그래밍

### 작업 선택 물어보기

> **TODO**

> <b>프로그램은 어떤 작업을 할지 사용자에게 물어본다. 사용자는 제시된 작업중 하나를 선택한다.</b>

이제 CLI 프로그램을 만들어볼 차례입니다.
사용자 스토리는 아주 단순합니다.
앞에서 출력값 테스트 코드를 활용해서 조금만 변경합니다.

```go
func ExampleInitMsg() {
	initMsg()
	// Output:
	//어떤 작업을 하시겠습니까?
	//1.환전, 2.더하기, 3.곱하기
}
```
소스 코드도 이에 맞게 변경합니다.
```god
func initMsg() {
	fmt.Println("어떤 작업을 하시겠습니까?")
	fmt.Println("1.환전, 2.더하기, 3.곱하기")
}
```
이렇게 하고나면 `main()`에 적어놓은 문장이 중복이 됩니다.
프로그램이 시작되면 프로그램의 시작을 알리는 메시지를 출력하도록 하는 것이 좋을 것 같습니다.
```go
func ExampleMain() {
	main()
	// Output:
	//Exchange 프로그램을 시작합니다.
}
```
소스 코드에 위문장과 동일한 메시지로 수정합니다.

### 파라미터 처리
현재까지 작성한 소스 코드는 `main()` 과 `initMsg()` 가 있지만 둘은 어떠한 연관도 없습니다.
`main()`에서 아무런 조건 없이 `initMsg()`를 호출하면 `ExampleMain`은 실패합니다.
`main()`은 시작 메시지를 제시하고 입력된 파라미터에 따라서 `initMsg()`를 호출해야 합니다.
CLI에서 파라미터를 처리하려면 `flag` 패키지를 이용합니다.
`flag` 패키지의 예제에는 `flag`를 테스트할 단서를 제공합니다.
`flag.Value` 인터페이스를 구현하면 `flag.Var()`에서 사용할 수 있습니다.
`flag.Value` 는 다음과 같이 정의되어 있습니다.
```go
type Value interface {
    String() string
    Set(string) error
}
```


```go
func
```




### Money 사용

이제 테스트를 바꾸어서 Money를 사용하도록 해야 합니다.
그런데 실행프로그램이 정확히 무슨 일을 하는지 정의되지 않았습니다.
사용자 스토리를 다시 참조해야 합니다.

> 사용자A는 가상의 거래 시스템을 이용해서 다중 통화의 가중평균을 구할 수 있다.

> 가중평균을 구하기 위해 통화는 더하기 및 곱하기 연산을 지원해야 한다.

> 거래 주식수에 주식의 가격을 곱하여 총 금액을 얻을 수 있어야 한다.

실행프로그램은 정확히 사용자가 스토리의 기능을 수행해야 합니다.

> **TODO**

> $5와 5,000원을 입력하면 10,000원이 출력됩니다.

이 기능을 구현하려면 좀 귀찮게 사용자에게 요구할 것이 많습니다.
너무 많은 기능들을 구현해야 하기 때문에 한번에 테스트하기는 힘들어 보입니다.

> **TODO**

> $5를 입력받기

위의 할일 목록과 같이 일을 단순한 형태로 진행합니다.
이제 입력받는 것을 어떻게 테스트할 것인지가 중요해 졌습니다.

### 입력값 테스트

입력은 기본적으로 사용자와의 상호작용을 전제로 합니다.
그래서 함수명을 좀 창의적으로 이름인 `dialogue`를 붙였습니다.
이 함수는 원하는 질의를 인수로 받아서 화면에 표시하고 사용자가 입력한 값을 반환합니다.


```go
func TestFiveDollarInput(t *testing.T) {
	answer := dialogue("원하는 화폐를 고르세요.")

	if answer != "1" {
		t.Error("입력값 오류")
	}

}
```

`fmt.Println()` 에서 `os.Stdout`을 사용한다고 했는데 실제 소스를 보면 다음과 같이 정의되어 있습니다.

```go
var Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
```




### Example 이름짓기
func Example() // package
func ExampleF() // function
func ExampleT() // type
func ExampleT_M() // method
func Example*_xyz()// more...


### Coverage 테스트

https://www.elastic.co/blog/code-coverage-for-your-golang-system-tests

https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742#.vv5tztg2k

https://splice.com/blog/lesser-known-features-go-test/

testing/iotest
testing/quick
net/http/httptest
