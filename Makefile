run:
	go run cmd/aoj/main.go
build:
	docker build -t apiq .
dockerrun:
	docker run --name=apiq -p 8088:8088 --rm apiq

.PHONY: run build dockerrun