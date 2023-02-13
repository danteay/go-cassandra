# go-cassandra

**This repo is on development**

This is a simple Query Builder for Apache Cassandra database. It uses `gocql` driver and `goclqx` Query Builder 
internally to create friendly layer to interact with the database trough struct bindings.

## Installing

Simple run next command to download the latest version.

```shell script
go get github.com/danteay/go-cassandra@latest
```

You can create a client with this:

```go
package main

import gocassandra "github.com/danteay/go-cassandra"

func main() {
    client, err := gocassandra.NewClient(gocassandra.DefaultConfig())
    if err != nil {
        panic(err)
    }

    // Do stuff...    

    client.Close()
}
```

