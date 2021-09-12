build:
	setx GOARCH wasm
	setx GOOS js
	go build -o web/app.wasm
	go build

run: build
	./hello