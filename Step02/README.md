# Step 2. 잃어버린 객체를 찾아서

## 객체를 보호하라.

### 가치가 있는 객체

앞에서 작성한 `times()` 는 하나의 문제를 가지고 있습니다. 이 함수를 호출하고 나면
객체의 값이 바뀐다는 점입니다. 내 지갑에 있는 돈의 명목 가치가 어떤 사건에 의해서
바뀐다는 것은 상상할 수 없는 이야기 입니다. 한번 생성된 객체의 값은 변하지 않기를 바랍니다.
사실 디자인 패턴 세상에서 이런 걸 방지하기 위한 묘안을 제시하고 있습니다.
Value Object 패턴입니다. VO와 DTO를 혼돈하는 개발자들이 있습니다.
자세한 설명은  마틴 파울러의 책 "엔터프라즈 소프트웨어 아키텍처" 를 참고하시기 바랍니다.

이런 문제점을 고치기 위해서 TODO  리스트에 할일을 추가 합니다.

> **TODO**

> Dollar를 VO 로 구현하기

그런데 이런 표현은 명확한 테스트 지침을 제공하지 않습니다.
VO라는 표현이 사람들을 헷갈리게 할 수도 있습니다.
사실 이번에 할일은 단순히 테스트가 아니라 리펙토링입니다.  

> **TODO**

> Dollar를 VO 로 리펙토링하기

### 뜯어고치기

리펙토링을 하려면 제대로 된 테스트 코드가 필요합니다.
다행이도 우리는 훌륭한 테스트 코드를 이미 작성했습니다.
테스트 코드를 변경해서 우리가 하고자 하는 일을 명확히 표현합니다.

```go
var ten = five.times(2)
if ten.amount != 10 {
    t.Errorf("Amount is expected 10, but actually %d ", ten.amount)
}
```

테스트 코드를 위와 같이 변경하면 바로 에러가 발생합니다. 컴파일러 에러입니다.
우리는 다시 한번 TDD 의 작업 단계를 확인할 필요가 있습니다.

> 1단계 : 컴파일러 에러나 테스트 실패 등을 무시하고 테스트를 작성한다.

> 2단계 : 테스트 실패는 무시하고 컴파일러 에러가 없도록 만든다.

> 3단계 : 테스트가 성공하도록 소스를 수정한다.

1단계는 완료했으니 2단계를 위해 소스를 아래처럼 수정합니다.

```go
func (dollar *Dollar) times(multiplier int) Dollar {
	dollar.amount = dollar.amount * multiplier

	return Dollar{0}
}
```

컴파일러 에러는 사라지지만 테스트는 실패합니다.

`Dollar{0}` 을 반환하는 코드를 보고 일부 흥분하는 개발자가 있을 수도 있습니다.
우리가 여기서 작업을 끝낸다면 제가 비난받아 마땅하지만 아직은 아닙니다.

3 단계로 이제 테스트가 성공하도록 만들 차례입니다.

```go
func (dollar *Dollar) times(multiplier int) Dollar {
	// dollar.amount = dollar.amount * multiplier
	return Dollar{dollar.amount * multiplier}
}
```
Atom 에디터 사용자라면 아래 그림과 같이 커버리지가 정상적으로 확인되는 표시를
볼 수 있습니다. 녹색으로 반전된 부분은 테스트에서 정상 동작할 것임을 알려줍니다.

![](https://docs.google.com/drawings/d/1G_nrGzmRV_1LO16as7FEgAUX5YfRym337hmr94tLmbA/pub?w=672&amp;h=217)

모든 것이 이제 정상입니다. TDD의 작업단계를 준수하는 일은 사실 천재들에게는 필요가 없습니다.
TDD는 정말 '천재같지 않은' 개발자를 위해 필요한 지침입니다.

## 그래도 불안하다.

그런데 아직 미심쩍은 부분이 남아 있습니다. amount 필드의 존재가 마음에 걸립니다.
앞에서 exported 와 unexported의 구분은 이름의 맨 앞글자가 대소문자인지로 구분된다고 했습니다.
이는 매우 단순한 이름 규칙이기 때문에 강력하기도 하지만 세심한 주의를 필요로 합니다.

### 동일 패키지 내에서

저는 좀더 과장을 해서 강도를 만났다고 가정했습니다. `robber.go`, `robber_test.go` 파일을 만들었습니다.

강도는 우리의 구현 대상에 포함되어 있지 않기 때문에 Robber 구조체를 구현할 필요는 없습니다.
언제가 미래에 이를 구현할 필요가 있을지도 모르지만 지금은 아닙니다.

아무튼 `robber_test.go` 에 아래와 같은 테스트를 작성했습니다.

```go
package money

import (
	"testing"
)

func TestUnexportedAmount(t *testing.T) {
	var five = Dollar{5}

	five.amount = 4

	if five.amount != 5 {
		t.Error("You're being robbed. %d", five.amount)
	}
	// getAmount()
}
```

우리는 테스트를 제대로 작성했습니다. 그래서 테스트가 실패합니다.

우선 unexported 의 개념부터 확인해야 겠습니다.

Java와 같은 private 선언이 없는 GO에서는 사실 패키지안에서 구조체의 접근을 막을 방법이 없습니다.
unexported는 Java의 protected 선언처럼 동작합니다.
GO는 OOP가 아니기 때문에 데이터의 보호에 그렇게 철저하지 않습니다.
같은 패키지 내에 선언된 구조체의 필드를 보호하려면 그냥 접근하지 않으면 됩니다.
Java 에 익숙한 제가 이런 말을 하고 있다니 놀랍도록 무책임해 보일 수 있습니다.
하지만 이런 기준으로 작성된 코드는 매울 깔끔합니다.

이제 테스트를 바꿔야 겠습니다. 강도는 통화 패키지 안에 들어갈 이유가 없습니다.

### 테스트 패키지

robber는 분명 테스트를 위한 것입니다. 패키지 명을 `money_test` 로 변경하고 테스트를 실행합니다.
패키지 명에도 `_test` 규칙이 적용됩니다.

만약 아무 상관없는 패키지명(예를들면 street)을 사용하면 아래와 같은 메시지를 만나게 됩니다.

```
can't load package: package AgileGO/money: found packages money (dollar.go) and street (robber.go) in /home/jacob/Projects/CodeGo/src/AgileGO/money
```

패키지 적재가 불가능한 이유는 2개의 패키지가 하나의 URL에 있을 수 없기 때문입니다.

GO 도움말을 읽어보니 다음과 같은 문구가 보입니다.

> 독립된 URL 당 단 하나의 Package 만 가능.

이제 쓸모없어 보였던 robber.go 파일을 지울 수 있습니다. robber.go 파일을 지웁니다.

이제 컴파일러 에러를 해결해야 합니다.
지금까지 만든 robber 테스트에는 Dollar를 찾을 수 없다는 메시지가 표시됩니다.

```
./robber_test.go:8: undefined: Dollar
FAIL	AgileGO/money [build failed]
```

테스트 패키지에서도 테스트 대상 패키지를 Import 해야 합니다.
```
import (
	"AgileGO/money"
	"testing"
)
```

그리고 패키지 외부에서 테스트하고 있기 때문에 money 패키지로부터 시작을 해야 된다는 것을 컴파일러 에러를 통해서 알수 있습니다.

```
imported and not used: "AgileGO/money" as money
undefined: Dollar
```
다시 테스트 코드를 수정합니다. 이제 해결해야 할 컴파일러 에러가 변경됩니다.

```go
var five = money.Dollar{5}
```

```
./robber_test.go:9: implicit assignment of unexported field 'amount' in money.Dollar literal
FAIL	AgileGO/money [build failed]
```

이 컴파일러 에러의 메시지는 amount 필드를 직접 접근하지 말라고 합니다.
우리는 amount 를 철저하게 보호하려는 의도로 지금 테스트 코드를 수정하고  있습니다.
그런데 객체를 신규로 생성할 때 에러가 발생합니다.
amount를 모르는 상태에서 어떻게 객체를 신규로 생성할 수 있을까요?

이 문제를 디자인 패턴 세상에서 이미 생성자 패턴으로 정리해 놓았습니다.

### 생성자를 만나다.

생성자 패턴은 객체의 생성을 보장합니다. 생성자 패턴에는 사실 여러 패턴의 집합입니다.
Singleton, Monostate, Factory 등 한번쯤 들어보았음직한 패턴이 여기에 속합니다.

다행히도 우리는 이 모든 패턴을 알 필요가 없고 지극히 단순한 생성자만 필요합니다.
앞에서 수정한 코드에서 객체 생성 메소드를 호출하는 것으로 바꿉니다.

```go
var five = money.Creator(5)
```

다시 한번 컴파일러 에러가 변경됩니다.

소스 코드에 다음 메소드를 추가합니다.

```go
// Construct Dollar 생성자
func Construct(amount int) Dollar {
	return Dollar{amount}
}
```

사실 이 메소드를 추가하는 과정도 컴파일러의 도움을 받았습니다.  
이제야 우리가 보고 싶은 컴파일 에러를 만났습니다.
```
five.amount undefined (cannot refer to unexported field or method amount)
```
Dollar 객체를 생성할 수는 있지만 Dollar 객체의 필드에 직접 접근할 수 없습니다.
결국 우리는 한단계 더 앞으로 나아갈 수 있게 되었습니다.

## 객체를 검증하는 테스트

원하는 것을 얻기는 했지만 컴파일러 에러로 테스트에 문제가 발생합니다.

`try ~ catch ` 구문같은 것으로 컴파일 에러를 처리할 수는 없습니다.
객체를 직접 검증할 방법을 찾아야 합니다.

최신 언어에는 `Reflection` 기능이 있습니다. 우리도 이 기능을 사용해야 합니다.
리플렉션을 사용하면  Run Time 에서 객체의 정보를 얻어낼 수 있습니다.   
우선 `reflect` 패키지를 임포트 합니다. 테스트 코드를 아래와 같이 변경합니다.

```go
func TestUnexportedAmount(t *testing.T) {
	var five = money.Construct(5)
	var value = reflect.ValueOf(five).FieldByName("amount")
	if value.CanSet() {
		t.Error("You're being robbed.", value)
	}
}
```
컴파일러는 아무 불만이 없고 테스트도 성공합니다. 이제 할 일을 끝냈습니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 x 2 = $10</s>

> <s>Dollar를 VO 로 리펙토링하기</s>

## 테스트를 통해 알게된 것

### 테스트 패키지
이번 장에서 테스트 패키지를 만들어 보았습니다.
대상 패키지를 블랙박스로 하고 테스트하고자 할 때 필요한 테스트 작성 방식입니다.
특히 GO 문서화에서 흔히 볼 수 있는 example 코드도 테스트 패키지에 포함합니다.
```
func ExampleDollar() {
	fmt.Println("var five=money.Construct(5)")
	fmt.Println("ten=five.times(5)")
}
```
위의 코드를 `robber_test.go` 에 추가하고 `godoc -http=:6060` 웹으로 확인해 보세요.
테스트 패키지는 GO 언어 자체에서도 많이 사용됩니다.
[GO 소스코드](https://golang.org/src/) 를 참조하세요. 좋은 소스가 많이 있습니다.

### 컴파일러와 대화하며 코딩하기
컴파일러의 도움을 받으면서 일하는 방식은 그 어떤 선배들의 지원보다도 훌륭한 품질을 보장합니다.
그렇게 되기 위해서 코딩 즉시 컴파일러의 에러를 볼 수 있어야 합니다.
제가 컴파일러를 적극(?) 활용하는 Atom 에디터를 사용하는 이유이기도 합니다.
여러분의 에디터가 이런 기능을 지원하도록 설정하는 것은 여러분의 몫입니다.

### 리플렉션의 탁월함
리플렉션에는 `ValueOf` 이외에도 `TypeOf` 와 같은 여러 메소드가 있습니다.
새로운 언어를 배울 때 저는 우선 테스트 코드를 어떻게 작성하는 지 문서를 찾아보고 다음으로 찾아보는 것이 리플렉션입니다. 그만큼 중요하고 강력한 코딩을 할 수 있기 때문이죠.
또한 객체의 동일성 비교를 위해 `DeepEqual` 메소드를 가지고 있는데 조만간에 여러분이 사용하게 될 메소드 입니다.
패키지 문서를 꼼꼼히 읽어 볼 것을 권합니다.
[reflect](https://golang.org/pkg/reflect/)
