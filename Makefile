# Nome do executável
BINARY_NAME := tg

# Caminho de instalação
INSTALL_PATH := /usr/local/bin

# Comando para compilar
build:
	@echo "Compilando o projeto..."
	go build -o $(BINARY_NAME)

# Comando para limpar os arquivos compilados
clean:
	@echo "Limpando os arquivos compilados..."
	rm -f $(BINARY_NAME)

# Comando para instalar o executável
install: build
	@echo "Instalando o executável..."
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)

# Comando para compilar para diferentes sistemas operacionais
build-linux:
	@echo "Compilando para Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)

build-windows:
	@echo "Compilando para Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe

build-mac:
	@echo "Compilando para macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)
