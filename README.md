# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta etapa, foi configurado o proxy reverso do NGINX por meio do arquivo http-server-projeto-korp.conf, encaminhando as requisições feitas para http://localhost:80 ao serviço http-server-projeto-korp na porta 8080.

## Tecnologias Utilizadas

- Golang
- Docker
- Docker Compose
- Nginx

## Estrutura do Projeto

```text
projeto-korp/
├── app/
├── nginx/
├── .gitignore
├── docker-compose.yaml
└── README.md
```

## Funcionalidades Implementadas

### Parte 1: Serviço e Arquitetura

#### 1. Serviço HTTP

- Servidor HTTP em Golang
- Endpoint `GET /projeto-korp`
- Retorno de resposta em formato JSON
- Dockerfile para build da aplicação
- Instalação e a configuração do Docker

#### 2. Instalação e Configuração do Docker

- Instalação e configuração do Docker em ambiente Linux
- Validação do funcionamento do Docker

#### 3. Configuração de Rede Docker

- Criação de uma rede Docker

#### 4. Docker Compose

- Configurar dois containers: `http-server-projeto-korp` e `nginx`
- Conectar ambos à mesma rede Docker
- Manter o http-server-projeto-korp sem porta exposta ao host
- Mapear a porta 80 do NGINX e montar o volume em `/etc/nginx/conf.d/`

#### 5. Configuração do Proxy Reverso

- Adição do arquivo `http-server-projeto-korp.conf` com a configuração do proxy reverso.
- Encaminhamento das requisições feitas para `http://localhost:80` ao serviço na porta `8080`.

#### 6. Teste de Funcionamento

- Teste o ambiente usando `curl http://localhost:80/projeto-korp `.

## Como Executar

Para criar e iniciar os containers, a rede e o volume, execute:

```
docker compose up --build
```

## Como Testar

Para validar se o serviço está funcionando, execute:

```
curl http://localhost:80/projeto-korp
```

A resposta esperada é:

```
{"nome":"Projeto Korp","horario":"2026-07-08T15:53:36Z"}
```
