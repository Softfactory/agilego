# Step 08. 정규화

## 상수대첩

### 통화기호 상수화

> **TODO**

> 통화기호를 상수로 바꾸기

통화기호를 상수로 바꾸는 것은 GO의 `const`를 사용해서 해결할 수 있습니다.
문제는 상수를 어떻게 테스트할 것인가입니다.  

간단하게 테스트를 작성하면 아래와 같이 하면 됩니다.
```go
func TestConstCurrency(t *testing.T) {
	if KRW != "KRW" {
		t.Error("money.KRW is not KRW, actually ", KRW)
	}
}
```
컴파일 에러가 분명하게 나타나서 소스를 수정합니다.

```go
const KRW = "KRW"
const USD = "USD"
```
테스트는 정상입니다.
이대로 정리하면 우리는 통화를 모두 위와 같은 `const` 형태로 정의해야 합니다.

GO에서 문자열을 배열이나 슬라이스에 담아서 상수화하는 방법이 없을까요?
아쉽게도 GO에서는 그러한 [요청](https://github.com/golang/go/issues/6386)을 수용하지 않습니다.
혹시라도 GO 차기버전인 GO2에서는 가능할지도 모릅니다.

문자열 슬라이스를 변수로 선언하는 방법도 있을 겁니다.

```go
var Currency = []string{"KRW", "USD"}
```
하지만 테스트 코드는 우리가 원하는 대로 작성될 수 없습니다.
`Currency[0] != "KRW"` 와 같이 슬라이스의 인덱스를 사용해야 하기 때문입니다.

소스에 151개의 통화기호를 추가해야 합니다.
깨끗한 코드를 원하는 개발자들이 보기에 너무 지저분해집니다.

별도 파일로 `currency.go`를 만들어서 통화기호에 관련된 상수를 모두 넣어 놓습니다.
이제 테스트 코드를 `example_test.go`로 이동시킵니다.
왜냐하면 'Exported'인지 정확한 테스트가 필요하기 때문이죠.

### 문자열 리펙토링

이제 소스 코드에 있는 모든 통화기호는 상수를 사용해서 수정되어야 합니다.
생성자에서만 통화기호를 사용했기 때문에 고칠 곳은 2군데 뿐입니다.

```go
func Dollar(amount float64) Money {
	return Money{amount, USD}
}
func Won(amount float64) Money {
	return Money{amount, KRW}
}
```

문제는 테스트 코드인데 수정할 부분이 많습니다.
테스트 코드를 깔끔하게 유지해야 하는가에 대한 논란도 있습니다.
가능한 중복을 없애는 것을 권장합니다.

### 새로운 생성자

> **TODO**

> <s>통화기호를 상수로 바꾸기</s>

> <b>유연한 생성자 만들기</b>

지금 우리가 작성한 코드는 새로운 통화가 추가될 때마다 생성자를 만들어야 하는 부담이 있습니다.
이 기능을 한번에 처리할 수 있는 새로운 생성자를 만들어야 합니다.

```go
func TestNewConstruct(t *testing.T) {
	var fiveWon = Construct(5, KRW)
	if !reflect.DeepEqual(fiveWon, Won(5)) {
		t.Error("Construct Error", fiveWon)
	}
}
```

테스트 코드를 작성하고 소스 코드에 `Construct(amount float64, currency string)`을 추가하니 문제가 생깁니다.
`Construct`가 중복해서 선언되었다는 군요.
함수의 시그니처가 다르지만 GO에서는 이름을 중복해서 사용할 수 없습니다.
기존 코드와의 호환성 유지를 위해서 `currency`에 기본값을 넣어주는 방법이 있습니다.

```go
func Construct(amount float64, args ...string) Money {
	currency := USD
	if len(args) == 1 {
		currency = args[0]
	}
	if len(args) > 1 {
		panic("Too many arguments!")
	}
	return Money{amount, currency}
}
```
이제 주석에서 Deprecated 표식을 지울 수 있습니다.

마무리를 위해서 인수를 3개 이상 받는 경우 실패를 테스트하는 코드를 작성해야 합니다.
```go
func TestFailedConstruct(t *testing.T) {
	defer func() {
		s := recover()
		if s != "Too many arguments!" {
			t.Error("Failed Construct Test Error")
		}
	}()
	var fiveWon = Construct(5, KRW, USD)
	log.Println("Failed Construct test is success ", fiveWon)
}
```
모든 테스트가 정상입니다. 커버리지 표시도 만족스럽습니다.
다른 생성자도 수정할 수 있습니다.

```go
func Dollar(amount float64) Money {
	return Construct(amount, USD)
}
func Won(amount float64) Money {
	return Construct(amount, KRW)
}
```

### 통화 상수 제한하기

만약 우리가 원하는 통화가 아닌 다른 통화기호가 들어온다면 시스템에 오류가 발생할 수 있습니다.
'KRW' 가 아닌 'KRY'라는 통화를 생성하려는 위조지폐범에게 시스템은 사용자 오류를 표시해야 합니다.

```go
func TestCounterfeitCurrency(t *testing.T) {
	defer func() {
		s := recover()
		log.Println(s)
		if s != "Invalid Currency Code Error" {
			t.Error("Test for Detecting Counterfeit Failed")
		}
	}()
	var counterFeit = Construct(5, "KRY")
	log.Println("Failed Currency test is success ", counterFeit)
}
```
소스 코드를 수정하려고 보니 생성자의 통화기호가 일반적인 문자열로 되어 있습니다.
통화기호를 정확히 검증하기 위한 별도의 데이터형이 필요해 보입니다.
`Concurrency` 라는 데이터형을 선언하고 상수 선언도 바꿉니다.
```go
type Currency string

const (
	AFA = Currency("AFA")
	ALL = Currency("ALL")
	....
)
```
생성자에서 인수의 데이터형을 바꿉니다.
데이터형의 안정성 검증을 위해서 타입을 검사하고 다른 데이터형이면 패닉을 발생시킵니다.
```go
// Construct Dollar 생성자
func Construct(amount float64, args ...Currency) Money {
	currency := USD
	if len(args) == 1 {
		arg := args[0]
		s := reflect.TypeOf(arg)
		if s.Name() != "Currency" {
			panic("Invalid Currency Code Error")
		}
		currency = arg
	}
	if len(args) > 1 {
		panic("Too many arguments!")
	}
	return Money{amount, currency}
}
```
많은 컴파일 에러가 발생합니다.
컴파일 에러가 발생한 모든 메소드에서 `string`을 `Currency`로 바꾸어야 합니다.
테스트가 실패합니다.
우리가 예상했던 것과 달리 `Concurrency`는 `string`을 재정의한 것이고 따라서 문자열 값이 담겨지기 때문입니다.

GO에서 타입을 자동변환하는 기능이 있습니다.
그래서 문자열로 받아도 `Currency` 형으로 변환됩니다.
추가적으로 구조체로 선언할 수도 없습니다.

지금 우리에게 꼭 `string` 이 필요할까요? 사실은 아닙니다.
`string`인 이유는 처음 작성한 코드에 'KRW' 와 같은 표현을 가능하게 하려는 의도 때문입니다.
아마도 언제가는 맵에 담아두는 형태로 표현해야 할날이 올 것입니다.

```go
type Currency int

const (
	AFA Currency = iota + 6
	...
	MAX
)
```

테스트 코드를 명확하게 바꿀 수 있습니다.
```go
var counterFeit = Construct(5, MAX)
```

생성자도 이렇게 바꿀 수 있습니다.

```go
func Construct(amount float64, args ...Currency) Money {
	currency := USD
	if len(args) == 1 {
		arg := args[0]
		if arg >= MAX {
			panic("Invalid Currency Code Error")
		}
		currency = arg
	}
	if len(args) > 1 {
		panic("Too many arguments!")
	}
	return Money{amount, currency}
}
```
> **TODO**

> <s>통화기호를 상수로 바꾸기</s>

> <s>유연한 생성자 만들기</s>

## 테스트를 통해 알게 된 것

### iota 의 용법

`const` 선언 영역에서 `KRW = "KRW"` 같이 정의할 때 이 한줄을 `ConstSpec` 이라고 합니다.
 `=` 왼쪽 영역을 `IdentifierList` 라고 합니다.
오른쪽은 `ExpressionList` 라고 하고, 양쪽 모두 리스트이기 때문에 여러 `Identifier`와 `Expression`을 쉼표로 연결할 수 있습니다.

```go
const(
	KRW, USD = "KRW", "USD"
)
```

`iota`는 상수 선언부에서 미리 선언된 식별자를 제공하여 상수를 자동 생성하게 합니다.  
`const`라는 키워드를 소스에서 만날 때마다 0으로 리셋됩니다.
이때 `iota`를 사용하면 각 `ConstSpec` 마다 증가하게 됩니다.
아래는 bitmask 예제로 아주 이해하기 쉽게되어 있습니다.
세번째 `ConstSpec` 다음에 수가 증가하는 것을 확인하시기 바랍니다.

```go
const(
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	bit1, mask1                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3, mask3                           // bit3 == 8, mask3 == 7
)
```
자세한 내용은 [GO 블로그](https://blog.golang.org/constants)를 참조하세요.

상수를 선언할 때 상수의 데이터형을 구조체로 선언할 수 없다고 했는데  `const` 로 생성하려면 구조체가 생성되기 전에 정의되어야 하기 때문입니다.

### 담백한 함수

GO는 이름을 기준으로 중복여부를 판단합니다.
심지어 구조체와 함수가 같은 이름이어도 안됩니다.
즉, 타입과는 무관하게 이름값은 패키지 내에서 중복될 수 없습니다.
함수의 시그니처가 달라도 이름이 같다면 중복으로 판단합니다.

함수에는 또한 인수를 기본값으로 받을 수 있도록 정의된 방식이 없습니다.
작동하는 코드를 위해 위와 같이 정의해서 기본값을 사용할 수는 있지만 명확하게 보이지는 않습니다.
이 방식 말고도 기본값을 사용할 다양한 방법이 있기는 하지만 왜 이런 구차한 방식들을 사용하도록 했을까요?

프로그램의 확장성에서 고정된 인수는 귀찮은 녀석입니다.
다른 언어에서는 하위 호환성을 유지하기 위해서 동일한 함수명을 정의하고 인수를 추가해야 했습니다.
따라서 시그니처가 다르면 함수를 중복으로 인식해서는 안됩니다.
GO는 인터페이스도 하나의 타입으로 지원하기 때문에 모든 형을 인수로 받을 수 있습니다.
따라서 확장성을 위해 중복된 함수를 정의할 필요가 없어진 거죠.
그리고 앞에서 사용한 것처럼 인수가 전혀 없는 것에서부터 많은 수까지 받을 수 있도록 하는 `variadic`을 지원합니다.
C 언어를 제대로 승계한 것이라고 할 수 있습니다.
Variadic Function을 Wikipedia에서는 다음과 같이 정의합니다.

> 인수의 수가 정해지지 않은 함수

### panic에 대하여

앞에서 `Construct` 함수에서 `panic`을 호출합니다.
하지만 실제 프로그램에서 `panic`을 사용하면 안됩니다.
무책임하게 Exception을 사용자에게 전가하기 때문입니다.
소스코드가 안전하게 에러를 전달하도록 해야 합니다.
이 부분은 이글을 읽는 여러분의 몫입니다.
예제 소스를 확인하기 전에 `panic`을 사용하지 않도록 수정해 보시기 바랍니다.
이 부분은 다음 장의 예제 소스에 반영되어 있습니다. 
