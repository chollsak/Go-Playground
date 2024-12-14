package main

import (
	"fmt"

	"github.com/chollsak/go-example/chollsak"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Println("UUID:", id)

	chollsak.SayTest()
	chollsak.SayHello()
}
