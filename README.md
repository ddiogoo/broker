# broker

Rewriting an academic project to understand resilient, critical and scalable distributed systems.

## About

Project initially created at college, which will now have more services using several different languages.

## Microservices details

Application is divided as follows:

- `socket-server`: Service responsible for notifying the user in real time.
- `key-manager`: Service responsible for managing access permission keys to services.
- `auth-manager`: Service responsible for managing the authentication and authorization.
- `traffic-exchange`: Service responsible for buy or sell a specific asset made available by a broker.
- `transaction-service`: Service responsible for carrying out transactions through a purchase order and another sales order.
