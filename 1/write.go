package main

import (
	"io/ioutil"
	"log"

	pb "./tutorial"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./person.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write person:", err)
	}
}
