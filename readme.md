# Go MemShare

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-paypal-blue)](https://paypal.me/shaynejrtaylor?country.x=US&locale.x=en_US)

A simple go module for sharing memory pointers as encoded strings.

This module not only converts a pointer to a string and back again, but also verifies if the receiving end is converting to a valid pointer and to the correct var type.

You may optionally encrypt the stringified pointer to keep the memory address safe (uses aes-cfb encryption method).

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

  // get a new stringified pointer (also pushes the pointer to persistant memory)
  pointer := memshare.Stringify[string](&myVar)

  // optional: encrypt the pointer string with a key
  encryptedPointer := memshare.Stringify[string](&myVar, "myKey123")

  anotherInstance(pointer, encryptedPointer)

  fmt.Println(myVar) // "hello, there"

  // allow the garbage collector to remove this pointer automatically (removes it from the persistant list)
  memshare.Delete(&myVar)
}

func anotherInstance(pointer string, encryptedPointer string){
  // parse the stringified pointer to a new var (note: the garbage collector may ignore this var)
  myVar := memshare.Parse[string](pointer)
  if myVar == nil {
    panic("pointer was invalid or not converted to the correct type")
  }

  // you will need to pass your encryption key if you set one
  myVar := memshare.Parse[string](encryptedPointer, "myKey123")
  if myVar == nil {
    panic("pointer may have also failed to decrypt")
  }

  fmt.Println(*myVar) // "hello, world"

  *myVar = "hello, there"
}

```
