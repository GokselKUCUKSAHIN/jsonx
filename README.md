# jsonx

A high-performance JSON serialization and deserialization library built on top of `jsoniter`, offering efficient and flexible JSON handling in Go.

## Installation

```sh
go get github.com/GokselKUCUKSAHIN/jsonx
```

## Usage

### JSON Serialization and Deserialization

```go
import "github.com/GokselKUCUKSAHIN/jsonx"

// Struct to be serialized
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{Name: "John Doe", Email: "john@example.com"}

// Marshal to JSON
jsonData, err := jsonx.Marshal(user)
if err != nil {
    log.Fatal(err)
}

// Unmarshal from JSON
var parsedUser User
if err := jsonx.Unmarshal(jsonData, &parsedUser); err != nil {
    log.Fatal(err)
}
```

### Casting Helpers

```go
// Cast JSON response directly into a struct
rawData, err := fetchData()
user, err := jsonx.Cast[User](rawData, err)
if err != nil {
    log.Fatal(err)
}

// Cast JSON response into a slice of structs
rawList, err := fetchList()
users, err := jsonx.CastSlice[User](rawList, err)
if err != nil {
	log.Fatal(err)
}
```

## Features

- **Fast JSON encoding/decoding** powered by `jsoniter`
- **Convenient casting functions** for structured data
- **Minimal performance overhead** with efficient memory usage

