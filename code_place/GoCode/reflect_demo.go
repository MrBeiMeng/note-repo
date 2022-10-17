package main

import (
	"fmt"
	"reflect"
)

func PrintType2(struc reflect.Type) {
	num := struc.NumField()

	for i := 0; i < num; i++ {
		tmpTyp := struc.Field(i)
		if tmpTyp.IsExported() {
			if tmpTyp.Type.Kind() == reflect.Struct {
				PrintType2(tmpTyp.Type)
			}
			if tmpTyp.Tag.Get("lindorm") != "" {
				fmt.Printf("%v\t", tmpTyp.Name)
				fmt.Printf("%v\n", tmpTyp.Type)
			}
		}

	}
}

func PrintType(struc interface{}) {
	typ := reflect.TypeOf(struc)
	if typ.Kind() != reflect.Struct {
		print("不是结构体")
	}
	PrintType2(typ)
}

func main() {
	type Animal struct {
		Age int    `lindorm:"yes2"`
		Sex string `lindorm:"bug"`
	}

	type Human struct {
		Animal
		Id string `lindorm:"yes"'`
	}

	var human Human
	human.Id = "9527"
	human.Sex = "男"
	human.Age = 23

	// 尝试打印所有数据
	PrintType(human)
}
