package main

import (
	"go-smtp-study/lib"
)

func main() {

	res := lib.DeadOrLive("https://symbolnode.blockchain-authn.app:3001/node/health")

	if !res {
		lib.SendGmail([]string{"ym.u.ichiro@gmail.com"})
	} else {
		println("ok")
	}
}
