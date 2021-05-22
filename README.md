# Go Web Api

**WebApi REST created in GoLang**

This project was created in order to improve my knowledge in GoLang and REST API's

---

[![Go WebApi PR CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_pr_develop.yml/badge.svg?branch=develop)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_pr_develop.yml)
[![Go WebApi Push CI](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml/badge.svg?branch=develop)](https://github.com/ralvescosta/go_webapi/actions/workflows/ci_push_develop.yaml)


## Features

<p style='color:red'>Was Implemented</p> - <span class="text-gray">Work Now</span> - <span class="text-purple">Future feature</span> 

- <div class="text-green">User Management</div>
- <div class="text-gray">Session using JWT with asymmetric keys</div> 
- <div class="text-purple">Authentication Middleware</div> 
- <div class="text-purple">Authorization Middleware</div> 

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
