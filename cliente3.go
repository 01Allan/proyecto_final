package main

import (
	"fmt"
	"net"
	"encoding/gob"
)


type person struct {
	name string
	email []string // slices
}


func client(persona person) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {

	persona1 := person{name: "Allan Pineda", email: []string{"allanpinedag@gmail.com", "apinedag@unah.hn"}}

	go client(persona1)
	
	var input string
	fmt.Scanln(&input)
}


