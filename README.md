
# GoMeeting ![build badge](https://github.com/ecator/gomeeting/actions/workflows/build.yml/badge.svg) ![GitHub last commit](https://img.shields.io/github/last-commit/ecator/gomeeting) ![GitHub language count](https://img.shields.io/github/languages/count/ecator/gomeeting) ![GitHub repo size](https://img.shields.io/github/repo-size/ecator/gomeeting) ![GitHub](https://img.shields.io/github/license/ecator/gomeeting) 
It's a very simple meeting room booking system written by go language for learning.

## Build
### Required
 - go >= 1.18 (must)
 - curl (must)
 - unzip (must)
 - GNU Make (must)
 
### Steps
```shell
go get github.com/ecator/gomeeting
cd ~/go/src/github.com/ecator/gomeeting
make
```
Now the `bin/gomeeting` is generated and you can run it.

## Usage
### Dependencies
- MySQL(MariaDB) (must)
  - Now it's only support mysql
- A modern browser (must Chrome or Firefox)

### Start Server
If you are firstly running gomeeting you must install MySQL and create a database named `gomeeting`.

Creating the databse and tables are very simple, you just run this.
```shell
mysql -u user -p password < script/create.sql
```
The `user` and `password` is used to login to your local mysql service. It's recommended the database can be only accessed locally. And the tables `user` `org` `room` `meeting` will be also created.

To run it must have a configuration file named `config.yml`.There is a template and you can use it.
```shell
mv config.yml.sample config.yml
```
The file is also very simple because it's only related to the database and ldap and teams.

Like this !

```yaml
db:
  host: "localhost"
  port: 3306
  user: "user"
  password: "password"
ldap:
  enable: false
  placeholder:
    username: "AD account"
    password: "AD password"
  addr: "127.0.0.1:389"
  baseDN: "cn=users,dc=test,dc=local"
  level: 10
  orgID: 100
  attrMapKey:
    name: "displayName"
    email: "mail"
teams:
  enable: false
  entrypoint: "https://teams.microsoft.com/l/meeting/new"
```

It will use that to connect to mysql so be careful.

If you want to use ldap, set the `ldap.enable` to `true`. It can't support TLS now. The `ldap.placeholder` will be shown in login page if `ldap.enable` is `true`.

The `ldap.orgID` will insert into `user`table as `org_id`field to mark as a ldap user. So you should add a ldap org like this.

```sql
INSERT INTO org(id, name) VALUES (100,"ldap")
```

The `ldap.attrMapKey` can use below to decide which attribute mapping `name` and `email` of `user`table.

```shell
ldapsearch -x -b "cn=username,cn=users,dc=test,dc=local" -h 127.0.0.1 -p 389 -D "cn=username,cn=users,dc=test,dc=local" -w "password"
```

The teams icon will be shown if `teams.enable` is true and the `teams.entrypoint` will be used to decide how to open teams schedule.

Now you can run `bin/gomeeting -a 0.0.0.0` to start the server.The `-a 0.0.0.0` option makes sure you can access this service from remote,otherwise it only listens on local.You can run `bin/gomeeting -h` for more details.

```shell
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
The `script/create.sql` will add a root user for managing the service. So you can open `http://server_ip:7728` and login as root using `username:root;password:root` easily.

Otherwise there is another way to manage the service with no ui. You can use RESTful APIs directly within the super token.

Like this.

```
curl -X GET --cookie='auth=xxxx' localhost:7728/api/user/1000
```

The above will get the information of the user whose id is 1000.

> The super token will displayed when server is started successfully, so it will change everytime.

There are same scripts you can use in the `script` folder.

```shell
# You must export same environment variables
export GOMEETING_HOST=localhost:7728
export GOMEETING_TOKEN=xxxxx

# Add a new user
script/user.sh add 'username=test&password=123&level=10&org_id=1000&name=Martin&email=mail@example.com'
# Delete the user whose id is 1000
script/user.sh del 1000
# Modify the user's password whose is is 1000
script/user.sh mod 1000 'password=123'
# Get the information of user whose id is 1000
script/user.sh get 1000

# Add a new room named roo1
script/room.sh add name=room1
# Delete the room which id is 1000
script/room.sh del 1000
# Modify the room name of id 1000
script/room.sh mod 1000 name=room3
# Get the information of room which id is 1000
script/room.sh get 1000

# Add a new org named roo1
script/org.sh add name=org1
# Delete the org which id is 1000
script/org.sh del 1000
# Modify the org name of id 1000
script/org.sh mod 1000 name=org3
# Get the information of org which id is 1000
script/org.sh get 1000
```

### Using

After you add same `user` `org` `room`, open `http://server_ip:7728`  and have fun!

---
# Reference

- [httprouter](https://github.com/julienschmidt/httprouter)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [layDate](https://www.layui.com/laydate/)
- [YAML for Go](https://github.com/go-yaml/yaml)
- [go-ldap](https://github.com/go-ldap/ldap)
- [axios](https://github.com/axios/axios)
- [js-md5](https://github.com/emn178/js-md5)
- [Vue](https://github.com/vuejs/vue)
- [Font Awesome](https://fontawesome.com/)
- [Bulma](https://bulma.io/)
- [🐱 HTTP Cats API ](https://github.com/httpcats/http.cat)
- [markdown-js](https://github.com/evilstreak/markdown-js)
