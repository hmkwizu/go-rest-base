# GO REST BASE

<p align="center">
    <a href="https://golang.org/doc/install" target="_blank">
        <img src="https://img.shields.io/badge/Language-Go-green.svg?style=flat" alt="Go">
    </a>
    <a href="https://golang.org/doc/install" target="_blank">
        <img src="https://img.shields.io/badge/Platforms-OS%20X%20%7C%20Linux%20%7C%20Windows-green.svg?style=flat" alt="Platforms OS X | Linux | Windows">
    </a>
    <a href="https://github.com/hmkwizu/go-rest-base/blob/master/LICENSE.md" target="_blank">
        <img src="https://img.shields.io/badge/License-MIT-green.svg?style=flat" alt="License MIT">
    </a>
</p>

This is a REST API written in go language. It is meant to be boilerplate code to use when you are starting a new REST API project in go.

## Overview

The code includes the following

* Routing (using [go-chi/chi](https://github.com/go-chi/chi))
* Database (using [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql))
* Authentication (using [go-chi/jwtauth](https://github.com/go-chi/jwtauth))
* Configuration (using [viper](https://github.com/spf13/viper))

## How to install
Before you continue, please make sure you have installed go on your system. You can install it from [HERE](https://golang.org/doc/install).


Then Go to your $GOPATH/go/src, usually $HOME/go/src,

```
cd
cd go/src
```

Then clone this repository

```
git clone https://github.com/hmkwizu/go-rest-base
```

Run this command

```
go get
go run main.go
```
You are done!

Now go to [http://localhost:8080/v1/api/test](http://localhost:8080/v1/api/test)

If you wish to build the executable and run it instead, run

```
go build main.go 
./main
```
## Setup
For easy deployment in a microservice, all configurations are done in environment variables.

You will need to export CONNECTION_STRING, before you can connect to your database

```
export CONNECTION_STRING="username:password@tcp(dbhost:dbport)/dbname"
```

Other variables you can export

PORT - Port for the REST API to listen (Default: 8080)

SIGN_KEY - Signing Key for JWT Tokens (Default: 123456789)

## How to use
The code is well commented to help you adapt it to your needs. Here is an explanation of some files you might need the most. The important files are in the folder api.

**repository.go** - this file keeps the code for database interactions. So anything to do with the database eg SELECT, UPDATE, DELETE etc should be in this file.

**controller.go** - receives calls from the router and does some application logic and/or calls the repository to get data from the database.

**routes.go** - this is where the actual routing happens. add your routes to this file. Each route calls a certain method in controller, so you might need to add the method in controller first.

**global.go** - all the global variables and initializations are in this file. If you want to access the Configuration struct you can use Config variable. 

example, Config.Port, Config.ConnectionString etc

Have fun coding!