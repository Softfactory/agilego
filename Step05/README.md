# Step 5. 정상을 향하여

## 진짜를 위한 준비   

### 그냥 더하기
이제 우리가 정말 만들고 싶은 기능을 구현할 수 있을 것 같습니다. 더하기 기능 말이죠.
그런데 더하기를 하려면 `$5+$5` 가 먼저 가능해야 합니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <b>$5 + $5 = $10 </b>

테스트 코드로 이를 작성하면 아래처럼 표현되겠죠.

```go
func TestSimplePlus(t *testing.T) {
	var five = Dollar(5)
	var ten = five.plus(Dollar(5))
	var isTen, _ = ten.equals(Dollar(10))
	if !isTen {
		t.Errorf("Expected $10, actual value is %s", ten)
	}
}
```
`plus` 메소드를 작성합니다.
```go
func (money Money) plus(addend Money) Money {
	return Money{money.amount, money.currency}
}
```
물론 테스트는 실패합니다. 이제 테스트가 성공하도록 소스 코드를 작성할 차례입니다.
```go
func (money Money) plus(addend Money) Money {
	var amount = money.amount + addend.amount
	return Money{amount, money.currency}
}
```
예제를 보지 않고도 직접 수정할 수 있을 정도로 쉽게 수정이 가능할 것입니다.

### 두번째 준비
`plus` 메소드를 보면 우리 모두가 아는 문제가 있습니다.
통화가 서로 다른 화폐를 더하면 어떤 단일 화폐로 환전이 되어야 한다는 점입니다.
우리가 가진 객체들 중에는 이런 일을 처리할 적임자가 보이지 않습니다.
켄트 벡은 이 부분에서 'Expression' 과 'Bank' 메타포를 사용했습니다.
연산을 처리하고자 하는 입장에서는 특정 표현식을 받아서 단일 화폐로 바꾸어주는 'Bank' 객체를 생각해서 적용했습니다.
'Bank' 와 'Expression' 은 분명 그에게 저작권이 있습니다.
저는 불필요한 논란을 없애고 더 분명하게 하기 위해 환전 시장(`Market`) 객체를 사용합니다.  

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 + $5 = $10 </s>

> 환전상 객체를 이용해서 환전이 가능하도록 구현하기

따라서 `Market` 구조체를 바로 사용하는 것으로 테스트 코드를 정리할 수 있습니다.
```go
func TestSimplePlus(t *testing.T) {
	var five = Dollar(5)
	var ten = five.plus(Dollar(5))
	var market = new(Market)
	var changed = market.exchange(ten, "USD")
	var isTen, _ = changed.equals(Dollar(10))
	if !isTen {
		t.Errorf("Expected $10, actual value is %s", ten)
	}
}
```

소스코드를 다음과 같이 바꾸면 현재까지 문제들은 어느 정도 해결된 것으로 보이지만 아직 제대로 환전 테스트를 하지는 않았습니다.
```go
// Market 환율을 반영한 환전을 담당한다.
type Market struct {
}

func (market Market) exchange(money Money, currency string) Money {
	return Money{money.amount, currency}
}
```

### 세번째 준비
두가지 일을 준비했는데 준비가 또 남았다구요? 이번 장은 아무래도 순서가 무언가 꼬여 있는 느낌이 강합니다. 하지만 이것도 여러분에게 도움이 될 것 같습니다. 왜냐하면 개발자마다 구현하는 순서가 서로 달라도 같은 결론에 다다를 것이라는 희망을 가질 수 있기 때문입니다.

할일에서 `$5+5000KRW` 를 주목하면 한가지 특징을 발견할 수 있습니다. `+` 기호를 중심으로 동등한 객체로 대우를 받아야 한다는 것이죠. `5$` 와  `5KRW`는 동등한 객체입니다. 우리는 앞장에서 `Dollar`와 `Won` 객체를 지워버렸고 `Money`로 통합했기 때문에 동등한 객체라는 것을 의심하지 않았습니다. 하지만 향후에 이 프로그램에 엄청난 양의 소스를 갖게되고 생성자 부분에 버그가 생긴다면 예전에 의심하지 않았던 자신을 질책할 지도 모릅니다. 아무튼 발 뻗고 자는 개발자가 되기 위해서 아래 테스트 코드를 추가합니다.

```go
func TestMoneyEquals(t *testing.T) {
	var fiveWon = Won(5)
	var fiveDollars = Dollar(5)
	var market = new(Market)
	if !reflect.DeepEqual(market.exchange(fiveWon, "USD"), fiveDollars) {
		t.Errorf("Deep Equality Test Failed %s %s", fiveWon, fiveDollars)
	}
}
```
여기에는 함축적인 의미가 있는데 우선 환율을 모르면 1:1 교환이 이루어 진다고 가정한 것입니다.
그리고 객체를 비교하기 위해서 `DeepEual` 이라는 리플렉션 메소드를 사용했습니다.  

이제 진짜 할일을 해결할 준비가 끝났습니다.

## 환전 시장을 완성하기

### 어느 것을 먼저 할 것인가?
환전 시장에서 환전을 할 수 있으려면 환율 정보를 입력받아서 특정 환율로 변환하는 역할을 수행할 수 있어야 합니다.
때론 테스트 주도 개발에서 좀더 먼 길로 들어설 수 있는데 지금 보시는 이 부분에서 저는 몇번 길을 잃었습니다.
테스트 케이스로 환율이 바뀌는 지를 먼저 테스트할 것인지 아니면 환율정보처리에 집중할 것인지 결정해야 합니다.
어느 것을 먼저 시작하나 결론은 같은 길로 나아가지만 빠르게 결정할 수 있는 실천적인 지침이 필요합니다.

> * 테스트 실패 보다 컴파일 에러가 나는 테스트 케이스를 먼저 작성하라.

두가지 테스트 케이스가 있습니다. 먼저 `exchange` 메소드의 기능을 테스트하는 코드입니다.
```go
func TestExchange(t *testing.T) {
	var five = Dollar(5)
	var market = new(Market)
	var changed = market.exchange(five, "KRW")
	var is5ThousandWon, _ = changed.equals(Won(5000))
	if !is5ThousandWon {
		t.Errorf("Expected 5000KRW, actual value is %s", changed)
	}
}
```
그리고 다음 테스트는 환전 시장에서 환율정보를 제대로 처리할 수 있음을 보장하는 테스트입니다.
```go
func TestExchangeRate(t *testing.T) {
	var market = new(Market)
	var rate float64 = 1000

	market.setRate("USD", "KRW", rate)
	if !floatEquals(rate, market.getRate("USD", "KRW")) {
		t.Errorf("Expected 1000, actual value is %f", rate)
	}
	var defaultRate = market.getRate("KRW", "CHF")

	if !floatEquals(1.0, defaultRate) {
		t.Errorf("Expected 1.0, actual value is %f", defaultRate)
	}
}
```
어떤 테스트를 먼저 만들어 넣을 것인지는 여러분이 결정할 몫입니다.
저는 `TestExchangeRate` 부터 추가했습니다. 바로 컴파일 에러를 만날 수 있기 때문이죠.

소스 코드를 수정합니다.

```go
func (market *Market) setRate(to string, from string, rate float64) {

}

func (market Market) getRate(to string, from string) float64 {
	return 1000
}
```
`getRate` 를 단순히 1000을 반환하는 것으로 작성했기 때문에 테스트는 실패합니다.

### 환율 정보 담기
이제 소스 코드를 제대로 수정해서 테스트 실패를 없앨 생각입니다.
우리가 구현하는 환전상에는 환율정보를 담을 필드가 필요합니다.
GO에서는 `Map`을 사용해서 이런 정보를 담는 것이 편리합니다.
Java에서 흔히 사용되는 `HashTable` 과 유사하다고 보시면 됩니다.
`Map`은 `필드명 map[키_데이터형]값_데이터형` 으로 선언합니다.
`Map` 에 접근할 때는 `필드명[키]=값`  과 같이 익숙한 형태로 사용할 수 있습니다.
슬라이스는 GO에서 Bag입니다. 우리에게는 Set 이 필요합니다.
```go
// Market 환율을 반영한 환전을 담당한다.
type Market struct {
	rates map[pair]float64
}

type pair struct {
	from string
	to   string
}

func (market *Market) setRate(from string, to string, rate float64) {
	if market.rates == nil {
		market.rates = make(map[pair]float64)
	}
	market.rates[pair{from, to}] = rate
}

func (market Market) getRate(from string, to string) float64 {
	var rate = 1.0
	if market.rates != nil {
		var saved, isRated = market.rates[pair{from, to}]
		if isRated {
			rate = saved
		}
	}
	return rate
}
```
GO에서 `nil`은 Null 값을 의미합니다. 포인터형 변수를 선언하면 초기값은 언제나 `nil` 입니다.
또한 동작중에 `nil`이 될 수도 있습니다. 자세한 설명은 뒤에서 합니다.
소스 코드를 한줄 한줄 작성해 나갈때 마다 테스트를 실행해서 작성중인 코드가 이상이 없는지 확인해야 합니다.
`pair` 는 환전될 통화들의 기호를 의미하는 것이어서 구조체를 작성하는 과정이 바로 떠오르지 않을 수 있습니다.
통화 시장에서 'pair'라는 용어를 사용한다는 점을 알아야 하기 때문입니다.
아무튼 `pair`를 `unexported`로 선언했고 `Money` 패키지만이 `pair`를 알 수 있습니다.

### 모든 것의 정상에서
작성한 테스트는 모두 성공합니다. 이제 `exchange`를 테스트할 차례입니다.
위에 `TestExchnage`를 추가합니다. 테스트는 '당연히' 실패할 것입니다.
`exchange` 메소드를 우리가 수정하지 않았기 때문입니다.

```go
func (market Market) exchange(money Money, currency string) Money {
	var rate = market.getRate(money.currency, currency)
	var amount = float64(money.amount) * rate
	return Money{int(amount), currency}
}
```
여기서 테스트를 성공시키기 위해서 몇가지 범죄를 저질렀습니다.

* Money와 Market 가 제공하는 필드의 데이터형이 다름에도 곱셈연산을 수행했습니다.
* 소수점 이하를 버림으로써 고객의 돈을 갈취했습니다.

비록 모든 것이 정상인 지점에 있지만 다시 올라가야 할 산을 만났습니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 + $5 = $10 </s>

> <s>환전상 객체를 이용해서 환전이 가능하도록 구현하기</s>

> 데이터 형을 `Float64`로 표준화하기

## 테스트를 통해 알게 된 것

### Floating 연산
Float 데이터형을 연산하면서 equals를 특별한 함수를 작성해서 사용했습니다.
테스트에 대한 설명이 초점을 잃고 흐를까봐서 그냥 지나쳤습니다.
```go
var epsilon = math.Nextafter(1, 2) - 1.0
func floatEquals(a, b float64) bool {
	if math.Abs(a-b) <= epsilon {
		log.Printf("Epsilon is %f ", epsilon)
		return true
	}
	log.Printf("Input Value is %f , %f ", a, b)
	log.Printf("(estimated value) %f > %f (epsilon on this machine) ", math.Abs(a-b), epsilon)
	return false
}
```
첫줄은 계산기 엡실론(ε)을 구합니다.
계산기 엡실론은 부동소수점 연산에서 반올림으로 인해 발생하는 오차의 상한입니다.
`float64`를 사용하기 때문에 `Nextafter`를 사용했고, `float32`를 사용한다면 `Nextafter32`를 적용해야 합니다.
그리고 3장에서 언급했던 `log` 패키지를 이번에는 제대로 사용했습니다.
`log`는 엔터프라이즈 시스템에서는 필수로 들어가는 부분입니다.

### new와 make
지금까지 언급하지 않았던 `new()` 함수를 사용했습니다.
`new()`는 지정된 데이터형의 zero값을 메모리에 할당해서 변수를 사용할 수 있도록 해줍니다.
그러나 할당된 메모리를 해제하는 것은 GO의 가비지 컬랙션의 역할이기에 해제 코드를 작성할 필요가 없습니다. 이때 할당된 것은 값은 zero이고 메모리 주소만 할당됩니다. 즉 메모리를 초기화 하지 않습니다.
```go
func TestMemory(t *testing.T) {
	var car = new(Car)
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car, *&car, unsafe.Sizeof(car))
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car, *&car, reflect.TypeOf(car).Size())
}
```
출력되는 내용은 `Address &i=0xc82002e038 Address Value=&{} Size=8` 으로 동일합니다.
`unsafe.Sizeof()` 또는 `reflect.TypeOf().Size()` 모두 동일한 결과를 얻습니다.
애석하게도 여기서 보이는 크기는 메모리에 실제 할당된 크기가 아닙니다.
순전히 이론적인 바이트 크기일 뿐입니다.

`make()` 는 슬라이스, 맵, 채널을 생성할 때 사용합니다. `new()`와 다르게 초기화한 값을 반환합니다.
그리고 포인터를 반환하지 않습니다. 필요시 명확하게 포인터를 사용하겠다고 해야 합니다.
슬라이스,맵,채널은 내부 자료에 대한 참조를 가지고 있어야 사용할 수 있기 때문입니다.  
특히 `nil` 인 채널에 데이터를 전송하거나 받으려고 하면 영원히 블러킹되는 문제가 발생합니다.

```go
var p = new([]int)       // 슬라이스를 할당하지만 *p == nil 이다.
var v = make([]int, 100) // 100의 int에 대한 참조를 갖는다.

fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &p, *&p, reflect.TypeOf(p).Size())
fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &v, *&v, reflect.TypeOf(v).Size())

*p = make([]int, 100, 100)
fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &p, *&p, reflect.TypeOf(p).Size())

v1 := make([]int, 100, 100)
fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &v1, *&v1, reflect.TypeOf(v1).Size())
```
자세한 설명은 [Effective GO](https://golang.org/doc/effective_go.html#allocation_new)를 참조하세요.

### reflect.DeepEqual
객체의 일치여부를 확인하는 간단한 방법으로 GO는 `reflect.DeepEqual()`  을 제공합니다.
이 메소드를 사용하면 모든 필드를 조사해서 동치성 여부를 반환합니다.
