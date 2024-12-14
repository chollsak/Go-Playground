package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main(){
	id := uuid.New()

	fmt.Println("UUID:",id)
	fmt.Println("Hello World 3")
}