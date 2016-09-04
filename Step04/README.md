# Step 4.  갈림길에서

## 쉬지말고 의심하라

### times()의 동치성
동치성에 대한 테스트(`$5=$5`)는 완료되었지만 `times()` 로 반환된 객체가 동치성을 보장하는지는 아직 미지수 입니다.

`times()`를 위한 동치성 테스트를 추가합니다.

```go
func TestTimesEquals(t *testing.T) {
	var five = Dollar{5}
	var ten = five.times(2)
	isTen, _ := ten.equals(Dollar{10})
	if !isTen {
		t.Errorf("Dollar equals method is invalid")
	}
}
```
다행히 테스트가 성공하고 아무 문제도 발생하지 않습니다.

### 테스트가 많을수록 품질이 좋은가?
테스트에서 가장 중요한 이슈는 아마도 얼마나 많은 테스트를 작성해야 하는가에 대한 논쟁입니다.
"Code Complete"의 저자 스티브 맥코넬은 테스트를 더 많이 한다고 해서 품질이 올라가는 것은 아니라고 했습니다.
여러 다양한 실천 방법들을 사용해서 적용하라고 합니다.

꼭 필요한 테스트 범위 이외에 더 많은 테스트를 작성하는 것은 불필요하거나 자원을 낭비하는 일입니다.
테스트 범위가 제대로 수행되고 있는지 확인하는 좋은 지표가 있습니다.
테스트 코드의 라인수와 소스 코드의 라인수를 비교해보시기 바랍니다.
소스 코드에 비해 테스트 코드의 라인수가 합리적 이유없이 적거나 너무 큰 경우 프로젝트가 위험에 빠졌다는 신호일 것입니다.  
구체적인 기준은 프로젝트 마다 달라집니다.
현재 우리가 작성하고 있는 코드는 소스 코드보다 테스트 코드가 1.26배 더 많습니다.
최종으로 작성하게 될 코드도 2배가 약간 넘는 수준입니다.

## 추상화의 길목

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 x 2 = $10</s>

> <s>Dollar를 VO 로 리펙토링하기</s>

> <s>$5=$5 인 equals 구현</s>

이제 할일에 남은 것은 다중 통화를 위한 기능입니다.
지금까지는 `Dollar`를 사용했는데 더이상 여기에 머무를 수가 없게 되었습니다.

### 바로가는 길(?)
다중통화를 생각하다 보면 우리는 객체 지향적 관점에서 상속을 떠올리게 됩니다.
GO에서 상속관계에 대한 지정은 `is a`와 `has a`의 관계를 지정할 수 있습니다.
```go
func TestCar(t *testing.T) {
	var car1 = SuperCar{Car{"SUV"}}
	car1.car.start()
	var car2 = MyCar{Car{"Truck"}}
	car2.start()
}

type Car struct {
	model string
}
type SuperCar struct {
	car Car
}
type MyCar struct {
	Car
}

func (car Car) start() {
	fmt.Printf(" %s : 부르릉 ~~~\n", car.model)
}
```
위에서 `MyCar`는 `Car`를 이름없는 필드로 지정해서 그 자신이 `Car`에 속함을 정의하고 있고,
`SuperCar`는 필드이름을 정의함으로써 Car를 소유하는 관계임을 표현합니다.

그렇다면 우리가 바로 이런 객체 지향적 관계로 소스 코드를 수정할 수 있을까요?
아닙니다. 현재의 소스 코드를 그렇게 고칠 수 있다고 자신할 수 없습니다.
좀더 먼 길로 돌아가야 할 수 있을 지 판단할 수 있습니다.

### KRW 통화 테스트
우선 `Dollar` 가 아닌 화폐가 정상적으로 동작할 것인지를 우리는 보장해야 합니다.
따라서 할일을 다음과 같이 추가합니다.

> **TODO**

> 5원 = 5원

```go  
func TestKRW(t *testing.T) {
	var five = Won{5}
	var isFive, _ =five.equals(Won{5})
	if !isFive{
		t.Errorf("Amount is expected 10, but actually %d ", five.amount)
	}
}
```
컴파일 에러를 해결하기 위해 소스 코드를 수정합니다.
별다른 고민없이 Dollar에서 구현 코드를 그대로 복사하고 일부를 수정하면 됩니다.

```go
type Won struct {
	amount int
}

func (won *Won) equals(object interface{}) (bool, error) {
	defer func() {
		fmt.Println(object, reflect.TypeOf(object))
	}()

	switch v := object.(type) {
	case nil:
		return false, nil
	case int, int32, int64:
		return won.amount == v, nil
	case Dollar:
		return won.amount == v.amount, nil
	default:
		var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcuableError
	}
}
```
테스트는 이제 정상입니다.

### 중복제거
이렇게 짠 코드에는 문제가 있는데 equals의 중복이 보입니다.
이는 제거해야 될 요소입니다. 이제 여러분의 개발 능력에 따라 갈림길에 서게 될 것입니다.
하지만 걱정할 필요는 없습니다. 문제가 생기면 다시 갈라졌던 위치로 돌아오면 됩니다.
할일은 다음과 같이 바뀝니다.

> **TODO**

> <s>5원 = 5원</s>

> equals() 중복 제거

중복을 제거하기 위해서는 두가지 방법이 있어 보입니다.
`Java`와 같은 객체지향 언어에서는 당연히 상속관계를 사용하는 것이 중복을 제거하는 올바른 길이라고 생각할 수 있습니다. 하지만 GO는 객체지향 언어가 아닙니다.

일단 잘못될 가능성을 염두에 두고 앞에서 설명한 `is a` 로 소스를 수정합니다.
테스트는 일단 정상이므로 소스 코드를 수정합니다.
```go
type Money struct {
	amount int
}
type Dollar struct {
	Money
}
type Won struct {
	Money
}
```
이렇게 수정하고 나면 컴파일러에 의해서 우리가 얼마나 많은 부분을 수정해야 하는지 알 수 있습니다.
우선은 `Dollar`를 반환하는 모든 메소드에서 `Dollar{Money{5}}` 와 같이 혐오스럽게 보이도록 코드를 수정해야 합니다. 테스트 코드 역시 그대로 사용할 수가 없는데 우리가 해왔던 많은 객체 생성 부분을 수정해야 합니다.
대략 10 줄 정도가 수정되어야 한다고 컴파일러가 알려줍니다.

기존의 테스트를 최소한으로 수정하면서 중복 기능을 제거할 방법을 찾아야 합니다.

### 생성자 패턴을 활용하기
사실 우리가 통화에 대한 개념에서 잘못 알고 있는 것이 있습니다.
`equals`와 `times` 는 통화(즉, 돈)이 갖는 기본 기능입니다.
`Dollar`라는 것은 이 통화의 속성 중에 하나일 뿐입니다. 통화기호에 따라 `equals` 와 `times`의 논리적 구현이 바뀌지는 않을 것이기 때문입니다.
즉 `Money`에 어느 나라의 통화인지 여부만 알 수 있으면 우리가 만들었던 테스트는 여전히 유효합니다.

이후에 우리가 주목해야 할 것은 객체 생성시점에 통화기호를 입력해야 한다는 점입니다.
누가 보더라도 명확하게 특정 통화 기호(KRW 나 USD 같은)를 지닌 `Money` 객체를 갖고 싶다는 것을 표현합니다.

또한 통화기호를 잊어버리지 않도록 할일을 추가합니다.

> **TODO**

> 5원의 통화기호는 'KRW' , $5 의 통화기호는 'USD'

> <b>equals() 중복 제거</b>

하고자 하는 것은 여전히 'equals() 중복 제거'입니다. 테스트 코드에서 `Dollar{5}` 와 같은 표현을 생성자를 사용하는 것으로 변경합니다.
테스트 코드를 수정했으니 이제 소스 코드를 리팩토링해서 컴파일 에러를 없애는 과정을 설명할 것입니다.
첫번째 순서는 `cannot convert 5 (type int) to type Dollar` 와 같이 객체 생성이 아닌 생성자 메소드 호출로 변경하는 작업입니다.
우리는 `Construct`라는 생성자 메소드를 만들었기 때문에 쉽게 `Dollar` 구조체를 대체할 `Dollar()` 생성자를 작성할 수 있습니다. 이때 조금더 과감해져야 하는데 `Dollar` 와 `Won` 구조체를 삭제합니다. 테스트가 안정적으로 뒷받침 해줄 것이기 때문에 혹시라도 무언가 잘못되더라도 돌아오면 됩니다.

```go
func Dollar(amount int) Money {
	return Money{amount}
}

func Won(amount int) Money {
	return Money{amount}
}

func Construct(amount int) Money {
	return Money{amount}
}
```
생성자를 추가하면 다음 순서로 `equals()` 와 `times()` 메소드를 수정하는 일입니다. 메소드의 리시버와 반환 타입을 수정합니다.

```go
func (money *Money) times(multiplier int) Money {
	return Money{money.amount * multiplier}
}

func (money *Money) equals(object interface{}) (bool, error) {
	defer func() {
		fmt.Println(object, reflect.TypeOf(object))
	}()

	switch v := object.(type) {
	case nil:
		return false, nil
	case int, int32, int64:
		return money.amount == v, nil
	case Money:
		return money.amount == v.amount, nil
	default:
		var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcuableError
	}
}
```
`Won` 구조체를 리시버로 하는 `equals()` 는 이제 삭제할 수 있습니다. 소스 코드를 저장하면 커버리지 표시가 뜹니다. 할일에서 또하나를 삭제할 수 있게 되었습니다.
> **TODO**

> <s>equals() 중복 제거</s>

### 통화 기호 넣기
이제 통화기호를 넣을 차례입니다. 통화기호는 [환율정보](http://kr.fxexchangerate.com)를 참조합니다.
지금까지 테스트 코드는 통화를 미리 지정하기 때문에 통화기호를 테스트할 수 없습니다.

> **TODO**

> <b> 5원의 통화기호는 'KRW' , $5 의 통화기호는 'USD' </b>

테스트를 추가합니다.
```go
func TestCurrency(t *testing.T) {
	var fiveWon = Won(5)
	if fiveWon.currency != "KRW" {
		t.Errorf("Expected currency is KRW, Actual currency is %s", fiveWon.currency)
	}
	var fiveDollar = Dollar(5)
	if fiveDollar.currency != "USD" {
		t.Errorf("Expected currency is USD, Actual currency is %s", fiveDollar.currency)
	}
}
```

`currency` 필드가 없기 때문에 컴파일 에러가 발생합니다. `Money` 구조체를 수정합니다.
```go
type Money struct {
	amount   int
	currency string
}
```
`Money`를 사용하는 생성자는 물론 `times()`에서 컴파일 에러가 보입니다. 하지만 테스트를 믿고 쉽게 수정할 수 있습니다. 모두 `return Money{amount, "USD"}` 형태로 변경하면 됩니다.  `times()`를 예로들면 다음과 같이 변경됩니다.

```go
func (money *Money) times(multiplier int) Money {
	return Money{money.amount * multiplier, money.currency}
}
```

모두 정상으로 바뀝니다. 하지만 아직 끝나지 않았습니다. `Construct` 생성자 부분에서 우리는 중복을 발견했습니다. `Dollar` 와 동일합니다. 중복을 피하기 위해서라도 `Dollar(amount)` 로 변경합니다.  

지금까지 우리가 해온 많은 일들을 한번 살펴볼 때가 되었습니다.

> **TODO**

> $5 + 5000원 = $10 (1:1000 환율일때 )

> <s>$5 x 2 = $10</s>

> <s>Dollar를 VO 로 리펙토링하기</s>

> <s>$5=$5 인 equals 구현</s>

> <s>5원 = 5원</s>

> <s>equals() 중복 제거</s>

> <s>5원의 통화기호는 'KRW' , $5 의 통화기호는 'USD' </s>

## 테스트를 통해 알게 된 것

### GO의 객체지향성
이 장에서 상속이 가능하다는 사실을 알았습니다. GO에서 다중 상속은 물론 다형성도 가능합니다.
뿐만 아니라 서브 클래스를 통해서 상위 클래스를 호출할 수도 있습니다.
```go
func TestRentCar(t *testing.T) {
	var car = RentCar{Car{"Truck"}, Rental{}}
	car.getPrice()
	car.start()
	car.Rental.getPrice()
}

type Rental struct {
}

type RentCar struct {
	Car
	Rental
}

func (rental Rental) getPrice() {
	fmt.Printf("  $10/1h \n")
}

func (special RentCar) getPrice() {
	fmt.Printf("  $5/1h \n")
}
```

### 테스트는 비계입니다.
국어사전에 비계는 다음과 같이 정의됩니다.

> [건축] 높은 건물을 지을 때 디디고 서도록 긴 나무 따위를 종횡으로 엮어 다리처럼 걸쳐 놓은 설치물.

우리가 만드는 테스트는 이 비계와 같습니다. 비계도 테스트와 마찬가지로 건물이 한층한층 올라갈때 마다 커지기 때문입니다. 무작정 비계를 높이 올리지도 건물만 높이 올리지도 않습니다. 한층씩 올라갈때 마다 비계를 올리고 올라간 층을 마무리하고 다시 다음 층에서 반복적인 작업을 수행합니다.
그런데 이번 장에서 작성한 코드를 켄트 벡의 책에서는 좀더 호흡이 길게 여러 장으로 이어 갑니다. 아마도 "통화기호라는 것이 Money의 속성중 하나이다."를 전제하기 보다는 객체 지향적 접근으로 소스를 어떻게 지속적으로 성장시켜 나가는지 보여주려는 의도같습니다. GO의 언어적 특성과 통화에 대한 고찰을 통해 벡의 객체지향적 사고 과정을 거쳐가지 않았고 따라서 좀더 짧은 호흡으로 코드를 성장시켰습니다. 호흡의 차이만 있을 뿐 결론은 같습니다.
