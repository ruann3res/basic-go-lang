# Defer

- Adia uma função para ser executada após o retorno da função que a chamou.
- O defer é executado em ordem reversa, ou seja, o primeiro defer a ser chamado e o ultimo a ser executado.

- Estrutura do defer:

```go 

func main() {
    x := doDefer()
    fmt.Println(x)
}

func doDefer() int {
    defer fmt.Println("world")
    fmt.Println("Hello")
    return 10
}
```
PS: Ordem das saidas: Hello world 10

O defer e executado após o retorno da função que o chamou, ou seja, o defer e executado após o `return 10`, e antes do `fmt.Println(x)`.

------------

- LIFO (Last In First Out):

```go 

func main() {
    doDefer() // 1 2 3
}

func doDefer() int {
    defer fmt.Println(3)
	defer fmt.Println(2)
    fmt.Println(1)
}
```

Note que o defer e executado em ordem reversa, ou seja, o defer que foi chamado por ultimo e o primeiro a ser executado.

------------

- O escopo do defer e o mesmo do escopo da função que o chamou:

```go
func main () {
    doDefer()
}

var arquivos []string{"foo.txt", "bar.txt"}

func doDefer() {
    for _, arquivo := range arquivos {
        file, err := os.Open(arquivo)
        if err != nil { }
        defer file.Close()
    }
    // Note que isso ficaria inviável caso eu tenha varios arquivos mesmo que eu tenha defer pois ele esta atrelado ao escopo da função.
    // Uma soluçao para isso seria criar uma funçao anonima e deixar o defer dentro desse escopo.

    for _, arquivo := range arquivos {
        func (){
            file, err := os.Open(arquivo)
            if err != nil { }
            defer file.Close()
        }()
    }
}
```
