package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

func client()  {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// m1:  mensaje 1
	m1 := "debo estar bien, por mi bien"
	fmt.Println(m1)
	err = gob.NewEncoder(c).Encode(m1)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go client()
	
	var input string
	fmt.Scanln(&input)
}


