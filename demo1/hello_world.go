package main

import "fmt"

/*
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
		count := 5
		for count < 10 {
			println(count)
			count++
		}
	}
*/
/*
func temp() ([]int, int) {
	var temp_slice []int = make([]int, 2, 4)
	temp_slice = append(temp_slice, 3)
	temp_slice = append(temp_slice, 5)
	fmt.Println(temp_slice)
	sort.Slice(temp_slice, func(i, j int) bool {
		return temp_slice[i] > temp_slice[j]
	})
	q := temp_slice[len(temp_slice)-1]
	fmt.Println(reflect.TypeOf(temp_slice))
	return temp_slice, q
}
*/
/*
func temp(temp_map map[int]string) map[int]string {
	temp_map[2] = "char"
	temp_map[3] = "int"
	temp_map[1] = "string"
	temp_map[0] = "nil"
	return temp_map
}
*/
type Temp interface {
	sell()
}

type Book struct {
	Name  string
	Price int
}

type NeoBook struct {
	Book
	Selled bool
}

func (book Book) sell() {
	fmt.Println("Book sell success")
}

func main() {
	var temp Temp
	temp = new(Book)
	temp.sell()
}
