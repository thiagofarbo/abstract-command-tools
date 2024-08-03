package main

import (
	"fmt"
	"os"
)

func create2(text []byte) {

	file, err := os.Create("answers.json")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(text)
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
	}
	fmt.Println("File created successfully")
}
