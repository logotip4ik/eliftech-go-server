# First version of server in go

Also is my first time creating something in go

## Getting started

1. Setup your database (i am using docker):

```shell
$ docker-compose up -d
```

2. Then just start up the server, if you have [go](https://golang.org/) installed:

```shell
$ go run main.go
```

## Tech used:

- [go](https://golang.org/) as backend language
- [gin](https://gin-gonic.com/) as web server
- [gorm](https://gorm.io/) as orm, for connecting to db (postgres in my case)

### Finally

As a JavaScript developer, this was a bit **painful**,
especially after using some of really awesome tech like
[sails.js](https://sailsjs.com/) or
[express.js](https://expressjs.com/) in connection with the best orm
(in my opinion) called [prisma](https://www.prisma.io/)
