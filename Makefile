test:	
	@echo "MAESTRIA SOFTWARE - Verifico dependencias "
	go mod tidy
	go mod vendor
	go mod verify
	@echo "MAESTRIA SOFTWARE - Creo y ejecuto compilacion"
	go build -o tp.exe
	.\tp.exe
	@echo "MAESTRIA SOFTWARE - Ejecuto pruebas"
	go test