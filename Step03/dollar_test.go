package money

import (
	"testing"
)

// FIXED : $5x2=$10
func TestMultiplication(t *testing.T) {

	var five = Dollar{5}
	var ten = five.times(2)
	if ten.amount != 10 {
		t.Errorf("Amount is expected 10, but actually %d ", ten.amount)
	}
}

// FIXED : Equals
func TestEquals(t *testing.T) {
	var five = Dollar{5}
	isFive, _ := five.equals(Dollar{5})

	if !isFive {
		t.Errorf("Dollar equals method is invalid")
	}
	// Value 이므로 Int 형을 제공하더라도 일치성을 유지해야 한다.
	var six = Dollar{6}
	isSix, _ := six.equals(6)
	if !isSix {
		t.Errorf("Dollar equals method is invalid")
	}
	isNil, _ := six.equals(nil)
	if isNil {
		t.Errorf("Dollar equals method is invalid")
	}
}

func TestEqualsFail(t *testing.T) {

	var six = Dollar{6}
	isSix, err := six.equals("InSane")
	if err == nil {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	if isSix {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	// fmt.Printf(err.Error())
}
