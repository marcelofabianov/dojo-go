# Documentação de Endpoints - Dojo Go API

Este documento descreve os endpoints disponíveis na Dojo Go API, com exemplos de como interagir com eles usando `curl`.

## 1. Health Check

Verifica o status da aplicação para garantir que ela está online e respondendo.

- Endpoint: `GET /healthz`
- Descrição: Retorna o status operacional da API.

**Comando**

```bash
curl 'http://localhost:8080/healthz'
```

**Resposta de Sucesso (`200 OK`)**

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
    "status": "OK"
}
```

## 2. Criação de novo Curso

Cria um novo registro de curso no banco de dados.

- Endpoint: `POST /api/v1/courses`
- Descrição: Cria um novo curso.

**Comando**


```bash
curl -i -X POST http://localhost:8080/api/v1/courses \
-H "Content-Type: application/json" \
-d '{
    "title": "Domain-Driven Design in Go",
    "description": "Applying DDD principles in Go applications."
}'
```

**Resposta de Sucesso (`201 Created`)**

```bash
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8

{
    "id": "01997b1a-c2a8-7d8e-b123-abcdef123456",
    "title": "Domain-Driven Design in Go",
    "description": "Applying DDD principles in Go applications.",
    "created_at": "2025-09-24 00:26:18.336917285 +0000 UTC"
}
```

**Resposta de Erro (`400 Bad Request`)**

```bash
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8

{
    "message": "Request validation failed",
    "code": "invalid_input",
    "details": [
        {
            "message": "validation failed on field 'Title'",
            "code": "invalid_input",
            "context": {
                "field": "Title",
                "param": "",
                "tag": "required"
            }
        }
    ]
}
```

## 3. Buscar Curso por ID

Busca um curso específico pelo seu ID.

- Endpoint: `GET /api/v1/courses/{id}`
- Descrição: Retorna um curso com base em seu ID.

**Comando**

```bash
# Substitua <COURSE_ID> por um UUID válido de um curso existente
curl -i http://localhost:8080/api/v1/courses/<COURSE_ID>
```

**Resposta de Sucesso (`200 OK`)**

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
    "id": "01997b1a-c2a8-7d8e-b123-abcdef123456",
    "title": "Domain-Driven Design in Go",
    "description": "Applying DDD principles in Go applications.",
    "created_at": "2025-09-24 00:26:18.336917285 +0000 UTC"
}
```

**Resposta de Erro (`404 Not Found`)**

Ocorre se o curso com o ID especificado não for encontrado.

```bash
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8

{
    "message": "course not found",
    "code": "not_found"
}
```

4. Atualizar Curso

Atualiza um curso existente

- Endpoint: `PUT /api/v1/courses/{id}`
- Descrição: Atualiza um curso com base em seu ID.

**Comando**

```bash
# Substitua <COURSE_ID> por um UUID válido de um curso existente
curl -i -X PUT http://localhost:8080/api/v1/courses/<COURSE_ID> \
-H "Content-Type: application/json" \
-d '{
    "title": "Advanced Domain-Driven Design in Go",
    "description": "Updated course with advanced DDD concepts."
}'
```

**Resposta de Sucesso (`200 OK`)**

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
    "id": "01997b1a-c2a8-7d8e-b123-abcdef123456",
    "title": "Advanced Domain-Driven Design in Go",
    "description": "Updated course with advanced DDD concepts.",
    "created_at": "2025-09-24 00:26:18.336917285 +0000 UTC"
}
```

_Nota: A resposta de erro 404 Not Found também se aplica aqui._

## 5. Deletar Curso

Remove um curso existente

- Endpoint: `DELETE /api/v1/courses/{id}`
- Descrição: Remove um curso

**Comando**

```bash
# Substitua <COURSE_ID> por um UUID válido de um curso existente
curl -i -X DELETE http://localhost:8080/api/v1/courses/<COURSE_ID>
```

**Resposta de Sucesso (`204 No Content`)**

```bash
HTTP/1.1 204 No Content
```

_Nota: A resposta de erro 404 Not Found também se aplica aqui._
