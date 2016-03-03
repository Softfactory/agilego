package money

import (
	"testing"
)

// TODO : $5+5000원=$10(환율 1:1000 일 경우)

// FIXED : $5x2=$10
func TestMultiplication(t *testing.T) {

	var five = Dollar{5}
	five.times(2)
	if five.amount != 10 {
		//GO 에는 Assert가 없다. Assert를 구현한 외부 패키지도 있지만,
		//GO 의 Design 사상에는 맞지 않는다.
		t.Errorf("Amount is expected 10, but actually %d ", five.amount)
	}

	var ten = Dollar{10}
	ten.times(3)
	if ten.amount != 30 {
		t.Errorf("Amount is expected 30, but actually %d ", ten.amount)
	}
}
