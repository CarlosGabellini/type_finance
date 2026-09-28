package ctr

import "fmt"

func Saudacao() {
	
	for i := 0; i < 30; i++ {
		fmt.Printf("=")
	}

	fmt.Printf("\n")
	fmt.Printf("Seja bem vindo ao controle financeiro!\n")

	fmt.Printf("Quais opcoes deseja usar para cadastro?\n")
	fmt.Printf("01 - Cadastro\n")
	fmt.Printf("02 - Escolher mes\n")
	fmt.Printf("03 - Escolher data\n")
}