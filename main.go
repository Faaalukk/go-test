package main

import (
	"fmt"
	// "github.com/Faaalukk/go-test/customer"
	"github.com/google/uuid"
)

func init() {
	fmt.Println("test ... 30/05/2002")
}
func main() {
	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}
