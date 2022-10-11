package main

import "fmt"

func main() {
	println("Hello World!")
	fmt.Println("Hello" + "luhanmin")

	var age int
	age = 23
	println("luhanmin is ", age)

	s := fmt.Sprintf("%d -> luhanmin", age)
	println(s)

	var year, month, day int
	year = 2022
	month = 4
	day = 12
	println(fmt.Sprintf("luhanmin born in %d %d %d", year, month, day))

	for i := 0; i < 10; i++ {
		println(i)
	}

}
