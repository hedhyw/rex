package main

import (
	"log"
	"os"

	"github.com/hedhyw/rex/internal/generator"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatalln("wrong amount of arguments")
	}

	if len(args[1]) == 0 {
		log.Fatalln("given regex is empty")
	}

	regex := args[1]

	result, err := generator.GenerateCode(regex)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(result + "\n")
}
