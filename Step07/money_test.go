package main

import (
	"reflect"
	"testing"
)

// TODO : $5+10CHF=$10(환율 2:1 일 경우)

// FIXED : Dollar 부작용?
func TestValueMultiplication(t *testing.T) {
	var five = new(Money).Dollar(5)
	var product = five.times(2)
	if product.amount != 10 {
		t.Errorf("Amount is expected 10, but actually %d ", product.amount)
	}
}

// FIXED : Equals
func TestEquals(t *testing.T) {
	var five = new(Money).Dollar(5)
	isFive, _ := five.equals(new(Money).Dollar(5))

	if !isFive {
		t.Errorf("Dollar equals method is invalid %s", five)
	}
	// 구현 오류를 검증할 수 있어야 한다. [경계값 분석]
	var ten = new(Money).Dollar(10)
	isTen, _ := ten.equals(new(Money).Dollar(1))
	if isTen {
		t.Errorf("Dollar equals method is invalid")
	}

	// Value 이므로 Int 형을 제공하더라도 일치성을 유지해야 한다.
	var six = new(Money).Dollar(6)
	isSix, _ := six.equals(6)
	if !isSix {
		t.Errorf("Dollar equals method is invalid")
	}
	isNil, _ := six.equals(nil)
	if isNil {
		t.Errorf("Dollar equals method is invalid")
	}

}

// FIXED : Successful Failure Test
func TestEqualsFail(t *testing.T) {
	var six = new(Money).Dollar(6)
	_, err := six.equals("InSane")
	if err == nil {
		t.Errorf("Dollar equals method is not implemented with successful failure")
	}
	// fmt.Printf(err.Error())
}

// 배열,슬라이스,맵, 구조체의 필드 등 객체의 모든 요소에 대해서 동치성을 검사.
// 어느 한쪽이 nil 이면 실패한다.
// Equality Check from Reflection, DeepEqual() scan elements of Dollar
func TestObjectEquality(t *testing.T) {
	var five = new(Money).Dollar(5)
	if !reflect.DeepEqual(new(Money).Dollar(10), five.times(2)) {
		t.Errorf("Amount is expected 10, but actually %s ", five.times(2))
	}
}

// FIXED: KRW 추가
func TestKRW(t *testing.T) {
	var five = new(Money).Won(5)
	if !reflect.DeepEqual(new(Money).Won(10), five.times(2)) {
		t.Errorf("Amount is expected 10, but actually %s ", five.times(2))
	}
}

func TestMoney(t *testing.T) {
	var five = new(Money).Dollar(5)
	if reflect.DeepEqual(five, new(Money).Won(5)) {
		t.Error("Two value objects is not equal, but Money.equals() method return true")
	}
	if new(Money).Won(5).currency == five.currency {
		t.Error("Currency of two value objects is equal")
	}
}

func TestMoneyEquals(t *testing.T) {
	var five = new(Money).Dollar(5)
	var isEqual, _ = five.equals(new(Money).Won(5))

	if isEqual {
		t.Error("Two value objects is not equal, but Money.equals() method return true ")
	}
}
