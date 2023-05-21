check_env:
ifeq ($(origin WEBHOOK), undefined)
	@echo "WEBHOOK is not set."
else
	@echo "WEBHOOK is set to: $(WEBHOOK)"
endif

link: check_env
	@go run generate.go

build:
	@go build -o bin/webhook main.go

run:
	@go run main.go

deploy:
	@flyctl deploy