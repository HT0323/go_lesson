package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func anything(a interface{}) {
	fmt.Println(a)
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("i dont,t know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)

	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	i2 := 0
	for {
		i2++
		if i2 == 3 {
			break
		}
		fmt.Println("loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr2 := [3]int{1, 2, 3}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	anything("aaa")
	anything(1)

	var x2 interface{} = 3
	i4 := x2.(int)
	fmt.Println(i4 + 2)

	if x2 == nil {
		fmt.Println("None")

	} else if i, isInt := x2.(int); isInt {
		fmt.Println(i, "s is int")
	} else if s, isString := x2.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("i not know")
	}
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("start")
	// 				break Loop
	// 			}
	// 			fmt.Println("No func")
	// 		}
	// 	}
	// 	fmt.Println("end")

	// for i5 := 0; i5 < 3; i++ {
	// 	for j := 1; j < 3; j++ {
	// 		fmt.Println(i5, j, i5*j)
	// 	}
	// }

	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")

	}()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error")
	}
	defer file.Close()
	file.Write([]byte("Hello"))

	// panic("runtime error")
	// fmt.Println("start")

	// go sub()
	// go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	fmt.Println("Main")
}
