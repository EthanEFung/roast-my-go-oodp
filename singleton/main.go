package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {}

var instance *single

func getInstance() *single {
  if instance == nil {
    lock.Lock()
    defer lock.Unlock()
  }
  if instance == nil {
    /*
      Why would we need to check again if the instance is nil?
      This is because multiple routines could have seen that
      the instance was not created, and be simultaneously trying
      to create the instance. This ensures that only the first
      routine that has successfully executed the lock will
      create the instance.
    */
    fmt.Println("creating instance")
    instance = &single{}
    return instance
  } 
  fmt.Println("Single instance created")
  return instance
}

func main() {
  for i := 0; i < 30; i++ {
    go getInstance()
  }
  fmt.Scanln()
}

