## Makefile simples e robusto para compilar o projeto principal e o submódulo
.PHONY: help build build-mestres build-all run clean test fmt vet tidy

BIN_DIR := bin
BINARY ?= mestre_do_mux
MASTERS_BIN := $(BIN_DIR)/mestres-do-mux

help:
	@printf "Makefile - targets disponíveis:\n"
	@printf "  make build          -> compila o binário principal ($(BIN_DIR)/$(BINARY))\n"
	@printf "  make build-mestres  -> compila participantes/mestres-do-mux ($(MASTERS_BIN))\n"
	@printf "  make build-all      -> compila ambos\n"
	@printf "  make run            -> compila e executa o binário principal\n"
	@printf "  make test           -> roda go test ./...\n"
	@printf "  make fmt            -> go fmt ./...\n"
	@printf "  make vet            -> go vet ./...\n"
	@printf "  make tidy           -> go mod tidy (root + participantes/mestres-do-mux)\n"
	@printf "  make clean          -> remove $(BIN_DIR)/\n"

# Compila o binário principal (assume main.go ou package main na raiz)
build:
	@echo "Building root binary -> $(BIN_DIR)/$(BINARY)"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY) ./ || (echo "go build falhou"; exit 1)

# Compila apenas o subprojeto participantes/mestres-do-mux
build-mestres:
	@echo "Building participantes/mestres-do-mux -> $(MASTERS_BIN)"
	@mkdir -p $(BIN_DIR)
	@cd participantes/mestres-do-mux && go build -o ../../$(MASTERS_BIN) ./ || (echo "go build (mestres) falhou"; exit 1)

build-all: build build-mestres

run: build
	@echo "Running $(BIN_DIR)/$(BINARY)"
	@$(BIN_DIR)/$(BINARY)

clean:
	@echo "Removing $(BIN_DIR)/"
	@rm -rf $(BIN_DIR)

test:
	@echo "Running tests"
	@go test ./...

fmt:
	@echo "Formatting code (go fmt)"
	@go fmt ./...

vet:
	@echo "Running go vet"
	@go vet ./...

tidy:
	@echo "Running go mod tidy (root and participantes/mestres-do-mux)"
	@go mod tidy
	@cd participantes/mestres-do-mux && go mod tidy
