# Makefile

# Path to your environment files
ENV_DEV=/Users/ekoagustina/Documents/Project/BE/go/Portfolio/env_go_ms_portfolio/.env.development
ENV_PROD=/Users/ekoagustina/Documents/Project/BE/go/Portfolio/env_go_ms_portfolio/.env.production

# Docker Compose files
DOCKER_COMPOSE_DEV=docker-compose -f docker-compose.development.yml --env-file $(ENV_DEV)
DOCKER_COMPOSE_PROD=docker-compose -f docker-compose.production.yml --env-file $(ENV_PROD)

.PHONY: dev prod

dev:
	$(DOCKER_COMPOSE_DEV) up -d --build

prod:
	$(DOCKER_COMPOSE_PROD) up -d --build

