
APPLICATION=app1
.PHONY=build runlocal deploy

build:
	go build -o bin/$(APPLICATION)

deploy:
	@echo Deploying to Azure

runlocal:
	func start --prefix az-functions


