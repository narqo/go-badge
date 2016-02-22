package main

import (
	"flag"
	"log"
	"os"

	"github.com/narqo/go-badge/badge"
)

var (
	subject = flag.String("subject", "", "Badge subject")
	status = flag.String("status", "", "Badge status")
	//color = flag.String("color", "red", "Badge color")
)

func main() {
	flag.Parse()
	err := badge.Render(*subject, *status, badge.ColorBlue, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
