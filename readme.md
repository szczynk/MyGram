# Final Project - MyGram

Kelas: Scalable Web Services with Go (Programming Language) - Digitalent ✕ Hacktiv8

MyGram is a free photo sharing app written in Go. People can share, view, and comment photos by everyone. Anyone can create an account by registering an email address and selecting a username.

## Getting Started

To start running this project locally,

```bash
git clone https://github.com/szczynk/MyGram.git
```

Open MyGram folder and install all required dependencies

```bash
cd MyGram && go mod tidy
```

Copy the example config file and adjust the config file

```bash
cp config.example.yaml config.yaml
```

Start the server

```bash
go run main.go
```
