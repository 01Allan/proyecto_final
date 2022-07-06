package main 

import (
	"fmt"
	"net"
	"encoding/gob"
)

// estructura que se va enviar

type person struct {
	name string
	email []string // slices
}

// creando una funcion para el servidor

func server() {
	// definiendo el protocolo y el puerto a utilizar.
	// ins: instancia del servidor
	// err: error
	ins, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// c: cliente y se acepta lo que el cliente va a enviar.
		c, err := ins.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Por cada cliente que se conecte se ejecuta esta concurrencia:
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var persona person 
	err := gob.NewDecoder(c).Decode(&persona)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Texto:", persona)
	}

}

func main() {
	// ejecutando el servidor con una concurrencia que reciba datos: 
	go server()

	var input string 
	fmt.Scanln(&input)

}