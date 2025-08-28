package main

import "fmt"

func main() {
	fmt.Println("🎉 Olá! Vamos brincar com Go!")
	
	// Brincando com variáveis
	nome := "Programador"
	idade := 25
	altura := 1.75
	
	fmt.Printf("👋 Olá, %s!\n", nome)
	fmt.Printf("📊 Você tem %d anos e %.2fm de altura\n", idade, altura)
	
	// Brincando com if/else
	if idade >= 18 {
		fmt.Println("✅ Você é maior de idade!")
	} else {
		fmt.Println("❌ Você é menor de idade!")
	}
	
	// Brincando com for
	fmt.Println("\n🔢 Contando de 1 a 5:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("🟢 %d é par\n", i)
		} else {
			fmt.Printf("🔴 %d é ímpar\n", i)
		}
	}
	
	// Brincando com strings
	fmt.Println("\n🎨 Brincando com texto:")
	palavra := "GoLang"
	fmt.Printf("Palavra: %s\n", palavra)
	fmt.Printf("Tamanho: %d letras\n", len(palavra))
	
	// Brincando com números
	fmt.Println("\n🧮 Brincando com matemática:")
	numero := 7
	fmt.Printf("Tabuada do %d:\n", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
	
	// Brincando com booleanos
	fmt.Println("\n🎯 Brincando com verdadeiro/falso:")
	chovendo := false
	quente := true
	
	if chovendo && quente {
		fmt.Println("🌧️ Está chovendo e quente (estranho!)")
	} else if chovendo {
		fmt.Println("🌧️ Está chovendo, leve um guarda-chuva!")
	} else if quente {
		fmt.Println("☀️ Está quente, beba água!")
	} else {
		fmt.Println("❄️ Está frio, vista um casaco!")
	}
	
	// Brincando com arrays
	fmt.Println("\n📦 Brincando com listas:")
	frutas := []string{"🍎", "🍌", "🍊", "🍇", "🍓"}
	fmt.Println("Minhas frutas favoritas:")
	for i, fruta := range frutas {
		fmt.Printf("%d. %s\n", i+1, fruta)
	}
	
	// Brincando com funções
	fmt.Println("\n🎭 Brincando com funções:")
	soma := somar(10, 20)
	fmt.Printf("10 + 20 = %d\n", soma)
	
	quadrado := calcularQuadrado(5)
	fmt.Printf("5² = %d\n", quadrado)
	
	// Brincando com switch
	fmt.Println("\n🎲 Brincando com switch:")
	dia := "quarta"
	switch dia {
	case "segunda":
		fmt.Println("😴 Segunda-feira... que preguiça!")
	case "terça":
		fmt.Println("😐 Terça-feira... ainda não é sexta!")
	case "quarta":
		fmt.Println("😊 Quarta-feira! Metade da semana!")
	case "quinta":
		fmt.Println("😎 Quinta-feira! Quase lá!")
	case "sexta":
		fmt.Println("🎉 SEXTA-FEIRA! Fim de semana chegando!")
	default:
		fmt.Println("😴 Fim de semana! Descanso!")
	}
	
	fmt.Println("\n🎊 Fim da brincadeira! Divirta-se programando em Go!")
}

// Função para somar dois números
func somar(a, b int) int {
	return a + b
}

// Função para calcular o quadrado de um número
func calcularQuadrado(x int) int {
	return x * x
}
