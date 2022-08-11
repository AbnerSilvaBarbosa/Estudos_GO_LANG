package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {

	return fmt.Sprintln("String")
}

func f4(p1, p2 string) string {
	return fmt.Sprintln(p1, p2)
}

func f5() (string, float64) {
	return "String", 123.00
}

func main() {
	resultado3, resultado4 := f3(), f4("Teste4", "Teste4.1")
	resultado51, resultado52 := f5()

	f1()
	f2("Abner", "Teste")
	fmt.Println(resultado3)
	fmt.Println(resultado4)
	fmt.Println(resultado51, resultado52)

}
