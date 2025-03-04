package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
	"io"
	"os"
	"strings"
)

var (
	exprFlag = flag.String("e", "", "Вираз у командному рядку")
	fileFlag = flag.String("f", "", "Файл з виразом")
	outFlag  = flag.String("o", "", "Файл для збереження результату (необов’язковий)")
)

func main() {
	flag.Parse()

	// Перевірка неправильного використання аргументів
	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Помилка: не можна використовувати одночасно -e і -f")
		os.Exit(1)
	}

	var inputSource io.Reader
	var err error

	// Якщо є -e, створюємо reader з цього рядка
	if *exprFlag != "" {
		inputSource = strings.NewReader(*exprFlag)
	} else if *fileFlag != "" {
		inputSource, err = os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка відкриття файлу: %v\n", err)
			os.Exit(1)
		}
		defer inputSource.(*os.File).Close()
	} else {
		fmt.Fprintln(os.Stderr, "Помилка: потрібно передати -e або -f")
		os.Exit(1)
	}

	// Визначаємо місце для запису результату
	var outputDestination io.Writer = os.Stdout
	if *outFlag != "" {
		outputDestination, err = os.Create(*outFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка створення файлу: %v\n", err)
			os.Exit(1)
		}
		defer outputDestination.(*os.File).Close()
	}

	// Створюємо ComputeHandler
	handler := lab2.ComputeHandler{
		Input:  inputSource,
		Output: outputDestination,
	}

	// Виконуємо обчислення
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
