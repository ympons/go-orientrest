orientrest - Orientdb client for Go
===========================

Package `orientrest` is a [Go](http://golang.org) client library providing access to
the [OrientDB](http://orientdb.com/) document/graph database via its REST API. Orientrest 
was inspired by [Neoism](https://github.com/jmcvetta/neoism).


# Requirements

[Go 1.1](http://golang.org/doc/go1.1) or later is required.

Tested against OrientDB Server v2.0.9.  


# Installation

```
go get -v github.com/ympons/go-orientrest
```


# Documentation
-

# Usage

## Init the client
There are some ways initialize the client

```go
client, err := orientrest.OrientDB("http://localhost:2480/", orientrest.Options{
	DbName: dbname,
	DbUser: user,
	DbPass: pass,
	Conn: true,
})
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
err := client->Command("create class Person extends V")
```
## Make a query

```go
result, err := client.Query("select * from V")
```

This is Free Software, released under the terms of the [GPLv3](http://www.gnu.org/copyleft/gpl.html).
