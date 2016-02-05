package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"flag"
	"time"
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
		//        var message string
		//        _ , _ = fmt.Scanln("%s", &message)
		//        fmt.Fprintf(conn, message + "\n")
		fmt.Fprintf(conn, "client hello \n")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.SetFlags(log.Flags()|log.Lshortfile)

	meetAddrFmt := "http://%s/meet/"
	meetAddr := "localhost:8080"
	var clientPairId string

	flag.StringVar(&meetAddr, "meetAddr", meetAddr, "Server for discover other")
	flag.StringVar(&clientPairId, "pair", "", "key to macth service")
	flag.Parse()

	meetAddr = fmt.Sprintf(meetAddrFmt, meetAddr)

	log.Printf("Discover server is %s\n", meetAddr)
	log.Printf("unique ID for the client pair: %s\n", clientPairId)
	if len(clientPairId) == 0{
		log.Fatal("Must config the pair")
	}

	response, err := http.Get(meetAddr + clientPairId)
	if err != nil{
		log.Fatal(err)
	}
	address, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Dial %s\n", string(address))
	conn, err := net.Dial("tcp", string(address))
	if err != nil{
		log.Fatal(err)
	}

	go readSocket(conn)
	writeSocket(conn)
}
