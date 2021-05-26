# Go Web Api

**WebApi REST created in GoLang**

This project was created in order to improve my knowledge in GoLang and REST API's

---

[![Go WebApi PR CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_pr_develop.yml/badge.svg?branch=develop)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_pr_develop.yml)
[![Go WebApi Push CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml/badge.svg?branch=develop)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml)


## Features

**LEGEND**

📝 Was Implemented

🧑‍💻 Work Now

💻 Future feature 

- 📝 User Management
- 📝 Session using JWT with asymmetric keys
- 🧑‍💻 Authentication Middleware

## Installation

### Create your own RSA keys

- To the keys make sure you have the **openssl** CLI installed and run this command bellow:

```bash
make private-key
```

```bash
make public-key
```

### Get all Go packages we need

```bash
go get -u
```

### Create a PostgreSQL container

- Make sure you have **docker** and **docker-compose** installed:

```bash
docker-compose up -d
```

- Before running the container, create your own database and put the database name into the *.env* files
