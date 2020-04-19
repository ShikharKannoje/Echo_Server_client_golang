package main

import (
	
	"io"
	"log"
	"net"
	"os"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err !=nil{
		log.Print(err)
	}
	defer conn.Close()
	go mustcopy(os.Stdout, conn)
	mustcopy(conn, os.Stdin)
}

func mustcopy(dst io.Writer, src io.Reader){
	_, err := io.Copy(dst, src)
	if err != nil{
		log.Print(err)
	}
}