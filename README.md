pacotes -> grupo de arquivos dentro do mesmo diretório e são compiladas juntos
- uma variável dentro de um arquivo, fica visivel dentro de todos os outros arquivos

Executar o arquivo go
```sh
go run first.go
```


### Termos
- Ponteiros: 
- Map: maps é uma coleção integrada e não ordenada de pares chave-valor, ideal para buscas, inserções e exclusões rápidas. Os mapas são essencialmente referências a uma tabela hash.
  - Permite armazenar dados usando uma chave única para acesso rápido ao valor correspondente. É ideal para situações onde a busca por um item é feita através de um identificador
  - Muito utilizado para associar um tipo de dado a outro, como, por exemplo, mapear IDs de usuários para seus nomes, códigos de moedas para suas taxas de câmbio, ou nomes de países para suas capitais
  - É excelente para contar a ocorrência de elementos, como a frequência de palavras em um texto. 
- Loops: o Golang só utiliza o **for** como estrutura de Loop, sem While, for in, etc.
- Funções Variáticas: funções que aceitam receber um número variado de parâmetros. 
  - O Variático é sempre o último no parâmetro, e só pode haver 1
- defer: com a função defer, você pode chamar uma função, mas impedir sua execução até que outras funções sejam executadas. Ela adia funções e trechos no código até o final da execução de um programa
- Panic: para que sua aplicação Go entre em pânico (panic), ela precisa chegar a um ponto em que não consegue mais executar porque não sabe o que fazer. Isso pode ser acionado pelo compilador ou manualmente em seu código.
- Recover: Quando um programa entra em estado de pânico, temos a opção de ajudá-lo a se recuperar, usando a palavra-chave recover.

### Código
- slice: fatia de um array (interno).
  - O slice tem uma capacidade máxima, mas quando essa capacidade máxima é atingida, ela é dobrada. 
- make: função que me permite alocar um valor na memória
- len: mostra o tamanho de um slice/array
- cap: mostra a capacidade máxima de um slice/array

### Detalhes

#### Mais sobre o Panic
**O que ele faz**
- Interrompe imediatamente o fluxo normal do programa.
- Desempilha a stack (stack unwinding).
- Executa todos os defer pendentes.
- Se não for recuperado, derruba o programa.

**Quando usar**
- Erros irrecuperáveis (bug de programação)
  - Algo que nunca deveria acontecer se o código estivesse correto.
- Falha crítica na inicialização
  - Se a aplicação não pode subir sem aquilo.
- Em funções Must...
  - Convenção idiomática do Go.

Ex:
```go
func MustParseInt(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return n
}
```

O panic é diferente de um erro. O panic interrompe o fluxo do programa se não for recuperado. 

Resumindo:

| Tipo | Como tratar |
| ---  |     ---     |
| Erro esperado | `error` |
| Erro inesperado, mas recuperável | panic + recover |
| Erro impossível / bug | panic |