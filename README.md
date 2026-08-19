# Go-Chat

[![English](https://img.shields.io/badge/lang-English-blue)](README.md)
[![Português](https://img.shields.io/badge/lang-Portugu%C3%AAs-green)](README.pt-BR.md)

A real-time chat app built in Go, using WebSockets for the messaging part.

I made this project to practice. It's not meant to be a full chat platform, it's just a working example of rooms + live messages without using any JS framework.

## About the project

Each user joins by typing the name of a room. If the room doesn't exist yet, the server creates it on the spot; if it already exists, the user joins the existing room with whoever else is in there.

On connecting, each client gets a random guest name, there's no login or sign-up. Messages travel over WebSocket and get broadcast by the server to everyone in the same room. Since there's no persistence, closing the tab wipes the history and the user's identity.

## Tech stack

- Go
- Gorilla WebSocket
- HTML
- CSS
- JavaScript

## Project structure

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

## How to run the project

### Prerequisites

- Go installed (check the version in `go.mod`)

### Steps

```bash
git clone https://github.com/GustavoKS412/go-chat.git
cd go-chat
go mod download
go run .
```

The app runs on `http://localhost:8080`. Just open it in your browser, type a room name, and join the chat.