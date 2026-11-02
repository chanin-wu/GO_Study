package main

import (
	"fmt"
	"reflect"
)

type Age int

// 读取类型
func ReadType() {
	var age Age = 18

	t := reflect.TypeOf(age)
	fmt.Println(t)
	fmt.Println(t.Name()) // 读取类型名称 Age
	fmt.Println(t.Kind()) // 读取底层类型 int

	if t.Kind() == reflect.Int {
		fmt.Println("绝对是整数类型")
	}
}

func (a Age) GetAge() int {
	return int(a)
}

// 修改值
func ModifiedValue() {
	var age Age = 18
	value := reflect.ValueOf(&age).Elem() // 获取指针指向的值
	if value.CanSet() && value.Kind() == reflect.Int && value.IsValid() {
		value.SetInt(20) // 修改值
	}
	fmt.Println(age, age.GetAge()) // Age

	// 类型断言.(类型)
	v := value.Interface().(Age)
	fmt.Println(value, v.GetAge()) // reflect.Value
}

func (a Age) GetPerson(name string, age int) {
	fmt.Println(name, age)
}

func CallMethod() {
	var age Age = 18
	value := reflect.ValueOf(age)
	person := value.MethodByName("GetPerson")
	if person.IsValid() {
		fmt.Println(age)
		person.Call([]reflect.Value{reflect.ValueOf("xiaoman"), reflect.ValueOf(26)})
	}
}

// 1、 读取类型
// 2、 修改值
func main() {
	ReadType()
	ModifiedValue()
	CallMethod()
}
