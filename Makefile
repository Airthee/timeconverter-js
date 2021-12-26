MODULE_NAME=timeconverter
SOURCE_FILE=./${MODULE_NAME}.go
WASM_OUTPUT_DIRECTORY=./build/wasm
WASM_OUTPUT_FILE=${WASM_OUTPUT_DIRECTORY}/${MODULE_NAME}.wasm
GO_COMMAND=GOOS=js GOARCH=wasm go

install:
	$(GO_COMMAND) get ./...

build: install
	rm -rf ${WASM_OUTPUT_DIRECTORY}
	mkdir -p ${WASM_OUTPUT_DIRECTORY}
	$(GO_COMMAND) build -o ${WASM_OUTPUT_FILE} ${SOURCE_FILE}

example-update: build
	cp -f ${WASM_OUTPUT_FILE} ./example/dist

example-serve: example-update
	cd ./example && yarn serve