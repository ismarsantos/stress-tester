# GO Stress Tester

## Descrição

Este projeto implementa uma ferramenta CLI em Go para realizar testes de carga em um serviço web. O usuário deve fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Funcionalidades

- **Realizar requests HTTP:** O sistema realiza requests HTTP para a URL especificada.
- **Distribuição de requests:** Distribui os requests de acordo com o nível de concorrência definido.
- **Garantia de requests:** Garante que o número total de requests seja cumprido.

## Geração de Relatório

Após a execução dos testes, o sistema gera um relatório contendo:
- Tempo total gasto na execução
- Quantidade total de requests realizados
- Quantidade de requests com status HTTP 200
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.)

## Requisitos

- [Go](https://golang.org/dl/) 1.18 ou superior
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração

### 1. Clonar o Repositório

Clone o repositório para o seu ambiente local:

```sh
git clone <URL_DO_REPOSITORIO>
cd stress-tester
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo .env na raiz do projeto e defina as variáveis de ambiente:

```
URL=http://google.com
REQUESTS=1000
CONCURRENCY=10
```

### 3. Construir e Executar a Imagem Docker

Remova qualquer contêiner e imagem anterior para garantir uma construção limpa. Construa a imagem Docker:

```sh
docker-compose build --no-cache
```

Inicie os contêineres:

```sh
docker-compose up
```

### 4. Executar Manualmente com Docker Run

Você pode executar o contêiner manualmente passando as variáveis de ambiente:

```sh
docker run --env-file .env stress-tester_stress-tester /root/stress-tester --url=$(grep URL .env | cut -d '=' -f2) --requests=$(grep REQUESTS .env | cut -d '=' -f2) --concurrency=$(grep CONCURRENCY .env | cut -d '=' -f2)
```

### Entrada de Parâmetros via CLI

A ferramenta aceita os seguintes parâmetros via CLI:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

Exemplo de execução:

```sh
docker run --env-file .env stress-tester_stress-tester /root/stress-tester --url=$(grep URL .env | cut -d '=' -f2) --requests=$(grep REQUESTS .env | cut -d '=' -f2) --concurrency=$(grep CONCURRENCY .env | cut -d '=' -f2)
```

## Estrutura do Projeto

```
stress-tester/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
└── .env
```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## Contato

Para mais informações ou dúvidas, entre em contato com [ismar.san@outlook.com](mailto:ismar.san@outlook.com).
