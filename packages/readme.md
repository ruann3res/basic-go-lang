# Packages

- Todo arquivo GO deve estar dentro de um pacote (na primeira linha do arquivo), Apenas comentarios podem estar antes da declaraçao do pacote. 

- Pacotes podem ser importados de outros pacotes, para isso basta usar a palavra reservada `import` seguida do caminho do pacote.

- Para utilizar um pacote, basta usar o nome do pacote seguido de um ponto e o nome da função ou variável que deseja utilizar.

- Por convenção, o nome do pacote ele e sempre o ultimo e elemento do caminho do pacote.

- Para renomear um import por qualquer motivo basta definir o novo nome + o caminho do pacote. `import (novoNome "fmt")`

- E possivel importar um pacote utilizando `.` para que nao seja necessario usar o nome do pacote para acessar suas funçoes e variaveis (Nao recomendado). `import . "fmt"` 