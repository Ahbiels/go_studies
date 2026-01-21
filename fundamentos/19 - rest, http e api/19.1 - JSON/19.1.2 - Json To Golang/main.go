package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Especificacoes struct {
	Processador   string `json:"processador"`
	Ram           uint   `json:"ram"`
	Armazenamento string `json:"armazenamento"`
}

type Produto struct {
	Id             int32          `json:"id"`
	Produto        string         `json:"produto"`
	Preco          float64        `json:"preco"`
	Disponivel     bool           `json:"disponivel"`
	Especificacoes Especificacoes `json:"especificacoes"`
	Tags           []string       `json:"tags"`
}

func main() {
	dataProduto := openFile("file.json")
	dataName := openFile("name.json")
	var produto Produto
	name := make(map[string]string)

	// fmt.Printf("%T\n",data)

	if err := json.Unmarshal([]byte(dataProduto), &produto); err != nil {
		log.Fatal(err)
	}

	fmt.Println(produto)

	if err := json.Unmarshal([]byte(dataName), &name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)

}

func openFile(nameFile string) string {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
