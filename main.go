package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("no args")
		os.Exit(1)
	}

	runes := []rune(args[0])
	for _, r := range runes {
		var byteString []string
		for _, b := range []byte(string(r)) {
			byteString = append(byteString, fmt.Sprintf("%x", b))
		}

		fmt.Printf("%s", string(r))
		fmt.Print("|-|")
		fmt.Printf("%d", r)
		fmt.Print("|-|")
		fmt.Printf("%s\n", strings.Join(byteString, "|"))
	}

}
