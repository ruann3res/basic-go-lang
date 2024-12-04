# Slices

- Slices são considerados arrays infitos, porem isso nao e verdade, por baixo dos panos estamos trabalhando com arrays, os slices sao divididos em 3 partes: 

    - Ponteiro: aponta para o primeiro elemento do array
    - Tamanho: e o tamanho do array
    - Capacidade: e o tamanho maximo do array

- Por outro lado os slices sao dinamicos, ou seja, podemos adicionar ou remover elementos do slice, e o tamanho do array sera ajustado automaticamente.

- Estrutura Slice:

```go 
func main(){
    var mySlice []int = []int{1,2,3,4,5}
    fmt.Println(mySlice)
}
```

------------

- Para entendermos melhor vou dar exemplos de codigo e explicarei o comportamento de cada um:

```go

func main(){
    var mySlice []int = []int{1,2,3,4,5}
    fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice), cap(mySlice), mySlice)
}
```
Estamos imprimindo o tamanho do slice, a capacidade do slice e o slice em si.

------------

```go 
func main(){
    var mySlice []int = []int{1,2,3,4,5}
    fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])
}
```
Nesse caso estamos passando um slice [:0] que significa que estamos pegando um slice do inicio ate o indice 0, ou seja, estamos pegando um slice vazio. Porem e possivel notar que a Capacidade mantida `5` no caso. 

------------

```go 
func main(){
    var mySlice []int = []int{1,2,3,4,5}
    fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])
}
```

Nesse caso estamos passando um slice [:4] que significa que estamos pegando um slice do inicio ate o indice 4, ou seja, estamos pegando um slice com 4 elementos. Porem e possivel notar que a Capacidade mantida `5` no caso e o tamanho dele é `4`.

------------

```go 
func main(){
    var mySlice []int = []int{1,2,3,4,5}
    fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])
}
```

Eu ignorei os 2 primeiros valores do slice, entao agora minha capacidade é `3` e o tamanho é `3`.

------------

- Aumentando a capacidade do Slice

```go 
func main(){
    var mySlice []int = []int{1,2,3,4,5}
	
	mySlice = append(mySlice, 6)

    fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice), cap(mySlice), mySlice)
}
```

Voce nao necessariamente vai aumentar a capacidade desse slice, pois ele sempre vai apontar para um array e esse array tem a capacidade fixa.
Nesse caso acima o slice aponta para um array de tamanho 5, porem quando eu adicionei um novo elemento, o slice apontou para um novo array de tamanho 10, o tamanho dele foi para 6 porem a capacidade dele é 10 pois ele dobrou a capacidade do array anterior.


------------