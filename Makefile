.PHONY: all test dep compile build push checkenv deploy kubefile setup migrate

export GO111MODULE ?= on

all: compile start

mod:
	go mod download

test:
	go test -cover -coverprofile=cover.out ./...

compile: $(ODIR)
	@$(foreach svc, $(VAR_SERVICES), \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(svc)/$(svc) app/$(svc)/main.go;)

start:
	go run app/api/main.go

cover:
	go tool cover -html=cover.out

cover-html:
	go tool cover -html=cover.out -o cover.html
	open cover.html
