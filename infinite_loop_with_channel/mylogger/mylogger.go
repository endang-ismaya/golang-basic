package mylogger

import "log"

func ListenForLog(ch chan string) {

	for {
		msg := <-ch
		log.Printf("We got your message: %s\n", msg)
	}
}
