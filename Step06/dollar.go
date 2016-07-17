package money

import (
	"fmt"
	"reflect"
)

// Market 환율을 반영한 환전을 담당한다.
type Market struct {
	rates map[pair]float64
}
type pair struct {
	from string
	to   string
}

//Operator 연산자 인터페이스
type Operator interface {
	exchange(Market, string) Money
	times(int) Operator
	plus(Operator) Operator
	equals(interface{}) (bool, error)
}

// Sum Summation Operator
type Sum struct {
	augend Operator
	addend Operator
}

// Money 통화를 나타낸다.
type Money struct {
	amount   int
	currency string
}

// Dollar USD 통화를 나타낸다.
func Dollar(amount int) Money {
	return Money{amount, "USD"}
}

// Won KRW 통화를 나타낸다.
func Won(amount int) Money {
	return Money{amount, "KRW"}
}

// Construct Dollar 생성자
func Construct(amount int) Money {
	return Dollar(amount)
}

func (market Market) exchange(operator Operator, currency string) Money {
	return operator.exchange(market, currency)
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

func (sum Sum) exchange(market Market, currency string) Money {
	var augend = sum.augend.exchange(market, currency).amount
	var addend = sum.addend.exchange(market, currency).amount

	var amount = augend + addend

	return Money{amount, currency}
}

func (money Money) exchange(market Market, currency string) Money {
	var rate = market.getRate(money.currency, currency)
	var amount = float64(money.amount) * rate
	return Money{int(amount), currency}
}

func (sum Sum) times(multiplier int) Operator {
	var augend = sum.augend.times(multiplier)
	var addend = sum.addend.times(multiplier)
	return Sum{augend, addend}
}

// times 는 Dollar의 곱셈 연산을 합니다.
// VO를 구현하기 위해서 `Operator`를 반환합니다.
func (money Money) times(multiplier int) Operator {
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
		var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcuableError
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
		return money.amount == v, nil
	case Money:
		return money.amount == v.amount, nil
	default:
		var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcuableError
	}
}
