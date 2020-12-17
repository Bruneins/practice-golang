package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

type vuelta int

var m vuelta
var n int

func main() {

	fmt.Println("--------------Ejercicio practico 1")
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("1. ", x, y, z)
	fmt.Println("a.", x)
	fmt.Println("b.", y)
	fmt.Println("c.", z)

	fmt.Println("--------------Ejercicio practico 2")
	fmt.Println("a.", a)
	fmt.Println("b.", b)
	fmt.Println("c.", c)

	fmt.Println("--------------Ejercicio practico 3")

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

	fmt.Println("--------------Ejercicio practico 4")

	fmt.Println(m)
	fmt.Printf("tipo de dato %T \n", m)
	m = 42
	fmt.Println(m)

	fmt.Println("--------------Ejercicio practico 5")
	fmt.Println(m)
	fmt.Printf("El tipo de dato es: %T \n", m)
	m = 69
	n = int(m)

	fmt.Printf("El valor es %v y su tipo de dato es %T \n", n, n)

}
