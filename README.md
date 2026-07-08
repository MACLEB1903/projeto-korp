# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta etapa, foi construído um serviço HTTP em Golang e a estrutura inicial necessária para sua execução em container Docker.

## Tecnologias Utilizadas

- Golang
- Docker

## Estrutura do Projeto

```text
projeto-korp/
├── app/
├── .gitignore
└── README.md
```

## Funcionalidades Implementadas

### Parte 1: Serviço e Arquitetura

- Servidor HTTP em Golang
- Endpoint GET /projeto-korp
- Retorno de resposta em formato JSON
- Dockerfile para build da aplicação

## Como Executar

Para executar a aplicação, primeiro construa a imagem Docker:

```
docker build -t http-server-projeto-korp ./app
```

Em seguida, execute o container:

```
docker run -d --name http-server-projeto-korp -p 8080:8080 http-server-projeto-korp
```

## Como Testar

Para testar a aplicação, execute:

```
curl http://localhost:8080/projeto-korp
```

A resposta esperada é:

```
{
"nome": "Projeto Korp",
"horario": "<horário_atual>"
}
```
