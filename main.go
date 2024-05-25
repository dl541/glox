package main

import (
	"bufio"
	"fmt"
	"lox/scanner"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	run(string(content))

	return nil
}

func runPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		if err := run(scanner.Text()); err != nil {
			return nil
		}
	}
	return nil
}

func run(source string) error {
	fmt.Printf("run line %v\n", source)
	scanner := scanner.NewScanner(source)

	for _, token := range scanner.ScanTokens() {
		fmt.Printf("parsing token %+v\n", token)
	}
	return nil
}
