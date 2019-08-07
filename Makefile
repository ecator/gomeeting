all: clean
	-mkdir bin
	go build -o bin/gomeeting -ldflags "-w" main.go
clean:
	-rm -rf bin