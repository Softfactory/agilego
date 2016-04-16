package money

import (
	"log"
	// "math"
	"reflect"
	"testing"
)

// 0.000000
// Moved to source
// var epsilon = math.Nextafter(1, 2) - 1.0

// FIXED : $5x2=$10
func TestMultiplication(t *testing.T) {

	var five = Dollar(5)
	var ten = five.times(2)
	if ten.(Money).amount != 10 {
		t.Errorf("Amount is expected 10, but actually %d ", ten.(Money).amount)
	}
}

// FIXED : Equals
func TestEquals(t *testing.T) {
	var five = Dollar(5)
	isFive, _ := five.equals(Dollar(5))

	if !isFive {
		t.Error("Dollar equals Money method is invalid", five)
	}
	// Value 이므로 Int 형을 제공하더라도 일치성을 유지해야 한다.
	var six = Dollar(6)
	isSix, _ := six.equals(6)
	if !isSix {
		t.Error("Dollar equals Integer method is invalid", six)
	}
	isNil, _ := six.equals(nil)
	if isNil {
		t.Error("Dollar equals Nil method is invalid", six)
	}
	isFloat64, _ := six.equals(6.0)
	if !isFloat64 {
		t.Error("Dollar equals Float method is invalid", six)
	}
	isFloat32, _ := six.equals(float32(6.0))
	if !isFloat32 {
		t.Error("Dollar equals Float method is invalid", six)
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
	if fiveWon.currency != KRW {
		t.Errorf("Expected currency is KRW, actual currency is %s", fiveWon.currency)
	}
	var fiveDollar = Dollar(5)
	if fiveDollar.currency != USD {
		t.Errorf("Expected currency is USD, actual currency is %s", fiveDollar.currency)
	}
}

func TestSimplePlus(t *testing.T) {
	var five = Dollar(5)
	var ten = five.plus(Dollar(5))
	var market = new(Market)
	var changed = market.exchange(ten, USD)
	var isTen, _ = changed.equals(Dollar(10))
	if !isTen {
		t.Errorf("Expected $10, actual value is %s", ten)
	}
}

func TestMoneyEquals(t *testing.T) {
	var fiveWon = Won(5)
	var fiveDollars = Dollar(5)
	var market = new(Market)
	if !reflect.DeepEqual(market.exchange(fiveWon, USD), fiveDollars) {
		t.Errorf("Deep Equality Test Failed %s %s", fiveWon, fiveDollars)
	}
}

func TestExchange(t *testing.T) {
	var five = Dollar(5)
	var market = new(Market)
	market.setRate(USD, KRW, 1000)
	var changed = market.exchange(five, KRW)
	var is5ThousandWon, _ = changed.equals(Won(5000))
	if !is5ThousandWon {
		t.Errorf("Expected 5000KRW, actual value is %s", changed)
	}
}

func TestExchangeRate(t *testing.T) {
	var market = new(Market)
	var rate float64 = 1000

	market.setRate(USD, KRW, rate)
	if !floatEquals(rate, market.getRate(USD, KRW)) {
		t.Errorf("Expected 1000, actual value is %f", rate)
	}
	var defaultRate = market.getRate(KRW, CHF)

	if !floatEquals(1.0, defaultRate) {
		t.Errorf("Expected 1.0, actual value is %f", defaultRate)
	}
}

// Moved to source
// func floatEquals(a, b float64) bool {
// 	if math.Abs(a-b) <= epsilon {
// 		log.Printf("Epsilon is %f ", epsilon)
// 		return true
// 	}
// 	log.Printf("Input Value is %f , %f ", a, b)
// 	log.Printf("(estimated value) %f > %f (epsilon on this machine) ", math.Abs(a-b), epsilon)
// 	return false
// }

//$5 + 5000원 = $10 (1:1000 환율일때 )
func TestSummation(t *testing.T) {
	var fiveDollars = Dollar(5)
	var fiveThousandWon = Won(5000)
	var market = new(Market)
	market.setRate(KRW, USD, 0.001)
	var sum = fiveDollars.plus(fiveThousandWon)

	if !reflect.DeepEqual(market.exchange(sum, USD), Dollar(10)) {
		t.Error("Expected $10", market.exchange(sum, USD))
	}
	if !reflect.DeepEqual(sum.exchange(*market, USD), Dollar(10)) {
		t.Error("Expected $10", sum.exchange(*market, USD))
	}
}

func TestSumEquals(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = Sum{fiveDollars, fiveDollars}
	var sum1 = fiveDollars.plus(fiveDollars)
	var isEquals, _ = sum.equals(sum1)
	if !isEquals {
		t.Error("Sum Equals failed ", sum)
	}
	var _, err = sum.equals(fiveDollars)
	if err == nil {
		t.Error("Sum have not to be equals to Money")
	}
	var isNullEqual, _ = sum.equals(nil)
	if isNullEqual {
		t.Error("Sum equals must failed with nil")
	}
}

//
func TestSumTimes(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = Sum{fiveDollars, fiveDollars}.times(2)
	var market = new(Market)
	if !reflect.DeepEqual(sum.exchange(*market, USD), Dollar(20)) {
		t.Error("Sum Times is failed ", sum)
	}

	//테스트에서도 중복을 최대한 줄여야 한다.
	//여기서는 중복이 아니라 처리 순서에 따른 확인이다.
	var sixDollars = Dollar(6)
	var newSum = Sum{fiveDollars, sixDollars}
	var sum1 = newSum.plus(newSum)
	var sum2 = sum1.times(2)
	var sum3 = Sum{Dollar(10), Dollar(12)}
	if !reflect.DeepEqual(sum2, Sum{sum3, sum3}) {
		t.Error("Sum is failed ", sum2)
	}
}

func TestSumPlus(t *testing.T) {
	var fiveDollars = Dollar(5)
	var sum = fiveDollars.plus(fiveDollars)
	var sum1 = sum.plus(sum)
	var sum2 = sum.plus(fiveDollars)
	var market = new(Market)
	if !reflect.DeepEqual(sum1.exchange(*market, USD), Dollar(20)) {
		t.Error("Sum Plus is failed ", sum1)
	}
	if !reflect.DeepEqual(sum2.exchange(*market, USD), Dollar(15)) {
		t.Error("Sum Money", sum2)
	}
}

func TestNewConstruct(t *testing.T) {
	var fiveWon, _ = Construct(5, KRW)
	if !reflect.DeepEqual(fiveWon, Won(5)) {
		t.Error("Costruct Error", fiveWon)
	}
}

func TestFailedConstruct(t *testing.T) {
	var fiveWon, err = Construct(5, KRW, USD)
	if err.Error() != "Too many arguments!" {
		t.Error("Failed Construct Test Error")
	}
	log.Println("Failed Construct test is success ", fiveWon)
}

func TestCounterfeitCurrency(t *testing.T) {
	var counterFeit, err = Construct(5, MAX)
	if err.Error() != "Invalid Currency Code Error" {
		t.Error("Test for Detecting Counterfeit Failed")
	}
	log.Println("Failed Currency test is success ", counterFeit)
}
