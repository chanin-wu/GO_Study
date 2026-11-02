package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Name string

func (n Name) PrefixString() Name {
	return "挂载方法-" + n
}

func (n Name) RandomString() Name {
	return Name(strconv.Itoa(rand.Intn(10))) + n
}

func MountingMethod() {

	var name Name = "前缀"

	newName := name.PrefixString()
	fmt.Println(newName)

	RandomName := name.RandomString()
	fmt.Println(RandomName)
}

// 1、挂载方法
// 2、类型安全
func main12() {
	MountingMethod()
	TypeSafety()
}

type UID int
type OID int

func CreateType(uid UID, oid OID) (UID, OID) {
	return uid, oid
}

func TypeSafety() {
	var uid UID = 1
	var oid OID = 2
	// CreateType(oid, uid) // 传入错误的类型
	CreateType(uid, oid)
}
