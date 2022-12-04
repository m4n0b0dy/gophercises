BINARY_NAME=MarkDown.app
APP_NAME=MarkDown
VERSION=1.0.0

## build: build binary and package app
## rm $(addsuffix .exe, ${APP_NAME})
build:
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release  --appID io.fyne.terminal

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...