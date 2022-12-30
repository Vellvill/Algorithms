new:
	@go run ./base/naming/main.go

ENABLE ?= $(shell bash -c 'read -p "enable url? y or n"')
