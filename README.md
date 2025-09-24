# Dojo Go API

Uma API RESTful em Go desenvolvida como parte de um dojo de programação, demonstrando conceitos de basicos de Go, injeção de dependência e boas práticas de desenvolvimento.

## ✨ Sobre o Projeto

Este projeto consiste em uma API para gerenciamento de cursos, construída com Go. A arquitetura foi desenhada para ser modular, escalável e testável, separando claramente as responsabilidades entre as camadas.

Utilizamos ferramentas modernas do ecossistema Go, como o router `Chi` para roteamento HTTP, `sqlx` para interação com o banco de dados, `uber-fx` para injeção de dependência e `goose` para gerenciamento de migrations. Todo o ambiente de desenvolvimento é containerizado com Docker.

---

## 🚀 Começando

Siga os passos abaixo para configurar e executar o ambiente de desenvolvimento localmente.

### Pré-requisitos

* **Docker**
* **Docker Compose**

### Instalação e Execução

1.  **Clone o repositório**
    ```bash
    git clone https://github.com/marcelofabianov/dojo-go
    cd dojo-go
    ```

2.  **Configure as variáveis de ambiente**
    Copie o arquivo de exemplo `.env.example` para criar seu arquivo de configuração local `.env`.
    ```bash
    cp .env.example .env
    ```

3. **Edite o .env nas variáveis de ambiente necessárias**
    ```bash
    HOST_UID=1001
    HOST_GID=1001
    ```

4.  **Inicie os serviços com Docker Compose**
    Este comando irá construir as imagens e iniciar os contêineres da API e do banco de dados (PostgreSQL) em background (`-d`).
    ```bash
    docker compose up -d
    ```

5.  **Execute as migrations do banco de dados**
    Para criar as tabelas necessárias, execute o `goose` dentro do contêiner da aplicação.
    ```bash
    docker exec -it dojo-api goose up
    ```

A API estará disponível em `http://localhost:8080`.

---

## 🛠️ Uso e Endpoints da API

Você pode interagir com a API utilizando uma ferramenta como `curl` ou Postman.

### 1. Health Check

Verifique se a API está online e respondendo.

* **Endpoint:** `GET /healthz`
* **Comando:**
    ```bash
    curl 'http://localhost:8080/healthz'
    ```
* **Resposta Esperada:**
    ```json
    {"status":"OK"}
    ```

### 2. Criar um Novo Curso

Crie um novo registro de curso no banco de dados.

* **Endpoint:** `POST /api/v1/courses`
* **Comando:**
    ```bash
    curl --location 'http://localhost:8080/api/v1/courses' \
    --header 'Content-Type: application/json' \
    --data '{
        "title": "Introduction to Go",
        "description": "A comprehensive course on Golang basics."
    }'
    ```
* **Resposta de Sucesso (Status `201 Created`):**
    ```json
    {
        "id": "019978f2-b2c4-7850-99e0-eff33bcda947",
        "title": "Introduction to Go",
        "description": "A comprehensive course on Golang basics.",
        "created_at": "2025-09-23 23:39:55.460545098 +0000 UTC"
    }
    ```

### Endpoints

Consule o arquivo [API.md](API.md) para mais detalhes.

---

## 💻 Tecnologias Utilizadas

* **Linguagem:** [Go](https://golang.org/)
* **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
* **Containerização:** [Docker](https://www.docker.com/)
* **Roteador HTTP:** [Chi](https://github.com/go-chi/chi)
* **Acesso ao Banco de Dados:** [sqlx](https://github.com/jmoiron/sqlx)
* **Injeção de Dependência:** [Uber FX](https://github.com/uber-go/fx)
* **Migrations:** [Goose](https://github.com/pressly/goose)
* **Configuração:** [Viper](https://github.com/spf13/viper) & [Godotenv](https://github.com/joho/godotenv)
* **Validação:** [go-playground/validator](https://github.com/go-playground/validator)

---

## 📄 Licença

Este projeto é distribuído sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
