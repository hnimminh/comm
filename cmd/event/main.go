package main

import (
	"fmt"
	"log"

	"github.com/hnimminh/comm/pkg/fsesl"
)

func main(){
	conn, err := fsesl.Dial("127.0.0.1:8021", "ramdomstr")
	if err != nil {
		log.Fatal(err)
	}

	conn.Send("events json ALL")
	for {
		event, err := conn.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nNew event")
		event.PrettyPrint()
	}
	conn.Close()
}
