# Maps (HashTables)

- Maps tambem sao uma estrutura de dados que armazena pares chave-valor, mas a diferenca entre eles e que os maps nao possuem ordem. Ou seja, a ordem dos elementos nao e garantida. Em Go, os maps sao implementados usando a estrutura de dados hash table.

- Estrutura do map:

```go
func main() {
	
	map1 := map[typeDaChave]typeDoValor{
        chave1: valor1,
        chave2: valor2,
    }
}

    salarios := map[string]float64{
    "José": 11325.45,
    }
    fmt.Println(salarios)
}
```

-------

- E possivel inicializar um map sem um valor definido:

```go
func main() {
	
	map1 := map[typeDaChave]typeDoValor{
        chave1: valor1,
        chave2: valor2,
    }
}

    map2 := make(map[string]float64)
	map3 := map[string]float64{}
}
```


-------
- Para remover um elemento de um map, usamos a funcao `delete` do proprio Go:

```go

func main() {
	
    salarios := map[string]float64{
    "Maria": 11325.45,
	"João": 11325.45,
    }
    fmt.Println(salarios["Maria"]) // Maria:11325.45]
	
	delete(salarios, "Maria")
	fmt.Println(salarios) // map[João:11325.45]
	
}
```

-------

- Para adicionar um elemento a um map, basta atribuir um valor a uma chave que ainda nao existe:

```go
 
func main() {
    
    salarios := map[string]float64{
    "Maria": 11325.45,
    "João": 11325.45,
    }
    fmt.Println(salarios["Maria"]) 
    
    salarios["Pedro"] = 1200.0
    fmt.Println(salarios)
    
}
```