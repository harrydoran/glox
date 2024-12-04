package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(source string) {
	fmt.Println(source)
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	run(string(bytes))
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("> ")
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		run(line)
		fmt.Printf("> ")
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v: %v", line, where, message)
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Hello, World!")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[0])
	} else {
		runPrompt()
	}
}
