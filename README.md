# Personalib API

## Description

RESTful API Web Server for simple personal library using Echo framework, GORM and PostgreSQL

## Installation

* Clone project:
```shell
git clone https://github.com/suksest/personalib-api.git
```

* Change current working directory:
```shell
cd personalib-api
```
* Get external libraries

```shell
go get "github.com/lib/pq"
go get "github.com/jinzhu/gorm"
go get "github.com/labstack/echo"
go get "github.com/garyburd/redigo/redis"
```

* Change environment settings

Change `host`, `user`, `password`, `dbname` file in `db.go`

* Build binaries:
``` 
go build
```

* run the server:

```
./personalib-api (linux & Mac)
personalib-api.exe (windows) 
```

## API Documentation

https://documenter.getpostman.com/view/1513031/RWTeWhw4