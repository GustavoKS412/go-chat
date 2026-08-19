# Go-Chat

[![Português](https://img.shields.io/badge/lang-Portugu%C3%AAs-green)](README.pt-BR.md)
[![English](https://img.shields.io/badge/lang-English-blue)](README.md)

Um app de chat em tempo real, feito em Go, usando WebSockets pra parte de mensagens.

Fiz esse projeto pra praticar. Não é pra ser uma plataforma de chat completa, é só um exemplo funcional de salas + mensagens ao vivo sem usar nenhum framework JS.

## Sobre o projeto

Cada usuário entra digitando o nome de uma sala. Se a sala ainda não existe, o servidor cria ela na hora; se já existe, o usuário entra na sala já existente com quem mais estiver lá.

Ao conectar, cada cliente recebe um nome de convidado aleatório, não tem login nem cadastro. As mensagens trafegam por WebSocket e são retransmitidas pro servidor pra todo mundo que tá na mesma sala. Como não tem persistência, ao fechar a aba, o histórico e a identidade do usuário somem.

## Tecnologias

- Go
- Gorilla WebSocket
- HTML
- CSS
- JavaScript

## Estrutura do projeto

```text
go-chat/
├── main.go
├── room.go
├── client.go
├── templates/
│   ├── index.html
│   └── chat.html
└── static/
    ├── css/
    │   └── styles.css
    └── js/
        └── chat.js
```

## Como executar o projeto

### Pré-requisitos

- Go instalado (confira a versão no `go.mod`)

### Passos

```bash
git clone https://github.com/GustavoKS412/go-chat.git
cd go-chat
go mod download
go run .
```

A aplicação sobe em `http://localhost:8080`. É só abrir no navegador, digitar o nome de uma sala e entrar no chat.
