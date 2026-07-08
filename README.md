# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta etapa, foi criada uma rede Docker no modo bridge para comunicação entre containers.

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

#### 1. Serviço HTTP

- Servidor HTTP em Golang
- Endpoint GET /projeto-korp
- Retorno de resposta em formato JSON
- Dockerfile para build da aplicação
- Instalação e a configuração do Docker

#### 2. Instalação e Configuração do Docker

- Instalação e configuração do Docker em ambiente Linux
- Validação do funcionamento do Docker

#### 3. Configuração de Rede Docker

- Criação de uma rede Docker

## Como Executar

Para criar uma rede, execute:

```
docker network create korp-network
```

## Como Testar

Para validar se a rede foi criada:

```
docker network ls
```

A resposta esperada é:

```
NETWORK ID     NAME           DRIVER    SCOPE
xxxxxxxxxxxx   korp-network   bridge    local
```
