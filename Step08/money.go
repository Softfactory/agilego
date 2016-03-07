/*
Package money 는 통화에 따른 금전의 환율 변환과 연산을 수행한다.
Money 는 Value Object 패턴을 따르고 있으며,
환율에 대한 정보는 http://www.fxexchangerate.com 을 참조하라.
 TODO: 작업에 필요한 사항을 명시합니다.
*/
package money

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

var epsilon = math.Nextafter(1, 2) - 1.0

// Market 환율을 반영해서 환전을 처리한다. 환율시장을 모델로 한다.
type Market struct {
	rates map[pair]float64
}
type pair struct {
	from Currency
	to   Currency
}

//Operator 수식에서 연산자 인터페이스
type Operator interface {
	exchange(Market, Currency) Money
	times(float64) Operator
	plus(Operator) Operator
	equals(interface{}) (bool, error)
}

// Sum 더하기 연산자
type Sum struct {
	augend Operator
	addend Operator
}

// Money 통화를 나타낸다. 통화기호는 국제 표준을 따른다.
type Money struct {
	amount   float64
	currency Currency
}

// Dollar USD 통화를 나타낸다.
func Dollar(amount float64) Money {
	return Construct(amount, USD)
}

// Won KRW 통화를 나타낸다.
func Won(amount float64) Money {
	return Construct(amount, KRW)
}

// Construct Dollar 생성자
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

// func Construct(amount float64, currency string= USD) Money {
// 	return Money{amount, currency}
// }

// exchange Unexported 메소드에 대한 설명은 웹으로 노출되지 않는다.
func (market Market) exchange(operator Operator, currency Currency) Money {
	return operator.exchange(market, currency)
}

func (market *Market) setRate(from Currency, to Currency, rate float64) {
	if market.rates == nil {
		market.rates = make(map[pair]float64)
	}
	market.rates[pair{from, to}] = rate
}

func (market Market) getRate(from Currency, to Currency) float64 {
	var rate = 1.0
	if market.rates != nil {
		var saved, isRated = market.rates[pair{from, to}]
		if isRated {
			rate = saved
		}
	}
	return rate
}

// exchange 중복된 내용을 제거한다.
// augend와 addend의 currency를 비교할 필요는 없어보인다.
// Money의 exchange가 안전하다면 동일한 currency를 반환할 것이기 때문이다.
// 즉, currency 가 동일하게 반환될 책임은 Sum의 exchange가 아니라 Money의 exchange 가 갖고 있다.
// currency string 통화기호를 의미한다.
func (sum Sum) exchange(market Market, currency Currency) Money {
	var augend = sum.augend.exchange(market, currency).amount
	var addend = sum.addend.exchange(market, currency).amount

	var amount = augend + addend

	return Money{amount, currency}
}

func (money Money) exchange(market Market, currency Currency) Money {
	var rate = market.getRate(money.currency, currency)
	var amount = money.amount * rate
	return Money{amount, currency}
}

//  BUG(Jacob): sum.augend.(Money).amount 같은 접근은 버그를 유발할 수 있다.
func (sum Sum) times(multiplier float64) Operator {
	var augend = sum.augend.times(multiplier)
	var addend = sum.addend.times(multiplier)
	return Sum{augend, addend}
}

// times 는 Money의 곱셈 연산을 수행한다.
// VO를 구현하기 위해서 `Operator`를 반환한다.
func (money Money) times(multiplier float64) Operator {
	return Money{money.amount * multiplier, money.currency}
}

func (sum Sum) plus(addend Operator) Operator {
	return Sum{sum, addend}
}

func (money Money) plus(addend Operator) Operator {
	return Sum{money, addend}
}

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

func (money Money) equals(object interface{}) (bool, error) {
	defer func() {
		fmt.Println(object, reflect.TypeOf(object))
	}()

	switch v := object.(type) {
	case nil:
		return false, nil
	case int, int32, int64:
		var intValue, _ = v.(int)
		return floatEquals(money.amount, float64(intValue)), nil
	case float32:
		return floatEquals(money.amount, float64(v)), nil
	case float64:
		return floatEquals(money.amount, v), nil
	case Money:
		log.Println("Value Type : ", money.amount, v)
		return money.amount == v.amount, nil
	default:
		var NotCalcualbleError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcualbleError
	}
}

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) <= epsilon {
		log.Printf("Epsilon is %f ", epsilon)
		return true
	}
	log.Printf("Input Value is %f , %f ", a, b)
	log.Printf("(estimated value) %f > %f (epsilon on this machine) ", math.Abs(a-b), epsilon)
	return false
}
