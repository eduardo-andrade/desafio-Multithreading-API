package models

// Estrutura para a resposta da BrasilAPI
type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

// Estrutura para a resposta da ViaCEP
type ViaCEPResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IBGE         string `json:"ibge"`
	GIA          string `json:"gia"`
	DDD          string `json:"ddd"`
	SIAFI        string `json:"siafi"`
}