package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf(
		`current time: %s
exact time: %s
`, time.Now(), t.Round(0))
}
