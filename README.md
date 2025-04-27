# Sistema de Entrega (y simulacion de calificacion) de Tareas

Este proyecto simula un sistema donde los estudiantes pueden enviar tareas, y los profesores las reciben y califican. Utiliza gRPC para la comunicación entre cliente y servidor, RabbitMQ para la mensajería asincrónica, y Docker para la contenerización de los servicios.

### Características:
- Cliente gRPC: Permite a los estudiantes enviar tareas.

- Servidor gRPC: Recibe las tareas y las publica en una cola de RabbitMQ.

- Worker (Profesor): Consume las tareas desde RabbitMQ y las califica aleatoriamente.

- Docker: Conteneriza el servidor y el worker para facilitar la ejecución.

### Estructura del proyecto:

```
SD20251-advanced-homework-system/
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

# Instrucciones para ejecutar el proyecto

Es posible que requiera instalar rabbit en su sistema (Linux):
```
sudo apt install rabbitmq-server
```

1. Clonar el repositorio:
```
git clone https://github.com/mxidev/SD20251-advanced-homework-system.git
cd SD20251-advanced-homework-system
```

2. Compilar protocol buffers para client y server:
```
cd client/proto
protoc --go_out=. --go-grpc_out=. ./homework.proto

cd ../../server/proto
protoc --go_out=. --go-grpc_out=. ./homework.proto
```

3. Inicializar modulo Go en cada entidad
```
go mod init <nombre_entidad>
go mod tidy
```

Por ejemplo, para client:
```
cd client
go mod init client
go mod tidy
```

4. Ejecutar cada entidad por separado en distintas terminales:
```
make docker-client -> Levanta un contenedor Docker para el cliente
make docker-server -> Levanta un contenedor Docker para el servidor
make run-worker -> Ejecuta el codigo GO asociado a la entidad del worker
```