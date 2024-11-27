# Arrays

- Estrutura de dados que existe em go, que representa um espaço continuo.
- Armazena varios elementos do mesmo tipo.

- Estrutura do array:
```go 
    func main(){
        arr := [3]int{1, 2, 3}
  }
  ```
------------

- Para criar um array e necessario saber o tamanho do array e o tipo de dado que ele vai armazenar.

- Nao e necessario inicializar os elementos do array.

- E possivel definir o valor de um indice especifico do array:

```go 
    func main(){
        arr := [13]int{5:400}
        fmt.Println(arr) // [0 0 0 0 0 400 0 0 0 0 0 0 0]
		
        arr2 := [13]string{5:"400", 7:"Ola Mundo"}
        fmt.Println(arr2) // [  400     Ola Mundo]
  }
  ```
------------
- Nao e possivel utilizar variaveis para definir o tamanho do array, porem e possivel utilizar constantes:

```go 
    func main(){
        var x int = 3
        arr := [x]int{1, 2, 3}
        fmt.Println(arr) // Erro: non-constant array bound x
		
		const y int = 3
        arr2 := [y]int{1, 2, 3}
        fmt.Println(arr2) // [1 2 3]
		
  }
  ```
------------

- Não e possivel aumentar ou diminuir o tamanho de um array.