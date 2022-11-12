package main

import (
	"log"
	"os"

	"github.com/hedhyw/rex/internal/generator"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("wrong amount of arguments")
	}

	if len(os.Args[1]) == 0 {
		log.Fatalln("given regex is empty")
	}

	regex := os.Args[1]

	result, err := generator.GenerateCode(regex)
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.WriteString(result + "\n")
}
