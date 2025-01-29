# Projeto Desafio Multithreading API

Este projeto em Go realiza consultas concorrentes a múltiplas APIs de CEP (BrasilAPI e ViaCEP) e retorna a resposta mais rápida.

## Estrutura de Diretórios
```
desafio_Multithreading_API/
├── models/
│   ├── address.go
├── services/
│   ├── cep_service.go
├── utils/
│   ├── http_client.go
├── main.go
├── go.mod
└── README.md
```

## Requisitos
- Go 1.20 ou superior
- Conexão com a internet para acessar as APIs de CEP:
  - [BrasilAPI](https://brasilapi.com.br/)
  - [ViaCEP](https://viacep.com.br/)

## Instalação

1. **Clone o repositório**:
   ```bash
   git clone <url-do-repositorio>
   cd desafio_Multithreading_API
   ```

2. **Baixe as dependências**:
   ```bash
   go mod tidy
   ```

## Execução

1. Execute o programa principal:
   ```bash
   go run main.go
   ```
   
2. O programa irá buscar o CEP definido no código e exibir o endereço retornado pela API que respondeu primeiro. Exemplo de saída:
   ```bash
   Endereço recebido da API: BrasilAPI
   {
     "cep": "01153-000",
     "street": "Rua Vitorino Carmilo",
     "neighborhood": "Barra Funda",
     "city": "São Paulo",
     "state": "SP"
   }
   ```

## Erros Comuns
- **"timeout: nenhuma resposta dentro do tempo limite"**:
  - Isso ocorre quando ambas as APIs demoram mais de 1 segundo para responder. Tente novamente ou aumente o timeout em `cep_service.go`.

- **"package not found"** ao importar módulos internos:
  - Certifique-se de estar no diretório raiz do projeto e execute:
    ```bash
    go mod tidy
    ```
## Licença
Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.