# uds

This repo contains go projects common utilities.

## Download and Import

```shell
go get github.com/HDS3991/uds
```

```go
package main

import (
    "fmt"

    "github.com/HDS3991/uds/utils"
)

func main() {
    list := []int{1,2,3,4,4,4,4,5}
    uniqueList := utils.UniqueArray(list)
    fmt.Println(uniqueList) // [1,2,3,4,5]
}
```
