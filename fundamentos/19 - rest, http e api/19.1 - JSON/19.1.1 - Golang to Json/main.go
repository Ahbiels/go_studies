package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type prices struct {
	TotalCost float64
	Credits   float64
	UsageCost float64
}

type services struct {
	Name    string   `json:"name"`
	Sku     string   `json:"sku"`
	Project []string `json:"project"`
	Prices  prices   `json:"prices"`
}

func main() {
	// Map ou Struct para json - Marshal
	// Json para struct ou map - Unmarshal
	structToJson()
	mapToJson()
}

func structToJson() {
	service := services{
		"Cloud SQL",
		"SUE-WID-AKDS",
		[]string{"project001", "project002", "project003"},
		prices{
			178.65,
			23.65,
			155.00,
		},
	}

	serviceToJSON, err := json.Marshal(service)
	if err != nil {
		log.Fatal(err)
	}
	// Buffer = uma área de memória temporária para guardar dados enquanto eles estão sendo produzidos ou consumidos.
	serviceToBuffer := bytes.NewBuffer(serviceToJSON)
	saveFile(serviceToBuffer, serviceToJSON, "services")
}

func mapToJson(){
	project := map[string]string {
		"projectName": "projeto001",
		"orgid": "org0032",
		"folderLocation": "folder001",
	}
	serviceToJSON, err := json.Marshal(project)
	if err != nil {
		log.Fatal(err)
	}
	// Buffer = uma área de memória temporária para guardar dados enquanto eles estão sendo produzidos ou consumidos.
	serviceToBuffer := bytes.NewBuffer(serviceToJSON)
	saveFile(serviceToBuffer, serviceToJSON, "projects")

}

func saveFile(jsonData *bytes.Buffer, bytesData []byte, fileName string) {
	format := ".json"
	fileName = fileName + format

	fmt.Println(jsonData)

	if err := os.WriteFile(fileName, bytesData, 0644); err != nil {
		log.Fatal(err)
	}
}
