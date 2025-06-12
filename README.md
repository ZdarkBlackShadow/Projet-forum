# Projet-forum

## Prerequisites

- Git
- Golang (version 1.23.8)

## Install 

```bash
git clone https://github.com/ZdarkBlackShadow/Projet-forum.git
```

## Start the project

```bash
go run main.go
```

##  Le .env à compléter avec vos données

```bash
PEPPER =
DB_NAME =
DB_PORT =
DB_USER =
DB_PWD =
DB_HOST =
JWT_SECRET =
```

## Read the documentation with godoc

### Install godoc

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

### Start the godoc serve

```bash
godoc -http=:6060
```

### The address to put in your browser

```bash
http://localhost:6060/pkg/projet-forum/
```
