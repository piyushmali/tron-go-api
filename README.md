# Tron API

This project provides a GoLang API to interact with the Tron blockchain, allowing users to fetch the balance and transactions of a given Tron address.

## Setup

### Tron Node

To set up a Tron node, you can use the Docker image provided by TronTools.

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

### GoLang API

1. **Clone the repository**:

   ```bash
   git clone <repository-url>
   cd tron-api
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Run the API**:

   ```bash
   go run main.go
   ```

### Endpoints

- **Get Balance**: Fetch the balance of a given Tron address.

  ```
  GET /balance/{address}
  ```

  Example:

  ```bash
  curl http://localhost:8080/balance/TF3bmiZU69FfDTc3sBv32v1MfYCnVpR4G9
  ```

- **Get Transactions**: Fetch all transactions of a given Tron address.

  ```
  GET /transactions/{address}
  ```

  Example:

  ```bash
  curl http://localhost:8080/transactions/TF3bmiZU69FfDTc3sBv32v1MfYCnVpR4G9
  ```

### Explanation of the Solution

The solution involves setting up a Tron node using Docker and creating a GoLang API to interact with the node. The API uses the `grpc` package to communicate with the Tron node and the `gin` framework to handle HTTP requests.

**Challenges Faced**:

- Understanding the Tron gRPC API: Familiarizing myself with the Tron gRPC API and the data structures used was initially challenging.
- Error handling: Ensuring that all possible errors were handled gracefully required careful consideration.

### Dependencies

- Docker
- GoLang
- `github.com/gin-gonic/gin`
- `google.golang.org/grpc`
- `github.com/tron-us/go-btfs-common/protos/protocol/api`

### License

This project is licensed under the MIT License.
