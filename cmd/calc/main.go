package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"vmrc/pkg/fraction"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func waitKeyPress() {
	fmt.Println("\n\nPressione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadRune()
}

func mainMenu() bool {
	clearScreen()
	fmt.Println("Calculadora")
	fmt.Println("---------------------------------")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("0 - Sair")
	operation := readNumber("a operação desejada")
	switch operation {
	case 1:
		sumMenu()
		return false
	case 2:
		subtractionMenu()
		return false
	case 3:
		multiplicationMenu()
		return false
	case 4:
		divisionMenu()
		return false
	default:
		return true
	}
}

func divisionMenu() {
	clearScreen()
	fmt.Println("Calculadora - Dividir")
	fmt.Println("---------------------------------")
	fmt.Println("Digite os valores da primeira fração:")
	f1 := readFraction()
	fmt.Println("Digite os valores da segunda fração:")
	f2 := readFraction()
	if f1 != nil && f2 != nil {
		result, err := f1.Divide(f2)
		if err != nil {
			fmt.Println("Erro ao dividir as frações:", err)
		}
		fmt.Println("Resultado: ", result.ToString())
	}
	waitKeyPress()
}

func multiplicationMenu() {
	clearScreen()
	fmt.Println("Calculadora - Multiplicar")
	fmt.Println("---------------------------------")
	fmt.Println("Digite os valores da primeira fração:")
	f1 := readFraction()
	fmt.Println("Digite os valores da segunda fração:")
	f2 := readFraction()
	if f1 != nil && f2 != nil {
		result, err := f1.Multiply(f2)
		if err != nil {
			fmt.Println("Erro ao multiplicar as frações:", err)
		}
		fmt.Println("Resultado: ", result.ToString())
	}
	waitKeyPress()
}

func subtractionMenu() {
	clearScreen()
	fmt.Println("Calculadora - Subtracao")
	fmt.Println("---------------------------------")
	fmt.Println("Digite os valores da primeira fração:")
	f1 := readFraction()
	fmt.Println("Digite os valores da segunda fração:")
	f2 := readFraction()
	if f1 != nil && f2 != nil {
		result, err := f1.Subtract(f2)
		if err != nil {
			fmt.Println("Erro ao subtrair as frações:", err)
		}
		fmt.Println("Resultado: ", result.ToString())
	}
	waitKeyPress()
}

func sumMenu() {
	clearScreen()
	fmt.Println("Calculadora - Soma")
	fmt.Println("---------------------------------")
	fmt.Println("Digite os valores da primeira fração:")
	f1 := readFraction()
	fmt.Println("Digite os valores da segunda fração:")
	f2 := readFraction()
	if f1 != nil && f2 != nil {
		result, err := f1.Add(f2)
		if err != nil {
			fmt.Println("Erro ao somar as frações:", err)
		}
		fmt.Println("Resultado: ", result.ToString())
	}
	waitKeyPress()
}

func readFraction() *fraction.Fraction {
	numerator := readNumber("o numerador")
	denominator := readNumber("o denominador")
	f, err := fraction.NewFraction(numerator, denominator)
	if err != nil {
		fmt.Println("Erro ao criar a fração:", err)
		panic("Error")
	}
	return f
}

func readNumber(message string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite " + message + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	valor, _ := strconv.Atoi(text)
	return valor
}

func main() {
	exit := false
	for !exit {
		exit = mainMenu()
	}
}
