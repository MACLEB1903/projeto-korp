# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta etapa, foram criados a rede Docker no modo bridge, os containers e o volume usando o docker-compose.yaml.

## Tecnologias Utilizadas

- Golang
- Docker

## Estrutura do Projeto

```text
projeto-korp/
├── app/
├── .gitignore
├── docker-compose.yaml
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

#### 4. Docker Compose

- Configurar dois containers: http-server-projeto-korp e nginx
- Conectar ambos à mesma rede Docker
- Manter o http-server-projeto-korp sem porta exposta ao host
- Mapear a porta 80 do NGINX e montar o volume em /etc/nginx/conf.d/

## Como Executar

Para criar os containers, a rede e o volume, execute:

```
docker compose up --build
```

## Como Testar

Para validar se os containers, a rede e o volume foram criados:

```
docker ps
docker network ls
docker volume ls
```

As respostas esperadas são:

```
CONTAINER ID   IMAGE                      ...   NAMES
xxxxxxxxxxxx   nginx:alpine               ...   projeto-korp-nginx-1
xxxxxxxxxxxx   http-server-projeto-korp   ...   http-server-projeto-korp
```

```
NETWORK ID     NAME           DRIVER    SCOPE
xxxxxxxxxxxx   korp-network   bridge    local
```

```
local     korp-volume
```
