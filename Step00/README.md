# Step 0. GO 언어 기초

## GO 언어
GO는 매우 강력한 기능을 지닌 언어입니다. TDD에 직접 뛰어들기 전에
조금은 따분한 내용으로 첫 시작을 하려고 하는데 우선 환경 설치입니다.

### GO 설치하기

저는 Ubuntu 환경에서 Fish Shell 과 Atom을 가지고 개발환경을 만들었습니다.
이 글에 있는 소스를 실행하기 위해서 저와 동일한 환경을 갖출 필요는 없습니다.
각자의 환경에 맞는 에디터와 쉡이 존재할 것입니다. 각자의 환경에 맞는 설치는 구글신에게 물어보시기 바랍니다.

아무리 환경이 다르더라도 저와 같은 환경을 꾸미고 싶으신 분들에게 제가 설정한 내용을
공유합니다.

Ubuntu 환경에서는 Go 설치가 아주 쉽습니다.

```
apt-get install golang

apt-cache policy golang
설치: 2:1.5.1-0ubuntu2
후보: 2:1.5.1-0ubuntu2
버전 테이블:
*** 2:1.5.1-0ubuntu2 0
      500 ~~~ / wily/universe amd64 Packages
      100 /var/lib/dpkg/status
```

Fish Shell의 설정은 다음과 같습니다. Bash의 export 가 아니어서 매우 당황스럽죠?
```json
$ cat ~/.config/fish/config.sh

set -x GOPATH $HOME/Projects/CodeGo
set -g -x PATH $GOPATH/bin $PATH
set -x TMPDIR $HOME/tmp
```

Atom 에서는 go-plus, go-rename 패키지를 설치해야 합니다. 설치 후에 설정은 다음과
같습니다.

```
GOPATH : ~/Projects/CodeGo
GO Installation Path : /usr/bin/go
```

## GO 언어의 특징

프로그램 언어를 다루는 각종 책과 마찬가지로 GO 언어의 특징을 이야기하려고 합니다.
너무 식상한 구성이라서 이 부분을 넣어야 할지 고민했습니다. GO 언어를 어느 정도
이해하고 계신 분들은 다음 내용을 건너뛰셔도 됩니다.

#### Compiled Language

** C, C++ ** 과 같이 Assembly 언어로 컴파일되는 언어입니다.
GO의 중요 설계 목표중의 하나가 컴파일 속도였다는 것을 볼때 GO는 \** C,C++ \** 을
확실히 승계하는 언어로 인정받을 수 있을 것 같습니다.
Java 에 익숙한 개발자인 저로서는 이점이 분명 강점으로 여겨졌습니다.

#### 정적 형식

변수는 반드시 특정한 데이터 형 중의 하나여야 합니다. (int, string, bool, []byte, etc.)
이는 변수의 선언시에 정의될 수도 있고, 컴파일러가 (Scala 에서 처럼) 데이터 형을
추정하여 지정할 수도 있습니다. 하지만, 추정이 불가능할 경우에 컴파일 에러를 반환합니다.
그럼에도 GO의 컴파일러는 그렇게 딱딱하게 굴지 않은데 좀더 유연한 데이터 형 체계를 가지고 있다.

#### C를 닮은 문법

Java, C# 등도 C를 닯은 언어이긴 하지만 이들 언어보다 더욱 간결하고 단순한 문법을 가지고 있습니다.

#### Garbage Collector

** Ruby, Python, Java, C# ** 와 같은 가비지 컬렉터를 가지고 있습니다.
가비지 컬렉터가 어느 정도의 오버헤드를 발생시킨다는 것은 인정하지만,
일일이 메모리 해제를 수행해야 하는 불편함을 일부러 겪을 이유도 없습니다.
GO는 Virtual Machine을 가지고 있지 않기 때문에 GO의 가비지 컬렉터는 컴파일 과정에서 실행파일에 내장됩니다.

#### 병행성 (Concurrency)

GO 가 갖는 장점중에 병행성 처리가 좀더 수월하다는 것이 가장 강조됩니다.
실행파일은 GO 함수들로 이루어진 GO 루틴으로 구성되는 각각의 GO 루틴은 필요한 스레드 수에 따라 스레드내에서 시분할로 처리됩니다.
이때 채널을 통해서 GO 루틴간 호출을 제어합니다.

#### 버전 관리 시스템 연동

GO 는 버전관리 시스템을 통해서 외부 패키지를 import 하는 기능을 가지고 있습니다.
`go get` 이나 `go install` 은 자주 사용하는 명령어입니다.

## 컴파일과 실행                                                                                                                                                                                                                                                                                                                                                                                                             
GO를 컴파일하고 실행하는 방법은 아주 간단합니다.

```json
go build main.go
./sample.org

```

좀더 간단하게는 아래와 같이 실행할 수 있습니다. `--work`는 임시파일의 위치를 알려줍니다.

```json
go run main.go
go run --work main.go
```

### main

GO 프로그램을 실행하려면 프로그램 안에 `main` 패키지와 그 패키지 안에 `main` 함수를 가지고 있어야 합니다.

```go
package main

func main() {
	println("it's over 9000!")
}
```

### import

필요한 패키지는 `import` 를 사용해 추가합니다.

```go
import (
	"fmt"
	"os"
)
```

### 변수 선언

#### 명시적 선언

아래와 같이 명시적으로 형식을 지정할 수 있습니다.

```go
func declareExplicit() {
	var power int
	power = 10000
	fmt.Printf("Value is  %d\n", power)
}
```

#### 형 추정 (Type Inference)

컴파일러가 추정을 통해 데이터의 형식을 지정하기도 합니다. 이를 위해 좀 이상하지만 `:=` 기호를 사용한다.

```go
func declareImplicit() {
	name, power := "한글날, ", 2015
	fmt.Printf("Value is %s %d\n", name, power)
}

```

#### 선언 오류

`:=` 은 때로는 오류를 발생시킬 수 있는데 `:=` 의 왼쪽은 항상 새로운 변수가 와야 합니다.
따라서 변수를 다시 지정하려 할 경우 오류를 발생시킵니다.
또한 묵시적으로 지정된 형식을 변경하려는 경우도 아주 점잖게 형을 변경할 수 없다고 알려줍니다.

```go
func declareError() {
	power := 9000
	fmt.Printf("Value is %d\n", power)

	//Compile Error
	//no new variables on left side of :=
	power := 9001
	fmt.Printf("Value is %d\n", power)

	//Compile Error
	//cannot use "Change Type" (type string) as type int in assignment
	name, power := "Test", "Change Type"
	fmt.Printf("Value is %s %s\n", name, power)
}
```

형 변경할 수 없는 경우에도 오류가 발생합니다.

```go
	//Compile Error
	// invalid operation: a + b (mismatched types int and float32)
	var a int = 1
	var b float32 = 1.3
	var c float32 = a + b
```

#### 예외적인 선언

중복해서 형을 지정하지 못하는 문제는 아래와 같은 트릭으로 변경할 수 있습니다.

```go
func declareNoError() {
	power := 1000
	fmt.Printf("default power is %d\n", power)
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}
```

### 함수 선언

함수 선언은 C 문법과는 약간 다릅니다.
`declareFunction()` 의 첫번째 줄은 반환된 값을 어떻게 받을 수 있는지 보여줍니다.

```go
func declareFunction() {

	value, exists := power("goku")
	if exists == false {
		fmt.Printf("Something is wrong! %b\n", exists)
	}
	fmt.Printf("Name is %d", value)

}

func power(name string) (int, bool) {
	fmt.Printf("Name is %s", name)
	return 0, false
}
```

#### Blank 인식자

GO 에서 Blank Identifier (`_`) 는 실재로는 값을 할당하지 않는 특별한 지시어가 존재합니다.
반환되는 데이터형에 상관없이 반복해서 사용할 수 있습니다.
저는 예전에 코딩 규범(Coding Convention)에 이런 것이 있으면 강력하게 반발하곤 했습니다.
GO 언어을 사용하면 적어도 이 부분에 대한 논쟁은 사라지게 되었습니다.

```go
_, exists := power("goku")
if exists == false {
	// handle this error case
}
```
#### 지연 처리와 에러 처리

GO에서는 지연 처리 개념이 있습니다.
아래에서 defer 키워드가 붙은 구문은 main 함수가 종료될 때 처리됩니다.
순서는 함수내의 위치에서 역순입다.
즉, 맨 마지막에 정의된 defer 구문이 그 보다 앞에서  정의된  defer 보다 먼저 실행됩니다.

```go
func main() {
	// Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

    ... 파일 작업 ....
}
```

에러는 바로 GO의 지연 처리를 이용해서 사용자 Exception을 처리할 수 있습니다.
사용자 Exception은 `panic()` 함수를 이용해 정의하면 됩니다.
아래 코드는 err를 반환하지 않고 바로 Exception을 호출합니다.

```go
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

```

아래는 `recover()` 함수의 사용법을 보여 줍니다.
```go
func Parse(input string) (numbers []int, err error) {
    defer func() {
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok {
                err = fmt.Errorf("pkg: %v", r)
            }
        }
    }()

    fields := strings.Fields(input)
    numbers = fields2numbers(fields)
    return
}

```

### 몇가지 실행해 보기

#### 설치 환경 점검

```json
$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/jacob/Projects/CodeGo"
GORACE=""
GOROOT="/usr/lib/go"
GOTOOLDIR="/usr/lib/go/pkg/tool/linux_amd64"
TERM="dumb"
CC="gcc"
GOGCCFLAGS="-g -O2 -fPIC -m64 -pthread"
CXX="g++"
CGO_ENABLED="1"
```

#### 설치 후 추가할 패키지들

GO 개발환경을 꾸밀 때 자주 사용하는 패키지들이 있습니다.
자세한 정보는 https://godoc.org/golang.org/x/tools 를 참조합니다.

```json
$ go get -u golang.org/x/tools/cmd/oracle
$ go get -u golang.org/x/tools/cmd/vet
$ go get -u golang.org/x/tools/cmd/goimports

```

#### GODOC 을 사용하기 (Ubuntu)

언어를 배우면서 가장 난감한 부분이 표준 라이브러리에 대한 도움말을 조회하려고 할 때 어디서 찾아야 할지 몰라 당황스러울 때가 있습니다.
GO는 이런 문제를 최소화하기 위해서 godoc 을 제공합니다.
GoDoc을 사용하기 위해서 아래와 같이 실행해야 합니다.

```json
$ sudo apt-get install golang-doc
$ sudo apt-get install golang-go.tools
$ godoc fmt
$ godoc -http=:6060
```

#### 외부에서 도움말 얻기

godoc 을 실행하지 않고 외부 사이트의 도움을 얻을 수도 있습니다.

-	https://godoc.org/
-	https://golang.org/pkg/
