package main 

import (
	"fmt"
	"net"
	"encoding/gob"
)

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
	var msg1 string
	err := gob.NewDecoder(c).Decode(&msg1)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Texto:", msg1)
	}

}

func main() {
	// ejecutando el servidor con una concurrencia que reciba datos: 
	go server()

	var input string 
	fmt.Scanln(&input)

}