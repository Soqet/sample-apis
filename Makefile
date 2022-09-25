ifeq ($(OS),Windows_NT)
    SHELL := powershell.exe #change shell for windows
    .SHELLFLAGS := -Command
    ending := exe
else
    ending := out
endif

cmdpath = $(mod)/cmd
files = $(cmdpath)/main.go


run:
	go run $(files)

build:
	go build -o ./builds/$(mod)/main.$(ending) $(files)

fmt: 
	go fmt $(cmdpath) $(mod)/internal/db $(mod)/internal/api
