package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mainTutorial() {
	fmt.Println("🚀 TUTORIAL COMPLETO DE GOROUTINES EM GO")
	fmt.Println("==========================================\n")

	// 1. CONCEITOS BÁSICOS
	conceitosBasicos()

	// 2. CHANNELS FUNDAMENTAIS
	channelsFundamentais()

	// 3. BUFFERED CHANNELS
	bufferedChannels()

	// 4. SELECT STATEMENT
	selectStatement()

	// 5. WORKER POOL
	workerPool()

	// 6. SINCRONIZAÇÃO COM WAITGROUP
	sincronizacaoWaitGroup()

	// 7. PROTEÇÃO COM MUTEX
	protecaoMutex()

	// 8. CONTEXT PARA CANCELAMENTO
	contextCancelamento()

	// 9. EXEMPLOS PRÁTICOS
	exemplosPraticos()

	fmt.Println("\n🎯 Parabéns! Você dominou as goroutines em Go!")
}

// ===== 1. CONCEITOS BÁSICOS =====
func conceitosBasicos() {
	fmt.Println("📚 1. CONCEITOS BÁSICOS DE GOROUTINES")
	fmt.Println("----------------------------------------")

	fmt.Println("• Goroutine é uma função que executa de forma concorrente")
	fmt.Println("• São leves (2KB de stack inicial)")
	fmt.Println("• Executam independentemente da função principal")
	fmt.Println("• Sintaxe: go nomeFuncao()")

	// Exemplo básico
	fmt.Println("\n🔍 Exemplo: Goroutines básicas")

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("  Goroutine 1: contando %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("  Goroutine 2: contando %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Aguarda as goroutines terminarem
	time.Sleep(1 * time.Second)
	fmt.Println("  ✅ Goroutines básicas concluídas!\n")
}

// ===== 2. CHANNELS FUNDAMENTAIS =====
func channelsFundamentais() {
	fmt.Println("📡 2. CHANNELS FUNDAMENTAIS")
	fmt.Println("----------------------------")

	fmt.Println("• Channels permitem comunicação entre goroutines")
	fmt.Println("• Operações são bloqueantes por padrão")
	fmt.Println("• Sintaxe: make(chan tipo)")
	fmt.Println("• Envio: canal <- valor")
	fmt.Println("• Recebimento: valor := <-canal")

	// Exemplo: Comunicação simples
	fmt.Println("\n🔍 Exemplo: Comunicação entre goroutines")

	canal := make(chan string)

	// Goroutine que envia mensagem
	go func() {
		fmt.Println("  📤 Enviando mensagem...")
		time.Sleep(500 * time.Millisecond)
		canal <- "Olá do canal!"
		fmt.Println("  ✅ Mensagem enviada!")
	}()

	// Goroutine principal recebe
	fmt.Println("  📥 Aguardando mensagem...")
	mensagem := <-canal
	fmt.Printf("  📨 Mensagem recebida: %s\n", mensagem)

	// Exemplo: Canal bidirecional
	fmt.Println("\n🔍 Exemplo: Canal bidirecional")

	canalBidirecional := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			canalBidirecional <- i
			fmt.Printf("  📤 Enviado: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(canalBidirecional) // Importante fechar o canal
	}()

	// Lê todos os valores
	for valor := range canalBidirecional {
		fmt.Printf("  📥 Recebido: %d\n", valor)
	}

	fmt.Println("  ✅ Channels fundamentais concluídos!\n")
}

// ===== 3. BUFFERED CHANNELS =====
func bufferedChannels() {
	fmt.Println("📦 3. BUFFERED CHANNELS")
	fmt.Println("------------------------")

	fmt.Println("• Podem armazenar múltiplos valores sem bloqueio")
	fmt.Println("• Sintaxe: make(chan tipo, tamanho)")
	fmt.Println("• Útil para controlar fluxo de dados")
	fmt.Println("• Evita deadlocks em certas situações")

	// Exemplo: Canal com buffer
	fmt.Println("\n🔍 Exemplo: Canal com buffer de 3")

	bufferedCanal := make(chan int, 3)

	// Envia múltiplos valores sem bloqueio
	go func() {
		for i := 1; i <= 5; i++ {
			bufferedCanal <- i
			fmt.Printf("  📤 Enviado: %d (buffer: %d/%d)\n", i, len(bufferedCanal), cap(bufferedCanal))
			time.Sleep(100 * time.Millisecond)
		}
		close(bufferedCanal)
	}()

	// Processa os valores
	time.Sleep(600 * time.Millisecond) // Aguarda alguns envios

	for valor := range bufferedCanal {
		fmt.Printf("  📥 Processando: %d\n", valor)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("  ✅ Buffered channels concluídos!\n")
}

// ===== 4. SELECT STATEMENT =====
func selectStatement() {
	fmt.Println("🎯 4. SELECT STATEMENT")
	fmt.Println("----------------------")

	fmt.Println("• Escolhe entre múltiplos canais não-bloqueantes")
	fmt.Println("• Útil para timeout e operações múltiplas")
	fmt.Println("• Executa o primeiro caso que estiver pronto")
	fmt.Println("• time.After() para implementar timeout")

	// Exemplo: Select com múltiplos canais
	fmt.Println("\n🔍 Exemplo: Select com timeout")

	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine 1: resposta lenta
	go func() {
		time.Sleep(2 * time.Second)
		canal1 <- "Resposta lenta"
	}()

	// Goroutine 2: resposta rápida
	go func() {
		time.Sleep(500 * time.Millisecond)
		canal2 <- "Resposta rápida"
	}()

	// Select escolhe o que responder primeiro
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Printf("  📨 Canal 1: %s\n", msg1)
		case msg2 := <-canal2:
			fmt.Printf("  📨 Canal 2: %s\n", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("  ⏰ Timeout!")
		}
	}

	fmt.Println("  ✅ Select statement concluído!\n")
}

// ===== 5. WORKER POOL =====
func workerPool() {
	fmt.Println("👷 5. WORKER POOL")
	fmt.Println("------------------")

	fmt.Println("• Múltiplas goroutines processando jobs")
	fmt.Println("• Controle de concorrência")
	fmt.Println("• Balanceamento de carga")
	fmt.Println("• Padrão muito usado em produção")

	// Exemplo: Worker pool
	fmt.Println("\n🔍 Exemplo: 3 workers processando 6 jobs")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Inicia 3 workers
	for w := 1; w <= 3; w++ {
		go workerTutorial(w, jobs, results)
	}

	// Envia 6 jobs
	for j := 1; j <= 6; j++ {
		jobs <- j
		fmt.Printf("  📋 Job %d enviado para fila\n", j)
	}
	close(jobs)

	// Coleta resultados
	fmt.Println("  📥 Coletando resultados...")
	for a := 1; a <= 6; a++ {
		resultado := <-results
		fmt.Printf("  ✅ Job %d processado: resultado = %d\n", a, resultado)
	}

	fmt.Println("  ✅ Worker pool concluído!\n")
}

// Função worker para o pool
func workerTutorial(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("    👷 Worker %d processando job %d\n", id, job)
		// Simula trabalho variável
		tempo := time.Duration(rand.Intn(500)+200) * time.Millisecond
		time.Sleep(tempo)
		results <- job * job // Retorna o quadrado do job
	}
}

// ===== 6. SINCRONIZAÇÃO COM WAITGROUP =====
func sincronizacaoWaitGroup() {
	fmt.Println("⏳ 6. SINCRONIZAÇÃO COM WAITGROUP")
	fmt.Println("-----------------------------------")

	fmt.Println("• Sincroniza múltiplas goroutines")
	fmt.Println("• Métodos: Add(), Done(), Wait()")
	fmt.Println("• Evita usar time.Sleep() para sincronização")
	fmt.Println("• Conta quantas goroutines estão ativas")

	// Exemplo: WaitGroup
	fmt.Println("\n🔍 Exemplo: 5 goroutines com WaitGroup")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Adiciona uma goroutine ao grupo

		go func(id int) {
			defer wg.Done() // Marca como concluída ao final

			tempo := time.Duration(rand.Intn(1000)+500) * time.Millisecond
			fmt.Printf("  🚀 Goroutine %d iniciada (tempo: %v)\n", id, tempo)
			time.Sleep(tempo)
			fmt.Printf("  ✅ Goroutine %d concluída!\n", id)
		}(i)
	}

	fmt.Println("  ⏳ Aguardando todas as goroutines...")
	wg.Wait() // Espera todas terminarem
	fmt.Println("  🎉 Todas as goroutines concluídas!")

	fmt.Println("  ✅ WaitGroup concluído!\n")
}

// ===== 7. PROTEÇÃO COM MUTEX =====
func protecaoMutex() {
	fmt.Println("🔒 7. PROTEÇÃO COM MUTEX")
	fmt.Println("--------------------------")

	fmt.Println("• Protege dados compartilhados entre goroutines")
	fmt.Println("• Evita race conditions")
	fmt.Println("• Métodos: Lock(), Unlock()")
	fmt.Println("• Útil para contadores e estruturas compartilhadas")

	// Exemplo: Contador compartilhado
	fmt.Println("\n🔍 Exemplo: Contador compartilhado com 10 goroutines")

	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	// Inicia 10 goroutines que incrementam o contador
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simula trabalho antes de acessar o contador
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			mutex.Lock() // Bloqueia acesso
			counter++
			fmt.Printf("    🔒 Goroutine %d: contador = %d\n", id, counter)
			mutex.Unlock() // Libera acesso
		}(i)
	}

	wg.Wait()
	fmt.Printf("  🎯 Valor final do contador: %d\n", counter)

	fmt.Println("  ✅ Proteção com Mutex concluída!\n")
}

// ===== 8. CONTEXT PARA CANCELAMENTO =====
func contextCancelamento() {
	fmt.Println("⏹️  8. CONTEXT PARA CANCELAMENTO")
	fmt.Println("----------------------------------")

	fmt.Println("• Cancela operações longas")
	fmt.Println("• Passa valores entre goroutines")
	fmt.Println("• Tipos: WithTimeout(), WithCancel(), WithValue()")
	fmt.Println("• Ideal para APIs e operações com timeout")

	// Exemplo: Context com timeout
	fmt.Println("\n🔍 Exemplo: Context com timeout de 2 segundos")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Canal para receber resultado
	resultado := make(chan string)

	go func() {
		// Simula trabalho que pode demorar
		select {
		case <-time.After(5 * time.Second):
			resultado <- "Trabalho concluído com sucesso!"
		case <-ctx.Done():
			resultado <- "Trabalho cancelado por timeout!"
		}
	}()

	// Aguarda resultado ou cancelamento
	select {
	case res := <-resultado:
		fmt.Printf("  📨 Resultado: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("  ⏹️  Contexto cancelado: %v\n", ctx.Err())
	}

	// Exemplo: Context com cancelamento manual
	fmt.Println("\n🔍 Exemplo: Context com cancelamento manual")

	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("  🚫 Cancelando contexto...")
		cancel2() // Cancela manualmente
	}()

	<-ctx2.Done()
	fmt.Printf("  ✅ Contexto cancelado: %v\n", ctx2.Err())

	fmt.Println("  ✅ Context para cancelamento concluído!\n")
}

// ===== 9. EXEMPLOS PRÁTICOS =====
func exemplosPraticos() {
	fmt.Println("💼 9. EXEMPLOS PRÁTICOS")
	fmt.Println("-------------------------")

	fmt.Println("• Simulação de cenários reais")
	fmt.Println("• Web scraping concorrente")
	fmt.Println("• Processamento de arquivos")
	fmt.Println("• API calls paralelas")

	// Exemplo 1: Web scraping simulado
	fmt.Println("\n🔍 Exemplo 1: Web scraping concorrente")

	urls := []string{
		"https://site1.com",
		"https://site2.com",
		"https://site3.com",
		"https://site4.com",
		"https://site5.com",
	}

	var wg sync.WaitGroup
	resultados := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			// Simula download de página
			tempo := time.Duration(rand.Intn(1000)+500) * time.Millisecond
			time.Sleep(tempo)

			resultados <- fmt.Sprintf("📄 %s baixado em %v", url, tempo)
		}(url)
	}

	// Fecha o canal quando todas terminarem
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Coleta resultados
	for resultado := range resultados {
		fmt.Printf("  %s\n", resultado)
	}

	// Exemplo 2: Processamento de arquivos
	fmt.Println("\n🔍 Exemplo 2: Processamento de arquivos paralelo")

	arquivos := []string{"arquivo1.txt", "arquivo2.txt", "arquivo3.txt"}

	var wg2 sync.WaitGroup
	processados := make(chan string, len(arquivos))

	for _, arquivo := range arquivos {
		wg2.Add(1)
		go func(nome string) {
			defer wg2.Done()

			// Simula processamento
			tempo := time.Duration(rand.Intn(800)+300) * time.Millisecond
			time.Sleep(tempo)

			processados <- fmt.Sprintf("📁 %s processado em %v", nome, tempo)
		}(arquivo)
	}

	go func() {
		wg2.Wait()
		close(processados)
	}()

	for processado := range processados {
		fmt.Printf("  %s\n", processado)
	}

	fmt.Println("  ✅ Exemplos práticos concluídos!\n")
}

// Função main para executar o tutorial
func main() {
	mainTutorial()
}
