package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// declaração de variáveis
	tentativas := 15

	// leitura de arquivo
	frutas := leituraArquivo()

	// sorteando a fruta
	fruta := sorteandoFruta(frutas)

	// mostrando quantidade de letras da palavra
	quantLetras := monteArry(fruta)

	// iniciando o jogo
	intro()
	fmt.Println(quantLetras)

	// looping do jogo
	loopGame(fruta, quantLetras, tentativas)

}

func loopGame(fruta []string, quantLetras []string, tentativas int) {
	for {

		ttv := fmt.Sprintf("Você tem %d tentativa(s)", tentativas)
		fmt.Println(ttv)
		// leitura do chute
		chute := entrada()

		// verificando se o chute é igual a alguma letra da palavra
		quantLetras = verificandoLetra(fruta, chute, quantLetras)

		// imprimindo a quantidade de letras
		fmt.Println(quantLetras)
		fmt.Println("")

		//verificando se guanhou
		tentativas = winOrLose(quantLetras, tentativas, fruta)

	}
}

func winOrLose(quantLetras []string, tentativas int, fruta []string) int {

	tentativas--

	if strings.Join(quantLetras, "") == strings.Join(fruta, "") {
		fmt.Println("Parabéns, você acertou!")
		os.Exit(0)
	}

	if tentativas == 0 {
		fmt.Println("Você perdeu!")
		os.Exit(0)
	}

	return tentativas
}

func verificandoLetra(fruta []string, chute string, quantLetras []string) []string {

	for i := 0; i < len(fruta); i++ {
		if chute == fruta[i] {
			quantLetras[i] = chute
		}
	}

	return quantLetras
}

func entrada() string {
	var chute string

	fmt.Print("Digite uma letra: ")
	fmt.Scan(&chute)
	chute = strings.ToLower(chute)

	return chute
}

func monteArry(fruta []string) []string {

	var quantLetras []string

	for i := 0; i < len(fruta); i++ {
		quantLetras = append(quantLetras, "_")
	}

	return quantLetras
}

func sorteandoFruta(frutas []string) []string {

	// geração de número aleatório
	rand.Seed(time.Now().UnixNano())
	numAl := rand.Intn(len(frutas)-0) + 0

	fruta := strings.Split(frutas[numAl], "")

	return fruta
}

func leituraArquivo() []string {

	var frutas []string

	arquivo, err := os.Open("frutas.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
		os.Exit(0)
		return []string{""}
	}

	leitor := bufio.NewReader(arquivo)

	for {
		fruta, err := leitor.ReadString('\n')
		fruta = strings.TrimSpace(fruta)

		if err == io.EOF {
			break
		}

		frutas = append(frutas, fruta)
	}

	arquivo.Close()

	return frutas
}

func intro() {
	fmt.Println("Olá, bem vindo ao Adivinhe a Fruta")
	fmt.Println("")
}
