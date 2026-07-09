# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta etapa, objetivo a é adicionar monitoramento ao serviço `http-server-projeto-korp` usando Grafana e Prometheus.

## Tecnologias Utilizadas

- Golang
- Docker
- Docker Compose
- NGINX
- Prometheus
- Grafana

## Estrutura do Projeto

```text
projeto-korp/
├── app/
├── grafana/
├── nginx/
├── prometheus
├── .gitignore
├── docker-compose.yaml
└── README.md
```

## Funcionalidades Implementadas

### Parte 1: Serviço e Arquitetura

#### 1. Serviço HTTP

- Criar um servidor HTTP em Golang com o endpoint `GET /projeto-korp`.
- Retornar uma resposta em formato JSON.
- Criar o Dockerfile da aplicação.
- Instalar, configurar e validar o Docker em ambiente Linux.
- Criar uma rede Docker para comunicação entre containers.
- Configurar o Docker Compose com os containers `http-server-projeto-korp` e `nginx`.
- Configurar o NGINX como proxy reverso para encaminhar requisições à porta `8080`.
- Testar o ambiente com curl `http://localhost:80/projeto-korp`.

#### Etapa 2: Monitoramento com Prometheus e Grafana

- Implementar métricas no serviço
- Expor métricas no padrão Prometheus
- Configurar Prometheus
- Configurar Grafana
- Criar dashboard no Grafana
- Atualizar Docker Compose
- Testar coleta e visualização das métricas

## Como Executar

Para rodar a aplicaçâo, execute:

```
docker compose up --build
```

## Como Testar

Para validar se a aplicação está funcionando, execute:

```
curl http://localhost:80/projeto-korp
```

```
# A resposta esperada é:

{
    "nome":"Projeto Korp",
    "horario":"2026-07-08T15:53:36Z"
}
```

Para validar se o Grafana está funcionando, acesse:

```
http://localhost:3000/
```

```
# Use as seguintes credenciais:

username: korp
password: korp
```
