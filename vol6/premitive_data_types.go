package main

import (
	"fmt"
)

func main() {
	var bool_var bool = true
	bool_var = false == true
	fmt.Printf("%v %T \n", bool_var, bool_var)

	var init_int int8 = 42
	fmt.Printf("%v %T \n", init_int, init_int)
	fmt.Println("Hello World")
}

// Path: vol6/variables.go
/*
Primitive Data Types
1. Boolean
2. Numeric
	2.1 Integer
		unsigned integer
		signed integer
	2.2 Floating Point float32, float64
	2.3 Complex Complex64, Complex128 12.7e+18i
3. Text
	3.1 String
	3.2 Rune
	3.3 Byte
*/
