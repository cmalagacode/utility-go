package main

import "fmt"

func main() {
	// like modern languages assigns a zero value to variables not initialized
	var a int
	fmt.Println(a) // 0
	// declare and initialize a variable
	var b int = 10_000
	fmt.Println(b)
	// single quotes and double quotes are not interchangeable
	var c string = "Hello, World!"
	fmt.Println(c)
	var d rune = 'a'
	fmt.Println(string(d))
	// types
	e := true
	fmt.Println(e)
	f := 10.0
	fmt.Println(f)
	g := 10
	fmt.Println(g)
	// byte is a type alias to uint8
	var h byte = 10
	fmt.Printf("Byte: %d\n", h)
	// rune is a type alias to int32
	var i rune = 10
	fmt.Printf("Rune: %d\n", i)
	// Declaring multiple variables at once
	var j, k int
	fmt.Println(j, k)
	var (
		x      int
		y          = 20
		z      int = 200
		dd, ee     = 40, "Hello"
		ff, gg string
	)
	fmt.Println(x, y, z, dd, ee, ff, gg)
	// constants
	const pi float64 = 3.14159265359
	const zzz = "Hello"
	fmt.Println(pi, zzz)
}
