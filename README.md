# Go-Utils

Utility package for Internal Portalnesia

This package converted from [Javascript Version](https://github.com/portalnesia/portalnesia/tree/main/packages/utils) of Portalnesia Utils

## Install

```bash
go get github.com/portalnesia/go-utils
```

## Example

```go
package main

import (
  utils "github.com/portalnesia/go-utils"
  "fmt"
)

func main() {
  text := "Hello World"

  slug := utils.Slug(text)
  fmt.Printf("Slugify Format: %s",slug)
}
```

## Go References
[pkg.go.dev/github.com/portalnesia/go-utils](https://pkg.go.dev/github.com/portalnesia/go-utils)