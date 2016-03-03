package money

import (
	"testing"
)

// TODO : $5+10CHF=$10(환율 2:1 일 경우)

// TODO : amount를 private으로 만들기

// FIXED : $5x2=$10
func TestMultiplication(t *testing.T) {

	var five = Dollar{5}
	var ten = five.times(2)
	if ten.amount != 10 {
		t.Errorf("Amount is expected 10, but actually %d ", ten.amount)
	}
}
