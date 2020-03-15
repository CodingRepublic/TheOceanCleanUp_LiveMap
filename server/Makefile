GREEN        := $(shell tput -Txterm setaf 2)
BOLD         := $(shell tput bold)

RESET := $(shell tput -Txterm sgr0)

dev: clean env-up build
	APP_ENV=local ./server

build:
	go build

cleanbuild:
	@rm -f server

env-up:
	@echo "${GREEN}${BOLD}Starting development environement${RESET}"
	@cp -r config/dev/dataSaved config/dev/data
	@cd config/dev/ && docker-compose up -d
	@echo "${GREEN}${BOLD}Environment up${RESET}"

clean: cleanbuild
	@echo "${GREEN}${BOLD}Shuting development environement${RESET}"
	@cd config/dev/ && docker-compose down -v
	@rm -rf config/dev/data
	@echo "${GREEN}${BOLD}Environement down${RESET}"
