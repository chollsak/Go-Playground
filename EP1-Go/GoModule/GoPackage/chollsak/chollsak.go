package chollsak

import (
	"fmt"
	"github.com/chollsak/go-example/chollsak/internal/chollsak"
)

func setName() string { // This function is private function
	const name string = "Hello, Chollsak!"
	return name
}

func SayHello() {
	fmt.Println(setName())
	SayTest()

	chollsak.SayCholl()
}
