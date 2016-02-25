package main

import (
	"flag"
	"log"
	"os"

	"github.com/narqo/go-badge"
)

var (
	subject = flag.String("subject", "", "Badge subject")
	status = flag.String("status", "", "Badge status")
	color = flag.String("color", "blue", "Badge color")
)

func main() {
	flag.Parse()
	err := badge.Render(*subject, *status, badge.Color(*color), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
