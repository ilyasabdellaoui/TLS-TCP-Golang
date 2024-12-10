# TLS-TCP Golang Project
This repo is a simple implementation of TLS over TCP in Golang.

## Prerequisites

- Go 1.16 or higher
- OpenSSL (for generating certificates)

## Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/ilyasabdellaoui/tls-tcp-golang.git
    cd tls-tcp-golang
    ```

2. **Generate TLS certificates:**
    - Generate a private key:
    ```sh
        openssl genrsa -out server.key 2048
    ```
    - Generate a certificate signing request:
    ```sh
        openssl req -new -key server.key -out server.csr
    ```
    - Generate a self-signed certificate:
    ```sh
        openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

3. **Run the server:**
    ```sh
    go run server.go
    ```

4. **Run the client:**
    ```sh
    go run client.go
    ```

## Project Structure

- `server.go`: Contains the server implementation.
- `client.go`: Contains the client implementation.
- `server.crt` and `server.key`: TLS certificates.

## How It Works

1. The server listens for incoming TCP connections and uses TLS for secure communication.
2. The client connects to the server using TLS and exchanges messages securely.

## Notes

- Ensure that the `server.crt` and `server.key` files are in the same directory as `server.go` and `client.go`.
- Modify the server address and port in the code if needed.