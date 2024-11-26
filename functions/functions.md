# Functions

- Funçoes sao blocos que podem ser repetidos no codigo.

- Estrutura da funçao:
- ```go 
    func {name}({params}) { funçao }
  ```

------------
- E possivel declarar se a funçao tem retorno ou nao: 
- ```go
    func sum(a int, b int) int {  return a + b }
   ```
------------
- E possivel declarar o nome do retorno ex:
- ```go
  func division(a, b int) (res int, rem int) {
  res = a / b
  rem = a % b
  return
  }
  ```

------------
- E possivel criar funçoes anonimas:
```go
func hiOrderSum(a int) func(int) int {
   		return func(b int) int {
			   return a + b
   		}
   }
```
------------
- E criar funçoes com varios parametros do mesmo tipo:
- ```go
    func variaticSum(nums ...int) int {
    var out int
    for _, num := range nums {
    out += num
    }
    return out
    } 
  ```

