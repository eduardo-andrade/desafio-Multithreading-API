package main

import (
	"desafio-Multithreading-API/services"
	"desafio-Multithreading-API/models"
	"fmt"
	"encoding/json"
)

func main() {
	cep := "14091060"
	result, source, err := services.GetFastestCEP(cep)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	var jsonData []byte

	if source == "BrasilAPI" {
		data, _ := result.(models.BrasilAPIResponse)
		jsonData, _ = json.MarshalIndent(data, "", "  ")
	} else {
		data, _ := result.(models.ViaCEPResponse)
		jsonData, _ = json.MarshalIndent(data, "", "  ")
	}

	fmt.Println("Endereço recebido da API:", source)
	fmt.Println(string(jsonData))
}