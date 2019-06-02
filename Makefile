ContainerName := iblog

bash:
	container_name=$(ContainerName)  docker-compose -f docker/docker-compose.yml up -d	$(ContainerName) 
	docker exec -it $(ContainerName) /bin/bash
depensure:
	rm -rf vendor
	dep ensure
rm:
	docker rm -f $(ContainerName)
parse:
	go run cli.go
