package main

import "fmt"

var idade int

func main() {
	var nome, sobrenome = "Ruan", "Neres"

	var profissaoPessoa = "Desenvolvedor"

	var (
		linguaguensPessoal = "Python"
		cursoPessoa        = "Ciencia Da computaçao"
	)
	// declarando e definindo a variavel
	// Apenas em escopo de funçao em escopo de pacote nao funciona. SHortSintaxe
	statusCivil := "Solteiro"

	// declarar variavel de tipo diferente na mesma linha e possivel apenas quando a variavel e inicializada
	// var baz, qux int,string
	// declaraçao short o type e sempre inferido
	foo, bar := "Foo", 2
	fmt.Println(nome, sobrenome, profissaoPessoa, linguaguensPessoal, cursoPessoa, statusCivil, foo, bar, baz, qux)

}
