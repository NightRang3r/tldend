package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var tldDomain bool
	flag.BoolVar(&tldDomain, "d", false, "top level domain")
	flag.Parse()

	domain := flag.Arg(0)
	if domain == "" {
		fmt.Println("Please specify top level domain to match")
		return
	}
	domain = strings.ToLower(domain)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.HasSuffix(strings.ToLower(scanner.Text()), domain) {
			fmt.Println(strings.ToLower(scanner.Text()))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
