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
git clone https://github.com/ismarsantos/stress-tester
cd stress-tester
```

### 2. Construir e Executar a Imagem Docker

Remova qualquer contêiner e imagem anterior para garantir uma construção limpa. Construa a imagem Docker:

```sh
docker-compose build --no-cache
```

### 3. Executar com Docker Run


```sh
docker run stress-tester —url=http://google.com —requests=1000 —concurrency=10
```

### 4. Resultado esperado


```sh
Total time taken: 48.505189456s
Total requests: 1000
200 OK: 1000
```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## Contato

Para mais informações ou dúvidas, entre em contato com [ismar@diamantinastudio.com](mailto:ismar@diamantinastudio.com).
