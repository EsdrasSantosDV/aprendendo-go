package main

import (
	"fmt"
	"time"
)

// FUNÇÃO PRINCIPAL - ponto de entrada do programa
func main() {
	fmt.Println("🎉 Olá! Vamos aprender Go do básico!")
	
	// ===== VARIÁVEIS E TIPOS =====
	fmt.Println("\n--- VARIÁVEIS E TIPOS ---")
	
	// Declaração explícita de tipos
	var nome string = "Esdras"
	var idade int = 25
	var altura float64 = 1.75
	var ativo bool = true
	

    var peso int =1000
	var name string = "Valor inicial"
	fmt.Println(name)
	fmt.Println(peso)



	// Declaração com inferência de tipo (:=)
	cidade := "São Paulo"
	ano := 2024

	baby :=`
		Oi
		Como vai?
	`
	fmt.Println(baby)
	


	fmt.Printf("Nome: %s, Idade: %d, Altura: %.2f, Ativo: %t\n", nome, idade, altura, ativo)
	fmt.Printf("Cidade: %s, Ano: %d\n", cidade, ano)

	// Mudando o valor da variável
   nome = "João"
   fmt.Println(nome)
	
	// ===== CONSTANTES =====
	fmt.Println("\n--- CONSTANTES ---")
	const PI = 3.14159
	const NOME_APP = "Meu App Go"
	fmt.Printf("PI: %f, Nome do App: %s\n", PI, NOME_APP)
	
	
	// ===== ARRAYS E SLICES =====
	fmt.Println("\n--- ARRAYS E SLICES ---")
	
	// Array fixo
	var numeros [5]int = [5]int{1, 2, 3, 4, 5}

    numeros = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", numeros)
	
	// Slice (array dinâmico)
	frutas := []string{"maçã", "banana", "laranja"}
	frutas = append(frutas, "uva") // Adiciona elemento
	fmt.Printf("Frutas: %v\n", frutas)
	frutas = append(frutas, "uva")
	fmt.Printf("Frutas: %v\n", frutas)
	fmt.Printf("Tamanho da slice: %d\n", len(frutas))
	fmt.Printf("Capacidade da slice: %d\n", cap(frutas))

	frutas = append(frutas, "uva")
	fmt.Printf("Frutas: %v\n", frutas)
	fmt.Printf("Tamanho da slice: %d\n", len(frutas))
	fmt.Printf("Capacidade da slice: %d\n", cap(frutas))
	
	// ===== MAPS (DICIONÁRIOS) =====
	fmt.Println("\n--- MAPS ---")
	
	pessoa := map[string]interface{}{
		"nome":   "João",
		"idade":  30,
		"cidade": "Rio de Janeiro",
	}
	fmt.Printf("Pessoa: %v\n", pessoa)
	
	// Acessando valores do map
	fmt.Printf("Nome da pessoa: %s\n", pessoa["nome"])
	
	// ===== ESTRUTURAS DE CONTROLE =====
	fmt.Println("\n--- ESTRUTURAS DE CONTROLE ---")
	
	// IF/ELSE
	numero := 42
	if numero > 50 {
		fmt.Println("Número é maior que 50")
	} else if numero == 42 {
		fmt.Println("Número é 42 - resposta para tudo! 🤖")
	} else {
		fmt.Println("Número é menor que 50")
	}
	
	// FOR loop
	fmt.Println("Contando de 1 a 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// FOR range (para slices, arrays, maps)
	fmt.Println("Frutas com índice:")
	for indice, fruta := range frutas {
		fmt.Printf("%d: %s\n", indice, fruta)
	}
	
	// SWITCH
	dia := "segunda"
	switch dia {
	case "segunda":
		fmt.Println("Boa segunda-feira! 💪")
	case "sexta":
		fmt.Println("Sextou! 🎉")
	default:
		fmt.Println("Bom dia!")
	}
	
	// ===== FUNÇÕES =====
	fmt.Println("\n--- FUNÇÕES ---")
	
	resultado := soma(10, 20)
	fmt.Printf("Soma de 10 + 20 = %d\n", resultado)
	
	// Função com múltiplos retornos
	quociente, resto := divisao(17, 5)
	fmt.Printf("17 ÷ 5 = %d com resto %d\n", quociente, resto)
	
	// Função com retorno nomeado
	area, perimetro := retangulo(5, 3)
	fmt.Printf("Retângulo 5x3 - Área: %d, Perímetro: %d\n", area, perimetro)
	
	// ===== PONTEIROS =====
	fmt.Println("\n--- PONTEIROS ---")
	
	valor := 100
	ponteiro := &valor // & obtém o endereço
	fmt.Printf("Valor: %d, Endereço: %p\n", valor, ponteiro)
	
	*ponteiro = 200 // * desreferencia o ponteiro
	fmt.Printf("Novo valor: %d\n", valor)
	
	// ===== GOROUTINES (CONCORRÊNCIA) =====
	fmt.Println("\n--- GOROUTINES ---")
	
	fmt.Println("Iniciando goroutines...")
	go funcaoAssincrona("Goroutine 1")
	go funcaoAssincrona("Goroutine 2")
	
	// Aguarda um pouco para as goroutines executarem
	time.Sleep(2 * time.Second)
	
	// ===== INTERFACES =====
	fmt.Println("\n--- INTERFACES ---")
	
	var forma Forma = Retangulo{largura: 4, altura: 6}
	fmt.Printf("Área da forma: %d\n", forma.Area())
	
	// ===== TRATAMENTO DE ERROS =====
	fmt.Println("\n--- TRATAMENTO DE ERROS ---")
	
	resultado, err := divisaoSegura(10, 0)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Resultado: %d\n", resultado)
	}
	
	// ===== EXEMPLO PRÁTICO =====
	fmt.Println("\n--- EXEMPLO PRÁTICO ---")
	
	// Calcula fatorial
	fatorial5 := fatorial(5)
	fmt.Printf("Fatorial de 5: %d\n", fatorial5)
	
	// Fibonacci
	fmt.Println("Sequência Fibonacci (primeiros 10 números):")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
	
	fmt.Println("\n🎯 Parabéns! Você aprendeu o básico do Go!")
}

// ===== FUNÇÕES AUXILIARES =====

// Função simples com um retorno
func soma(a, b int) int {
	return a + b
}

// Função com múltiplos retornos
func divisao(a, b int) (int, int) {
	return a / b, a % b
}

// Função com retornos nomeados
func retangulo(largura, altura int) (area, perimetro int) {
	area = largura * altura
	perimetro = 2 * (largura + altura)
	return // retorna os valores nomeados
}

// Função assíncrona (goroutine)
func funcaoAssincrona(nome string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: contando %d\n", nome, i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Interface
type Forma interface {
	Area() int
}

// Struct
type Retangulo struct {
	largura, altura int
}

// Método para Retangulo
func (r Retangulo) Area() int {
	return r.largura * r.altura
}

// Função com tratamento de erro
func divisaoSegura(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil
}

// Função recursiva (fatorial)
func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

// Função recursiva (fibonacci)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}