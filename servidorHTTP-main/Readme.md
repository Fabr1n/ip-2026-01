# Servidor HTTP em Go · Cadastro de pacientes

Projeto acadêmico com operações de criação, consulta, atualização e exclusão de pacientes, usando Go, PostgreSQL e páginas HTML.

## Implementação atual

| Parte | Arquivo |
| :--- | :--- |
| Inicialização e rotas | [app/main.go](./app/main.go) |
| Handlers de pacientes | [app/handlers/patientHandler.go](./app/handlers/patientHandler.go) |
| Conexão e criação do banco/tabela | [app/utils/connectToDB.go](./app/utils/connectToDB.go) |
| Interface usada pelo servidor | [app/static](./app/static/) |

A implementação atual usa a tabela `patients`. Os arquivos legados de contas em `static/` não correspondem às rotas registradas no `app/main.go` atual.

## Pré-requisitos

- Go compatível com a versão declarada em [go.mod](./go.mod), atualmente 1.24.1.
- PostgreSQL em execução.
- Usuário do banco com permissão para criar o banco de dados, como exigido por `ensureDatabase`.

## Configuração

Clone e entre no módulo:

```bash
git clone https://github.com/Fabr1n/ip-2026-01.git
cd ip-2026-01/servidorHTTP-main
go mod download
```

Configure as variáveis no ambiente do processo:

| Variável | Finalidade |
| :--- | :--- |
| `DB_HOST` | Host do PostgreSQL |
| `DB_PORT` | Porta do PostgreSQL |
| `DB_USER` | Usuário do banco |
| `DB_PASSWORD` | Senha do banco |
| `DB_NAME` | Nome do banco da aplicação |

O arquivo [.env.example](./.env.example) lista as variáveis. O código usa `os.Getenv`; ele **não carrega um arquivo `.env` automaticamente**.

Exemplo de configuração em PowerShell:

```powershell
$env:DB_HOST = "localhost"
$env:DB_PORT = "5432"
$env:DB_USER = "postgres"
$env:DB_PASSWORD = "<sua senha local>"
$env:DB_NAME = "servidorhttp"
go run ./app/main.go
```

A conexão inicial cria o banco quando necessário e executa `CREATE TABLE IF NOT EXISTS patients`.

## Execução

Na pasta `servidorHTTP-main`, execute:

```bash
go run ./app/main.go
```

Acesse **http://localhost:8080/**. O terminal também mostra o endereço da interface de rede local.

## Rotas registradas

| Rota | Operação |
| :--- | :--- |
| `/` | Página inicial |
| `/createPatient` | Formulário e criação de paciente |
| `/patients` | Listagem |
| `/updatePatient?id=...` | Edição |
| `/deletePatient?id=...` | Exclusão |
| `/static/` | Arquivos em `app/static/` |

Use dados fictícios nos exercícios e nas demonstrações. Este README descreve os arquivos e as rotas atuais; não representa validação de produção.

---

[Voltar ao repositório acadêmico](../README.md) · [Perfil do autor](https://github.com/Fabr1n)
