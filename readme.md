# Go MemShare

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-paypal-blue)](https://paypal.me/shaynejrtaylor?country.x=US&locale.x=en_US)

A simple go module for sharing memory pointers as encoded strings.

This module not only converts a pointer to a string and back again, but also verifies if the receiving end is converting to a valid pointer and to the correct var type.

## Installation

```shell script
go get github.com/AspieSoft/go-memshare
```

## Usage

```go

import (
  "github.com/AspieSoft/go-memshare"
)

func main(){
  myVar := "hello, world"

  pointer := memshare.Stringify[string](&myVar)

  anotherInstance(pointer)

  fmt.Println(myVar) // "hello, there"
}

func anotherInstance(pointer string){
  myVar := memshare.Parse[string](pointer)
  if myVar == nil {
    panic("pointer was invalid or not converted to the correct type")
  }

  fmt.Println(*myVar) // "hello, world"

  *myVar = "hello, there"
}

```
