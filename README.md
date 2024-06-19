Sure! Here's a comprehensive and polished README file for your GitHub repository [tron-go-api](https://github.com/piyushmali/tron-go-api):

---

# Tron Go API

A GoLang API for fetching the balance and transactions of a given Tron address, with a full Tron node setup using Docker.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [Tron Node Setup](#tron-node-setup)
  - [GoLang API Setup](#golang-api-setup)
- [API Endpoints](#api-endpoints)
- [Explanation of the Solution](#explanation-of-the-solution)
- [Challenges Faced](#challenges-faced)
- [Dependencies](#dependencies)
- [License](#license)

## Introduction

This project provides a GoLang-based RESTful API that interacts with the Tron blockchain to retrieve the balance and transaction history of a specified Tron address. The project includes setting up a Tron node using Docker and implementing the API using the Gin framework and gRPC for communication with the Tron node.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Docker
- GoLang

## Setup

### Tron Node Setup

1. **Install Docker**:

    ```bash
    sudo apt-get update
    sudo apt-get install -y docker.io
    ```

2. **Run Tron Node**:

    ```bash
    docker run -it -p 8090:8090 -p 50051:50051 --name tron --rm trontools/quickstart
    ```

3. **Verify Node Connection**:

    ```bash
    docker logs -f tron
    ```

### GoLang API Setup

1. **Clone the repository**:

    ```bash
    git clone https://github.com/piyushmali/tron-go-api.git
    cd tron-go-api
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Run the API**:

    ```bash
    go run main.go
    ```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Get Balance

Fetch the balance of a given Tron address.

```
GET /balance/{address}
```

Example:

```bash
curl http://localhost:8080/balance/TF3bmiZU69FfDTc3sBv32v1MfYCnVpR4G9
```

### Get Transactions

Fetch all transactions of a given Tron address.

```
GET /transactions/{address}
```

Example:

```bash
curl http://localhost:8080/transactions/TF3bmiZU69FfDTc3sBv32v1MfYCnVpR4G9
```

## Explanation of the Solution

The solution involves:

1. **Tron Node Setup**: Using Docker to set up a full Tron node with the official `trontools/quickstart` image.
2. **GoLang API Implementation**:
   - **Gin Framework**: Used for creating the RESTful API.
   - **gRPC**: Used to communicate with the Tron node.
   - **Tron Client**: Implements methods to fetch balance and transactions from the Tron node.
   - **API Handlers**: Define endpoints to serve the balance and transactions data.

## Challenges Faced

1. **Understanding the Tron gRPC API**: Initial challenges in understanding the Tron gRPC API and its data structures were overcome through documentation and examples.
2. **Package Availability**: Identifying the correct gRPC package for Tron took some research. The package `github.com/tron-us/go-btfs-common/protos/protocol/api` was used.
3. **Error Handling**: Ensuring robust error handling for various scenarios, including node downtime and invalid addresses.

## Dependencies

- Docker
- GoLang
- `github.com/gin-gonic/gin`
- `google.golang.org/grpc`
- `github.com/tron-us/go-btfs-common/protos/protocol/api`

## License

This project is licensed under the MIT License.

---

This README file should give users a clear understanding of the project, how to set it up, and how to use it. It also provides insights into the challenges faced during development and the dependencies required for the project.
