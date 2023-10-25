package main

import (
    "fmt"
    "sync"
)

var (
    singleInstance *single
    once           sync.Once
)

type single struct{}

func getInstance() *single {
    once.Do(func() {
        fmt.Println("Creating single instance now.")
        singleInstance = &single{}
    })

    if singleInstance != nil {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}

func main() {
    for i := 0; i < 30; i++ {
        go getInstance()
    }

    fmt.Scanln()
}