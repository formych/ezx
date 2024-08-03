package main

import (
	"fmt"

	"github.com/fsm-xyz/ezx"
)

var C = Config{}

type Config struct {
	Hello string `json:"hello"`
}

func main() {
	e := ezx.New(&C)
	fmt.Println(C.Hello)
	e.Run()
}
