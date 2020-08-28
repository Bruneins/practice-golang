package main

import "fmt"

//variable declarada en scoope del package
//no es buena practica tener demasiadas variables con scope global
var variablez = "hola func numero"

func main() {
	//println recive una catidad variable de parametros ..interface{}
	fmt.Println("hello world")

	//operador de declaracion de variable corta
	variablex := 2 + 4
	variabley := 42
	fmt.Println(variablex, variabley)
	funcion1()
	funcion2()
	tiposString()
}

func funcion1() {
	fmt.Println(variablez)
}

var variablew bool

//varible igual a 0, false cuando no se la inicializa
func funcion2() {
	fmt.Println(variablew)
}

//
func tiposString() {
	variablea := "hola como estas" //interpreted string literal
	variableb := `hola que tal como estas
	estes es un row string literal` //row string literal
	fmt.Println(variablea, variableb)
	variablec := 40
	variabled := "programa"
	fmt.Printf("El valor de la variablec es: %v \n", variablec)       //
	fmt.Printf("El valor de la variabled es: %v \n", variabled)       //
	fmt.Printf("El tipo de dato d ela variablec es %T \n", variablec) //para saber el tipo de dato
	fmt.Printf("El tipo de dato d ela variabled es %T \n", variabled)

	string1 := fmt.Sprint("El ", variabled, " dice hola mundo")
	fmt.Println(string1)
}
//clase siguiente 13
