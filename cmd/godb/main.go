package main

import (
	"flag"
	"fmt"
	"godb"
	"log"
	"os"
)

var (
	user     = ""
	password = ""
	version  bool
)

func init() {
	flag.StringVar(&user, "user", "", "username")
	flag.StringVar(&password, "password", "", "password for the user")
	flag.BoolVar(&version, "v", false, "print version")
}

func main() {
	flag.Parse()

	if version {
		fmt.Println("v0.2.0")
		os.Exit(0)
	}

	fmt.Printf("user=%s, password=%s\n", user, password)
	if err := godb.Run(user, password); err != nil {
		log.Fatalln(err)
	}
}
