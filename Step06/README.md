# Step 6. 참회록

## 인터페이스에 대한 고백

### 원죄 더하기

지금 할일 목록에는 적어도 두가지가 남아 있습니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> 데이터 형을 `Float64`로 표준화하기

첫번째 문장은 우리가 해온 절차에 따라서 테스트 코드를 작성할 수 있습니다.
테스트 코드는 다음과 같습니다.

```go
//$5 + 5000원 = $10 (1:1000 환율일때 )
func TestSummation(t *testing.T) {
	var fiveDollars = Dollar(5)
	var fiveThousandWon = Won(5000)
	var market = new(Market)
	market.setRate("KRW", "USD", 0.001)
	var sum = fiveDollars.plus(fiveThousandWon)

	if !reflect.DeepEqual(market.exchange(sum,"USD"), Dollar(10)) {
		t.Error("Expected $10", market.exchange(sum,"USD"))
	}
}
```
테스트는 보기 좋게 실패합니다. 테스트 결과를 확인하니 값이 $5005 입니다.
어떻게 된 것인지 우리 모두 알고 있습니다.
`plus` 메소드에서 환율을 적용하지 않았습니다. `Market`의 `exchange` 메소드를 호출하지 않았습니다.
현재까지 `Money`의 `plus`는 `Market`의 존재를 모릅니다.
지금 당장 `Market`을 인수로 받아서 테스트를 성공시킬 수 있지만 다른 복잡한 문제를 더 쌓게 됩니다.
테스트의 많은 부분은 새로 작성해야 할 것입니다.
주어진 문제인 `$5+5000원`에서 덧셈연산을 단순히 `plus`라는 메소드로 처리하려고 했습니다.
이것이 씻기 힘든 원죄가 되어 우리를 괴롭히고 있는 것입니다.

### 인터페이스 사고
벡이 자신의 책에서 이 문제의 해결책을 제시할 때 저는 다소 혼란스럽고, 대충 이야기하고 지나가는 인상을 받았습니다.
갑자기 `Expression` 이 튀어나왔습니다.
하지만 우리가 작성하는 개발 코드들 대부분이 불현듯 이렇게 도출되는 경우가 많습니다.
정확히 왜 있어야 하는지 알 수 없지만 '있어야 할 것 같아서' 구현했더니 정말로 필요한 코드가 되는 경우가 있습니다.
직감에 대한 문제는 벡의 책을 인용해 맡기고 쉽게 이글을 마무리하고 싶었습니다.
하지만 그렇게 되면 이글을 읽으시는 분들에게 GO를 더 재미있게 알려 드리지 못할 것 같습니다.

사실 앞에서 단초를 제공하기는 했습니다.
테스트를 전부 수정해야 할 일이 발생할 때 우리는 새로운 객체를 생각해내서 객체가 우리를 보호할 것이라고 생각했습니다.

이번에도 비슷한 흐름으로 전개할 수 있을 것 같습니다.
이번에 우리를 지켜줄 것은 새로운 구조체 뿐만이 아니라 메소드의 집합인 인터페이스입니다.
무언가 새로운 행위자가 필요하다는 것을 알 수 있습니다.
왜냐구요? 지금까지 작성한 테스트는 모두 정상입니다.
코드를 잘못 작성해온 것이 아닙니다.
그럼에도 더하기 연산 기능(행위)을 추가해야 하지만 더이상 진행을 할 수 없습니다.
그렇다면 행위를 할 수 있는 행위자를 추가해야 하고 지금 우리에게는 새로운 인터페이스가 필요합니다.

### 통화 수식을 위한 인터페이스
'$5+5000원' 은 그 자체가 수식을 표현하고 있습니다.
연산자(Operator)는 '+'이고 피연산자(Operand)로 피가산수($5)와 가산수(5000원)가 보입니다.
'Sum'을 구조체로 만든다면 구조체의 멤버인 피가산수(augend)와 가산수(addend)는 `Money` 가 되어야 합니다.
이와 유사하게 피감수(minuend)와 감수(subtrahend)를 멤버로 갖는 'Subtraction'이 추가될 수 있습니다.

그리고 'Sum' 같은 연산자는 정말로 우리가 구현하고자 하는 것들 중 핵심적인 부분입니다.
'Sum'과 같은 연산자는 `Market`이 계산할 수 있도록 간략화(reduce)되어야 합니다.
향후 확장성을 고려해서 모든 연산자를 간략화할 행위자(인터페이스)를 추가합니다.

연산자에 피연산자가 포함되어 있는 구조체를 떠올리면 연산자는 모두 `Operator`라는 인터페이스를 구현해야 할 것이 분명해 보입니다.
그러고 멤버인 `Money`는 `Operator` 인터페이스를 구현해야 할까요? 아니면 `Operand`라는 인터페이스를 추가해야 할까요?
지금 이 문제에 대한 답은 테스트전에는 확실하지 않아 보입니다.

이제 인터페이스 `Operator` 를 작성하는데 별다른 주저함이 없습니다.
우리가 갖는 연산자가 할 행위는 `Market`의 `exchange`와 유사하게 정의됩니다.
지금 우리에게 필요한 연산자는 'Sum'입니다. 적어도 지금까지는 그렇습니다.
`Sum` 구조체도 정의합니다. `Money`의 `plus` 도 `Operator`를 반환하는 것으로 수정합니다.

```go
type Sum struct {
	augend Money
	addend Money
}

type Operator interface {
	exchange(Market, string) Money
}

func (money Money) plus(addend Money) Operator {
	return Sum{money, addend}
}
```

### 컴파일러가 응답하다.

테스트 코드를 실행하면 어떤 결과를 보게 될까요?
컴파일 에러를 만나게 됩니다.

```
./dollar_test.go:88: cannot use ten (type Operator) as type Money in argument to market.exchange: need type assertion
FAIL	AgileGO/money [build failed]
```

아주 쉽게 작성했던 `TestSimplePlus`에서 발생한 에러입니다.

```
var ten = five.plus(Dollar(5))
var market = new(Market)
var changed = market.exchange(ten, "USD")
```

이 테스트 코드를 수정하지 않고 컴파일 오류를 해결할 방법이 있을까요?
앞에서 `Money`가 `Operator` 인터페이스를 구현해야 하는가에 대해 명확한 답을 제공하지 못했습니다.
지금은 테스트를 통해서 명확한 답변을 할 수 있습니다.
`Money` 도 `Operator` 인터페이스를 구현해야 합니다.

`exchange`를 다음과 같이 변경합니다.

```go
func (market Market) exchange(operator Operator, currency string) Money {
	return operator.exchange(market, currency)
}
```

## 리팩토링

### 다음 단계
우리는 순식간에 가장 어려운 일을 처리했습니다.
그것도 아주 우아한 결과물을 얻을 수 있었습니다.
그러나 원저에 있는 메타포를 희생하면서 얻은 결과여서 흐름은 뒤섞여 있고 아직 할일이 남아 있습니다.

> **TODO**

> <s>$5 + 5000원 = $10 (1:1000 환율일때 )</s>

> 데이터 형을 `Float64`로 표준화하기

지금의 테스트로도 충분하지만 아직 완벽하지는 않습니다.
`Sum`은 `Money` 가 가진 동치성 확인이나 곱셈연산, 더하기 연산을 제공하지 않습니다.
`Sum`도 이러한 메소드를 제공해야 다음과 같은 사용자 스토리가 가능합니다.

> **TODO**

> ($5+5000원)x2=$20

잊지 말아야 할 것은 처음부터 제대로 된 구현을 생각하지 않고 작성하기 입니다.
불안해 하지 말고 일단 테스트를 성공하도록 한 후에 리펙토링 작업을 하면서 개선해 나아가야 합니다.

```go
func TestSumEquals(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = Sum{fiveDollars, fiveDollars}
	var sum1 = fiveDollars.plus(fiveDollars)
	var isEquals, _ = sum.equals(sum1)
	if !isEquals {
		t.Error("Sum Equals failed ", sum)
	}
	var _, err = sum.equals(fiveDollars)
	if err == nil {
		t.Error("Sum have not to be equals to Money")
	}
	var isNullEqual, _ = sum.equals(nil)
	if isNullEqual {
		t.Error("Sum equals must failed with nil")
	}
}
```

`Sum`구조체를 위한 `equals` 는 구현하기 쉽습니다. 이미 우리는 `Money`에서 유사한 것을 작성했습니다.

```go
func (sum Sum) equals(object interface{}) (bool, error) {
	switch v := object.(type) {
	case nil:
		return false, nil
	case Sum:
		var isAugend, _ = sum.augend.equals(v.augend)
		var isAddEnd, _ = sum.addend.equals(v.addend)
		return isAugend && isAddEnd, nil
	default:
		var NotCalcualbleError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcualbleError
	}
}
```

이제 `times`를 테스트할 코드를 추가합니다.

```go
func TestSumTimes(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = Sum{fiveDollars, fiveDollars}.times(2)
	var market = new(Market)
	if !reflect.DeepEqual(sum.exchange(*market, "USD"), Dollar(20)) {
		t.Error("Sum Times is failed ", sum)
	}

	var sum1 = fiveDollars.plus(fiveDollars)
	var sum2 = sum1.times(2)
	if !reflect.DeepEqual(sum2, Sum{Dollar(10), Dollar(10)}) {
		t.Error("Sum is failed ", sum2)
	}
}
```
소스 코드에 `Sum`의 `times`를 추가합니다.
```go
func (sum Sum) times(multiplier int) Sum {
	var augend = Money{sum.augend.amount * multiplier, sum.augend.currency}
	var addend = Money{sum.addend.amount * multiplier, sum.addend.currency}
	return Sum{augend, addend}
}
```
이제 많은 부분이 진전되었습니다. 마지막 `plus`에서 우리는 곤란을 겪을 것입니다.

### 진정한 리팩토링
```go
func TestSumPlus(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = fiveDollars.plus(fiveDollars)
	var sum1 = sum.plus(sum)
	var market = new(Market)
	if !reflect.DeepEqual(sum1.exchange(*market, "USD"), Dollar(20)) {
		t.Error("Sum Plus is failed ", sum1)
	}
}
```
이에 대한 소스코드는 위해 `plus` 쉽게 작성할 수 없습니다.
```go
func (sum Sum) plus(addend Operator) Sum {
	return Sum{sum, addend}
}
```
컴파일 에러가 발생합니다. `augend` 필드에 `Sum` 을 넣을 수가 없습니다.
이번 작업은 제대로된 리팩토링을 해볼 수 있는 기회입니다.
지금까지의 테스트는 모두 성공을 했습니다.
이제 `Sum`을 재정의해서 우리가 원하는 기능을 추가하려고 합니다.

```go
type Sum struct {
	augend Operator
	addend Operator
}
```
그리고 인터페이스에 메소드를 추가합니다. 나머지는 컴파일러를 따라가면 됩니다.
```go
type Operator interface {
	exchange(Market, string) Money
	times(int) Operator
	plus(Operator) Operator
}
```

완성된 소스코드는 예제에 충분하게 있습니다.
하지만 중간 과정은 최대한 간략하게 언급했습니다.
이 부분은 벡의 책에 자세히 나와 있지만 저는 이글을 읽는 분들이 직접 리팩토링을 수행하기 원합니다.

> **TODO**

> <s>($5+5000원)x2=$20</s>

여기까지는 벡의 책을 참조로 작성되었습니다. 다음 장부터는 추가적인 작업들을 진행합니다.

## 테스트를 통해 알게된 것

### 인터페이스 이름짓기
GO에서 인터페이스 이름은 `Writer`와 같이 단일 메소드명에 '-er' 을 붙이는 것이 관례입니다.
인터페이스 이름은 굉장히 혼란스러울 수 있는데 어떤 언어에서는 접두어로 대문자 'I' 를 붙이는 경우도 있습니다.
저는 GO의 코딩 규칙을 준수하면서도 가능한 인터페이스는 행위자이기를 원합니다.
벡이 사용한 `Expression`을 사용하지 않은 이유가 이 때문이기도 합니다.
'수식'은 아무 행위도 하지 않습니다.
하지만 `Operator`를 사용하면서 `Money` 도 `Operator`가 되는 이상한 형상을 만들어 냈습니다.
벡은 `Operator`와 `Operand`를 아우르는 `Expression`이 필요했을 지도 모릅니다.
그럼에도 `Operator`로서 `Money`의 자격은 충분합니다.
`Money`가 '스스로'  곱하기 연산과 더하기 연산을 하기 때문입니다.

### Type Assertion 다시보기
'진정한 리팩토링' 을 수행할 때 만나게 될 컴파일 에러 중 중요한 메시지가 Type Assertion 실패입니다.
눈치가 빠른 분들은 Type Assertion을 명확히 해 주어서 컴파일을 성공시킬 것입니다.
```go
var augend = Money{sum.augend.(Money).amount * multiplier, sum.augend.(Money).currency}
```
중요한 것은 인터페이스를 반환하면 컴파일러는 명확하게 제시된 타입의 실패를 가정한다는 것입니다.
이 부분은 버그 가능성을 내포하고 있습니다.
우리는 `TestSumTimes` 에 아래 코드를 추가해야 안심할 수 있습니다.

```go
var newSum = Sum{fiveDollars, fiveDollars}
var sum1 = newSum.plus(newSum)
var sum2 = sum1.times(2)
var sum3 = Sum{Dollar(10), Dollar(10)}
if !reflect.DeepEqual(sum2, Sum{sum3, sum3}) {
	t.Error("Sum is failed ", sum2)
}
```
이제 예상치 못했던 버그를 알게 되었고 수정할 수 있게 되었습니다.
여러분의 소스코드에 `.(Money).` 같은 표현이 보인다면 예제를 보기전에 어떻게 수정할 것이지 고민해 보시기 바랍니다. 
