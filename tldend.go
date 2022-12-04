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

	var tldDomain string
	flag.StringVar(&tldDomain, "d", "", "top level domain")
	flag.Parse()

	tldDomain = strings.ToLower(tldDomain)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		subdomain := scanner.Text()
		if strings.HasSuffix(strings.ToLower(subdomain), "."+tldDomain) || subdomain == tldDomain {
			fmt.Println(strings.ToLower(subdomain))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
