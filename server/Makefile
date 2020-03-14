GREEN        := $(shell tput -Txterm setaf 2)
BOLD         := $(shell tput bold)

RESET := $(shell tput -Txterm sgr0)

all:

build:
	go build

run: build
	APP_ENV=local ./server

env-up:
	@echo "${GREEN}${BOLD}Starting development environement${RESET}"
	@cd config/local/ && docker-compose up -d
	@echo "${GREEN}${BOLD}Environment up${RESET}"

clean:
	@echo "${GREEN}${BOLD}Shuting development environement${RESET}"
	@cd config/local/ && docker-compose down
	@echo "${GREEN}${BOLD}Environement down${RESET}"
