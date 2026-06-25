# desafio_korp

Projeto de observabilidade com um serviço HTTP em Go, nginx na frente, coleta de métricas com Prometheus e painel no Grafana. O ambiente pode ser provisionado manualmente com Docker Compose ou automaticamente com Ansible.

## Visão geral

O sistema é composto por:

- `http-server`: aplicação Go que expõe o endpoint `/projeto-korp` e as métricas em `/metrics`
- `nginx`: proxy reverso que publica a aplicação na porta 80
- `prometheus`: coleta métricas da aplicação a cada 5 segundos
- `grafana`: dashboard para visualização das métricas

## Estrutura do projeto

- `http-server/`: código-fonte e Dockerfile da aplicação Go
- `nginx/`: configuração do proxy reverso
- `prometheus/`: configuração de scraping do Prometheus
- `grafana/`: provisionamento de datasource e dashboard
- `ansible/`: playbook e templates para provisionar o ambiente completo

## Requisitos

- Docker e Docker Compose
- ou Ansible, caso queira provisionar tudo automaticamente
- uma rede Docker externa chamada `korp-network`

## Como executar

### Opção 1: subir manualmente com Docker Compose

1. Crie a rede Docker externa, se ela ainda não existir:

	```bash
	docker network create korp-network
	```

2. Suba os containers a partir da raiz do projeto:

	```bash
	docker compose up -d --build
	```

3. Acesse os serviços:

	- aplicação: `http://localhost/projeto-korp`
	- métricas: `http://localhost/metrics`
	- Prometheus: `http://localhost:9090`
	- Grafana: `http://localhost:3000`

### Opção 2: provisionar com Ansible

O playbook em `ansible/playbook.yml` instala Docker, cria a rede `korp-network`, gera as configurações necessárias e sobe a stack.

Exemplo:

```bash
cd ansible
ansible-playbook -i inventory.ini playbook.yml
```

## Endpoints

### `GET /projeto-korp`

Retorna um JSON com o nome do projeto e o horário atual em UTC.

Exemplo de resposta:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-25T12:00:00Z"
}
```

### `GET /metrics`

Disponibiliza as métricas no formato do Prometheus, incluindo:

- `http_requests_total`
- `service_up`

## Observabilidade

O Prometheus faz scraping do serviço `http-server-projeto-korp:8080` a cada 5 segundos.

O Grafana já vem com datasource do Prometheus provisionado automaticamente e com dashboard carregado a partir de `grafana/dashboard/http-server-projeto-korp-dashboard.json`.

## Observações

- O nginx encaminha `/projeto-korp` e `/metrics` para o serviço Go.
- O container da aplicação expõe a porta interna `8080`, mas o acesso externo deve ser feito pelo nginx na porta `80`.
- Se você alterar nomes de container ou da rede, lembre-se de atualizar também os arquivos de configuração e o playbook.
