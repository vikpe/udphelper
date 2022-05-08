# UDPH

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

Respond to any request with given response.

```go
response := []byte("pong")
udphelper.New(":8000").Respond(response)
```
