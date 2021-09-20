package main

import (
	"fmt"
	"log"

	"github.com/hnimminh/comm/pkg/fsesl"
)

const audioFile = "/opt/media/BigGirlsCry-Sia-1630748496.wav"

func main() {
	fsesl.ListenAndServe("127.0.0.1:9021", handler)
}

func handler(conn *fsesl.Connection) {
	fmt.Println("new client:", conn.RemoteAddr())
	conn.Send("connect")
	conn.Send("myevents")
	conn.Execute("answer", "", false)
	ev, err := conn.Execute("playback", audioFile, true)
	if err != nil {
		log.Fatal(err)
	}
	ev.PrettyPrint()
	for {
		ev, err = conn.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nNew event")
		ev.PrettyPrint()
		if ev.Get("Application") == "playback" {
			if ev.Get("Application-Response") == "FILE PLAYED" {
				conn.Send("exit")
			}
		}
	}
}
