package main

import (
	"fmt"
	"server-watcher/lib"
	"time"
)

var (
	target     = "https://symbolnode.blockchain-authn.app:3001/node/health"
	recipients = []string{"ym.u.ichiro@gmail.com"}
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
