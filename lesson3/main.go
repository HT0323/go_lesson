package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int型
	var i int = 100
	fmt.Println(i + 50)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", int32(i))

	// float
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// bool
	var t, f bool = true, false
	fmt.Println(t, f)

	// string
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)
	fmt.Println(`test
	test
	d`)
	fmt.Println(string(s[0]))

	// byte
	byteA := []byte{49}
	fmt.Println(byteA)
	fmt.Println(string(byteA))
	c := []byte("1")
	fmt.Println(c)

	// array
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)
	fmt.Println(arr1[0])
	arr2[2] = "C"
	fmt.Println(arr2)
	fmt.Println(len(arr1))

	// interface
	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)

	var i1 int = 1
	f64 := float64(i1)
	fmt.Println(f64)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", f64)
	i2 := int(f64)
	fmt.Printf("%T\n", i2)
	var s1 string = "100"
	fmt.Printf("%T\n", s1)
	i3, _ := strconv.Atoi(s1)
	fmt.Println(i3)
	fmt.Printf("%T\n", i3)
	var i4 int = 200
	s3 := strconv.Itoa(i4)
	fmt.Println(s3)
	fmt.Printf("%T\n", s3)
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)
}
