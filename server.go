package main

import (
	"log"
	"os"
	"time"

	qp "github.com/quic-s/quics-protocol"
)

var host string
var port string

func init() {
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
}

func RestClientMessage(msgtype string, message []byte) {
	// initialize client
	quicClient, err := qp.New()
	if err != nil {
		log.Panicln(err)
	}

	// start client
	err = quicClient.Dial(host + ":" + port)
	if err != nil {
		log.Panicln(err)
	}

	// send message to server
	quicClient.SendMessage(msgtype, message)

	// delay for waiting message sent to server
	time.Sleep(3 * time.Second)
	quicClient.Close()
}

func RestClientFile(filepath string) {

	// initialize client
	quicClient, err := qp.New()
	if err != nil {
		log.Panicln(err)
	}

	// start client
	err = quicClient.Dial(host + ":" + port)
	if err != nil {
		log.Panicln(err)
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Panicln(err)

	}
	quicClient.SendFile(LocalAbsToRoot(filepath, getAccelerDir()), file)

	// delay for waiting message sent to server
	time.Sleep(3 * time.Second)
	quicClient.Close()
}
