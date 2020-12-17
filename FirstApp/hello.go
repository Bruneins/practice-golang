// para ejecutar se debe ejecutar go run hello.go
package main

import "fmt"

/*Un TIPO de dato "primitivo" es uno que es construido internamente en el lenguaje o
tipo de dato BÁSICO que es construido internamente en el lenguaje. *
*/

//nombres de paquetes short,consis and evocative

//codigo de go idiomatico, escrito con los estandares puestos por la comunidad

//para no usar el valor retornado por la una funcion se usa el _ underscore

//variable declarada en scoope del package
//no es buena practica tener demasiadas variables con scope global
var variablez = "hola func numero"

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
//creando variable con propio tipo de dato
var a int

type dinero int //creacion de un tipo de dato
var b dinero    //declaracion de variable con tipo de dato creado

func main() {
	//println recive una catidad variable de parametros ..interface{}
	fmt.Println("hello world")

	fmt.Println("------------operador de asignacion corta y println recive cantidad variable de parametros")

	//operador de declaracion de variable corta :=
	variablex := 2 + 4
	variabley := 42
	fmt.Println(variablex, variabley)
	funcion1()
	funcion2()

	fmt.Println("-------------tipo de datos y string")
	tiposString()

	fmt.Println("-------------creacion de un propio tipo de dato")
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//a = b no se puede asignar un tipo de datos diferente
	//convercion de variable
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
