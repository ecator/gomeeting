pack_file = GoMeeting-$(shell git tag -l | tail -n 1)-$(shell uname -s)-$(shell uname -p).tar.xz
all: clean
	-@mkdir bin
	go build -o bin/gomeeting -ldflags "-w" main.go
pack: all
	-@mkdir TEST
	tar cJf TEST/$(pack_file) assets bin script config.yml.sample
clean:
	-@rm -rf bin