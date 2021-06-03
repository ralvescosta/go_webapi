# Go Web Api

**WebApi REST created in GoLang**

This project was created in order to improve my knowledge in GoLang and REST API's

---

[![Go WebApi Develop CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml/badge.svg?branch=develop)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml)
[![Go WebApi Main CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_main.yaml/badge.svg?branch=main)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_main.yaml)


## Features

**LEGEND**

📝 Was Implemented

🧑‍💻 Work Now

💻 Future feature 

- 📝 User Management
- 📝 Session using JWT with asymmetric keys
- 📝 Basic Elastic APM Configuration
- 🧑‍💻 Elastic Metricbeat and Filebeat
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

### Create containers

- *PostgreSQL*
- *Elasticsearch*
- *Kibana*
- *Elastic APM*

- Make sure you have **docker** and **docker-compose** installed:

```bash
docker-compose up -d
```

- Before running the container, create your own database and put the database name into the *.env* files

### Run Migration

```bash
make migrate
```

## Articles

- [1] - [GoLang WebApi - pt-BR](https://ralvescosta.medium.com/como-estruturar-webapi-em-golang-e2a41502d809)