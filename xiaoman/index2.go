package main

import "fmt"

func main2() {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)

	b := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(b)

	c := "鑫翔七迅"

	/* for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	} */

	/* for key, value := range b {
		fmt.Println(key, value)
	} */

	/* for byIndex, uniCode := range c {
		fmt.Println(byIndex, uniCode)
	} */

	for key, value := range []rune(c) {
		fmt.Println(key, string(value))
	}
}
