package main

import "fmt"

func main() {
	fmt.Println("🎉 Olá! Vamos brincar com Go!")
	
	esdras := "esdras"


	fmt.Println(esdras)
	


	teste := 1

	if teste == 1 {
		fmt.Println("teste é 1")
	} else {
		fmt.Println("teste é 2")
	}



	fatorialde5 := fatorial(5)
	fmt.Println(fatorialde5)

	
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}