# Sistema de Entrega (y simulacion de calificacion) de Tareas

Este proyecto simula un sistema donde los estudiantes pueden enviar tareas, y los profesores las reciben y califican. Utiliza gRPC para la comunicación entre cliente y servidor, RabbitMQ para la mensajería asincrónica, y Docker para la contenerización de los servicios.

### Características:
- Cliente gRPC: Permite a los estudiantes enviar tareas.

- Servidor gRPC: Recibe las tareas y las publica en una cola de RabbitMQ.

- Worker (Profesor): Consume las tareas desde RabbitMQ y las califica aleatoriamente.

- Docker: Conteneriza el servidor y el worker para facilitar la ejecución.

### Estructura del proyecto:

```
homework-system/
├── docker-compose.yaml
├── Makefile
├── proto/
│   └── homework.proto
├── client/
│   └── main.go
├── server/
│   ├── Dockerfile
│   └── main.go
└── worker/
    ├── Dockerfile
    └── main.go
```
