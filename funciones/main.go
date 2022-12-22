package main

import (
	"errors"
	"fmt"
)

func main() {
	//Func1()
	//Func2("Juan")
	resultado, err := Divicion(2, 1)
	if err != nil {
		// aca termino porque tenemos erorroes
		fmt.Print(err)
	}
	fmt.Println(resultado)
}

func Func1() {
	fmt.Println("start")
}

func Func2(name string) {
	fmt.Println("Hola " + name)
}

func Test(a int) (int, int) {
	return a, a + 3
}

func Divicion(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No ingreses un cero, no sea pendejo")
	}
	return a / b, nil
}
