package money_test

import (
	"fmt"
	"reflect"
	"AgileGO/money"
	"testing"
	"unsafe"
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

func TestCar(t *testing.T) {
	var car1 = SuperCar{Car{"SUV"}}
	car1.car.start()
	var car2 = MyCar{Car{"Truck"}}
	car2.start()
}

func TestRentCar(t *testing.T) {
	var car = RentCar{Car{"Truck"}, Rental{}}
	car.getPrice()
	car.start()
	car.Rental.getPrice()
}

type Car struct {
	model string
}

type Rental struct {
}
type SuperCar struct {
	car Car
}
type MyCar struct {
	Car
}

type RentCar struct {
	Car
	Rental
}

func (rental Rental) getPrice() {
	fmt.Printf("  $10/1h \n")
}
func (special RentCar) getPrice() {
	fmt.Printf("  $5/1h \n")
}
func (car Car) start() {
	fmt.Printf(" %s : 부르릉 치타 ~~~\n", car.model)
}

func TestMemoryAllocation(t *testing.T) {
	var car = new(RentCar)
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car, *&car, unsafe.Sizeof(car))
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car, *&car, reflect.TypeOf(car).Size())

	car.model = "Genesis"
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car, *&car, reflect.TypeOf(car).Size())
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &car.model, *&car.model, reflect.TypeOf(car.model).Size())

	var p = new([]int)       // 슬라이스를 할당하지만 *p == nil 이다.
	var v = make([]int, 100) // 100의 int에 대한 참조를 갖는다.

	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &p, *&p, reflect.TypeOf(p).Size())
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &v, *&v, reflect.TypeOf(v).Size())

	*p = make([]int, 100, 100)
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &p, *&p, reflect.TypeOf(p).Size())

	v1 := make([]int, 100, 100)
	fmt.Printf("Address &i=%p Address Value=%v Size=%d\n", &v1, *&v1, reflect.TypeOf(v1).Size())

}
