![Golang](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) ![Gin](https://img.shields.io/badge/Gin-008ECF.svg?logo=Gin&logoColor=white) ![Nats.io](https://img.shields.io/badge/NATS.io-27AAE1.svg?logo=natsdotio&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1.svg?logo=PostgreSQL&logoColor=white)

# broker

Rewriting an academic project to understand resilient, critical and scalable distributed systems.

## About

Project initially created at college, which will now have more services using several different languages.

## Services

Application is divided as follows:

- `socket-server`: Service responsible for notifying the user in real time.
- `key-manager`: Service responsible for managing access permission keys to services.
- `auth-manager`: Service responsible for managing the authentication and authorization.
- `traffic-exchange`: Service responsible for buy or sell a specific asset made available by a broker.
- `transaction-service`: Service responsible for carrying out transactions through a purchase order and another sales order.

## Socket Server

### Description

This service is responsible for launching a WebSocket server that communicates with NATS.

### Communication flow

When a client makes a connection to the WebSocket server, it subscribes to a subject (queue) that will read data from it, the read data will be sent to the client in the Frontend.

![Fluxo de comunicação Socket Server](./.github/socket-server.png)

Note that communication goes from NATS to the server and from the server to the client, NATS will consume the data that will follow this flow from other services that will be the Producer, such as `traffic-exchange` and `transaction-service`.

## Reference

- [Golang](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [NATS.io](https://nats.io/)
- [PostgreSQL](https://www.postgresql.org/)
