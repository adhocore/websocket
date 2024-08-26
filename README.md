# Gorilla WebSocket

[![GoDoc](https://godoc.org/github.com/gorilla/websocket?status.svg)](https://godoc.org/github.com/gorilla/websocket)
[![CircleCI](https://circleci.com/gh/gorilla/websocket.svg?style=svg)](https://circleci.com/gh/gorilla/websocket)

Gorilla WebSocket is a [Go](http://golang.org/) implementation of the
[WebSocket](http://www.rfc-editor.org/rfc/rfc6455.txt) protocol.


### Documentation

* [API Reference](https://pkg.go.dev/github.com/gorilla/websocket?tab=doc)
* [Chat example](https://github.com/gorilla/websocket/tree/main/examples/chat)
* [Command example](https://github.com/gorilla/websocket/tree/main/examples/command)
* [Client and server example](https://github.com/gorilla/websocket/tree/main/examples/echo)
* [File watch example](https://github.com/gorilla/websocket/tree/main/examples/filewatch)

### Status

The Gorilla WebSocket package provides a complete and tested implementation of
the [WebSocket](http://www.rfc-editor.org/rfc/rfc6455.txt) protocol. The
package API is stable.

### Installation

```sh
go get github.com/gorilla/websocket
```

then in `go.mod`, add 

```
replace github.com/gorilla/websocket v1.5.1 => github.com/adhocore/websocket v1.5.1
```

then run

```sh
go mod tidy
```

### Usage

```go
import "github.com/gorilla/websocket"

buf := []byte
// netConn is net.Conn instance
conn := websocket.NewConn(netConn, true, 2048, 2048, nil, nil, buf)
```

### Protocol Compliance

The Gorilla WebSocket package passes the server tests in the [Autobahn Test
Suite](https://github.com/crossbario/autobahn-testsuite) using the application in the [examples/autobahn
subdirectory](https://github.com/gorilla/websocket/tree/main/examples/autobahn).
