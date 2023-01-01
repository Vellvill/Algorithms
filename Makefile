new:
	@go run ./base/core/naming/main.go

ENABLE ?= $(shell bash -c 'read -p "enable url? y or n"')
