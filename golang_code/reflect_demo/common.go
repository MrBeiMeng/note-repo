package reflect_demo

import (
	"fmt"
	"reflect"
)

func main() {

	a := 1000

	type Animal struct {
		Age int
		Sex string
	}

	type Human struct {
		Animal
		Id string
	}

	var human Human
	human.Id = "9527"
	human.Sex = "男"

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Printf("t :%v\n", t)
	fmt.Printf("v :%v\n", v)

}
