package main

import (
	"net"
	"bufio"
	"fmt"
	"time"
	"flag"
	"github.com/lunny/log"
)

func readSocket(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		s, _ := reader.ReadString('\n')
		fmt.Print(s)
	}
}

func writeSocket(conn net.Conn) {
	for {
//		var message string
//		_, _ = fmt.Scanf("%s", &message)
		fmt.Fprintf(conn, "message from server \n")
		time.Sleep(2 * time.Second)
	}
}

func handleConnection(conn net.Conn) {
	go readSocket(conn)
	go writeSocket(conn)
}

func main() {
	log.SetFlags(log.Flags()|log.Lshortfile)

	var clientPairId string
	meetAddr := "localhost:8080"

	flag.StringVar(&meetAddr, "meetAddr", meetAddr, "Server for discover other")
	flag.StringVar(&clientPairId, "pair", "", "key to macth service")
	flag.Parse()

	log.Printf("Discover server is %s\n", meetAddr)
	log.Printf("unique ID for the client pair: %s\n", clientPairId)
	if len(clientPairId) == 0{
		log.Fatal("Must config the pair")
	}

	firstConn, err := net.Dial("tcp", meetAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(firstConn, "GET /meet/"+clientPairId+" HTTP/1.0\r\n\r\n")
	firstConn.Close()

	// Block to let the OS free the port.
	time.Sleep(5 * time.Second)

	log.Println("firstConn.LocalAddr", firstConn.LocalAddr().String())
	ln, err := net.Listen("tcp", firstConn.LocalAddr().String())
	if err != nil {
		log.Println("Could not open listening socket! " + err.Error())
		return
	}

	log.Println("Listening on address " + firstConn.LocalAddr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleConnection(conn)
	}
}
