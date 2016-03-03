package money

import (
	"log"
	"math"
	"reflect"
	"testing"
)

// 0.000000
var epsilon = math.Nextafter(1, 2) - 1.0

// FIXED : $5x2=$10
func TestMultiplication(t *testing.T) {

	var five = Dollar(5)
	var ten = five.times(2)
	if ten.amount != 10 {
		t.Errorf("Amount is expected 10, but actually %d ", ten.amount)
	}
}

// FIXED : Equals
func TestEquals(t *testing.T) {
	var five = Dollar(5)
	isFive, _ := five.equals(Dollar(5))

	if !isFive {
		t.Errorf("Dollar equals method is invalid")
	}
	// Value 이므로 Int 형을 제공하더라도 일치성을 유지해야 한다.
	var six = Dollar(6)
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

	var six = Dollar(6)
	isSix, err := six.equals("InSane")
	if err == nil {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	if isSix {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	// fmt.Printf(err.Error())
}

func TestTimesEquals(t *testing.T) {
	var five = Dollar(5)
	var ten = five.times(2)
	isTen, _ := ten.equals(Dollar(10))
	if !isTen {
		t.Errorf("Dollar equals method is invalid")
	}
}

func TestKRW(t *testing.T) {
	var five = Won(5)
	var isFive, _ = five.equals(Won(5))
	if !isFive {
		t.Errorf("Amount is expected 10, but actually %d ", five.amount)
	}
}

func TestCurrency(t *testing.T) {
	var fiveWon = Won(5)
	if fiveWon.currency != "KRW" {
		t.Errorf("Expected currency is KRW, actual currency is %s", fiveWon.currency)
	}
	var fiveDollar = Dollar(5)
	if fiveDollar.currency != "USD" {
		t.Errorf("Expected currency is USD, actual currency is %s", fiveDollar.currency)
	}
}

func TestSimplePlus(t *testing.T) {
	var five = Dollar(5)
	var ten = five.plus(Dollar(5))
	var market = new(Market)
	var changed = market.exchange(ten, "USD")
	var isTen, _ = changed.equals(Dollar(10))
	if !isTen {
		t.Errorf("Expected $10, actual value is %s", ten)
	}
}

func TestMoneyEquals(t *testing.T) {
	var fiveWon = Won(5)
	var fiveDollars = Dollar(5)
	var market = new(Market)
	if !reflect.DeepEqual(market.exchange(fiveWon, "USD"), fiveDollars) {
		t.Errorf("Deep Equality Test Failed %s %s", fiveWon, fiveDollars)
	}
}

func TestExchange(t *testing.T) {
	var five = Dollar(5)
	var market = new(Market)
	market.setRate("USD", "KRW", 1000)
	var changed = market.exchange(five, "KRW")
	var is5ThousandWon, _ = changed.equals(Won(5000))
	if !is5ThousandWon {
		t.Errorf("Expected 5000KRW, actual value is %s", changed)
	}
}

func TestExchangeRate(t *testing.T) {
	var market = new(Market)
	var rate float64 = 1000

	market.setRate("USD", "KRW", rate)
	if !floatEquals(rate, market.getRate("USD", "KRW")) {
		t.Errorf("Expected 1000, actual value is %f", rate)
	}
	var defaultRate = market.getRate("KRW", "CHF")

	if !floatEquals(1.0, defaultRate) {
		t.Errorf("Expected 1.0, actual value is %f", defaultRate)
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
