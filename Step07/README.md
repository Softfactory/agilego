# Step 7. 새로운 시작

## 유산 정리

### 데이터 형 통일

우리는 앞장까지 정리된 내용에서 멈추지 않고 앞으로 나아가려 합니다.
아직 해야 할 일이 남았기 때문입니다.

> **TODO**

> <s>$5 + 5000원 = $10 (1:1000 환율일때)</s>

> 데이터 형을 `Float64`로 표준화하기

데이터형을 `Float64` 로 정리해야 겠다고 생각한 것은 이 글의 초안 작업을 하는 중이었습니다.
환율과 통화의 값이 정수형으로 되어 있다는 것이 석연치 않았습니다.
여기서 이런 의심을 더 이상하지 않도록 하고 싶습니다.
어떻게 시작해야 할까요?
우선 테스트 코드에서 정수형을 표시하는 부분을 수정하고 소스를 수정해야 할까요?
아닙니다. 우선 소스에서 정수형을 `float64` 로 바꾸는 것으로 시작합니다.

```go
type Money struct {
	amount   float64
	currency string
}
```
수정하자 마자 우리는 컴파일 에러를 만나게 됩니다.
컴파일러를 따라 나머지  소스를 수정합니다.
생성자 메소드와 `Money`를 리시버로 받는 메소드들을 우선 수정해야 합니다.
이렇게 하고 나면 인터페이스에 정의된 `times()` 도 수정해야 합니다.
수정된 소스 코드 블럭입니다.
```go
type Operator interface {
	exchange(Market, string) Money
	times(float64) Operator
	plus(Operator) Operator
	equals(interface{}) (bool, error)
}
...
func Dollar(amount float64) Money {
	return Money{amount, "USD"}
}
...
func Won(amount float64) Money {
	return Money{amount, "KRW"}
}
...
func Construct(amount float64) Money {
	return Dollar(amount)
}
...
func (money Money) exchange(market Market, currency string) Money {
	var rate = market.getRate(money.currency, currency)
	var amount = money.amount * rate
	return Money{amount, currency}
}
func (money Money) times(multiplier float64) Operator {
	return Money{money.amount * multiplier, money.currency}
}
```
전보다 확실히 소스가 깔끔해졌습니다. `float64`로 형변환하는 코드가 사라졌습니다.
테스트는 `TestEquals`에서 실패합니다. 정수형을 받는 부분이 문제인 것 같습니다.
```go
case int, int32, int64:
    return money.amount == v, nil
```
빨리 작업을 끝내고 싶어서 혐오하던 형변환을 시도합니다.
```go
case int, int32, int64:
    return money.amount == float64(v), nil
```
컴파일러가 우리를 막아섰습니다. Type Assertion이 필요하다는 군요.
```go
case int, int32, int64:
    var value, isFloat = v.(float64)
    if isFloat {
        return money.amount == value, nil
    }
    return false, nil
```
테스트를 다시 실행해도 여전히 실패입니다. 테스트 코드의 출력이 마음에 들지 않습니다.
좀더 상세한 정보를 제공하도록 바꿉니다.
```go
if !isFive {
    t.Error("Dollar equals Money method is invalid", five)
}
....
if !isSix {
    t.Error("Dollar equals Integer method is invalid", six)
}
...
if isNil {
    t.Error("Dollar equals Nil method is invalid", six)
}
```
이제야 컴파일러의 불평이 어떤 내용인지 알겠습니다.
```go
case int, int32, int64:
    var intValue, _ = v.(int)
    return money.amount==float64(intValue), nil
```
그래도 테스트가 실패합니다. 우리의 비교값에 문제가 있나 봅니다.
`floatEquals` 함수를 기억하시나요?
테스트 코드에 작성해 두었던 것으로 소스에는 필요없을 줄 알았던 함수입니다.
이제 우리에게 절실한 코드가 되었습니다. 소스 코드에 복사합니다. `epsilon` 도 잊지 마세요.
`equals`를 다시 수정합니다. float도 처리해야 합니다.
```go
case int, int32, int64:
    var intValue, _ = v.(int)
    return floatEquals(money.amount, float64(intValue)), nil
case float32:
    return floatEquals(money.amount, float64(v)), nil
case float64:
    return floatEquals(money.amount, v), nil
```
정상입니다. Float에 대한 테스트 코드를 조금 더 강화해야 합니다.
```go
isFloat64, _ := six.equals(6.0)
if !isFloat64 {
    t.Error("Dollar equals Float method is invalid", six)
}
isFloat32, _ := six.equals(float32(6.0))
if !isFloat32 {
    t.Error("Dollar equals Float method is invalid", six)
}
```
가장 중요한 구조체인 `Money`의 필드를 `float64` 로 변경했습니다.
걱정했던 것과는 다르게 쉽게 끝났습니다.
생각보다 너무도 간단히 끝나서 어리둥절 하지만 우리는 할일을 끝냈습니다.

> **TODO**

> <s>$5 + 5000원 = $10 (1:1000 환율일때)</s>

> <s>데이터 형을 `Float64`로 표준화하기</s>

### 문자열을 지닌 생성자

소스 코드에서 한가지 아쉬운 점이 있습니다.
`Money{amount, "USD"}` 같이 생성자에 문자열이 들어갑니다.
이 문자열들은 'Externalization' 대상입니다.
하지만 지금 당장 그렇게 할 필요는 보이지 않습니다.
오히려 지금 필요한 것은 이 문자열들이 제각각 생성되지 않도록 통제하는 것입니다.
누군가는 'KRW' 를 사용하고 누군가는 'KW'을 사용한다면 큰 문제가 발생할 수 있습니다.

> **TODO**

> 통화기호를 상수로 바꾸기



## 패키징 전에

### 문서화 작업

이미 언급했듯이 문서화는 대단히 중요한 부분입니다.
자신이 작성한 소스를 설명하고 인정받을 수 있는 것이 문서입니다.
문서화 작업을 작업시간에 반드시 포함해야 합니다.

#### 패키지 개요

패키지 개요는 문서의 최상단에 속합니다.
money 패키지 선언 바로 위에 기술해야 합니다.

#### 구조체 설명

`exported` 구조체는 설명을 가능한 상세히 달것을 권고합니다.
구조체의 역할과 책임에 대해서 분명히 해주어야 합니다.

#### 버그/TODO/Deprecated 기록하기
버그는 `BUG(who)` 이렇게 기록할 수 있습니다.

```go
//  BUG(-): 버그내용을 기술
```
Deprecated는 `Deprecated:` 로 표기합니다.
godoc에서 `Deprecated` 로 표시된 문서를 조회할 수 있습니다.

```go
//  Deprecated:
```
TODO는 `TODO:` 로 표기합니다.
```go
// TODO: 작업에 필요한 사항을 명시합니다.
```

Unexported 메소드나 필드에 달린 주석은 노출되지 않습니다.
```go
// exchange Unexported 메소드에 대한 설명은 웹으로 노출되지 않는다.
```
한가지 주의 할점은 버그/TODO/Deprecated 같은 주석을 기록할 때 항상 앞 단락보다 한칸 더 들여쓰기를 해야 합니다.
자세한 설명은 [공식 GODOC문서](https://godoc.org/golang.org/x/tools/cmd/godoc)를 참조하세요.

#### 패키지 명과 파일명 일치시키기
지금 우리가 작업하는 파일명은 'dollar.go'입니다.
패키지 명은 `money`인데 파일명을 수정하지 않고 그냥 지나쳐 왔습니다.
이제라도 수정하면 됩니다.
'robber_test.go'도 정규화된 이름('example_test.go')으로 바꾸기 바랍니다.

### 메소드의 제공 범위

메소드를 선언하면서 생성자 메소드를 제외하고 모든 메소드를 'Unexported'로 선언했습니다.
작성된 패키지를 어떤 용도로 사용할 것인지에 따라 메소드의 노출 범위를 지정해야 합니다.

### 소스 코드와 테스트 코드의 비율
지금 현재 이 비율은 정확히 1보다는 크고 2보다는 작습니다.
항상 이 비율을 유지하려고 노력해야 합니다.
테스트 코드도 말끔하게 정리할 것을 권장하지만 지금 당장 필요하지 않습니다.


## 1부를 마치며

여기까지가 money 패키지의 내용입니다.
2부에서는 money패키지를 실행 프로그램화 하는 방법을 설명합니다.

### 실행 프로그램

패키지를 실행하려면 main 패키지와 main 함수가 필요합니다.
메인함수는 명령줄 옵션의 사용법과 입출력 기능이 추가되어야 합니다.

### 실시간 환율정보

실시간 환율정보를 얻으려면 웹상에서 환율 정보를 조회해서 업데이트 해야 합니다.
http://www.webservicex.net/CurrencyConvertor.asmx/ConversionRate?FromCurrency=string&ToCurrency=string
