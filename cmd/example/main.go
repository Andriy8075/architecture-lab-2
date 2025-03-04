package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	exprFlag = flag.String("e", "", "Вираз у командному рядку")
	fileFlag = flag.String("f", "", "Файл з виразом")
	outFlag  = flag.String("o", "", "Файл для збереження результату (необов’язковий)")
)

func processFileInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("помилка відкриття файлу: %v", err)
	}
	return strings.TrimSpace(string(data)), nil
}

func main() {
	flag.Parse()

	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Помилка: не можна використовувати одночасно -e і -f")
		os.Exit(1)
	}

	var inputSource io.Reader
	var err error

	if *exprFlag != "" {
		inputSource = strings.NewReader(*exprFlag)
	} else if *fileFlag != "" {
		content, err := processFileInput(*fileFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		inputSource = strings.NewReader(content)
	} else {
		fmt.Fprintln(os.Stderr, "Помилка: потрібно передати -e або -f")
		os.Exit(1)
	}

	var outputDestination io.Writer = os.Stdout
	if *outFlag != "" {
		outputDestination, err = os.Create(*outFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка створення файлу: %v\n", err)
			os.Exit(1)
		}
		defer outputDestination.(*os.File).Close()
	}

	handler := lab2.ComputeHandler{
		Input:  inputSource,
		Output: outputDestination,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
