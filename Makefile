all:
	-mkdir bin
	go build -o bin/meeting -ldflags "-w" main.go
clean:
	-rm -rf bin