# Comandos Basicos

- `go version` exibe a versao do go instalada
- `go env` exibe as variaveis de ambiente do go
- `go mod init {module name}` cria um arquivo go.mod que gerencia as dependencias do projeto
- `go mod tidy` verifica as dependencias do projeto e remove as que nao sao utilizadas
- `go mod vendor` cria uma pasta vendor com as dependencias do projeto
`go build {filename}.go` compila o codigo e cria um executavel
- `go build -o {output compiler}` o parametro -o permite definir o nome do executavel 
No windows nao e permitido declarar as variaveis inline, para isso e necessario setar as variaveis com o comando
- `go env -w GOOS=windows && go env -w GOARCH=amd64` seta a variavel de ambiente para windows
- `GOOS=windows GOARCH=amd64 go build -o {output compiler}` compila o codigo para windows
- `GOOS=linux GOARCH=amd64 go build -o {output compiler}` compila o codigo para linux

`go run {filename}.go` compila o codigo, roda o codigo imediatamente e apaga o que foi compilado