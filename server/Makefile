GREEN        := $(shell tput -Txterm setaf 2)
BOLD         := $(shell tput bold)

RESET := $(shell tput -Txterm sgr0)

dev: clean env-up build
	APP_ENV=local ./app

build:
	go build -o app

cleanbuild:
	@rm -f app

env-up:
	@echo "${GREEN}${BOLD}Starting development environment${RESET}"
	@cp -r config/dev/basicData config/dev/data
	@cd config/dev/ && docker-compose up -d
	@echo "${GREEN}${BOLD}Environment up${RESET}"

clean: cleanbuild
	@echo "${GREEN}${BOLD}Shutting development environment${RESET}"
	@cd config/dev/ && docker-compose down -v
	@rm -rf config/dev/data
	@echo "${GREEN}${BOLD}Environment down${RESET}"
