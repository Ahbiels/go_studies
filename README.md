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

### Código
- slice: fatia de um array (interno).
  - O slice tem uma capacidade máxima, mas quando essa capacidade máxima é atingida, ela é dobrada. 
- make: função que me permite alocar um valor na memória
- len: mostra o tamanho de um slice/array
- cap: mostra a capacidade máxima de um slice/array
