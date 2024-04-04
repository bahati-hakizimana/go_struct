package main

import "fmt"

func main() {

	fmt.Println("Welcome to Struct Phase")

	bahati := User{"bahati", "bahati@gmail.com", true, 24}

    fmt.Println(bahati)
	fmt.Printf("my details are: %+v\n", bahati)

}



type User struct {
	Name string
	Email string
	Status bool
	Age int
}