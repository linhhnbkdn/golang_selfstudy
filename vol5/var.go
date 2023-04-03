package main

import (
	"fmt"
	"strconv"
)

// Global variables
var g_name string = "John"
var (
	g_name1  string = "John"
	g_name2  string = "John"
	g_name3  string = "John"
	g_shadow int    = 42
	G_N             = 42
)

func main() {
	// Declaring a variable
	// Block variable
	var i int
	i = 42
	var j int = 27
	k := 99.3
	var z string = strconv.Itoa(i)
	var g_shadow int = 99

	fmt.Println(i, j, k)
	fmt.Printf("%T %T %T", i, j, k)
	fmt.Println()
	fmt.Printf("%v %T", g_name, g_name)
	fmt.Println()
	fmt.Println(g_shadow)
	fmt.Println(z)
}

// Declaring a variable with var is the same as declaring a variable with let in JavaScript.
// The only difference is that var is a keyword in Go, and let is a keyword in JavaScript.

// priority of var
// 1. inside a function
// 2. outside a function
// 3. package level

// Declaring and not using a variable is an error in Go.

// If global variable start with a capital letter, it is exported. If it starts with a small letter, it is not exported.
// Rule of name in golang is called camelCase.
