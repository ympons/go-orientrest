[![Build Status](https://travis-ci.org/ympons/go-orientrest.png?branch=master)](https://travis-ci.org/ympons/go-orientrest)
orientrest - Orientdb client for Go
===========================

![Orientrest Logo](https://raw.github.com/ympons/go-orientrest/master/logo/orientrest.png)

Package `orientrest` is a [Go](http://golang.org) client library providing access to
the [OrientDB](http://orientdb.com/) document/graph database via its REST API.  


# Requirements

[Go 1.4](http://golang.org/doc/go1.3) or later is required.

# Installation

```
go get -v github.com/ympons/go-orientrest
```

# Usage

## Init the client

```go
client, err := orientrest.New("http://localhost:2480/")
```

```go
db, err := orientrest.OrientDB("http://localhost:2480/")
client, err := db.Connect(orientrest.Options{
	DbName: dbname,
	DbUser: user,
	DbPass: pass,
})
```

```go
db, err := orientrest.OrientDB("http://localhost:2480/")
db.Configure(orientrest.Options{
	DbName: dbname,
	DbUser: user,
	DbPass: pass,
})
client, err := orientrest.Connect()
```

## Create Database 

```go
err := client.DbCreate("my_new_database", orientrest.STORAGE_TYPE_PLOCAL, orientrest.DB_TYPE_GRAPH)
```

## Drop Database

```go
err := client.DbDrop(client.Name)
```

## Get the list of databases

```go
list, err := client.DbList()
```

## Get the available languages

```go
langs, err := client.DbAvailableLangs(client.Name)
```
## Close Database

```go
client.Close()
```
## Send a command

```go
_, err := client.Command("create class Person extends V")
```
## Make a query

```go
result, err := client.Query("select * from V")
```

# License

MIT License, see [LICENSE.md](./LICENSE).
