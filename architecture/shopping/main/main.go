package main

import (
	"fmt"
	"shopping/good"
	"shopping/order"

	"github.com/fatih/color"
)

func main() {
	good.GetGood()
	fmt.Println(good.Name)
	order.GetOrder()
	fmt.Println("hello world")

	color.Red("hello world")
}
