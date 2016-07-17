package money

import (
	"fmt"
	"reflect"
)

// Dollar 통화를 나타낸다.
type Dollar struct {
	// 통화금액
	amount int // amount 통화금액이다. 반드시 0 이상이어야 할 필요는 정의하지 않았다.
}

// Construct Dollar 생성자
func Construct(amount int) Dollar {
	return Dollar{amount}
}

// times 는 Dollar의 곱셈 연산을 합니다.
// VO를 구현하기 위해서 `Dollar`를 반환합니다.
func (dollar *Dollar) times(multiplier int) Dollar {
	// dollar.amount = dollar.amount * multiplier
	return Dollar{dollar.amount * multiplier}
}

func (dollar *Dollar) equals(object interface{}) (bool, error) {
	defer func() {
		fmt.Println(object, reflect.TypeOf(object))
	}()

	switch v := object.(type) {
	case nil:
		return false, nil
	case int, int32, int64:
		return dollar.amount == v, nil
	case Dollar:
		return dollar.amount == v.amount, nil
	default:
		var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
		return false, NotCalcuableError
	}
	// if object == nil {
	// 	return false, nil
	// } else if v, isInt := object.(int); isInt {
	// 	return dollar.amount == v, nil
	// } else if v, isDollar := object.(Dollar); isDollar {
	// 	return dollar.amount == v.amount, nil
	// }
	// var NotCalcuableError = fmt.Errorf("This value is not calcuable.")
	// return false, NotCalcuableError
	//
	// return false, NotCalcuableError
}
