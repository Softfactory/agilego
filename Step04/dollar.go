package money

import (
	"fmt"
	"reflect"
)

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

// times 는 Dollar의 곱셈 연산을 합니다.
// VO를 구현하기 위해서 `Dollar`를 반환합니다.
func (money *Money) times(multiplier int) Money {
	// dollar.amount = dollar.amount * multiplier
	return Money{money.amount * multiplier, money.currency}
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
