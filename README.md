
# GoMeeting [![Build Status](https://www.travis-ci.org/ecator/gomeeting.svg?branch=master)](https://www.travis-ci.org/ecator/gomeeting) ![GitHub last commit](https://img.shields.io/github/last-commit/ecator/gomeeting) ![GitHub language count](https://img.shields.io/github/languages/count/ecator/gomeeting) ![GitHub repo size](https://img.shields.io/github/repo-size/ecator/gomeeting) ![GitHub](https://img.shields.io/github/license/ecator/gomeeting) 
It's a very simple meeting room booking system written by go language for learning.

## Build
### Required
 - go >= 1.12 (must)
 - GNU Make (optional)
 
### Steps
```
go get go get github.com/ecator/gomeeting
cd ~/go/src/github.com/ecator/gomeeting
make
```
Now the `bin/gomeeting` is generated and you can run it.

If you don't have `make` command,you can run `mkdir bin && go build -o bin/gomeeting -ldflags "-w" main.go` to build it.It's the same to `make`.

## Usage
### Dependencies
- MySQL(MariaDB) (must)
  - Now it's only support mysql
- A modern browser (Chrome or Firefox is recommended)

### Start Server
If you are firstly running gomeeting you must install MySQL and create a database named `gomeeting`.

Creating the databse and tables are very simple, you just run this.
```
mysql -u user -p password < script/create.sql
```
The `user` and `password` is used to login to your local mysql service. It's recommended the database can be only accessed locally. And the tables `user` `org` `room` `meeting` will be also created.

To run it must have a configuration file named `config.yml`.There is a template and you can use it.
```
mv config.yml.sample config.yml
```
The file is also very simple because it's only related to the database now.

Like this !

```
db:
  host: "localhost"
  port: 3306
  user: "user"
  password: "password"
```

It will use that to connect to mysql so be careful.

Now you can run `bin/gomeeting -a 0.0.0.0` to start the server.The `-a 0.0.0.0` option makes sure you can access this service from remote,otherwise it only listens on local.You can run `bin/gomeeting -h` for more details.

```
Usage of bin/gomeeting:
  -a string
    	The listen address (default "localhost")
  -c string
    	The config file (default "config.yml")
  -f string
    	The assets folder includes html/css/js (default "assets")
  -h	Show usage
  -l string
    	The log file (default "server.log")
  -p uint
    	The listen port (default 7728)
```
### Administrator
Now there is no ui for managing the service,but you can use RESTful APIs directly within the super token.

Like this.

```
curl -X GET --cookie='auth=xxxx' localhost:7728/api/user/1000
```

The above will get the information of the user whose id is 1000.

> The super token will displayed when server is started successfully, so it will change everytime.

There are same scripts you can use in the `script` folder.

```
script/user.sh add xxx
script/user.sh del xxx
script/user.sh mod xxx
script/user.sh get xxx

script/room.sh add xxx
script/room.sh del xxx
script/room.sh mod xxx
script/room.sh get xxx

script/org.sh add xxx
script/org.sh del xxx
script/org.sh mod xxx
script/org.sh get xxx
```

### Using

After you add same `user` `org` `room`, open `http://server_ip:7728`  and have fun!

---
# Reference

- [httprouter](https://github.com/julienschmidt/httprouter)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [layDate](https://www.layui.com/laydate/)
- [YAML for Go](https://github.com/go-yaml/yaml)
- [axios](https://github.com/axios/axios)
- [js-md5](https://github.com/emn178/js-md5)
- [Vue](https://github.com/vuejs/vue)
- [Font Awesome](https://fontawesome.com/)
- [Bulma](https://bulma.io/)
- [🐱 HTTP Cats API ](https://github.com/httpcats/http.cat)
