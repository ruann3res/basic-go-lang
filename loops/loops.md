# Loops

- Loops são estruturas de repetição que permitem executar um bloco de código várias vezes.

## Loop: For
- Estrutura do for:
```go

    for i :=0 (Init statement); i < 10 (Condition); i++ (Post statement) {
        // Bloco de código
    }

// init statement: Executado antes da primeira iteração do loop

// condition: É avaliado antes de cada iteração do loop

// post statement: Executado no final de cada iteração do loop
	
  ```
------------

- Todos os parametros do For são opcionais, por exemplo:
```go
var resultFor1 int
var resultFor2 int
var resultFor3 int


for ; i < 10 (Condition); i++ (Post statement) {
    resultFor1++
}
fmt.Println(resultFor1) // 10

for i < 10 (Condition) {
    resultFor2++
}
fmt.Println(resultFor2) // 10

for {
    resultFor3++
}
fmt.Println(resultFor3) // Infinito 
// (Semelhante ao While True em outras linguagens)
	

 ```
------------
- Iterando por um array:
```go

func main(){
    arr := [3]int{1, 2, 3}

    for range arr {
        fmt.Println("Iterando") // 
    }
}
```
------------
- E possivel ver o indice e o valor do array:
```go
func main(){
    arr := [3]int{1, 2, 3}

    for  index, element  := range arr {
        fmt.Println(index, element) // 0 1, 1 2, 2 3
    }
}
```

PS: Os loops em range e possível declarar nenhuma variavel, 1 variavel ou 2 variaveis 

------------
- É possivel ver apenas o valor com o Blank Identifier `_`:

```go

func main(){
    arr := [3]int{1, 2, 3}

    for _, element  := range arr {
        fmt.Println(element) // 1, 2, 3
    }
}
```

### Mudanças no go 1.22^

- O time de GO percebeu que a declaração do loop era muito repetitiva e deciram simplificar a declaraçao:
 

```go

func main(){

    for range 10 {
        fmt.Println("Passei aqui") // Passei aqui 10x
    }
}
```
------------

- A declaraçao do for range sobre um inteiro e possivel apenas declarar uma variavel.

```go
    func main() {

	for i, v := range 10 {
		fmt.Println(i, v) // range over 10 (untyped int constant) permits only one iteration variable
	}
	
	for i := range 10 {
        fmt.Println(i) // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	}
    }
```
------------
- Ponteiros no for:

Nas versoes anteriores de GO os ponteiros dentro de for tinham sempre o mesmo endereco de memoria, porem a partir da versao 1.22^ isso alterado:

```go 
// Go 1.21.7
func main() {
    arr := [3]int{1,2,3}
    for index,element := range arr{
        fmt.Println(&index, &element) // 0xc000012148 0xc000012140 // 0xc000012148 0xc000012140 // 0xc000012148 0xc000012140
    }
}
```

```go 
// Go 1.22^
func main() {
    arr := [3]int{1,2,3}
    for index,element := range arr{
        fmt.Println(&index, &element) // 0xc000012148 0xc000012140 // 0xc000012168 0xc000012160 // 0xc000012178 0xc000012170
    }
}
```

------------

###### OBS: Essas entre outras mudanças ocorreram na versão 1.22^ do GO, para mais informações acesse o [link](https://tip.golang.org/doc/go1.22)