# Step 3. 정말 실패하는 자만이 성공한다.

## 동치성(Equality) 문제

Value Object는 생성된 이후에 값이 변경되지 않는다는 특징 이외에 연산을 위해 동치성을 구현해야 합니다.
1 장에서 사용자 스토리로 기술했던 부분 중에서

> 가중평균을 구하기 위해 통화는 더하기 및 곱하기 연산을 지원해야 한다.

이 부분이 관련된 스토리입니다.
스토리를 충족시키기 위해서 목록에 할일을 하나 추가합니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 x 2 = $10</s>

> <s>Dollar를 VO 로 리펙토링하기</s>

> $5=$5 인 equals 구현

우리가 추가한 문장은 멍청한 테스트 같습니다.
이 문장이 그렇게 보이는 이유는 너무도 당연한 결과를 테스트하는 것이라고 생각되기 때문입니다.

### '과학적' 테스트

일상 생활에서 흔히 과학적이라는 말을 사용할 때 어떤 의미로 사용하시나요?
제 경우는 '검증 가능성이 있는' 의 의미로 사용합니다. 물론 국어사전에 나오는 정의와는 다릅니다.
다음사전에 이렇게 정의되어 있네요.

> 과학적科學的 [--쩍]

> ①사물의 현상에 관한 보편적 원리 및 법칙을 알아내고 해명하는 것을 목적으로 하는 학문에 근거한

> ②사물의 현상에 관한 보편적 원리 및 법칙을 알아내고 해명하는 것을 목적으로 하는 학문에 근거한 것

> ③또는 그러한 학문과 관련된

얕은 지식으로 설명하면 검증이란 제시된 이론이나 가설을 어떤 누구라도 동일한 환경에서 동일한 결과가 나오는지 시험하는 것입니다.
가능하다는 말은 검정을 할 수 있어야 제대로 된 이론으로 대접받게 된다는 것이죠.
이렇게 이야기하는 이유는 이번 장에서 엉뚱하게도 실패하는 테스트 케이스를 좀더 고민해야 하기 때문입니다.
검증을 하려면 실패할 수도 있어야 하는데 '원은 둥글다.'와 같은 명제는 실패할 가능성이 없는 명제입니다.
우리가 테스트 하려는 '$5=$5' 도 너무나 당연해서 테스트가 필요없어 보이지만 그렇지 않습니다.
이런 경우에는 실패하는 테스트 코드를 짜야 역설적이게도 소스 코드가 정상임을 보장할 수 있습니다.

### 실패를 위하여

앞 장에서 테스트 순서를 밝혔듯이 이번 장부터는 순서를 별다른 설명없이 진행합니다.
먼저 테스트 코드입니다.

```go
func TestEquals(t *testing.T) {
	var five = Dollar{5}
	isFive, _ := five.equals(Dollar{5})

	if !isFive {
		t.Errorf("Dollar equals method is invalid")
	}
}
```
코드에서 `_` 은 Blank 인식자입니다. 'GO 언어 기초' 편을 참조하세요.
컴파일 에러를 해결하기 위해서 equals를 만듭니다.

```go
func (dollar *Dollar) equals(object Dollar) (bool, error) {
	return true, nil
}
```
일단 컴파일은 문제가 없습니다. 테스트는 우리의 구현에 문제가 있을 것이므로 실패해야 합니다.
테스트를 수행하면 테스트는 문제없이 성공이라고 합니다. 처음에 이를 보면 무슨 일이 일어난 것인지 얼떨떨합니다.

사실 우리는 테스트를 하나 더 작성해야 했는데 이를 하지 않았습니다.
TDD의 저자 켄트 벡은 이를 '삼각측량' 이라는 비유로 설명합니다.
개인 생각으로 이런 비유는 적절하지 않습니다.

```go
func TestEqualsFail(t *testing.T) {
    var six = Dollar{6}
	isSix, err := six.equals("InSane")
	if err == nil {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	if isSix {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
}
```
다시 컴파일러 오류가 발생합니다. 소스 코드를 수정합니다.

```go
func (dollar *Dollar) equals(object interface{}) (bool, error) {
    ...
}
```
메소드의 인수를 `interface{}` 타입으로 바꾸었습니다. 이렇게 하면 메소드가 실제로 모든(?) type을 받을 수 있게 됩니다. 이제 테스트는 확실히 실패해야 합니다.
추가한 테스트 코드는 정말 멍청한 오류를 검증합니다.  

### equals() 구현

이제 테스트를 성공하도록 만들어야 합니다. 우선 멍청한 테스트부터 해결해야 합니다.

```go
var NotCalcualbleError = fmt.Errorf("This value is not calcuable.")
return false, NotCalcualbleError
```
소스코드의 반환값을 바꾸었습니다. 직접 만든 에러를 생성하기 위해서 `fmt`를 임포트 했습니다.
이제 실패하는 테스트는 `TestEquals` 가 됩니다.

`equals()` 메소드가 받는 인수는 여러 타입으로 받을 수 있다고 했습니다.
따라서 각 타입에 대한 동치성 비교를 해야 합니다.
예를 들면 돈이아닌 정수를 인수로 받을 수도 있습니다.
어떤 타입이 올지 모르기 때문에 타입을 검사해야 합니다.
GO에서는 이를 위해서 `Type Assertion`을 제공합니다.
여러분이 익숙한 대부분의 언어에서는 Boxing과 Unboxing을 지원합니다.
GO는 이를 더 안전하게 사용할 수 있도록 만들었습니다.
`Type Assertion` 은 [GO Specification](https://golang.org/ref/spec#Type_assertions) 을 참조하기 바랍니다.

소스 코드를 아래와 같이 변경합니다.
```go
if object == nil {
    return false, nil
} else if v, isInt := object.(int); isInt {
    return dollar.amount == v, nil
} else if v, isDollar := object.(Dollar); isDollar {
    return dollar.amount == v.amount, nil
}
var NotCalcualbleError = fmt.Errorf("This value is not calcuable.")
return false, NotCalcualbleError
```

소스는 두가지 타입에 대해서 동치성을 검사합니다.
정수형과 Dollar 타입일 때만 비교를 해서 동치성 여부를 판단해 줍니다.
그 이외의 경우는 모두 실패 또는 에러를 표시합니다.

모든 테스트가 성공입니다. 이제 할일을 끝났다고 생각하면 될까요? 아직 아닙니다.

## 테스트 얼마나 작성해야 할까?

Atom을 사용하는 경우 소스를 저장하면 붉은 색 블럭으로 표시된 테스트 범위를 보실 수 있습니다.
아마 다른 에디터도 이런 기능을 제공할 것입니다.

![](https://docs.google.com/drawings/d/1RjfaNnxNKbMDdCYwOP-KQ4gfY51N22Eb4-gxB5-PQyo/pub?w=650&h=267)

### 테스트 범위
테스트 범위에 정수형과 `nil` 에 대한 내용이 포함되지 않음을 알려 줍니다.
개발자가 안심하고 저녁에 퇴근하기 위해서 이 부분에 대한 테스트를 추가해야 합니다.

테스트 코드에 다음을 추가합니다.
```go
// Value 이므로 Int 형을 제공하더라도 일치성을 유지해야 한다.
var six = Dollar{6}
isSix, _ := six.equals(6)
if !isSix {
    t.Errorf("Dollar equals method is invalid, int ")
}
isNil, _ := six.equals(nil)
if isNil {
    t.Errorf("Dollar equals method is invalid, nil")
}
```
테스트는 여전히 정상이고 전체 코드가 녹색 바탕으로 바뀝니다.

### 모든 경우의 수를 한번에?
소프트웨어 공학에서 이야기하는 테스트 기법에는 경계값 분석이라는 것도 존재합니다.
사실 경계값을 [-10,-9,0,9,10] 이런식으로 설명하는 책들이 많습니다.
이런식의 테스트 기법이 정말 효율적인지 저는 의심스럽습니다.
이들 기법은 자신들이 효율적인 테스트 기법임을 증명해야 합니다.

테스트 코드를 작성할 때 미리 모든 경계값을 경우의 수에 따라 작성하는 것은 매우 비효율적입니다.
우선 한가지 경우만 테스트를 작성하고 소스 코드를 전개한 후에 나머지 테스트 범위를 채워나가는 것이 좋습니다.

하지만 여러분이 받은 개발 문서에는 모든 경우의 수를 고려해서 테스트를 작성하라고 되어있을 것입니다.
조용히 무시하고 본인의 테스트를 작성해 나가서 단계적으로 요구사항을 충족하시면 됩니다.
여러분 뒤에 서서 모든 경계값 테스트를 하라고 간섭하는 관리자가 있다면 퇴사를 고민해야 할 시점입니다.
경계값 분석같은 기법은 전문 테스터의 영역이지 개발자 테스트에서 고민해야 하는 일이 아닙니다.

관리자는 경계값을 밝히고 필요한 조건의 범위를 요구사항에 기술하는 것만으로 자신의 역할을 충분히 했습니다.
소스상에 정의한 경우의 수를 설계에 반영하는 것은 개발자의 책임입니다.
모든 경우의 수를 고려한 테스트를 설계하는 것은 전문 테스터의 역할입니다.

이제 할일 목록은 다음과 같이 변경할 수 있습니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 x 2 = $10</s>

> <s>Dollar를 VO 로 리펙토링하기</s>

> <s>$5=$5 인 equals 구현</s>

## 테스트를 통해 알게 된 것

### 인터페이스
GO에서 인터페이스는 중요한 위치를 차지합니다.
다른 언어에 익숙한 사람은 GO의 인터페이스를 바로 사용할 수 있을 거라 생각합니다.
하지만 조금 더 들어가면 난해해지기 시작합니다.
GO에서 인터페이스는 메소드의 집합이라는 정의 이외에 타입의 하나라는 점을 명심해야 합니다.
인터페이스를 정의할 때 다음과 같이 하기 때문이죠.

```go
type Animal interface {
}
```
메소드가 아무것도 정의되지 않은 인터페이스를 빈 인터페이스(empty interface)라고 합니다.
우리가 사용한 그 인터페이스죠.  
애석하게도 GO에는 `implements` 라는 키워드가 없습니다.
따라서 모든 타입은 빈 인터페이스에 속하게 됩니다.
이것이 인터페이스가 난해해지는 근본적인 이유입니다.

우리가 작성한 코드를 보면 아래와 같이 정의했습니다.
```go
func (dollar *Dollar) equals(object interface{}) (bool, error) {
	...
}
```
`equals()` 의 인수는 '모든 타입'이 아닙니다.
'interface{}' 타입으로 모든 타입을 GO가 알아서 변경합니다.
이를 메소드에서는 다시 원래의 타입으로 돌려놓아야 사용할 수 있습니다.
한가지 더 알아야 할 것은 인터페이스는 타입과 값이라는 두가지 정보를 가지고 있다는 점입니다.
아래 코드를 `equals()` 에 포함시켜 보세요.

```go
defer func() {
	fmt.Println(object, reflect.TypeOf(object))
}()
```

따라서 포인터를 리시버로 정의한 경우에는 포인터를 제공해야 합니다.
그리고 여기에 더해 받고자 하는 타입이 슬라이스라면 좀더 지저분한 코드를 짜야 합니다.
우리가 작성하게 될 코드에는 인터페이스가 많이 포함될 것이므로 상세한 내용은 그때그때 다시 설명하겠습니다.
지금은 일단 이렇게 정리하고 넘어갑니다.

* 인터페이스는 여러 데이터형의 메소드를 추상화하기 위해 사용합니다. 필드가 아닙니다.
* `interface{}` 는 어떤 타입도 상관없다는 것이 아닙니다.. `interface{}` 타입을 받습니다.
* 인터페이스는 타입과 값 정보 두 가지를 가지고 있습니다.
* `interface{}`	 값을 반환하는 것보다는 `interface{}`를 인수로 받는 것이 효율적입니다.
* 값 형식에서 포인터 형식을 호출할 수는 없지만 포인터 형에서는 값 형식을 호출할 수 있습니다.
* 모든 것은 값으로 전달됩니다. 메소드의 리시버라고 예외는 아닙니다.
* 메소드 내에서 값을 변경하려면 포인터 지시자(`*`)를 사용합니다.

자세한 설명은 [Jordan Orelli의 블로그](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go) 를 참조하세요.

### Type Switch
이제 인터페이스에 타입과 값이라는 정보가 있다는 것을 알았으니 형변환에 대한 좀더 쉬운 사용법을 알아보려고 합니다.
우리는 데이터형을 검사하기 위해서 `Type Assertion`을 사용했습니다.
우리는 코드를 읽기 쉽게 작성해야 합니다.
이 코드는 아래와 같이 바꿀 수 있습니다.
여기서 `빈인터페이스명.(type)` 은 switch 구문에만 사용할 수 있습니다.

```go
switch v := object.(type) {
case nil:
	return false, nil
case int, int32, int64:
	return dollar.amount == v, nil
case Dollar:
	return dollar.amount == v.amount, nil
default:
	var NotCalcualbleError = fmt.Errorf("This value is not calcuable.")
	return false, NotCalcualbleError
}
```

`Switch` 구문에서 특이한 것이 있습니다.
`break` 가 보이지 않습니다. GO에서는 알아서 `break` 처리를 합니다.
중간에 흐름을 끊어야 할 필요가 있을 때만 `break` 를 직접 사용하면 됩니다.
또한 `break` 가 되지 말아야 할 부분에는 `fallthrough` 를 사용합니다.

###  커스텀 에러와 로깅
소스 수준에서 Run Time에 발생시킬 에러를 작성해야 할 때 우리는 `Errorf` 를 사용했습니다.
`fmt` 패키지에서 기본 제공하며 error 를 반환합니다.

만약 에러 타입을 직접 선언하고자 할 경우에는 어떻게 해야 할 까요?
힌트는 `error` 에 있습니다.

```go
type error interface {
    Error() string
}
```

인터페이스 `error` 는 `Error()`  메소드만 구현하면 됩니다. 다음과 같이 만들 수 있습니다.

```go
type CustomError struct {
	position int
}

func (cutomError CustomError) Error() string {
	return fmt.Sprintf("Expected value is 10, Actual value is %d", cutomError.position)
}
```

또한 에러를 기록하려면 `log` 패키지를 사용합니다.

```go
log.Fatal(err)
```

error 와 같이 자주 사용되는 것이 `panic` 과 `recover` 입니다.
