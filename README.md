<div align="center">
  <img src="assets/logo.png" alt="grd logo" width="200"/>

# grd

[![Go Reference](https://pkg.go.dev/badge/github.com/imadselka/grd.svg)](https://pkg.go.dev/github.com/imadselka/grd)
[![Go Report Card](https://goreportcard.com/badge/github.com/imadselka/grd)](https://goreportcard.com/report/github.com/imadselka/grd)
[![codecov](https://codecov.io/gh/imadselka/grd/branch/main/graph/badge.svg)](https://codecov.io/gh/imadselka/grd)
[![Build Status](https://github.com/imadselka/grd/workflows/Go/badge.svg)](https://github.com/imadselka/grd/actions)
[![Go Version](https://img.shields.io/github/go-mod/go-version/imadselka/grd)](https://github.com/imadselka/grd)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

</div>

**grd** is a lightweight Go library providing monadic-style error handling with chainable operations. It helps eliminate verbose error checks and improves code readability with composable pipelines.

## Features

* **🔗 Chainable Operations**: Write fluent error-handling pipelines
* **⚡ Zero Dependencies**: Pure Go, lightweight implementation
* **🔄 Generics Support**: Flexible, type-safe API
* **🛡️ Short-Circuit Logic**: Skips steps on failure automatically
* **🏁 Finally Blocks**: Guaranteed cleanup logic execution

## Installation

```bash
go get github.com/imadselka/grd
```

## Quick Start

```go
import "github.com/imadselka/grd"

result := grd.Try(func() (string, error) {
    return fetchUserData(userID)
}).Then(func(data string) (string, error) {
    return processUserData(data)
}).Then(func(processed string) (string, error) {
    return saveToDatabase(processed)
}).Catch(func(err error) string {
    log.Printf("Error occurred: %v", err)
    return "default_value"
})
```

## API Reference

### `Try[T](fn func() (T, error)) *TryResult[T]`

Initializes a chainable result from a function returning a value and error.

### `Then(fn func(T) (T, error)) *TryResult[T]`

Continues the chain on success; skipped if an error occurred.

### `Catch(fn func(error) T) T`

Handles any error that occurred and returns a fallback value.

### `Finally(fn func()) *TryResult[T]`

Runs regardless of success or failure. Common for logging or cleanup.

## Examples

### File Processing

```go
func ProcessConfig(path string) string {
    return grd.Try(func() ([]byte, error) {
        return os.ReadFile(path)
    }).Then(func(data []byte) (map[string]any, error) {
        var config map[string]any
        return config, json.Unmarshal(data, &config)
    }).Then(func(cfg map[string]any) (string, error) {
        return json.Marshal(cfg)
    }).Finally(func() {
        log.Println("Processing done")
    }).Catch(func(err error) string {
        log.Printf("Failed: %v", err)
        return "{}"
    })
}
```

### HTTP Client

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func GetUser(userID int) User {
    return grd.Try(func() (*http.Response, error) {
        return http.Get(fmt.Sprintf("/api/users/%d", userID))
    }).Then(func(resp *http.Response) ([]byte, error) {
        defer resp.Body.Close()
        if resp.StatusCode != 200 {
            return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
        }
        return io.ReadAll(resp.Body)
    }).Then(func(data []byte) (User, error) {
        var u User
        return u, json.Unmarshal(data, &u)
    }).Then(func(u User) (User, error) {
        u.Name = strings.Title(strings.ToLower(u.Name))
        return u, nil
    }).Finally(func() {
        metrics.IncrementAPICall("user")
    }).Catch(func(err error) User {
        log.Printf("Error: %v", err)
        return User{ID: -1, Name: "Unknown"}
    })
}
```

### DB Transaction

```go
func Transfer(from, to int, amount decimal.Decimal) bool {
    return grd.Try(func() (*sql.Tx, error) {
        return db.Begin()
    }).Then(func(tx *sql.Tx) (*sql.Tx, error) {
        _, err := tx.Exec("...")
        return tx, err
    }).Then(func(tx *sql.Tx) (*sql.Tx, error) {
        _, err := tx.Exec("...")
        return tx, err
    }).Then(func(tx *sql.Tx) (bool, error) {
        return true, tx.Commit()
    }).Finally(func() {
        auditLog.Record("tx")
    }).Catch(func(err error) bool {
        log.Printf("Failed: %v", err)
        return false
    })
}
```

## Traditional vs grd

### Without grd

```go
val, err := fetch()
if err != nil {
    return "", err
}
...
```

### With grd

```go
gr.Try(fetch).Then(...).Catch(...)
```

## Contribution

Open to PRs and ideas — feel free to contribute!

## License

MIT © [LICENSE](https://github.com/imadselka/grd/blob/main/LICENSE) [imadselka](https://github.com/imadselka)
