.PHONY: html
html:
	GOOS=js GOARCH=wasm go build -o html/asteroids.wasm

