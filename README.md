# Projeto Korp

Projeto desenvolvido para o processo seletivo de estágio na Korp ERP, com o objetivo de avaliar habilidades com Docker, programação, redes, servidores e automação em ambiente Linux.

Nesta última etapa, o objetivo é automatizar toda a configuração do ambiente descrita nas Partes 1 e 2 usando Ansible.

## Tecnologias Utilizadas

- Golang
- Docker
- Docker Compose
- NGINX
- Prometheus
- Grafana
- Ansbile

## Estrutura do Projeto

```text
projeto-korp/
├── ansible
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

### Etapa 3: Automação com Ansible

Usando Ansible:

- Instalar Docker em ambiente Linux
- Criar rede Docker
- Fazer build da imagem da aplicação
- Criar e executar containers com Docker Compose
- Configurar NGINX como proxy reverso
- Configurar Prometheus e Grafana
- Validar endpoint com requisição HTTP
- Exibir resposta da aplicação no console
- Provisionar tudo com um único comando Ansible

## Como Executar

Verifique se você tem o Ansible instalado:

```
ansible-playbook --version
```

Caso não tenha, execute:

```
sudo apt update
sudo apt install -y ansible
ansible-playbook --version
```

Para executar o playbook, execute:

```
cd projeto-korp
ansible-playbook ansible/main.yaml
```

Você deverá ver um resultado semelhante a este:

```
ok: [localhost] => {
    "msg": {
        "horario": "2026-07-09T04:53:25Z",
        "nome": "Projeto Korp"
    }
}
```
