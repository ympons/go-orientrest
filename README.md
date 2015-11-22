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

## Connect to the server to manage databases

Authenticate to the server

```go
admin, err := client.Auth("user", "pass")
```

Close admin

```go
admin.Close()
```

Create Database

```go
_, err := admin.DbCreate("database_name", orientrest.DB_TYPE_GRAPH, orientrest.STORAGE_TYPE_PLOCAL)
```

Drop Database

```go
err := admin.DbDrop("database_name")
```

Get the list of databases

```go
list, err := admin.DbList()
```

Get the available languages

```go
langs, err := admin.DbAvailableLangs("database_name")
```

## Open a database

```go
db, err := client.Open("database_name", "dbuser", "dbpassword")
```

Send a command

```go
_, err := db.Command(NewCommandSQL("create class Person extends V"))
```

Make a query

```go
result, err := db.Command(NewQuerySQL("select * from V"))
```

Close database

```go
db.Close()
```

# License

MIT License, see [LICENSE.md](./LICENSE).
