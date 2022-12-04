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

	var parentDomain string
	flag.StringVar(&parentDomain, "d", "", "The parent domain to match against")
	flag.Parse()

	parentDomain = strings.ToLower(parentDomain)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		subdomain := scanner.Text()
		if strings.HasSuffix(strings.ToLower(subdomain), "."+parentDomain) || subdomain == parentDomain {
			fmt.Println(strings.ToLower(subdomain))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
