# mTLS-TCP Golang Demo
This repo is a simple implementation of mTLS over TCP in Golang.

## Prerequisites

- Go 1.16 or higher
- OpenSSL (for generating certificates)

## Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/ilyasabdellaoui/tls-tcp-golang.git
    cd tls-tcp-golang
    ```

2. **Generate mTLS certificates:**
    - Generate the private key for the CA
        ```sh
        openssl genpkey -algorithm RSA -out ca.key -aes256
        ```
    - Generate the root certificate
        ```sh
        openssl req -x509 -new -key ca.key -out ca.crt -days 3650 -subj "/C=MA/ST=RABAT/L=RABAT/O=ENSIAS"
        ```
    - Generate the private key for the server
        ```sh
        openssl genpkey -algorithm RSA -out server.key
        ```
    - Generate the CSR for the server using the custom config file (including SAN)
        ```sh
        openssl req -new -key server.key -out server.csr -config openssl.cnf
        ```
    - Sign the server certificate with the CA
        ```sh
        openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extensions v3_ca -extfile openssl.cnf
        ```
    - Generate the private key for the client
        ```sh
        openssl genpkey -algorithm RSA -out client.key
        ```
    - Generate the CSR for the client
        ```sh
        openssl req -new -key client.key -out client.csr -subj "/C=MA/ST=RABAT/L=RABAT/O=ENSIAS/CN=client" -config openssl.cnf
        ```
    - Sign the client certificate with the CA
        ```sh
        openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extensions v3_ca -extfile openssl.cnf
        ```

3. **Run the server:**
The server requires both the server certificate and key, as well as the CA certificate to validate client certificates:
    ```sh
    go run server.go
    ```

4. **Run the client:**
The client requires its own certificate, key, and the CA certificate to verify the server's certificate. It will connect securely to the server:
    ```sh
    go run client.go
    ```

## Project Structure

- `server.go`: Contains the server implementation with mTLS configuration.
- `client.go`: Contains the client implementation with mTLS configuration.
- `server.crt` and `server.key`: Server's TLS certificates.
- `client.crt` and `client.key`: Client's TLS certificates.
- `ca.crt`: Certificate Authority (CA) certificate used for mutual authentication.

## How It Works

1. Server Side:
- The server listens for incoming TCP connections.
- It requires the client to present a valid certificate during the TLS handshake. The server uses the CA certificate to verify the client's certificate.
- The server also presents its own certificate to the client for mutual authentication.

2. Client Side:
- The client connects to the server and presents its certificate during the TLS handshake.
- The client uses the CA certificate to verify the server's certificate and establish a secure connection.

## Notes

- Ensure that the `server.crt`, `server.key`, `client.crt`, `client.key`, and `ca.crt` files are placed correctly in the same directory as `server.go` and `client.go` for ease of access.
- Modify the server address and port in the code if needed.