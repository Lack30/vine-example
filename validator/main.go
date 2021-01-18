package main

import (
	"log"

	pb "github.com/lack-io/vine-example/validator/proto"
)

func main() {
 	p := pb.Person{}
	p.Age = 1
	p.Email = "11"
 	err := p.Validate()
 	log.Printf("%v\n", err)
}
