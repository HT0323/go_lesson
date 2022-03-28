package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("i,m a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}

}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i4, _ := Div(9, 3)
	fmt.Println(i4)

	i5 := Double(1000)
	fmt.Println(i5)

	Noreturn()

	f := func(x, y int) int {
		return x + y
	}
	i6 := f(1, 2)
	fmt.Println(i6)

	i7 := func(x, y int) int {
		return x + y
	}(1, 4)
	fmt.Println(i7)

	f2 := ReturnFunc()
	f2()

	CallFunction(func() {
		fmt.Println("I,m a function")
	})

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("Hello2"))
	fmt.Println(f3("Hello3"))
	fmt.Println(f3("Hello4"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

	fmt.Println(ints())
}
