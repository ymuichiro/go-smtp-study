package main

import (
	"fmt"
	"os"
	"server-watcher/lib"
	"time"
)

var (
	target     = os.Getenv("TARGET_URL")
	recipients = []string{os.Getenv("GMAIL_TO")}
)

func main() {

	ticker := time.NewTicker(30 * time.Minute)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:

			fmt.Printf("check server health: %s\n", target)
			res := lib.DeadOrLive(target)

			if !res {
				fmt.Printf("server is down: %s\n", target)
				lib.SendGmail(recipients)
			} else {
				fmt.Printf("server is up: %s\n", target)
			}

		}
	}
}
