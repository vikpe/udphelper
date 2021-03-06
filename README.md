# udphelper [![Go Reference](https://pkg.go.dev/badge/github.com/vikpe/udphelper.svg)](https://pkg.go.dev/github.com/vikpe/udphelper)  [![Test](https://github.com/vikpe/udphelper/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vikpe/udphelper/actions/workflows/test.yml) [![codecov](https://codecov.io/gh/vikpe/udphelper/branch/main/graph/badge.svg)](https://codecov.io/gh/vikpe/udphelper) [![Go Report Card](https://goreportcard.com/badge/github.com/vikpe/udphelper)](https://goreportcard.com/report/github.com/vikpe/udphelper)

> UDP test helper for Go

## Usage

### Listen

Listen to given address without responding to requests.

```go
udphelper.New(":8000").Listen()
```

### Echo requests

Respond to requests with `ok:` prepended to the request packet.

```go
udphelper.New(":8000").Echo()
```

### Respond

Respond to any request with given response(s).

**Single**
```go
response := []byte("pong")
udphelper.New(":8000").Respond(response)
```

**Multiple**
```go
response1 := []byte("pong1")
response2 := []byte("pong2")
udphelper.New(":8000").Respond(response1, response2)
```
