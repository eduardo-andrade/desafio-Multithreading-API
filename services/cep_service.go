package services

import (
	"encoding/json"
	"errors"
	"desafio-Multithreading-API/models"
	"desafio-Multithreading-API/utils"
	"fmt"
	"time"
)

var (
	brazilAPIURL = "https://brasilapi.com.br/api/cep/v1/%s"
	viaCEPURL    = "https://viacep.com.br/ws/%s/json/"
)

func GetFastestCEP(cep string) (interface{}, string, error) {
	results := make(chan struct {
		data   interface{}
		source string
		err    error
	}, 2)

	go fetchCEP(fmt.Sprintf(brazilAPIURL, cep), "BrasilAPI", results)
	go fetchCEP(fmt.Sprintf(viaCEPURL, cep), "ViaCEP", results)

	select {
	case res := <-results:
		if res.err != nil {
			return nil, "", res.err
		}
		return res.data, res.source, nil
	case <-time.After(1 * time.Second):
		return nil, "", errors.New("timeout: nenhuma resposta dentro do tempo limite")
	}
}

func fetchCEP(url, source string, results chan<- struct {
	data   interface{}
	source string
	err    error
}) {
	body, err := utils.Fetch(url)
	if err != nil {
		results <- struct {
			data   interface{}
			source string
			err    error
		}{err: err}
		return
	}

	if source == "BrasilAPI" {
		var response models.BrasilAPIResponse
		if err := json.Unmarshal(body, &response); err != nil {
			results <- struct {
				data   interface{}
				source string
				err    error
			}{err: err}
			return
		}
		results <- struct {
			data   interface{}
			source string
			err    error
		}{data: response, source: source}
	} else {
		var response models.ViaCEPResponse
		if err := json.Unmarshal(body, &response); err != nil {
			results <- struct {
				data   interface{}
				source string
				err    error
			}{err: err}
			return
		}
		results <- struct {
			data   interface{}
			source string
			err    error
		}{data: response, source: source}
	}
}