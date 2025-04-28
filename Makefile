# Dockerizar server
docker-server: 
	sudo docker-compose up --build server 

# Dockerizar worker
docker-worker:
	sudo docker-compose up --build worker

# Dockerizar client
docker-client:
	sudo docker-compose up --build client

# Parar todo
docker-turnoff:
	@echo "🛑 Parando toda la infraestructura..."
	sudo docker-compose down