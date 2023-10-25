package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    once.Do(func() {
        singleInstance = &single{}
        fmt.Println("Creating single instance now.")
    })

    return singleInstance
}

func main() {
    for i := 0; i < 30; i++ {
        go getInstance()
    }

    fmt.Scanln()
}