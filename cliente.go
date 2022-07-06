package main

import (
	"fmt"
	"net"
)

func client()  {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// m1:  mensaje 1
	m1 := "Te amo mucho Lucía, aunque ya no estemos juntos, y no podas leer esto :("
	fmt.Println(m1)
	c.Write([]byte(m1))
	c.Close()
}


func main() {
	go client()
	
	var input string
	fmt.Scanln(&input)
}


