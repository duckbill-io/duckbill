ContainerName := iblog

parse:
	go run cli.go
serve:
	go run server.go
logs:
	docker logs -f $(ContainerName)

up:
	container_name=$(ContainerName)  docker-compose -f docker/docker-compose.yml up -d	$(ContainerName) 
bash:
	docker exec -it $(ContainerName) /bin/bash
restart:
	container_name=$(ContainerName)  docker-compose -f docker/docker-compose.yml restart	$(ContainerName)
rm:
	docker rm -f $(ContainerName)

