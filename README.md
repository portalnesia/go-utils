[![Go Reference](https://pkg.go.dev/badge/github.com/portalnesia/go-utils.svg)](https://pkg.go.dev/github.com/portalnesia/go-utils) ![Go](https://github.com/portalnesia/go-utils/actions/workflows/utils_test.yml/badge.svg)

# Go-Utils

Utility package for Internal Portalnesia

This package converted from [Javascript Version](https://github.com/portalnesia/portalnesia-utils) of Portalnesia Utils

## Install

```bash
go get go.portalnesia.com/utils
```

## Example

```go
package main

import (
  utils "go.portalnesia.com/utils"
  "fmt"
)

func main() {
  text := "Hello World"

  slug := utils.Slug(text)
  fmt.Printf("Slugify Format: %s",slug)
}
```

## Go References
[pkg.go.dev/go.portalnesia.com/utils](https://pkg.go.dev/go.portalnesia.com/utils)