// Package money TDD를 위한 IMoney 패키지 구현
package main

import "fmt"

//Money 구조체
//Times, Equals 의 중복을 제거하기 위해..
type Money struct {
	amount   int
	currency string
}

// //Dollar 구조체
// type Dollar struct {
// 	Money
// }
//
// //Won 구조체
// type Won struct {
// 	Money
// }

// Dollar 생성자
func (money *Money) Dollar(amount int) Money {
	// money :=
	return Money{amount, "USD"}
}

// Won 생성자
func (money *Money) Won(amount int) Money {
	// money :=
	return Money{amount, "KRW"}
}

// Value Object Pattern으로 구현
func (money *Money) times(multiplier int) Money {
	return Money{money.amount * multiplier, money.currency}
}

// Successful Failure 를 추가한 Object 에 대한 동치성 구현
func (money Money) equals(i interface{}) (bool, error) {
	if i == nil {
		return false, nil
	} else if v, isInt := i.(int); isInt {
		return money.amount == v, nil
	} else if v, isMoney := i.(Money); isMoney {
		return money.amount == v.amount && money.currency == v.currency, nil
		// return money.amount == v.amount, nil
	}
	var NotCalculableError = fmt.Errorf("This value is not calculable.")
	return false, NotCalculableError
	// return true, nil
	// panic(NotCalcuableError)
}
