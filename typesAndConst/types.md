# Types GoLang & Constantes

## Types
- Como Go é uma linguagem fortemente tipada, é necessário definir o tipo assim que a variável é criada, seja de forma implícita ou explícita.
- Tipos basicos:

Bool
```go 
// Bool (Valores: verdadeiros ou falsos)

    func main() {
        var isAdmin bool = true
        var isUser bool
        fmt.Println(isAdmin) // true 
        fmt.Println(isAdmin) // false
    }
```

------------
String
```go 
// String (Sequência de caracteres)

func main() {
    var sobrenome string
    var nome string = "Ruan"
    
    fmt.Println(nome) // Ruan
    fmt.Println(sobrenome) // ""
}
```

------------

Valores inteiros (Positivos e Negativos):


```go
// int, int8 int16 int32 int64 

// Note que exceto o `int` todos possuem o seu tamanho em memória previamente definido.

func main() {
    var idade int = 20
    var idade8 int8 = 127
    var idade16 int16 = 32767
    var idade32 int32 = 2147483647
    var idade64 int64 = 9223372036854775807

    fmt.Println(idade) // 20
    fmt.Println(idade8) // 127
    fmt.Println(idade16) // 32767
    fmt.Println(idade32) // 2147483647
    fmt.Println(idade64) // 9223372036854775807
}

```
------------

Valores Inteiros (Apenas positivos)


```go
// uint, uint8, uint16, uint32, uint64, uintptr

// Note que exceto o `uint`, todos possuem o seu tamanho em memória previamente definido. 
// uintptr é usado para armazenar o endereço de memória de um ponteiro.

func main() {
    var idade uint = 20
    var idade8 uint8 = 255
    var idade16 uint16 = 65535
    var idade32 uint32 = 4294967295
    var idade64 uint64 = 18446744073709551615
    var endereco uintptr = 0x0000000000000000

    fmt.Println(idade) // 20
    fmt.Println(idade8) // 255
    fmt.Println(idade16) // 65535
    fmt.Println(idade32) // 4294967295
    fmt.Println(idade64) // 18446744073709551615
    fmt.Println(endereco) // 0
}

```
------------

Byte e Rune

```go

// byte e rune são tipos de inteiros sem sinal, onde byte é um alias para uint8 e rune é um alias para int32.

func main() {
    var r int32 = 32
    var b uint8 = 8
	
    takeByte(b)
    takeRune(r)
	
    fmt.Println(b) // 8
    fmt.Println(r) // 32
}

func takeByte(b byte) {
    fmt.Println(b)
}

func takeRune(r rune) {
    fmt.Println(r)
}

```
------------

Valores decimais

```go

// float32 e float64

func main() {
    var preco float32 = 42.99
    var preco64 float64 = 42.99

    fmt.Println(preco) // 42.99
    fmt.Println(preco64) // 42.99
}

```

------------

Complexos

```go
// complex64 e complex128

func main() {
    var complexo complex64 = 1 + 2i
    var complexo128 complex128 = 1 + 2i

    fmt.Println(complexo) // (1+2i)
    fmt.Println(complexo128) // (1+2i)
}

```
------------


- Coversao de variaveis

Em alguns desses tipos "Basicos" e possivel realizar a conversao de um para outro, por exemplo:

```go
func main() {
    var x int = 64
	
    f := float64(x)
	
    fmt.Println(f) // 64 type float64
	
}
```

Porem nao sao todos que podem ser convertidos, por exemplo:

```go
func main() {
    var x int = 64
	
    b := bool(x)
	
    fmt.Println(b) // Erro: Cannot convert an expression of the type 'int' to the type 'bool
}

```

PS: Nao numa conversão direta.

------------
- Conversao para Strings

A conversao de strings e diferente pois ao mudar um valor inteiro para string, o valor e convertido para o valor ASCII correspondente.

```go

func main() {
    var x int = 64
    
    s := string(x)
    
    fmt.Println(s) // @
}

```

Para converter um valor inteiro para string, e necessario usar lib `strconv`:

```go

import (
    "fmt"
    "strconv"
)

func main() {
    var x int = 64
    
    s := strconv.FormatInt(int64(x),10)
    
    fmt.Println(s) // 64
}

```
------------


## Constantes

- Constantes sao valores que nao podem ser alterados durante a execucao do programa.
- Para declarar uma constante é utilizado o mesmo padrao que uma variavel, porem com a palavra reservada `const`:

```go

func main() {
    const idade int = 20 // Diferente de variáveis isso nao dara um erro de compilação apenas de lint, pois é possível definir constantes sem utilizá las. 
    var nome string = "Ruan" // Erro de compilaçao pois a variavel nao foi utilizada.
}
```
------------
- Nem todos os tipos podem ser constantes em go apenas: String, Bytes, Runas, Booleanos e Valores numericos.

- Não e possivel utilizar short declaration para declarar constantes.

- E possivel omitir o tipo da constante por exemplo:

Sem omissão de tipo:
```go

func main() {
	const x int = 10
	takeInt32(x) //  cannot use x (constant 10 of type int) as int32 value in argument to takeInt32
}

func takeInt32(x int32) {
	fmt.Println(x)
}
```

Omitindo o tipo:

```go
func main() {
	const x = 10
	takeInt32(x) // 10
}

func takeInt32(x int32) {
	fmt.Println(x)
}
```
    
Nome técnico: Untyped constant

PS: Constante e uma maneira de `Enganar` o sistema de types em go, pois o compilador ira inferir o tipo da constante baseado no seu uso. 

------------

- Constantes podem ser declaradas em bloco:

```go

func main() {
    const (
        idade int    = 20
        nome  string = "Ruan"
    )

    fmt.Println(idade, nome)
}
```

------------

- Constante Literal

```go
func main() {
    const x = 10
    takeInt32(10)// 10
    takeInt32(x) // 10
}

func takeInt32(x int32) {
    fmt.Println(x)
}
```


