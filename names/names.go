package names

import (
	"fmt"
	"names/internal/mine"
)

func main() {
	fmt.Println(mine.Mine)  // Variavel Publica
	fmt.Println(mine.Atest) // Variavel Publica
	fmt.Println(mine.Btest) // Variavel Publica
	// fmt.Println(mine.c) Variavel privada (Não pode ser acessada)
}

func Main() {}
