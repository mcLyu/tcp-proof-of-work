build:
	docker-compose build --no-cache

start:
	docker-compose up

start_ddos:
	docker-compose up --scale client=$(n)