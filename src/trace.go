package src

import (
	"log"
	"time"
)

func Trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
