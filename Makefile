.PHONY: html
html:
	GOOS=js GOARCH=wasm go build -ldflags "-w" -o html/asteroids.wasm

