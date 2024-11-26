# Variables

- Declaração de variáveis:

```go 
var nomeDaVariavel tipoDaVariavel
```

------------

- E possivel declarar mais de uma variavel na mesma linha:

```go 
var nomeDaVariavel1, nomeDaVariavel2 tipoDaVariavel
```

------------

- E possivel inicializar a variavel sem o tipo pois o compilador consegue inferir o tipo da variavel:

```go 
var nomeVariavel1, nomeVariavel2 = "valor1", "valor2"

// or 

var (
nomeVariavel1 = "valor1"
nomeVariavel2  = "valor2"
)

```

------------

- E possivel declarar e inicializar variaveis sem a palavra reservada `var`:

```go
nomeVariavel1, nomeVariavel2 := "valor1", "valor2"
```

ps: A forma de declarar e inicializar variaveis sem a palavra reservada `var` so pode ser utilizada dentro de funçoes.

```go 

package main

import "fmt"

var idadePessoa int 
nomePessoa := "Lucas" // Erro de compilação

func main() {
	var sobrenome string
	nome, idade := "Lucas", 20
	
	fmt.Println(nome, idade, sobrenome)
}

```

------------

- Declarar variavel de tipo diferente na mesma linha e possivel apenas quando a variavel e inicializada:

```go
package main

import "fmt"

func main() {
    var baz, qux int,string 
	
	fmt.Println(baz, qux) // Erro de compilação
}
``` 
- Declaraçao correta 
```go 

package main

import "fmt"

func main() {
	foo, bar := "Foo", 2
	
	fmt.Println(foo, bar) // Foo 2
}

```





