test:
#ifeq ($(DB_TYPE),windows)
	
	@echo "SPB - Creo y ejecuto compilacion"
	go build -o tp.exe
#timeout /T 3 /NOBREAK
	.\tp.exe