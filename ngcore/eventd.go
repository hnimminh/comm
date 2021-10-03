package ngcore

import (
	"fmt"
	"log"
	"time"

	"github.com/hnimminh/comm/pkg/fsesl"
)

func Eventd(){
	conn, err := fsesl.Dial("127.0.0.1:8021", "ramdomstr")
	for {
		if err != nil {
			fmt.Println("there is an error ", err, "wait for 20")
			time.Sleep(20 * time.Second)
			conn, err = fsesl.Dial("127.0.0.1:8021", "ramdomstr")
		} else {
			break
		}
	}

	conn.Send("events json ALL")
	for {
		event, err := conn.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("\nNew event")
		eventname := event.Get("Event-Name")
		legid := event.Get("Unique-ID")

		fmt.Println("----------------------------------------------\nReceive new event with name ", eventname, "id", legid)

		switch eventname {
		case "CHANNEL_PARK":
			r,e := conn.ExecuteUUID(legid, "answer", "")
			fmt.Println("answer ", r, "err", e)
			//event.PrettyPrint()
		case "CHANNEL_ANSWER":
			r,e := conn.ExecuteUUID(legid, "playback", "/opt/media/BigGirlsCry-Sia-1630748496.wav")
			fmt.Println("play ", r, "err", e)
			//event.PrettyPrint()
		default:
			fmt.Print("")
		}
	}
	conn.Close()
}
