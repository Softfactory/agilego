package money_test

import (
	"fmt"
	"reflect"
	"AgileGO/money"
	"testing"
)

func TestUnexportedAmount(t *testing.T) {
	var five = money.Construct(5)
	var value = reflect.ValueOf(five).FieldByName("amount")
	if value.CanSet() {
		t.Error("You're being robbed.", value)
	}
}

func ExampleDollar() {
	fmt.Println("var five=money.Construct(5)")
	fmt.Println("ten=five.times(5)")
}
