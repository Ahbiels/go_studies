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

### Concorrência
A concorrência em Go (Golang) é um recurso nativo poderoso e simples, baseado no modelo CSP (Processos Sequenciais Comunicantes), que usa Goroutines (funções leves e assíncronas) e Canais (para comunicação segura entre elas) para gerenciar múltiplas tarefas de forma eficiente, permitindo lidar com muitas coisas ao mesmo tempo e, com o agendador do Go, executar tarefas em paralelo em máquinas multi-core, diferentemente de threads pesadas de outros sistemas. 

#### Concorrência vs paralelismo
- **Concorrência**: refere-se à execução aparentemente simultânea de múltiplas tarefas em um mesmo ambiente. Isso permite que diferentes partes de um programa sejam executadas em paralelo, compartilhando recursos e melhorando a eficiência do sistema.
- **Paralelismo**: envolve a execução real simultânea de tarefas em recursos de processamento distintos. Ao dividir uma tarefa em partes menores, chamadas de threads ou processos, essas partes podem ser executadas independentemente em paralelo, aproveitando ao máximo os recursos disponíveis.

#### Goroutines
As goroutines são responsáveis por permitir a execução assíncrona e leve de tarefas, sem uma correnpondência direta com as thread do sistema operacional: uma goroutine é mapeada para (ou multiplexada por) uma ou mais threads, conforme a disponibilidade. O escalonador de goroutines do Go cuida dessa execução.

Ou seja, mesmo se um processador for de um thread, as gourotines ainda conseguem executar blocos de código de forma assíncrona (ou seja, simultaneamente). As goroutines são executadas em um único thread chamado “máquina de execução de goroutines” (goroutine scheduler). Essa máquina de execução é responsável por escalonar e distribuir as goroutines nos núcleos de processamento disponíveis, garantindo uma execução eficiente e concorrente.

As goroutines são executadas no mesmo espaço de endereços, o que significa que o acesso à memória compartilhada deve ser sincronizado. Embora seja possível usar o **pacote sync para sincronização**, essa abordagem é altamente <span style="color: red">desencorajada</span>. Em vez disso, o Go utiliza ***“channels”*** para sincronizar goroutines.

- **runtime.Goexit()**: uma goroutine em execução pode interromper a si mesma chamando esse método, embora isso raramente seja necessário
- **runtime.Gosched()**: Quando uma goroutine é muito intensiva em processamento, você pode chamar periodicamente runtime.Gosched() em seus loops de cálculo. Isso cede o processador, permitindo que outras goroutines sejam executadas. Isso não suspende a goroutine atual, então a execução é retomada automaticamente

#### Channels

Para podermos tirar o máximo proveito das goroutines, é preciso que elas se comuniquem, já que as goroutines não se comunicam por padrão. Isso pode ser feito por meio do compartilhamento de variáveis, mas isso é extremamente complexo, já que essa forma de trabalhar apresenta todas as dificuldades com memória compartilhada em multithreading. 

Para isso, o Go introduz o que chamamos de **channel**, que permite conectar canais concorrentes para ser possível enviar dados e lidar com a comunicação entre as goroutines, dessa forma evitamos os problemas da memória compartilhada. Já que o ato de comunicação mediante um canal garante a sincronização; não há necessidade de sincronização explícita por meio de bloqueio de nada.
- Um canal é uma fila de mensagens tipadas: dados podem ser transmitidos por meio dele. Trata-se de uma estrutura <span style="color:yellow">FIFO(First in First Out)</span>, o que significa que preserva a ordem dos itens enviados para ele.
- Além disso, um canal é uma referência, o que significa que devemos usar a função make() para alocar memória para ele. Aqui está um exemplo de uma declaração de canal chamado ch1 para strings, seguido pela sua instanciação:
  ```go
  var ch1 chan string
  ch1 = make(chan string)

  //ou

  ch1 := make(chan string)
  ```

Visualmente, um canal pode se parecer com isso
```css
Goroutine A  --->  [ channel ]  --->  Goroutine B
```

<span style="color: red">Enviar e receber em channel é BLOQUEANTE (na maioria dos casos)</span>

**Exemplo**
```golang
ch := make(chan int)
```

Não tem espaço interno. Isso significa que:
- Quem envia bloqueia até alguém receber
- Quem recebe bloqueia até alguém enviar

```golang
func main() {
	ch := make(chan int)

	go func() {
		ch <- 10 // BLOQUEIA até alguém receber
	}()

	valor := <-ch // desbloqueia o envio
	fmt.Println(valor)
}
```

Deadlock

```golang
func main() {
	ch := make(chan int)
	ch <- 10 // NINGUÉM está recebendo
}
```

#### Channels com buffer
O buffer permite dizer qual o "tamanho" da pilha dentro do canal, impedindo que um envio ou recebimento bloqueie o programa

```golang
package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // BLOQUEARIA (buffer cheio)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

