package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "./tutorial"
	"github.com/golang/protobuf/proto"
)

func main() {
	in, err := ioutil.ReadFile("./person.bin")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	p := &pb.Person{}
	if err := proto.Unmarshal(in, p); err != nil {
		log.Fatalln("Failed to parse person:", err)
	}

	fmt.Println(p)
}
