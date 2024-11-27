# Switchs

- Switchs sao usados para tomar decisoes baseadas em condiçoes. (Utilizado para simplificar if-else)


- Estrutura Switch:

```go 
func main(){
    var myCondition int = 1

	switchCase(myCondition)// Case 1
}

func switchCase(condition int){
    switch condition {
    case 1:
        fmt.Println("Case 1")
    case 2:
        fmt.Println("Case 2")
    default:
        fmt.Println("Default")
    }
}
```

------------
- Não e necessario usar a palavra chave `break` para finalizar o switch, pois o switch termina quando um caso e executado.

- Para execultar o proximo caso apos o caso selecionado, e necessario usar a palavra chave `fallthrough`:

```go
func main(){
    var myCondition int = 1
    
    switchCase(myCondition)//  Case 1  Case 2
}

func switchCase(condition int){
	
    switch condition {
        case 1:
            fmt.Println("Case 1")
            fallthrough
        case 2:
            fmt.Println("Case 2")
        default:
            fmt.Println("Default")
    }
}
```

------------

- E possivel fazer Switch sem condicao, e o switch ira executar o primeiro caso que for verdadeiro:

```go

func main(){
    var myCondition int = 1
    
    switchCase(myCondition)// Case 1
}

func switchCase(condition int){
    switch {
        case condition == 1:
            fmt.Println("Case 1")
        case condition == 2:
            fmt.Println("Case 2")
        default:
            fmt.Println("Default")
    }
}
```

------------

- E possivel declarar variaveis locais no switch:

```go

func main(){
    switch x := math.Sqrt(4); x {
        case 2:
            fmt.Println("Resultado é 2")
        default:
            fmt.Println("Algo esta errado")
    }
}
```

------------

- É possivel ter varios casos em um unico case:

```go

func main(){
    fmt.Println(isWeekend(time.Now())) // true
}

func isWeekend(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}
```

------------

- Switch Statement com variaveis de tipo:

```go

func main(){
    compareType(10)      // 10
    compareType("Hello") // Hello
}

func compareType(x any) {
switch t := x.(type) {
    case int:
        takeInt(t)
    case string:
        takeString(t)
    default:
        fmt.Println("Default")
    }
}

func takeString(s string) {
    fmt.Println(s)
}

func takeInt(x int) {
    fmt.Println(x)
}

```