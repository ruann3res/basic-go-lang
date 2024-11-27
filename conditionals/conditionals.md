# Conditionals

- Condiçoes sao usadas para tomar decisoes baseadas em condiçoes.

- Estrutura da condicionais:

```go
func main(){
    const condition = true

    if condition {
        // code
    } else if condition {
        // code
    } else {
        // code
    }
}
```

------------

- E possivel declarar variaveis dentro do bloco da condicional:

```go
package main

import "fmt"

func main() {
	if x := 10; x > 5 {
            fmt.Println("x e maior que 5")
	}
}
```

- Diferente do Loop a condicional e sempre obrigatória. 

```go
package main

import "fmt"

func main() {
	if x := 10; { // erro de compilação, falta a condição
            fmt.Println(x)
	}

	var resultFor1 int
	for {
		resultFor1++
		fmt.Print(resultFor1)
	}
}
```
