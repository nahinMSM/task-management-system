# Task Management System

Este é um sistema de gerenciamento de tarefas desenvolvido como um desafio onde adquiri conhecimentos em **GoLang** e **Docker**. O projeto segue a arquitetura **RESTful API** e utiliza **PostgreSQL** como banco de dados.


## Tecnologias Utilizadas

### Backend:
- **GoLang** - Linguagem principal do backend
- **Docker** - Para conteinerização do backend e banco de dados
- **PostgreSQL** - Banco de dados relacional
- **Gin** - Framework para desenvolvimento da API em Go
- **JWT** - Autenticação segura via JSON Web Token
- **GORM** - ORM para interação com o banco de dados
- **Thunder Client** - Para testes das APIs (Extensão do VS Code)
- **Jenkins** - Para automação de build e deploy
- **WebSockets** - Para comunicação em tempo real (futuramente)

### Frontend:
- **React.js** - Biblioteca para construção da interface
- **Next.js** - Framework para React com renderização otimizada
- **Axios** - Cliente HTTP para comunicação com a API

![](/frontend/public/Captura-de-tela.png)

<br>

## O que Aprendi com este Projeto

Durante o desenvolvimento deste sistema, aprofundei meus conhecimentos em:
- Programar em **GoLang**, criando APIs RESTful e gerenciando rotas com **Gin**.
- Configurar e utilizar **Docker** para conteinerizar o backend e banco de dados.
- Criar e gerenciar bancos de dados PostgreSQL com **GORM**.
- Implementar autenticação segura com **JWT**.
- Trabalhar com **Jenkins** para automação de CI/CD.
- Utilizar **WebSockets** para funcionalidades em tempo real (ainda em desenvolvimento).
- Desenvolver uma interface web moderna com **React.js** e **Next.js**.
- Integrar frontend e backend utilizando **Axios** para chamadas HTTP.

## Como Clonar e Executar o Projeto

### 1. Clonar o Repositório

```sh
 git clone https://github.com/nahinMSM/task-management-system.git
 cd task-management-system
```

### 2. Configurar as Variáveis de Ambiente
Crie um arquivo `.env` na raiz do projeto e configure as credenciais do banco de dados:

```
DB_HOST=go-db
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=task_db
DB_PORT=5432
JWT_SECRET=sua_chave_secreta
```

### 3. Subir os Containers com Docker

```sh
 docker-compose up -d
```

Isso iniciará os containers do backend e do banco de dados PostgreSQL.

### 4. Executar o Frontend

Acesse a pasta do frontend e instale as dependências:
```sh
 cd frontend
 npm install
```

Inicie o servidor de desenvolvimento:
```sh
 npm run dev
```
O frontend estará disponível em `http://localhost:3000`.

### 5. Testar a API

Para verificar se a API está rodando corretamente, use o Thunder Client ou faça uma requisição via cURL:

```sh
curl http://localhost:8080/ping
```

A resposta esperada:
```json
{"message": "pong"}
```

## Rotas da API

A API possui endpoints para gerenciamento de tarefas, incluindo:

- **POST /register** - Criar uma conta
- **POST /login** - Autenticação via JWT
- **GET /tasks** - Listar todas as tarefas
- **POST /tasks** - Criar uma nova tarefa
- **PUT /tasks/:id** - Atualizar uma tarefa
- **DELETE /tasks/:id** - Excluir uma tarefa

## Contribuição

Se quiser contribuir, siga os passos:
1. Fork o repositório
2. Crie uma branch (`feature/nova-funcionalidade`)
3. Faça suas alterações e comite
4. Envie um pull request

---

Desenvolvido por **Nahin Moreira** 🚀

